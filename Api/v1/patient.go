package v1

import (
	model "PatientM/Model"
	"net/http"

	"github.com/gin-gonic/gin"

)

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