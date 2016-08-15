package xml

type Statement struct {
	Id string `xml:"id,attr"`
	ParameterType string `xml:"parameterType,attr"`
	ResultType string `xml:"resultType,attr"`
	Sql string `xml:",chardata"`
}
