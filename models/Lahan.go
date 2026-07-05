package models


type Lahan struct {
    LahanId     uint    `gorm:"column:LahanId;primaryKey;autoIncrement" json:"LahanId"`
    NamaLahan   string  `gorm:"column:NamaLahan" json:"NamaLahan"`
    LuasTanah   float64 `gorm:"column:LuasTanah" json:"LuasTanah"`
    Kondisi     string  `gorm:"column:Kondisi" json:"Kondisi"`
    StatusLahan string  `gorm:"column:StatusLahan" json:"StatusLahan"`
}

func (Lahan) TableName() string {
    return "lahan"
}