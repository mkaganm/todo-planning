package data

type Developer struct {
	ID         uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name       string `gorm:"type:varchar(255);not null" json:"name"`
	Duration   string `gorm:"type:varchar(255);not null" json:"duration"`
	Difficulty string `gorm:"type:varchar(10);not null" json:"difficulty"`
	Estimation int    `gorm:"default:0" json:"estimation"`
}
