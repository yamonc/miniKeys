package main
type TKey struct {
	ID uint `gorm:"primaryKey;AUTO_INCREMENT"`
	KeyName string
	Url string
	Password string
	Account string
	Back string
	Category uint
	UserId uint
}
