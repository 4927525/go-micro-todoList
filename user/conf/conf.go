package conf

import (
	"fmt"
	"gopkg.in/ini.v1"
	"strings"
	"user/model"
)

var (
	AppMode    string
	HttpPort   string
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

func Init() {
	file, err := ini.Load("./config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径", err)
		return
	}

	// 加载配置
	LoadMysqlData(file)
	LoadService(file)

	//连接数据库
	path := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8&parseTime=true"}, "")
	model.Database(path)
}

func LoadService(f *ini.File) {
	AppMode = f.Section("service").Key("AppMode").String()
	HttpPort = f.Section("service").Key("HttpPort").String()
}

func LoadMysqlData(f *ini.File) {
	Db = f.Section("mysql").Key("DB").String()
	DbHost = f.Section("mysql").Key("DbHost").String()
	DbPort = f.Section("mysql").Key("DbPort").String()
	DbUser = f.Section("mysql").Key("DbUser").String()
	DbPassWord = f.Section("mysql").Key("DbPassWord").String()
	DbName = f.Section("mysql").Key("DbName").String()
}
