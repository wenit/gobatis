package xml

import (
	"testing"
	"fmt"
	"io/ioutil"
	//"log"
	"path/filepath"
	//"os"
	//"strings"
)

func TestParseXmlFile(t *testing.T) {
	var mapper Mapper
	var fileName = "D:/temp/mobileLogin.xml"
	ParseXmlFile(fileName, &mapper)
	fmt.Println(mapper)
}

func TestParseXmlString(t *testing.T) {
	var mapper Mapper
	var fileName = "D:/temp/mobileLogin.xml"
	bytes, _ := ioutil.ReadFile(fileName)
	ParseXmlString(string(bytes), &mapper)
	fmt.Println(mapper)

}

func TestParseDir(t *testing.T) {
	var fileDir = "D:/conf/project/bscz-mbank"
	mappers:=make(map[string]Mapper)

	//_ ,ok := mappers["a"]
	//fmt.Println(ok)

	ParseDir(fileDir,&mappers)

	//fmt.Println(mappers)

	//fmt.Println()
	//
	//files, err := ioutil.ReadDir(fileDir)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//for _, file := range files {
	//	//fileName:=file
	//	fmt.Println()
	//	fmt.Println(file)
	//
	//}
	//
	//filepath.Walk(fileDir, func(root string, info os.FileInfo,err error) error{
	//	if ( info == nil ) {
	//		return err
	//	}
	//	if info.IsDir() {
	//		return nil
	//	}
	//	fmt.Println(root)
	//	return nil
	//})
}

func TestName(t *testing.T) {
	var fileDir = "D:/temp//1.xml"
	fmt.Println(filepath.Clean(fileDir))
	fmt.Println(string(filepath.Separator))
	fmt.Println(filepath.Ext(fileDir))
}