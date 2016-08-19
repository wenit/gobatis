package xml


type Namespace struct {
	Id string
	Statements map[string]*Statement
}

func (n *Namespace) GetStatement(id string) *Statement {
	v, ok := n.Statements[id]
	if ok {
		return v
	} else {
		return nil
	}
}