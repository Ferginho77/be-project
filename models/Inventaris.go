package models

type Inventaris struct {
	IdBarang uint `json:"IdBarang" gorm:"primaryKey"`
	NamaBarang string `json:"NamaBarang"`
	Jenis string `json:"Jenis" gorm:"type:enum('Pupuk','Obat');default:'Obat'"` 
}

func (Inventaris) TableName() string {
	return "inventaris"
}