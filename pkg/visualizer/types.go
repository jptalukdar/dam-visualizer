package visualizer

import "github.com/go-echarts/go-echarts/v2/opts"

type visualizer struct {
	Name  string
	Nodes map[string]string
	Links map[string]map[string]float32
}

func (v *visualizer) AddNode(name string, value string) {
	v.Nodes[name] = value
}

func (v *visualizer) AddLink(source, destination string, value float32) {
	if v.Links[source] == nil {
		v.Links[source] = make(map[string]float32)
	}
	v.Links[source][destination] = value
}

func (v *visualizer) Init() {
	v.Nodes = make(map[string]string)
	v.Links = make(map[string]map[string]float32)
}

func Newvisualizer(name string) *visualizer {
	v := visualizer{
		Name: name,
	}
	v.Init()
	return &v
}

func (v *visualizer) GetSankeyNode() []opts.SankeyNode {
	var sankeyNode []opts.SankeyNode
	for k, v := range v.Nodes {
		sankeyNode = append(sankeyNode, opts.SankeyNode{Name: k, Value: v})
	}
	return sankeyNode
}

func (v *visualizer) GetSankeyLink() []opts.SankeyLink {
	var sankeyLink []opts.SankeyLink
	for k, v := range v.Links {
		for k2, v2 := range v {
			sankeyLink = append(sankeyLink, opts.SankeyLink{Source: k, Target: k2, Value: v2})
		}
	}
	return sankeyLink
}
