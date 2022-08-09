package resources

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	app "fybrik.io/fybrik/manager/apis/app/v1alpha1"
	"sigs.k8s.io/yaml"
)

func pprint(s interface{}) {
	a, _ := json.MarshalIndent(s, "", " ")
	fmt.Println(string(a))
}

func rprint(s interface{}) {
	fmt.Printf("%+v\n", s)
}

type Plotter struct {
	app.Plotter
	ClusterColors map[string]string
}

func ReadPlotter(filename string) *Plotter {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		// return
	}
	plotter := &Plotter{}
	err = yaml.Unmarshal(buf, &plotter)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		// return
	}
	return plotter
}
