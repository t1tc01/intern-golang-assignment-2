package model

type Properties struct {
	Mag     float64 `json:"mag"`
	Place   string  `json:"place"`
	Time    int64   `json:"time"`
	Updated int64   `json:"updated"`
	Tz      int32   `json:"tz"`
	URL     string  `json:"url"`
	Detail  string  `json:"detail"`
	Felt    int32   `json:"felt"`
	CDI     float64 `json:"cdi"`
	MMI     float64 `json:"mmi"`
	Alert   string  `json:"alert"`
	Status  string  `json:"status"`
	Tsunami int32   `json:"tsunami"`
	Sig     int32   `json:"sig"`
	Net     string  `json:"net"`
	Code    string  `json:"code"`
	IDs     string  `json:"ids"`
	Sources string  `json:"sources"`
	Types   string  `json:"types"`
	NST     int32   `json:"nst"`
	Dmin    float64 `json:"dmin"`
	RMS     float64 `json:"rms"`
	Gap     float64 `json:"gap"`
	MagType string  `json:"magType"`
	Type    string  `json:"type"`
}

type Geometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type Feature struct {
	Type       string     `json:"type"`
	Properties Properties `json:"properties"`
	Geometry   Geometry   `json:"geometry"`
	ID         string     `json:"id"`
}

type GeoJSON struct {
	Type     string    `json:"type"`
	Features []Feature `json:"features"`
}
