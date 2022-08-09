package main

import (
	res "plotter-gui/resources"
	"os"
)

func main() {
	plotterPath := os.Args[1] // "./resources/plotters/plotter-7.yaml"
	pipelineWritePath := "../src/pipelines/generated-pipeline.json"
	
	plotter := res.ReadPlotter(plotterPath)
	plotter.ConvertToGraph().ConvertToCanvasPipeline().ToFile(pipelineWritePath)

}
