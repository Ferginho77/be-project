package models


type Produksi struct {
    ProduksiId     uint    `gorm:"column:ProduksiId;primaryKey;autoIncrement" json:"ProduksiId"`
    TotalPanen   float64 `gorm:"column:TotalPanen" json:"TotalPanen"`
    Tanggal      string `gorm:"column:Tanggal" json:"Tanggal"`
    JumlahBuah     uint  `gorm:"column:JumlahBuah" json:"JumlahBuah"`
    PenanamanId uint  `gorm:"column:PenanamanId" json:"PenanamanId"`
}

func (Produksi) TableName() string {
    return "produksi"
}