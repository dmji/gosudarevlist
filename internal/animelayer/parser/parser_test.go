package animelayer_parser_test

import (
	animelayer_parser "collector/internal/animelayer/parser"
	"collector/pkg/model"
	"context"
	"encoding/json"
	"os"
	"testing"
)

func TestMainTest(t *testing.T) {

	items := make([]model.AnimeLayerItem, 0, 2000)
	ctx := context.Background()
	for i := range 120 {

		items = append(items, animelayer_parser.CollectBaseItemsFromAddress(ctx, "anime", i+1)...)

	}

	rankingsJson, _ := json.Marshal(items)
	_ = os.WriteFile("output.json", rankingsJson, 0644)

}
