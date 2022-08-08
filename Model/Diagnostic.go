package model

import (
	"time"

	"gorm.io/gorm"

	"github.com/google/uuid"
)

type Diagnostic struct {
	gorm.Model
	ID            uuid.UUID `gorm:"type:varchar(128);primaryKey"`
	PatientID     uuid.UUID `gorm:"type:varchar(128)" json:"patientID"   lable:"患者ID"`
	DiagnosticTime *time.Time `gorm:"type:datetime" json:"diagnosticTime"   lable:"就诊时间"`
	DoctorID        uuid.UUID `gorm:"type:varchar(128)" json:"doctorID"   lable:"医生ID"`
    OutpatientID   uuid.UUID `gorm:"type:varchar(128)" json:"outpatientID"   lable:"门诊ID"`
	IllnessID  uuid.UUID `gorm:"type:varchar(128)" json:"illnessID"   lable:"疾病ID"`
	MedicalM   string  `gorm:"type:long" json:"medicalm"   lable:"诊断数据"`
 }

 //添加信息
//分页查询
//删除信息
