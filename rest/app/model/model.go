package app

// SqlModel todo:дописать модель
type SqlModel struct {
	Id       int `gorm:"primaryKey;autoIncrement"`
	FullUrl  string
	ShortUrl string
}
