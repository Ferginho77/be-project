package models


type Penanaman struct {
	PenanamanId uint `json:"PenanamanId" gorm:"column:PenanamanId;primaryKey"`
	TanggalTanam string `json:"TanggalTanam"`
	RencanaPanen string `json:"RencanaPanen"`
	JumlahBibit uint `json:"JumlahBibit"`	
	TanamanId uint `json:"TanamanId"`
	LahanId uint `json:"LahanId"`
	Fase string  `json:"Fase" gorm:"type:enum('Vegetatif','Generatif','Panen');default:'Vegetatif'"`
}

func (Penanaman) TableName() string {	
	return "penanaman"
}