package model

type EarthquakeModel struct {
	ID          int
	Mag         float64 `json:"mag"`
	Time        int64   `json:"time"`
	Updated     int64   `json:"updated"`
	Tz          int32   `json:"tz"`
	URL         string  `json:"url"`
	Detail      string  `json:"detail"`
	Status      string  `json:"status"`
	Tsunami     int32   `json:"tsunami"`
	Sig         int32   `json:"sig"`
	Net         string  `json:"net"`
	Code        string  `json:"code"`
	NST         int32   `json:"nst"`
	Dmin        float64 `json:"dmin"`
	RMS         float64 `json:"rms"`
	Gap         float64 `json:"gap"`
	MagType     string  `json:"magType"`
	Type        string  `json:"type"`
	CreatedTime int64   `json:"createdTime"`
	UpdatedTime int64   `json:"updatedTime"`
	DeletedTime int64   `json:"deletedTime"`
}
