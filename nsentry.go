package main

import "time"

type NsEntry struct {
	Device  string
	OpenAps struct {
		Suggested struct {
			Temp             string    `json:"temp" bson:"temp"`
			Bg               float64   `json:"bg" bson:"bg"`
			Tick             float64   `json:"-" bson:"-"`
			EventualBG       float64   `json:"eventualBG" bson:"eventualBG"`
			TargetBG         float64   `json:"targetBG" bson:"targetBG"`
			InsulinReq       float64   `json:"insulinReq" bson:"insulinReq"`
			DeliverAt        time.Time `json:"deliverAt" bson:"deliverAt"`
			SensitivityRatio float64   `json:"sensitivityRatio" bson:"sensitivityRatio"`
			PredBGs          struct {
				IOB []float64 `json:"IOB"`
				ZT  []float64 `json:"ZT"`
				COB []float64 `json:"COB"`
				UAM []float64 `json:"UAM"`
			} `json:"predBGs"`
			COB       float64   `json:"COB"`
			IOB       float64   `json:"IOB"`
			Reason    string    `json:"reason"`
			Units     float64   `json:"units"`
			Rate      float64   `json:"rate"`
			Duration  int       `json:"duration"`
			Timestamp time.Time `json:"timestamp"`
		} `json:"suggested,omitempty" bson:"suggested,omitempty"`
		IOB struct {
			IOB      float64   `json:"iob" bson:"iob"`
			BasalIOB float64   `json:"basaliob" bson:"basaliob"`
			Activity float64   `json:"activity" bson:"activity"`
			Time     time.Time `json:"time" bson:"time"`
		} `json:"iob" bson:"iob"`
	} `json:"openaps" bson:"openaps"`
	Pump struct {
		Clock     time.Time `json:"clock"`
		Reservoir float64   `json:"reservoir"`
		Status    struct {
			Status    string `json:"status"`
			Timestamp int64  `json:"-" bson:"-"`
		} `json:"status"`
		Extended struct {
			Version               string  `json:"Version"`
			ActiveProfile         string  `json:"ActiveProfile"`
			TempBasalAbsoluteRate float64 `json:"TempBasalAbsoluteRate"`
			TempBasalPercent      int     `json:"TempBasalPercent"`
			TempBasalRemaining    int     `json:"TempBasalRemaining"`
		} `json:"extended"`
		Battery struct {
			Percent int `json:"percent"`
		} `json:"battery"`
	} `json:"pump"`
	User string `json:"-"`
}

type NsTreatment struct {
	CreatedAt time.Time `json:"created_at"`
	// Who entered the treatment.
	EnteredBy string `json:"enteredBy"`
	// The type of treatment event.
	// example: "BG Check", "Snack Bolus", "Meal Bolus", "Correction Bolus",
	// "Carb Correction", "Combo Bolus", "Announcement", "Note", "Question",
	// "Exercise", "Site Change", "Sensor Start", "Sensor Change",
	// "Pump Battery Change", "Insulin Change", "Temp Basal", "Profile Switch",
	// "D.A.D. Alert", "Temporary Target", "OpenAPS Offline", "Bolus Wizard"
	EventType string `json:"eventType"`
	// Amount of carbs given.
	Carbs int `json:"carbs,omitempty"`
	// Duration in minutes.
	Duration float32 `json:"duration,omitempty"`
	// Amount of insulin, if any.
	Insulin float64 `json:"insulin,omitempty"`
	IsSMB   bool    `json:"isSMB,omitempty"`
	// Description/notes of treatment.
	Notes string `json:"notes,omitempty"`
	// Eventual basal change in percent.
	Percent int `json:"percent,omitempty"`
	// Top limit of temporary target.
	TargetTop float64 `json:"targetTop,omitempty"`
	// Bottom limit of temporary target.
	TargetBottom float64 `json:"targetBottom,omitempty"`
	// For example the reason why the profile has been switched or why the temporary target has been set.
	Reason string  `json:"reason,omitempty"`
	Rate   float64 `json:"rate,omitempty"`
	// The units for the glucose value, mg/dl or mmol/l.
	// It is strongly recommended to fill in this field when glucose is entered.
	// example: "mg/dl", "mmol/l"
	Units string `json:"units,omitempty"`
	// Current glucose.
	Glucose string `json:"glucose,omitempty"`
	// Method used to obtain glucose, Finger or Sensor.
	// example: "Sensor", "Finger", "Manual"
	GlucoseType string `json:"glucoseType,omitempty"`
	User        string `json:"-"`
}

// "2026-01-31T20:00:49.000Z"      1769889649000   233     "Flat"  "unknown"
// "2026-01-31T19:55:48.000Z"      1769889348000   236     "FortyFiveUp"   "unknown"
//
//	{"_id":"697e72528e5493c015a6ad1b","sgv":206,
//	 "date":1769894450000,"dateString":"2026-01-31T21:20:50.000Z",
//	 "trend":5,"direction":"FortyFiveDown","device":"unknown",
//	 "type":"sgv","utcOffset":0,"sysTime":"2026-01-31T21:20:50.000Z","mills":1769894450000
//	}

// Blood glucose measurements and CGM calibrations
type NsMeasurement struct {
	TimestampStr string `json:"dateString,omitempty"`
	TimestampMs  int64  `json:"date,omitempty"`
	// The glucose reading. (only available for sgv types)
	SensorGlucose int `json:"sgv,omitempty"`
	// The units for the glucose value, mg/dl or mmol/l.
	// example: "mg", "mmol"
	Units string `json:"units,omitempty"`
	// Noise level at time of reading. (only available for sgv types)
	Noise int `json:"noise,omitempty"`
	Trend int `json:"trend,omitempty"`
	// Direction of glucose trend reported by CGM. (only available for sgv types)
	// example: "DoubleDown", "SingleDown", "FortyFiveDown", "Flat",
	// "FortyFiveUp", "SingleUp", "DoubleUp",
	// "NOT COMPUTABLE", "RATE OUT OF RANGE" for xdrip
	Arrow string `json:"direction,omitempty"`
	// The device from which the data originated (including serial number of the device, if it is relevant and safe).
	// example: dexcom G5
	Device string `json:"device,omitempty"`
	// "sgv", "mbg", "cal", "etc"
	Type string `json:"type,omitempty"`
	User string `json:"-"`
}

type Config struct {
	NsUri        string `json:"ns-uri,omitempty"`
	NsToken      string `json:"ns-token,omitempty"`
	MongoUri     string `json:"mongo-uri,omitempty"`
	MongoDb      string `json:"mongo-db,omitempty"`
	Limit        int64  `json:"limit,omitempty"`
	Skip         int64  `json:"skip,omitempty"`
	InfluxUri    string `json:"influx-uri,omitempty"`
	InfluxToken  string `json:"influx-token,omitempty"`
	InfluxOrg    string `json:"influx-org,omitempty"`
	InfluxBucket string `json:"influx-bucket,omitempty"`
	Imports      []struct {
		NsUri    string `json:"ns-uri,omitempty"`
		NsToken  string `json:"ns-token,omitempty"`
		MongoUri string `json:"mongo-uri,omitempty"`
		MongoDb  string `json:"mongo-db,omitempty"`
		User     string `json:"user"`
	} `json:"imports,omitempty"`
}
