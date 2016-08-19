package xml


var NSCache *NamespaceCache

type NamespaceCache struct {
	Namespaces map[string]*Namespace
}

func NewNamespaceCache() *NamespaceCache {
	NSCache=&NamespaceCache{
		Namespaces:make(map[string]*Namespace),
	}
	return NSCache
}


func (c *NamespaceCache) GetNameSpace(id string) *Namespace {
	v, ok := c.Namespaces[id]
	if ok {
		return v
	} else {
		return nil
	}
}

func (c *NamespaceCache) SetNameSpace(n *Namespace)  {

	c.Namespaces[n.Id]=n;
}

