package tests

import (
	res "plotter-gui/resources"
)

func PlotterToCanvas() {
	plotter := res.ReadPlotter("../examples/plotter-7.yaml")
	plotter.ConvertToGraph().ConvertToCanvasPipeline().
		ToFile("./generated-pipeline.json")
}
