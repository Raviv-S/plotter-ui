package resources

import "encoding/json"

var triggerImages = map[string]string{
	"init":  "data:image/png;base64,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",
	"timer": "data:image/png;base64,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",
}

func (g *Graph) ConvertToCanvasPipeline() *Canvas {
	var canvas Canvas
	var pipelines []*CanvasPipeline
	json.Unmarshal(PipelineTemplate, &canvas)

	var mainNodes []*CanvasNode

	mainPipeline := CanvasPipeline{ID: "main"}
	pipelines = append(pipelines, &mainPipeline)

	for _, snode := range g.Supernodes {
		pipeline := CanvasPipeline{ID: snode.Id}

		var nodes []*CanvasNode
		for _, node := range g.Nodes {
			if node.SupernodeParentId == pipeline.ID {
				nodes = append(nodes, node.ConvertToCanvasNode())
			}
		}
		pipeline.Nodes = nodes
		pipelines = append(pipelines, &pipeline)
		supernode := snode.ConvertToCanvasNode()
		supernode.Type = "super_node"
		supernode.SubflowRef.PipelineIdRef = snode.Id
		supernode.AppData.UIData.IsExpanded = true
		mainNodes = append(mainNodes, supernode)
	}

	// Adding the rest of the nodes to the main pipeline
	for _, node := range g.Nodes {
		if node.SupernodeParentId == "" {
			mainNodes = append(mainNodes, node.ConvertToCanvasNode())
		}
	}
	mainPipeline.Nodes = mainNodes
	canvas.Pipelines = pipelines

	// Adding color comments to main pipeline
	var comments []*CanvasComment
	for i, comment := range g.Comments {
		c := comment.ConvertToCanvasComment()
		c.XPos += i * (c.Width + 10)
		comments = append(comments, c)
	}
	mainPipeline.AppData.UIData.Comments = comments

	return &canvas
}

func (n Node) ConvertToCanvasNode() *CanvasNode {
	var canvasNode = CanvasNode{ID: n.Id, Type: "execution_node"}
	// canvasNode.SubflowRef.PipelineIdRef = n.SupernodeParentId
	canvasNode.AppData.UIData =
		UIData{Label: n.Title,
			Image:       n.Image,
			XPos:        0,
			YPos:        0,
			Description: n.Description,
			Color:       n.Color,
		}
	if len(n.Inputs) > 0 {
		canvasNode.Inputs = []InputOutput{{ID: n.Id + "-inp-1"}}
		inpLinks := []CanvasLink{}
		for _, link := range n.Inputs {
			inpLinks = append(inpLinks, CanvasLink{link.SourceNode.Id})
		}
		canvasNode.Inputs[0].Links = inpLinks
	}

	if n.IsOutput {
		canvasNode.Outputs = []InputOutput{{ID: n.Id + "-inp-1"}}
	}

	for i, trigger := range n.Triggers {
		deco := Decoration{Position: "topRight",
			XPos:      -25 + (-25 * i),
			YPos:      4,
			ID:        n.Id + trigger,
			Image:     triggerImages[trigger],
			ClassName: "default"}
		canvasNode.AppData.UIData.Decorations = append(canvasNode.AppData.UIData.Decorations, &deco)
	}
	if n.Status != "" {
		canvasNode.AppData.UIData.Messages = []Message{{Type: n.Status}}
	}
	return &canvasNode
}

func (c *Comment) ConvertToCanvasComment() *CanvasComment {
	var canvasComment = CanvasComment{Color: "comment-" + c.Color,
		ID:      c.Content,
		Content: c.Content,
		Height:  40,
		Width:   160,
		YPos:    0}
	return &canvasComment
}
