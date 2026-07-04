package models

type Tanaman struct {
	TanamanId   uint   `json:"TanamanId" gorm:"column:TanamanId;primaryKey"`
	NamaTanaman string `json:"NamaTanaman" gorm:"column:NamaTanaman"`
	UmurPanen uint  `json:"UmurPanen" gorm:"column:UmurPanen"`
	Deskripsi string `json:"Deskripsi" gorm:"column:Deskripsi"`
}

func (Tanaman) TableName() string {
	return "tanaman"
}