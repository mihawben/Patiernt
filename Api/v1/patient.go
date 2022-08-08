package v1

import (
	model "PatientM/Model"
	"net/http"

	"github.com/gin-gonic/gin"

)
///添加数据
func Addpatient(c *gin.Context){
var data model.Patient
_=c.ShouldBindJSON(&data)
code,str:=model.Addpatent(&data)

c.JSON(
	http.StatusOK, gin.H{
		"status":  200,
		"message": str+string(code),
	},
)


}
///查询单条
func  Getpatientone(c *gin.Context){

uid :=c.Query("patientid")
 data:=model.Getpatentone(uid)
c.JSON(
	http.StatusOK, gin.H{
		"status":  200,
		"data":data,
		"message": "",
	},

)

}
///删除单条
func Deletepatient(c *gin.Context){
uid :=c.Query("patientid")
_,data:=model.Deletpatient(uid)
c.JSON(
http.StatusOK,gin.H{
	"status":  200,
	
	"message": data,

},

)

}