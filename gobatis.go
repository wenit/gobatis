package gobatis

import (
	. "github.com/wenit/goutil/config"
	"github.com/wenit/goutil/log"
	"github.com/wenit/gobatis/xml"
	"github.com/wenit/gobatis/session"

)

var Properties Property

var logger *log.Logger

var nameSpaceCache *xml.NamespaceCache

func init() {
	logger = log.New()
	nameSpaceCache = xml.NewNamespaceCache()
}

func InitGobatis(configFile string) {
	Properties = *NewProperty(configFile)
	//初始化数据库连接
	url := Properties.GetString("db.url")
	session.GetDB(url)
	mapperDir := Properties.GetString("mybatis.mapper")
	xml.ParseDir(mapperDir, nameSpaceCache)
}

func CloseDB()  {
	session.CloseDB()
}


