package visualizer

import (
	"fmt"
	"io"
	"os"

	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/jptalukdar/dam-visualizer/pkg/types"
)

var Size float32 = 45

func Visualize(path string) {
	dam := CreateDAM(path)
	v := Newvisualizer(dam.Name)
	v.AddNode(dam.Name, Size)
	v.AddNodeType(dam.Name, "dam")
	for _, app := range dam.Application {
		appname := app.Name // fmt.Sprintf("%s [%s]",  app.Name, app.Type)
		v.AddNode(appname, Size)
		v.AddLink(dam.Name, appname, Size)
		v.AddNodeType(appname, app.Type)
		for _, dep := range app.Depends {
			depname := dep.Name // fmt.Sprintf("%s [%s]", dep.Name, dep.Type)
			v.AddNode(depname, Size)
			v.AddLink(appname, depname, Size)
			v.AddNodeType(depname, dep.Type)
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

func (v *visualizer) CreateVisualizerPage() {
	page := components.NewPage()
	page.AddCharts(
		v.GetSankeyBase(),
		v.GetGraphBase(),
	)

	path := fmt.Sprintf("html/%s.html", v.Name)
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
