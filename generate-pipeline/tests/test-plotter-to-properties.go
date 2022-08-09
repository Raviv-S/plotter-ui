package tests

import (
	res "plotter-gui/resources"
)

func TestPlotterFlatten() {
	plotter := res.ReadPlotter("./resources/plotters/plotter-2.yaml")
	plotter.PlotterToFlatten()
}

func TestAllParams() {
	plotter := res.ReadPlotter("./resources/plotters/plotter-2.yaml")
	pprint(plotter.AllParams())
}
