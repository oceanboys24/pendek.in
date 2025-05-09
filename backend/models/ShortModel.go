package models


type ShortModel struct {
	Id        string `json:"id" gorm:"primaryKey"`
	OriginUrl string `json:"origin_url"`
	ShortUrl  string `json:"short_url"`
}

type ShortModelPayload struct {
	Url string `json:"url"`
}