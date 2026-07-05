package model

type UrlVisitLog struct {
	ID        int    `json:"id"`
	ShortCode string `json:"short_code"`
	Ip        string `json:"ip"`
	Region    string `json:"region"`
	UserAgent string `json:"user_agent"`
	Device    string `json:"device"`
	Brower    string `json:"brower"`
	Os        string `json:"os"`
}

func (UrlVisitLog) TableName() string {
	return "url_visit_log"
}
