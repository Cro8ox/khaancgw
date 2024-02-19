package utils

type API struct {
	Url    string `json:"url"`
	Method string `json:"method"`
}

const (
	TimeFormatYYYYMMDDHHMMSS = "20060102150405"
	TimeFormatYYYYMMDD       = "20060102"
	HttpContent              = "application/json"
	XmlContent               = "application/xml"
)
