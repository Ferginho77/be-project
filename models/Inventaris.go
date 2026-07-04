package models

type Inventaris struct {
	InventarisId uint   `json:"InventarisId" gorm:"column:InventarisId;primaryKey"`
	NamaBarang   string `json:"NamaBarang"   gorm:"column:NamaBarang"`
	Jenis        string `json:"Jenis"        gorm:"column:Jenis;type:enum('Pupuk','Obat','Benih');default:'Obat'"`
	Stok         uint   `json:"Stok"         gorm:"column:Stok"`
}

func (Inventaris) TableName() string {
	return "inventaris"
}