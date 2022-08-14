package main

import (
	res "plotter-gui/resources"
	"os"
	"encoding/json"
)


func main() {
	plotterPath := os.Args[1]
	pipelineWritePath := "../src/generated-pipeline.json"
	paramsWritePath := "../src/generated-params.json"

	plotter := res.ReadPlotter(plotterPath)
	plotter.ConvertToGraph().ConvertToCanvasPipeline().ToFile(pipelineWritePath)

	params := plotter.AllParams()
	jsonParams, _ := json.MarshalIndent(params, "", " ")
	os.WriteFile(paramsWritePath, jsonParams, 0644)

}
