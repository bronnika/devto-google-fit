package models

const (
	HeightDataType    = "height"
	WeightDataType    = "weight"
	HeartRate = "heart_rate.bpm"
	HeartPoints = "heart_minutes"
	ActiveMinutes = "active_minutes"
	Steps = "step_count.delta"
	ActivitySegment = "activity.segment"
	CaloriesBurnt = "calories.expended"
	Hydration = "hydration"
	Nutrition = "nutrition"
)


var MealType = map[int]string {
	0: "Unknown",
	1: "Breakfast",
	2: "Lunch",
	3: "Dinner",
	4: "Snack",
}

const (
	MealTypeUnknown   = "Unknown"
	MealTypeBreakfast = "Breakfast"
	MealTypeLunch     = "Lunch"
	MealTypeDinner    = "Dinner"
	MealTypeSnack     = "Snack"

	MealTypeUnknownInt   = 0
	MealTypeBreakfastInt = 1
	MealTypeLunchInt     = 2
	MealTypeDinnerInt    = 3
	MealTypeSnackInt     = 4

	NutrientCalories           = "calories"        // Calories in kcal
	NutrientTotalFat           = "fat.total"       // Total fat in grams
	NutrientSaturatedFat       = "fat.saturated"   // Saturated fat in grams
	NutrientUnsaturatedFat     = "fat.unsaturated" // Unsaturated fat in grams
	// NutrientPolyunsaturatedFat NOTICE: it is strange that in taken data from Fit API it's exactly "polysaturated"
	// NutrientPolyunsaturatedFat NOTICE: not "polyUNsaturated" as it has to be
	NutrientPolyunsaturatedFat = "fat.polysaturated"   // Polyunsaturated fat in grams.
	NutrientMonounsaturatedFat = "fat.monounsaturated" // Monounsaturated fat in grams
	NutrientTransFat           = "fat.trans"           // Trans fat in grams
	NutrientCholesterol        = "cholesterol"         // Cholesterol in milligrams
	NutrientSodium             = "sodium"              // Sodium in milligrams
	NutrientPotassium          = "potassium"           // Potassium in milligrams
	NutrientTotalCarbs         = "carbs.total"         // Total carbohydrates in grams
	NutrientDietaryFiber       = "dietary_fiber"       //  Dietary fiber in grams
	NutrientSugar              = "sugar"               // Amount of sugar in grams
	NutrientProtein            = "protein"             // Protein amount in grams
)


var ActivitySegmentMap = map[int64]string{
	0: "In Vehicle",
	1: "Biking",
	3: "Still",
	4: "Unknown",
	5: "Tilting Sudden Device Gravity Change",
	7: "Walking",
	8: "Running",
	9: "Aerobics",
	10: "Badminton",
	11: "Baseball",
	12: "Basketball",
	13: "Biathlon",
	14: "Hand Biking",
	15: "Mountain Biking",
	16: "Road Biking",
	17: "Spinning",
	18: "Stationary Biking",
	19: "Utility Biking",
	20: "Boxing",
	21: "Calisthenics",
	22: "Circuit Training",
	23: "Cricket",
	24: "Dancing",
	25: "Elliptical",
	26: "Fencing",
	27: "Football American",
	28: "Football Australian",
	29: "Football Soccer",
	30: "Frisbee",
	31: "Gardening",
	32: "Golf",
	33: "Gymnastics",
	34: "Handball",
	35: "Hiking",
	36: "Hockey",
	37: "Horseback Riding",
	38: "Housework",
	39: "Jumping Rope",
	40: "Kayaking",
	41: "Kettlebell Training",
	42: "Kickboxing",
	43: "Kitesurfing",
	44: "Martial Arts",
	45: "Meditation",
	46: "Mixed Martial Arts",
	47: "P90X Exercises",
	48: "Paragliding",
	49: "Pilates",
	50: "Polo",
	51: "Racquetball",
	52: "Rock Climbing",
	53: "Rowing",
	54: "Rowing Machine",
	55: "Rugby",
	56: "Jogging",
	57: "Running On Sand",
	58: "Running Treadmill",
	59: "Sailing",
	60: "Scuba Diving",
	61: "Skateboarding",
	62: "Skating",
	63: "Cross Skating",
	64: "Inline Skating Rollerblading",
	65: "Skiing",
	66: "Back Country Skiing",
	67: "Cross Country Skiing",
	68: "Downhill Skiing",
	69: "Kite Skiing",
	70: "Roller Skiing",
	71: "Sledding",
	73: "Snowboarding",
	74: "Snowmobile",
	75: "Snowshoeing",
	76: "Squash",
	77: "Stair Climbing",
	78: "Stair Climbing Machine",
	79: "Stand Up Paddle Boarding",
	80: "Strength Training",
	81: "Surfing",
	82: "Swimming",
	84: "Swimming Open Water",
	83: "Swimming Swimming Pool",
	85: "Table Tennis",
	86: "Team Sports",
	87: "Tennis",
	88: "Treadmill",
	89: "Volleyball",
	90: "Volleyball Beach",
	91: "Volleyball Indoor",
	92: "Wakeboarding",
	93: "Walking Fitness",
	94: "Nording Walking",
	95: "Walking Treadmill",
	96: "Water Polo",
	97: "Weightlifting",
	98: "Wheelchair",
	99: "Windsurfing",
	100: "Yoga",
	101: "Zumba",
	102: "Diving",
	103: "Ergometer",
	104: "Ice Skating",
	105: "Indoor Skating",
	106: "Curling",
	108: "Other",
	113: "Crossfit",
	114: "HIIT",
	115: "Interval Training",
	116: "Walking Stroller",
	117: "Elevator",
	118: "Escalator",
	119: "Archery",
	120: "Softball",
	122: "Guided Breathing",
}

const (
	ActivitySegmentType = "activity_segment"
)