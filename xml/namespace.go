package xml

import "fmt"

type Namespace struct {
	Id string
	Statements map[string]Statement
}

func (n *Namespace) GetStatement(id string) Statement {
	fmt.Println(n.Statements)
	fmt.Println(id)
	fmt.Println(id,n.Statements[id].Sql)
	v, ok := n.Statements[id]
	if ok {
		return v
	} else {
		return nil
	}
}