package main

import (
	. "github.com/wenit/util/config"
	"github.com/wenit/util/log"
	"os/exec"
	"os"
	"path/filepath"
	"github.com/wenit/gobatis/xml"
	"github.com/wenit/util/json"
	"github.com/wenit/gobatis/executor"
)

var Properties Property

var logger *log.Logger

var nameSpaceCache *xml.NamespaceCache

const CONFIG_FILE = "conf/gobatis.properties"

func init() {
	logger = log.New()
	nameSpaceCache = xml.NewNamespaceCache()
}

func main() {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	dir := filepath.Dir(path)
	logger.Info("current workspace: %s", dir)
	var configFile = "d:/gobatis.properties"
	//var configFile =dir+string(filepath.Separator)+CONFIG_FILE
	logger.Info("config file path:%s", configFile)
	InitGobatis(configFile)
	//logger.Debug(Properties.Store)
	//fmt.Println(Properties)
	logger.Debug(json.JsonToString(nameSpaceCache))

	//params := make(map[string]interface{})
	//params["threshold"] = 10
	//info := executor.SelectOne("redis.queryRedis", &params)
	//list := executor.SelectList("redis.queryRedis", &params)
	//logger.Debug(info)
	//logger.Debug(list)

	//params["name"]="myhost"
	//params["id"]=100
	//rst:=executor.Update("redis.updateRedis",&params)
	//logger.Debug(rst)
}

func InitGobatis(configFile string) {
	Properties = *NewProperty(configFile)
	//初始化数据库连接
	url := Properties.GetString("db.url")
	executor.GetDB(url)
	mapperDir := Properties.GetString("mybatis.mapper")
	xml.ParseDir(mapperDir, nameSpaceCache)
}




