package visualizer

import (
	"fmt"
	"log"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func (v *visualizer) GetGraphNode() []opts.GraphNode {
	var graphNode []opts.GraphNode
	for k := range v.Nodes {
		if k == v.Name {
			log.Printf("%s is the root node", k)
			graphNode = append(graphNode, opts.GraphNode{Name: k, X: 0, Y: 0, SymbolSize: Size, Symbol: v.GetNodeSymbol(k), Category: 1})
		} else {
			graphNode = append(graphNode, opts.GraphNode{Name: k, Category: 0, SymbolSize: Size, Symbol: v.GetNodeSymbol(k)})
		}

	}
	return graphNode
}

func (v *visualizer) GetGraphLink() []opts.GraphLink {
	var graphLink []opts.GraphLink
	for k, v := range v.Links {
		for k2, v2 := range v {
			graphLink = append(graphLink, opts.GraphLink{Source: k, Target: k2, Value: v2})
		}
	}
	return graphLink
}

func (v *visualizer) GetGraphBase() *charts.Graph {
	graph := charts.NewGraph()
	graph.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: fmt.Sprintf("Data Application Model - %s", v.Name),
		}),
		charts.WithLegendOpts(opts.Legend{Show: true}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true, Trigger: "item", TriggerOn: "mousemove"}),
	)

	graph.AddSeries("graph", v.GetGraphNode(), v.GetGraphLink()).
		SetSeriesOptions(
			charts.WithGraphChartOpts(opts.GraphChart{
				Layout:             "circular",
				Roam:               true,
				FocusNodeAdjacency: true,
			}),
			charts.WithEmphasisOpts(opts.Emphasis{
				Label: &opts.Label{
					Show:     true,
					Color:    "black",
					Position: "left",
				},
			}),
			charts.WithLineStyleOpts(opts.LineStyle{
				Curveness: 0.3,
			}),
		)
	return graph
}
