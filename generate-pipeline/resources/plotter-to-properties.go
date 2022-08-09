package resources

import (
	"encoding/json"
	"strconv"
)

func Flatten(m map[string]interface{}) map[string]interface{} {
	o := make(map[string]interface{})
	for k, v := range m {
		switch child := v.(type) {
		case map[string]interface{}:
			nm := Flatten(child)
			for nk, nv := range nm {
				o[k+"."+nk] = nv
			}
		case []interface{}:
			for i, nv := range child {
				switch nv.(type) {
				case string:
					o[k] = strconv.Itoa(i) + "-" + nv.(string)
				case []interface{}:
					// pprint("bbbbbb")
					for i2, nv3 := range child {
						switch nv3.(type) {
						case string:
							o[k] = (strconv.Itoa(i2) + "-" + nv3.(string))
						default:
							nm := Flatten(nv.(map[string]interface{}))
							for nk4, nv4 := range nm {
								o[strconv.Itoa(i)+k+"."+nk4] = nv4
							}
						}
					}
				default:
					nm := Flatten(nv.(map[string]interface{}))
					for nk2, nv2 := range nm {
						o[strconv.Itoa(i)+k+"."+nk2] = nv2
					}
				}
			}
		default:
			o[k] = v
		}
	}
	return o
}

func (plotter Plotter) PToJson() map[string]ParamDef {
	main := make(map[string]ParamDef)
	// for asset, value := range plotter.Spec.Assets {
	// 	var m map[string]interface{}
	// 	marsh, _ := json.Marshal(value)
	// 	json.Unmarshal(marsh, &m)
	// 	main[asset] = ElementToParamDefFormat(asset, Flatten(m))
	// }
	var m map[string]interface{}
	marsh, _ := json.Marshal(plotter.Spec.Selector)
	json.Unmarshal(marsh, &m)
	workloadName := plotter.Spec.Selector.WorkloadSelector.MatchLabels["app"]
	main[workloadName] = ElementToParamDefFormat(workloadName, NewFlatten(m))
	return main
}

func NewFlatten(main map[string]interface{}) map[string]interface{} {
	output := make(map[string]interface{})
	for k, v := range main {
		switch child := v.(type) {
		case map[string]interface{}:
			subMap := Flatten(child)
			for sk, sv := range subMap {
				output[k+"."+sk] = sv
			}
		case []interface{}:
			for i, sv := range child {
				output[k+"-"+strconv.Itoa(i)] = sv
			}
		default:
			output[k] = v
		}
	}
	return output
}

func (plotter Plotter) AddWorkloadParams(allParams map[string]ParamDef) {
	var m map[string]interface{}
	marsh, _ := json.Marshal(plotter.Spec.Selector)
	json.Unmarshal(marsh, &m)
	workloadName := plotter.Spec.Selector.WorkloadSelector.MatchLabels["app"]
	allParams[workloadName] = ElementToParamDefFormat(workloadName, NewFlatten(m))
}

func (plotter Plotter) AddAssetsParams(allParams map[string]ParamDef) {
	for assetName, assetValue := range plotter.Spec.Assets {
		var m = make(map[string]interface{})
		m["status.ready"] = plotter.Status.Assets[assetName].Ready
		if plotter.Status.Assets[assetName].Error != "" {
			m["status.error"] = plotter.Status.Assets[assetName].Error
		}
		marsh, _ := json.Marshal(assetValue)
		json.Unmarshal(marsh, &m)
		allParams[assetName] = ElementToParamDefFormat(assetName, NewFlatten(m))
	}
}

func (plotter Plotter) AddTemplateParams(allParams map[string]ParamDef) {
	for templateName, templateValue := range plotter.Spec.Templates {
		var m = make(map[string]interface{})
		marsh, _ := json.Marshal(templateValue)
		json.Unmarshal(marsh, &m)
		allParams[templateName] = ElementToParamDefFormat(templateName, NewFlatten(NewFlatten(m)))
	}
}

// func MergeMaps()

func (plotter Plotter) AllParams() map[string]ParamDef {
	allParams := make(map[string]ParamDef)
	plotter.AddWorkloadParams(allParams)
	plotter.AddAssetsParams(allParams)
	plotter.AddTemplateParams(allParams)
	return allParams
}

func (plotter Plotter) PlotterToFlatten() {
	main := make(map[string]interface{})
	marsh, _ := json.Marshal(plotter.Status.Assets["fybrik-notebook-sample/paysim-csv"])
	// marsh, _ := json.Marshal(plotter.ObjectMeta)
	json.Unmarshal(marsh, &main)
	// pprint(NewFlatten(NewFlatten(NewFlatten(main))))
	// pprint(NewFlatten(main))
	pprint(plotter.Status.Assets["fybrik-notebook-sample/paysim-csv"].Ready)
	// rprint(plotter.Status.Assets["fybrik-notebook-sample/users"])
}
