package models

type Invenusage struct {
    UsageId            uint    `gorm:"column:UsageId;primaryKey;autoIncrement" json:"UsageId"`
	Jumlah 				uint   `gorm:"column:Jumlah" json:"Jumlah"`
    Tanggal            string `gorm:"column:Tanggal" json:"Tanggal"`
    PenanamanId 		uint   `gorm:"column:PenanamanId" json:"PenanamanId"`
    InventarisId 		uint   `gorm:"column:InventarisId" json:"InventarisId"`
}

func (Invenusage) TableName() string {
    return "invenusage"
}