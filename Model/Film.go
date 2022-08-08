package model

import (
	"gorm.io/gorm"
	"time"
	"github.com/google/uuid"
)

type Film struct {
	gorm.Model
	ID            uuid.UUID `gorm:"type:varchar(128);primaryKey"`
	DiagnosticID   uuid.UUID     `gorm:"type:varchar(128);not null" json:"patierntName"   lable:"就诊Id"`
	FilmtypeId     uuid.UUID   `gorm:"type:varchar(128);not null" json:"filmtypeid"   lable:"片子类型"`
	FilmtFile       string   `gorm:"type:varchar(288);not null" json:"filmtfile"   lable:"片子地址"`
	FilmtDiagnose      string   `gorm:"type:varchar(288);not null" json:"filmtdiagnose"   lable:"片子诊断说明"`
	FilmTime *time.Time `gorm:"type:datetime" json:"filmTime"   lable:"拍片子时间"`

}

//添加信息
func Addfilm(film *Film)(int,string){
	if film.ID != uuid.Nil {
		err := db.Save(&film).Error
		if err != nil {
			return 0, "修改失败"

		}
		return 1, "修改成功"

	}
	id := uuid.New()
	film.ID = id
	err := db.Create(&film).Error
	if err != nil {
		return 0, "保存失败"
	}

	return 1, "保存成功"

}
//分页查询

func Filmpage(DiagnosticID string ,page int,pagesize int)([]Film,pagecont int){
	Filmlist []Film
 err:=db.Where("DiagnosticID=?",DiagnosticID)Limit(pagesize).Offset((page - 1) * pagesize).Find(&Filmlist).Count(&pageCount).Error

}
//删除信息
