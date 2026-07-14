package models

//type Inventaris struct {
//	InventarisId uint   `json:"InventarisId" gorm:"column:InventarisId;primaryKey"`
//	NamaBarang   string `json:"NamaBarang"   gorm:"column:NamaBarang"`
//	Jenis        string `json:"Jenis"        gorm:"column:Jenis;type:enum('Pupuk','Obat','Benih');default:'Obat'"`
//	Stok         uint   `json:"Stok"         gorm:"column:Stok"`
//}

//func (Inventaris) TableName() string {
//	return "inventaris"
//}

type Inventaris struct {
	InventarisId uint   `json:"InventarisId" gorm:"column:InventarisId;primaryKey"`
	NamaBarang   string `json:"NamaBarang"   gorm:"column:NamaBarang"`
	Jenis        string `json:"Jenis"        gorm:"column:Jenis;type:enum('Pupuk','Obat','Alat-alat','Benih');default:'Obat'"`
	Stok         uint   `json:"Stok"         gorm:"column:Stok"`
}

func (Inventaris) TableName() string {
	return "inventaris"
}

// Di database ada 4 pilihan, tapi di codingan golang yang sebelumya cuman 3. sama aku ditambahin 'Alat-alat' biar sama kaya yang ada di database