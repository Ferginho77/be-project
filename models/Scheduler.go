package models

type Scheduler struct {
    SchedulerId   uint   `json:"SchedulerId" gorm:"column:SchedulerId;primaryKey"`
    NamaScheduler string `json:"NamaScheduler" gorm:"column:NamaScheduler"`
    Tanggal       string `json:"Tanggal" gorm:"column:Tanggal;type:date"`
    Status        string `json:"Status" gorm:"column:Status;type:enum('Pending','Selesai','Dibatalkan');default:'Pending'"`
    PenanamanId   uint   `json:"PenanamanId" gorm:"column:PenanamanId"`
    IsManual      bool   `json:"IsManual" gorm:"column:is_manual;type:tinyint(1);default:0"`
}

func (Scheduler) TableName() string {
	return "scheduler"
}

// Di database ada yang kelewat, bagian 'Dibatalkan'. Tapi sekarang di database udah di update buat nambahin ENUM 'Dibatalkan'