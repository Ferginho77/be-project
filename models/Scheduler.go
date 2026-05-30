package models

type Scheduler struct {
	SchedulerId uint `json:"SchedulerId" gorm:"column:SchedulerId;primaryKey"`
	NamaScheduler string `json:"NamaScheduler" gorm:"column:NamaScheduler"`
	Tanggal string `json:"Tanggal" gorm:"column:Tanggal"`
	Status string `json:"Status" gorm:"column:Status"`
}

func (Scheduler) TableName() string {
	return "scheduler"
}