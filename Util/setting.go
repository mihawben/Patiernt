package util
import(

"log"

"gopkg.in/ini.v1"

)
var (
	AppMode  string
	HttpPort string
	//JwtKey   string

	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	//AccessKey  string
	//SecretKey  string
	//Bucket     string
	//QiniuSever string
)

func init(){
	cfg, err := ini.Load("Config/config.ini")
	if err != nil {
	  log.Fatal("配置文件有误 ", err)
	
	
	}
	LoadData(cfg)
	LoadServer(cfg)
}
func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").String()
	HttpPort = file.Section("server").Key("HttpPort").String()
	
}

func LoadData(file *ini.File) {
	DbHost = file.Section("database").Key("DbHost").String()
	DbPort = file.Section("database").Key("DbPort").String()
	DbUser = file.Section("database").Key("DbUser").String()
	DbPassWord = file.Section("database").Key("DbPassWord").String()
	DbName = file.Section("database").Key("DbName").String()
}
