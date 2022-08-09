package resources

type Graph struct {
	Nodes      map[string]*Node
	Supernodes map[string]*Node
	Comments   []*Comment
}

type Node struct {
	Id                string
	Title             string
	Description       string
	Image             string
	Inputs            []*Link
	Outputs           []*Link
	IsOutput          bool
	Status            string	
	Color             string
	SupernodeParentId string
	Triggers          []string
}

type Link struct {
	SourceNode *Node
	Label      string
	Color      string
}

type Comment struct {
	Content        string
	Color          string
	ConnectedNodes []string
}

func (g *Graph) CreateSupernode(id string) *Node {
	_, exists := g.Supernodes[id]
	if exists {
		return g.Supernodes[id]
	}
	n := Node{Id: id, Title: id, Color: "default"}
	g.Supernodes[id] = &n
	return g.Supernodes[id]
}

func (g *Graph) CreateNode(id string) *Node {
	_, exists := g.Nodes[id]
	if exists {
		return g.Nodes[id]
	}
	n := Node{Id: id, Title: id, Color: "default"}
	g.Nodes[id] = &n
	return g.Nodes[id]
}

func (g *Graph) ConnectNodes(src *Node, dest *Node, label string) {
	l := Link{SourceNode: src, Label: label}
	dest.Inputs = append(dest.Inputs, &l)
	src.IsOutput = true
	//r := Link{SourceNode: dest, Label: label}

}

func (g *Graph) CreateComment(content string, color string) *Comment {
	c := Comment{Content: content, Color: color}
	g.Comments = append(g.Comments, &c)
	return &c
}
