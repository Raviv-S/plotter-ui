package resources

type ParamDef struct {
	Title         TitleDef               `json:"titleDefinition"`
	CurrentParams map[string]interface{} `json:"current_parameters"`
	Params        []Param                `json:"parameters"`
	UIHints       UIHints                `json:"uihints,omitempty"`
}

type TitleDef struct {
	Name     string `json:"title"`
	Editable bool   `json:"editable"`
}

type Param struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type UIHints struct {
	Id         string
	PropGroups []PropGroup
}

type PropGroup struct {
	Id       string
	Label    GroupLabel
	ParamRef []string
}

type GroupLabel struct {
	Default string
}

func ElementToParamDefFormat(nodeName string, parameters map[string]interface{}) ParamDef {
	pd := ParamDef{}
	pd.Title.Name = nodeName
	pd.Title.Editable = false
	pd.CurrentParams = parameters
	pd.Params = make([]Param, 0)
	for k := range pd.CurrentParams {
		param := Param{Id: k, Type: "string"}
		pd.Params = append(pd.Params, param)
	}
	return pd
}
