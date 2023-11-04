package app

type SqlModel struct {
	Id       int `gorm:"primaryKey;autoIncrement"`
	FullUrl  string
	ShortUrl string
}
