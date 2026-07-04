package models


type Penanaman struct {
	// Primary Key
	PenanamanId uint `json:"PenanamanId" gorm:"column:PenanamanId;primaryKey"`
	TanggalTanam string `json:"TanggalTanam" gorm:"column:TanggalTanam;type:date"`
	RencanaPanen string `json:"RencanaPanen" gorm:"column:RencanaPanen;type:date"`
	JumlahBibit uint `json:"JumlahBibit" gorm:"column:JumlahBibit;default:0"`
	TanamanId uint `json:"TanamanId" gorm:"column:TanamanId"`
	LahanId   uint `json:"LahanId" gorm:"column:LahanId"`
	Fase string `json:"Fase" gorm:"column:Fase;type:enum('Vegetatif','Generatif','Panen');default:'Vegetatif'"`
	Status string `json:"Status" gorm:"column:Status;type:enum('Aktif','Panen', 'Gagal', 'Selesai');default:'Aktif'"`
}

func (Penanaman) TableName() string {	
	return "penanaman"
}

