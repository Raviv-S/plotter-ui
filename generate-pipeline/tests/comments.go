package tests

import (
	res "plotter-gui/resources"
)

func AddCommentToPipeline() {
	pipe := res.Graph{Nodes: make(map[string]*res.Node), Supernodes: make(map[string]*res.Node)}
	pipe.CreateComment("hi", "blue")
	pprint(pipe)
}

func PlotterComments() {
	plotter := res.ReadPlotter("./resources/plotters/plotter-7.yaml")
	pprint(plotter.ConvertToGraph())
}

func CommentToCanvas() {
	plotter := res.ReadPlotter("./resources/plotters/plotter-7.yaml")
	plotter.ConvertToGraph().ConvertToCanvasPipeline().
		ToFile("./comments-testing.json")
}
