package models


type Lahan struct {
	LahanId uint `json:"LahanId" gorm:"column:LahanId;primaryKey"`
	NamaLahan string `json:"NamaLahan"`
	Lokasi string `json:"Lokasi"`
	LuasTanah uint `json:"LuasTanah"`
	Kondisi string `json:"Kondisi" gorm:"type:enum('Baik','Buruk');default:'Baik'"`
}

func (Lahan) TableName() string {
	return "lahan"
}