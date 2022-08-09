package tests

import (
	"encoding/json"
	"fmt"
)

func pprint(s interface{}) {
	a, _ := json.MarshalIndent(s, "", " ")
	fmt.Println(string(a))
}

func rprint(s interface{}) {
	fmt.Printf("%+v\n", s)
}

func CreatingSupernodes() {
	// pipe := res.Pipeline{Nodes: make(map[string]*res.Node)}
	// pipe.CreateNode("a")
	// pipe.CreateNode("b")
	// pipe.CreateNode("c")
	// grp := [3]string{"a", "b", "c"}
	// sn := pipe.CreateSuperNode("sn", grp[:])
	// pprint(sn.ConvertToCanvasNode())

}
