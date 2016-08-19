package xml

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"path/filepath"
	"cn/wenit/gobatis/consts"
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

func ParseDir(dirName string, nc *NamespaceCache) {
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fileName := filepath.Clean(dirName) + string(filepath.Separator) + file.Name()
		if file.IsDir() {
			ParseDir(fileName, nc)
		} else {
			ext := filepath.Ext(fileName)
			if ext == ".xml" {
				log.Println("mybaits mapper file :",fileName);
				var mapper Mapper
				ParseXmlFile(fileName, &mapper)
				nsId := mapper.Namespace
				if nsId == "" {
					nsId=consts.DEFAULT_NAMESPACE
				}

				namespace := (*nc).GetNameSpace(nsId)
				if namespace == nil {
					namespace=&Namespace{
						Id:nsId,
					}
				}

				putStatement(namespace,&mapper)

				nc.SetNameSpace(namespace)
			}
		}
	}
}

func putStatement(ns *Namespace, mapper *Mapper) {
	for _, s := range mapper.Selects {
		put(ns,&s,mapper)
	}
	for _, s := range mapper.Inserts {
		put(ns,&s,mapper)
	}
	for _, s := range mapper.Updates {
		put(ns,&s,mapper)
	}
	for _, s := range mapper.Deletes {
		put(ns,&s,mapper)
	}
}

func put(ns *Namespace,s *Statement, mapper *Mapper)  {
	v, ok := ns.Statements[s.Id]
	if ok {
		log.Printf("this namespace %s ,has statement id %s in file %s", ns.Id, s.Id, v.Mapper.FileName)
	} else {
		if(ns.Statements == nil){
			ns.Statements= make(map[string]*Statement)
		}
		ns.Statements[s.Id] = s
		s.Mapper = mapper
	}
}

