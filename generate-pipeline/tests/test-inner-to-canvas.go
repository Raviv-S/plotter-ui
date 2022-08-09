package tests

import (
	res "plotter-gui/resources"
)

func CanvasPipelineToJson() {
	pipe := res.Graph{Nodes: make(map[string]*res.Node), Supernodes: make(map[string]*res.Node)}
	a := pipe.CreateNode("a")
	b := pipe.CreateNode("b")
	c := pipe.CreateNode("c")
	d := pipe.CreateNode("d")
	e := pipe.CreateNode("e")
	f := pipe.CreateNode("f")
	z := pipe.CreateNode("z")
	pipe.CreateSupernode("sn")
	z.SupernodeParentId = "sn"
	pipe.ConnectNodes(a, c, "link")
	pipe.ConnectNodes(b, c, "link")
	pipe.ConnectNodes(c, d, "link")
	pipe.ConnectNodes(c, e, "link")
	pipe.ConnectNodes(e, f, "link")
	pipe.ConnectNodes(pipe.Nodes["e"], pipe.Nodes["f"], "l")
	pipe.ConnectNodes(z, z, "link")

	//jsonPipe, _ := json.Marshal(pipe.ConvertToCanvasPipeline())
	//fmt.Println(string(jsonPipe[:]))
	pipe.ConvertToCanvasPipeline().ToFile("./generated-pipeline.json")
}
