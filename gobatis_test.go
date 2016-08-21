package gobatis

import (
	"testing"

	"github.com/wenit/goutil/json"
	"github.com/wenit/gobatis/session"
	"github.com/wenit/goutil/file"
	"math/rand"
)

const CONFIG_FILE = "conf/gobatis.properties"


func TestName(t *testing.T) {

	logger.Info("current workspace: %s", file.AppDir())
	var configFile = "d:/gobatis.properties"
	//var configFile =dir+string(filepath.Separator)+CONFIG_FILE
	logger.Info("config file path:%s", configFile)
	InitGobatis(configFile)
	logger.Debug(Properties.Store)
	logger.Debug(json.JsonToString(nameSpaceCache))

	params := make(map[string]interface{})
	params["threshold"] = 10
	info := session.SelectOne("redis.queryRedis", &params)
	list := session.SelectList("redis.queryRedis", &params)
	logger.Debug(info)
	logger.Debug(list)

	params["name"]="myhos3"
	params["id"]=100
	rst:= session.Update("redis.updateRedis",&params)


	params["id"]=107
	params["name"]="myhos3"
	params["host"]="myhos3"
	params["port"]="myhos3"
	params["threshold"]="myhos3"
	rst=session.Delete("redis.deleteRedis",&params)
	logger.Debug(rst)
	params["id"]=107
	params["name"]="myhos3"
	params["host"]="myhos3"
	params["port"]=rand.Intn(100)
	params["threshold"]="myhos3"
	rst=session.Insert("redis.insertRedis",&params)
	logger.Debug(rst)
}