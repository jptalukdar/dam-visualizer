package icons

import (
	"embed"
	"encoding/base64"
)

type ICONS struct {
	Code        string `json:"code"`
	Application string `json:"application"`
	Datastore   string `json:"datastore"`
	DAM         string `json:"dam"`
}

var (

	//go:embed png/*.png
	i_e embed.FS

	ICONS_SET = map[string]string{
		"dam":       "png/dam.png",
		"app":       "png/app.png",
		"datastore": "png/datastore.png",
		"code":      "png/code.png",
	}
)

func NewIconSet() *ICONS {
	i := ICONS{}
	i.DAM = LoadIcon(ICONS_SET["dam"])
	i.Application = LoadIcon(ICONS_SET["app"])
	i.Datastore = LoadIcon(ICONS_SET["datastore"])
	i.Code = LoadIcon(ICONS_SET["code"])
	// log.Printf("%+v", i)
	return &i
}

func LoadIcon(path string) string {
	bytes, err := i_e.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var base64Encoding string
	base64Encoding += "image://data:image/png;base64,"

	// Append the base64 encoded output
	base64Encoding += toBase64(bytes)
	return base64Encoding
}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
