package resources

import (
	"encoding/json"
	"os"
)

type Canvas struct {
	DocType         string            `json:"doc_type"`
	Version         string            `json:"version"`
	JSONSchema      string            `json:"json_schema"`
	ID              string            `json:"id"`
	PrimaryPipeline string            `json:"primary_pipeline"`
	Pipelines       []*CanvasPipeline `json:"pipelines"`
	Schemas         []interface{}     `json:"schemas"`
}

type CanvasPipeline struct {
	ID      string        `json:"id"`
	Nodes   []*CanvasNode `json:"nodes"`
	AppData struct {
		UIData UIData `json:"ui_data"`
	} `json:"app_data"`
}

type CanvasNode struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	SubflowRef struct {
		PipelineIdRef string `json:"pipeline_id_ref,omitempty"`
	} `json:"subflow_ref,omitempty"`
	AppData struct {
		UIData UIData `json:"ui_data"`
	} `json:"app_data"`
	Inputs  []InputOutput `json:"inputs,omitempty"`
	Outputs []InputOutput `json:"outputs,omitempty"`
}

type UIData struct {
	Label       string           `json:"label"`
	Image       string           `json:"image,omitempty"`
	XPos        int              `json:"x_pos"`
	YPos        int              `json:"y_pos"`
	Description string           `json:"description,omitempty"`
	Messages    []Message        `json:"messages,omitempty"`
	Color       string           `json:"class_name"`
	IsExpanded  bool             `json:"is_expanded,omitempty"`
	Comments    []*CanvasComment `json:"comments,omitempty"`
	Decorations []*Decoration    `json:"decorations,omitempty"`
}

type InputOutput struct {
	ID    string       `json:"id"`
	Links []CanvasLink `json:"links"`
}

type CanvasLink struct {
	NodeIDRef string `json:"node_id_ref"`
}

type Message struct {
	Type string `json:"type"`
}

type CanvasComment struct {
	ID               string `json:"id"`
	XPos             int    `json:"x_pos"`
	YPos             int    `json:"y_pos"`
	Width            int    `json:"width"`
	Height           int    `json:"height"`
	Color            string `json:"class_name"`
	Content          string `json:"content"`
	AssociatedIDRefs []struct {
		NodeRef string `json:"node_ref"`
	} `json:"associated_id_refs"`
}

type Decoration struct {
	ID        string `json:"id"`
	Position  string `json:"position"`
	ClassName string `json:"class_name"`
	XPos      int    `json:"x_pos"`
	YPos      int    `json:"y_pos"`
	Image     string `json:"image"`
}

var PipelineTemplate = []byte(`
{
	"doc_type": "pipeline",
	"version": "3.0",
	"json_schema": "http://api.dataplatform.ibm.com/schemas/common-pipeline/pipeline-flow/pipeline-flow-v3-schema.json",
	"id": "script-generated",
	"primary_pipeline": "main",
	"pipelines": [],
	"schemas": []
}
`)

func (cp *Canvas) ToFile(filename string) {
	jsonPipeline, _ := json.MarshalIndent(cp, "", " ")
	os.WriteFile(filename, jsonPipeline, 0644)
}
