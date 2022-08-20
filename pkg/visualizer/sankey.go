package visualizer

import (
	"fmt"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func (v *visualizer) GetSankeyNode() []opts.SankeyNode {
	var sankeyNode []opts.SankeyNode
	for k, v := range v.Nodes {
		sankeyNode = append(sankeyNode, opts.SankeyNode{Name: k, Value: fmt.Sprintf("%f", v)})
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

func (v *visualizer) GetSankeyBase() *charts.Sankey {
	sankey := charts.NewSankey()
	sankey.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: fmt.Sprintf("DAM %s", v.Name),
		}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true, Trigger: "item", TriggerOn: "mousemove"}),
	)

	sankey.AddSeries(
		"Data Application Model",
		v.GetSankeyNode(),
		v.GetSankeyLink(),
		charts.WithLabelOpts(opts.Label{Show: true}),
	)
	return sankey
}
