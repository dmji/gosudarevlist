package main

import (
	"collector/pkg/model"
	_ "embed"
	"encoding/json"
	"os"
	"strings"
)

//go:embed test.json
var content []byte

type AnimeLayerItem struct {
	GUID string
	Name string
}

func main() {
	db := make([]AnimeLayerItem, 0)
	err := json.Unmarshal(content, &db)

	if err != nil {
		panic(err)
	}

	newarr := make([]model.AnimeLayerItem, len(db))
	for i, d := range db {
		newName, bFound := strings.CutSuffix(d.Name, " Complete")

		newarr[i].Completed = bFound
		newarr[i].Name = newName
		newarr[i].GUID = d.GUID
	}

	rankingsJson, _ := json.Marshal(newarr)
	_ = os.WriteFile("output.json", rankingsJson, 0644)

}
