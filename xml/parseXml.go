package xml

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"path/filepath"
)

func ParseXmlString(content string, mapper *Mapper) {
	err := xml.Unmarshal([]byte(content), mapper)
	if err != nil {
		log.Fatal(err)
	}
}

func ParseXmlFile(fileName string, mapper *Mapper) {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	err = xml.Unmarshal(content, mapper)
	if err != nil {
		log.Fatal(err)
	}
	mapper.FileName = fileName
	if mapper.Namespace != "" {
		log.Println("loaded xml:", fileName)
	}

}

func ParseDir(dirName string, mappers *map[string]Mapper) {

	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fileName := filepath.Clean(dirName) + string(filepath.Separator) + file.Name()
		if file.IsDir() {
			ParseDir(fileName, mappers)
		} else {
			ext := filepath.Ext(fileName)
			if ext == ".xml" {
				var mapper Mapper
				ParseXmlFile(fileName, &mapper)
				namespace := mapper.Namespace
				if namespace != "" {
					v, ok := (*mappers)[namespace]
					if ok {
						log.Println("has namespace:", namespace, "in file:", v.FileName)
					} else {
						(*mappers)[namespace] = mapper
					}
				}
			}
		}
	}
}
