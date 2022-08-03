package model

import (
	"gorm.io/gorm"

	"github.com/google/uuid"
)

type Patient struct {
	gorm.Model
	ID            uuid.UUID `gorm:"type:varchar(128);primaryKey"`
	PatierntName  string    `gorm:"type:varchar(128);not null" json:"patierntName"   lable:"患者名称"`
	PatierntAge   int64     `gorm:"size:3" json:"patierntage" lable:"年龄"`
	PatiernHeight float64   `gorm:"type:float" json:"patiernheight" lable:"身高"`
	MotherID uuid.UUID `gorm:"type:varchar(128)"`
	FatherID uuid.UUID `gorm:"type:varchar(128)"`
	
}

//添加数据或修改
func Addpatent(patent *Patient) (int, string) {
	if patent.ID != uuid.Nil {
		err := db.Save(&patent).Error
		if err != nil {
			return 0, "修改失败"

		}
		return 1, "修改成功"

	}
	id := uuid.New()
	patent.ID = id
	err := db.Create(&patent).Error
	if err != nil {
		return 0, "保存失败"
	}

	return 1, "保存成功"
}

//查询单条
func Getpatentone(Patientid string) Patient {
	var data Patient
	err := db.Where("ID=?", Patientid).First(&data).Error
	if err != nil {
		return data
	}       
	return data
}
