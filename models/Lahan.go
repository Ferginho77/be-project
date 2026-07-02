package models


type Lahan struct {
    LahanId    int     `gorm:"column:LahanId;primaryKey"`
    NamaLahan  string `json:"NamaLahan"   gorm:"column:NamaLahan"`
    LuasTanah  float64 `json:"LuasTanah"   gorm:"column:LuasTanah"`
    Kondisi    string  `json:"Kondisi"   gorm:"column:Kondisi"`
}

func (Lahan) TableName() string {
    return "lahan"
}