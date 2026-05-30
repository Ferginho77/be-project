package models

type Inventaris struct {
	InventarisId uint `json:"InventarisId" gorm:"column:InventarisId;primaryKey"`
	NamaBarang string `json:"NamaBarang"`
	Jenis string `json:"Jenis" gorm:"type:enum('Pupuk','Obat');default:'Obat'"` 
	Stok uint `json:"Stok"`
}

func (Inventaris) TableName() string {
	return "inventaris"
}