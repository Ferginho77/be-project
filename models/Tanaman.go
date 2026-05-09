package models

type Tanaman struct {
	TanamanId   uint   `json:"TanamanId" gorm:"primaryKey"`
	NamaTanaman string `json:"NamaTanaman"`
	Status string `json:"Status" gorm:"type:enum('Baik','Buruk');default:'Baik'"`
	LuasTanam uint  `json:"LuasTanam"`
}

func (Tanaman) TableName() string {
	return "tanaman"
}