package xml

type Mapper struct {
	FileName string
	Namespace string `xml:"namespace,attr"`
	Selects []Statement `xml:"select"` 
	Updates []Statement `xml:"update"`
	Deletes []Statement `xml:"delete"`
	Inserts []Statement `xml:"insert"`
}
