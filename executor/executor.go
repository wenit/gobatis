package executor

import (
	//"cn/wenit/gobatis/xml"
	"strings"
	"cn/wenit/gobatis/consts"
	"fmt"
	//"database/sql"
)

func Select(statement *string, params map[string]interface{}) {

}

func GetDB(url string)  {
	//sql.Open()
}

func getSQL(statement *string) {
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
	fmt.Println(s,n)
}