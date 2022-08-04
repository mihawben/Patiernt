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
	Nationality string  `gorm:"type:varchar(128);not null" json:"nationality"   lable:"国籍"`
	Nation string`gorm:"type:varchar(128);not null" json:"nation"   lable:"民族"`
	Address string`gorm:"type:varchar(128);not null" json:"address"   lable:"地址"`
    Email string`gorm:"type:varchar(128);not null" json:"email"   lable:"邮箱"`
	Phonecall string`gorm:"type:varchar(128);not null" json:"phonecall"   lable:"电话号码"`
	Ismarriage bool`gorm:"type:tinyint(1);not null" json:"ismarriage"   lable:"是否结婚"`
    IDNumber string`gorm:"type:varchar(128);not null" json:"idnumber"   lable:"身份证号"`
	Gender string `gorm:"type:varchar(128);not null" json:"gender"   lable:"性别"`
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

