package models

//type Aktivitas struct {
//	AktivitasId uint   `json:"AktivitasId" gorm:"column:AktivitasId;primaryKey"`
//	Tanggal   string `json:"Tanggal"   gorm:"column:Tanggal"`
//	JenisAktivitas        string `json:"JenisAktivitas"        gorm:"column:JenisAktivitas;type:enum('Pupuk','Obat','Benih');default:'Obat'"`
//	Keterangan         string   `json:"Keterangan"         gorm:"column:Keterangan"`
//	PenanamanId        uint     `json:"PenanamanId"        gorm:"column:PenanamanId"`
//}

//func (Aktivitas) TableName() string {
//	return "aktivitas"
//}

type Aktivitas struct {
	AktivitasId    uint   `json:"AktivitasId"    gorm:"column:AktivitasId;primaryKey"`
	Tanggal        string `json:"Tanggal"        gorm:"column:Tanggal"`
	JenisAktivitas string `json:"JenisAktivitas" gorm:"column:JenisAktivitas;type:enum('Pemupukan','Penyiraman','Pengobatan');default:'Pemupukan'"`
	Keterangan     string `json:"Keterangan"     gorm:"column:Keterangan"`
	PenanamanId    uint   `json:"PenanamanId"    gorm:"column:PenanamanId"`
	SchedulerId    uint   `json:"SchedulerId"    gorm:"column:SchedulerId"`
}

func (Aktivitas) TableName() string {
	return "aktivitas"
}

// JenisAktivitas nya beda sama yang ada di database, makanya diganti sama ada yang ditambahin. Soalnya di database enumnya cuman ada 2, terus sama aku ditambahin 'Pengobatan'