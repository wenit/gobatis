package session

import (
	"strings"
	"github.com/wenit/gobatis/consts"
	"fmt"
	"database/sql"
	"github.com/wenit/goutil/log"
	"github.com/wenit/gobatis/xml"
	_ "github.com/mattn/go-sqlite3"
	"regexp"
)

var REGEX_PARAM_KEY = regexp.MustCompile(`#{[a-zA-Z0-9_]+}`)

var REGEX_PARAM_PREFIX_SUFFIX, _ = regexp.Compile(`#|{|}`)

var db *sql.DB

var logger *log.Logger

func init() {
	logger = log.New()
}

func SelectOne(statement string, params *map[string]interface{}) map[string]interface{} {
	sql, newParams := buildSql(statement, params);
	list := Query(sql, newParams)
	if list == nil {
		return nil
	} else {
		return list[0]
	}
}

func SelectList(statement string, params *map[string]interface{}) []map[string]interface{} {
	sql, newParams := buildSql(statement, params);
	list := Query(sql, newParams)
	if list == nil {
		return nil
	} else {
		return list
	}
}

func Update(statement string, params *map[string]interface{}) bool {
	sql, newParams := buildSql(statement, params);
	return Exec(sql, newParams)
}

func Delete(statement string, params *map[string]interface{}) bool {
	sql, newParams := buildSql(statement, params);
	return Exec(sql, newParams)
}

func Insert(statement string, params *map[string]interface{}) bool {
	sql, newParams := buildSql(statement, params);
	return Exec(sql, newParams)
}

func Exec(sql string, params []interface{}) bool {
	stmt, _ := db.Prepare(sql)
	_, err := stmt.Exec(params...)

	if err != nil {
		logger.Error(err)
		return false
	} else {
		return true
	}
}

func Query(sql string, params []interface{}) []map[string]interface{} {
	stmt, _ := db.Prepare(sql)
	rows, err := stmt.Query(params...)
	if err != nil {
		fmt.Println("Unable to query foo table:", err)
	}
	defer rows.Close()

	cols, _ := rows.Columns();
	colSize := len(cols)

	list := make([]map[string]interface{}, 0)
	for rows.Next() {
		rst := make(map[string]interface{}, colSize)
		points := make([]interface{}, colSize)
		values := make([]string, colSize)
		for i, _ := range values {
			points[i] = &values[i]
		}
		rows.Scan(points...)

		for i, v := range cols {
			rst[v] = values[i]
		}
		list = append(list, rst)
	}
	return list
}

func GetDB(url string) *sql.DB {
	if db != nil {
		return db
	}

	d, err := sql.Open(consts.DB_TYPE_SQLITE3, url)
	if err != nil {
		logger.Error("get db url :", url, " error info :", err)
	}
	db = d
	return db
}

func getSQL(statement string) string {
	arr := strings.Split(statement, ".")
	var s string
	var n string
	if len(arr) == 2 {
		n = arr[0]
		s = arr[1]
	} else {
		n = consts.DEFAULT_NAMESPACE
		s = arr[0]
	}
	namespace := xml.NSCache.GetNameSpace(n)
	st := namespace.GetStatement(s)
	return strings.TrimSpace(st.Sql)
}

func buildSql(statement string, params *map[string]interface{}) (string, []interface{}) {
	sqltp := getSQL(statement)
	sqltp = strings.TrimSpace(sqltp)
	tempKey := REGEX_PARAM_KEY.FindAllString(sqltp, -1)
	sql := REGEX_PARAM_KEY.ReplaceAllString(sqltp, "?")
	keys := make([]interface{}, 0)
	for _, v := range tempKey {
		nv := REGEX_PARAM_PREFIX_SUFFIX.ReplaceAllString(v, "")
		value, ok := (*params)[nv]
		if ok {
			keys = append(keys, value)
		} else {
			fmt.Printf("value not exist %s\n", nv)
		}
		//fmt.Println(i, v)
	}
	logger.Debug("\nSQL:\n", sql, "\nparams:", keys)
	return sql, keys
}

func CloseDB()  {
	db.Close()
}