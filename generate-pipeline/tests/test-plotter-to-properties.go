package tests

import (
	res "plotter-gui/resources"
)

func TestPlotterFlatten() {
	plotter := res.ReadPlotter("../examples/plotter-2.yaml")
	plotter.PlotterToFlatten()
}

func TestAllParams() {
	plotter := res.ReadPlotter("../examples/plotter-2.yaml")
	pprint(plotter.AllParams())
}
