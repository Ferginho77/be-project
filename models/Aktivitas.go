package models

type Aktivitas struct {
	AktivitasId uint   `json:"AktivitasId" gorm:"column:AktivitasId;primaryKey"`
	Tanggal   string `json:"Tanggal"   gorm:"column:Tanggal"`
	JenisAktivitas        string `json:"JenisAktivitas"        gorm:"column:JenisAktivitas;type:enum('Pupuk','Obat','Benih');default:'Obat'"`
	Keterangan         string   `json:"Keterangan"         gorm:"column:Keterangan"`
	PenanamanId        uint     `json:"PenanamanId"        gorm:"column:PenanamanId"`
}

func (Aktivitas) TableName() string {
	return "aktivitas"
}