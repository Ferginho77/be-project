package models

type Scheduler struct {
	SchedulerId   uint   `json:"SchedulerId" gorm:"column:SchedulerId;primaryKey"`
	NamaScheduler string `json:"NamaScheduler" gorm:"column:NamaScheduler"`
	Tanggal       string `json:"Tanggal" gorm:"column:Tanggal;type:date"`
	Status        string `json:"Status" gorm:"column:Status;type:enum('Pending','Selesai','Dibatalkan');default:'Pending'"`
	PenanamanId   uint   `json:"PenanamanId" gorm:"column:PenanamanId"`
}

func (Scheduler) TableName() string {
	return "scheduler"
}