package visualizer

import (
	"fmt"
	"io"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/jptalukdar/dam-visualizer/pkg/types"
)

func Visualize(path string) {
	dam := CreateDAM(path)
	v := Newvisualizer(dam.Name)
	v.AddNode(dam.Name, "DAM")
	for _, app := range dam.Application {
		appname := fmt.Sprintf("%s [%s]", app.Name, app.Type)
		v.AddNode(appname, app.Type)
		v.AddLink(dam.Name, appname, 1)
		for _, dep := range app.Depends {
			depname := fmt.Sprintf("%s [%s]", dep.Name, dep.Type)
			v.AddNode(depname, dep.Type)
			v.AddLink(appname, depname, 1)
		}
	}
	v.CreateVisualizerPage()
}

func CreateDAM(path string) *types.DAM {
	dam := types.NewDAM()
	dam.ReadYaml(path)
	p := dam.GenerateYaml()
	fmt.Print(p)
	return dam
}

func SankeyLinks() {

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

func (v *visualizer) CreateVisualizerPage() {
	page := components.NewPage()
	page.AddCharts(
		v.GetSankeyBase(),
	)

	path := fmt.Sprintf("html/%s.html", v.Name)
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
