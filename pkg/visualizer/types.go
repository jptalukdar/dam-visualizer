package visualizer

import "github.com/jptalukdar/dam-visualizer/pkg/icons"

type visualizer struct {
	Name     string
	Nodes    map[string]float32
	Links    map[string]map[string]float32
	NodeType map[string]string
}

func (v *visualizer) AddNode(name string, value float32) {
	v.Nodes[name] += value
}

func (v *visualizer) AddNodeType(name string, typeName string) {
	if v.NodeType[name] == "" {
		v.NodeType[name] = typeName
	}
}
func (v *visualizer) AddLink(source, destination string, value float32) {
	if v.Links[source] == nil {
		v.Links[source] = make(map[string]float32)
	}
	v.Links[source][destination] += value
}

func (v *visualizer) Init() {
	v.Nodes = make(map[string]float32)
	v.Links = make(map[string]map[string]float32)
	v.NodeType = make(map[string]string)
}

func Newvisualizer(name string) *visualizer {
	v := visualizer{
		Name: name,
	}
	v.Init()
	return &v
}

func (v *visualizer) GetNodeType(name string) string {
	return v.NodeType[name]
}

func (v *visualizer) GetNodeCategory(name string) int {
	switch v.GetNodeType(name) {
	case "dam":
		return 0
	case "app":
		return 1
	case "dep":
		return 2
	case "datasource":
		return 3
	}

	return 10
}

func (v *visualizer) GetNodeSymbol(name string) string {
	icon := icons.NewIconSet()
	switch v.GetNodeType(name) {
	case "dam":
		return icon.DAM
	case "app":
		return icon.Application
	case "dep":
		return "triangle"
	case "datasource":
		return icon.Datastore
	case "datastore":
		return icon.Datastore
	case "code":
		return icon.Code
	default:
		return "circle"
	}
}
