package routes

import (
	v1 "PatientM/Api/v1"
	util "PatientM/Util"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(util.AppMode)
	r:=gin.New()
	r.Use(gin.Recovery())

router:=r.Group("api")
{

router.POST("patient/add",v1.Addpatient)
router.GET("patient/getpatient",v1.Getpatientone)
router.GET("patient/deletepatient",v1.Deletepatient)


}
_=r.Run(util.HttpPort)
}