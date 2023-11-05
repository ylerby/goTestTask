package app

type SqlModel struct {
	Id       int    `gorm:"primaryKey;autoIncrement"`
	FullUrl  string `gorm:"column:full_url"`
	ShortUrl string `gorm:"column:short_url"`
}
