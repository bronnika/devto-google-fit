package google_api

import (
	"context"
	"encoding/gob"
	"errors"
	"flag"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"hash/fnv"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

const (
	nanosPerMilli = 1e6
)

var (
	pwd, _ = os.Getwd()
	clientIDFile = flag.String("clientid-file", pwd+"/.secrets/clientid.dat" ,
		"Name of a file containing just the project's OAuth 2.0 Client ID from https://developers.google.com/console.")
	secretFile = flag.String("secret-file", pwd+"/.secrets/clientsecret.dat",
		"Name of a file containing just the project's OAuth 2.0 Client Secret from https://developers.google.com/console.")
	cacheToken = flag.Bool("cachetoken", true, "cache the OAuth 2.0 token")
	debug      = flag.Bool("debug", false, "show HTTP traffic")
)

// TimeToMillis converts time.Time to Unix millis
func TimeToMillis(time2 time.Time) int64 {
	return time2.UnixNano() / int64(nanosPerMilli)
}

// MillisToTime converts Unix millis to time.Time.
func MillisToTime(t int64) time.Time {
	return time.Unix(0, t*nanosPerMilli)
}

// NanosToTime converts Unix nanos to time.Time
func NanosToTime(t int64) time.Time {
	return time.Unix(0, t)
}

// TimeToNanos coverts time.Time to Unix nanos
func TimeToNanos(time2 time.Time) int64 {
	return time2.UnixNano()
}

// authClient returns HTTP client using Google token from cache or web
func authClient(ctx context.Context, config *oauth2.Config) *http.Client {
	cacheFile := tokenCacheFile(config)
	token, err := tokenFromFile(cacheFile)
	if err != nil {
		token = tokenFromWeb(ctx, config)
		saveToken(cacheFile, token)
	} else {
		log.Printf("Using cached token %#v from %q", token, cacheFile)
	}

	return config.Client(ctx, token)
}

// fileContent reads file
func fileContent(filename string) (string, error) {
	slurp, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(slurp)), nil
}

// returnConfig returns Config struct with clientID and clientSecret of Google Fit API depending on scopes
func returnConfig(scopes []string) (*oauth2.Config, error) {
	clientId, err := fileContent(*clientIDFile)
	if err != nil {
		return nil, errors.New("no such file")
	}
	clientSecret, err := fileContent(*secretFile)
	if err != nil {
		return nil, errors.New("no such file")
	}
	config := &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Endpoint:     google.Endpoint,
		Scopes: 	  scopes,
	}
	return config, nil
}

// tokenCacheFile returns path of the file with user's token
func tokenCacheFile(config *oauth2.Config) string {
	hash := fnv.New32a()
	hash.Write([]byte(config.ClientID))
	hash.Write([]byte(config.ClientSecret))
	hash.Write([]byte(strings.Join(config.Scopes, " ")))
	fn := fmt.Sprintf("upload-tok%v", hash.Sum32())
	return filepath.Join(".secrets/", url.QueryEscape(fn))
}

// tokenFromFile returns token from file
func tokenFromFile(file string) (*oauth2.Token, error) {
	if !*cacheToken {
		return nil, errors.New("--cachetoken is false")
	}
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	t := new(oauth2.Token)
	err = gob.NewDecoder(f).Decode(t)
	return t, err
}

// tokenFromWeb returns token if not found in cache
// it opens Google auth page in a browser
func tokenFromWeb(ctx context.Context, config *oauth2.Config) *oauth2.Token {
	ch := make(chan string)
	randState := fmt.Sprintf("st%d", time.Now().UnixNano())
	ts := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.Path == "/favicon.ico" {
			http.Error(rw, "", 404)
			return
		}
		if req.FormValue("state") != randState {
			log.Printf("State doesn't match: req = %#v", req)
			http.Error(rw, "", 500)
			return
		}
		if code := req.FormValue("code"); code != "" {
			fmt.Fprintf(rw, "<h1>Success</h1>Authorized.")
			rw.(http.Flusher).Flush()
			ch <- code
			return
		}
		log.Printf("no code")
		http.Error(rw, "", 500)
	}))
	defer ts.Close()

	config.RedirectURL = ts.URL
	authURL := config.AuthCodeURL(randState, oauth2.AccessTypeOffline)
	go openURL(authURL)
	log.Printf("Authorize this app at: %s", authURL)
	code := <-ch
	log.Printf("Got code: %s", code)

	token, err := config.Exchange(ctx, code)

	if err != nil {
		log.Fatalf("Token exchange error: %v", err)
	}
	return token
}

func openURL(url string) {
	try := []string{"xdg-open", "google-chrome", "open"}
	for _, bin := range try {
		err := exec.Command(bin, url).Run()
		if err == nil {
			return
		}
	}
	log.Printf("Error opening URL in browser.")
}

// saveToke saves token in a new generated file in .secrets/
func saveToken(file string, token *oauth2.Token) {
	f, err := os.Create(file)
	if err != nil {
		log.Printf("Warning: failed to cache oauth token: %v", err)
		return
	}
	defer f.Close()
	gob.NewEncoder(f).Encode(token)
}