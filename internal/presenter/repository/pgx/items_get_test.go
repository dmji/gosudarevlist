package repository_pgx_test

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/dmji/go-animelayer-parser"
	"github.com/dmji/gosudarevlist/internal/presenter/model"
	"github.com/dmji/gosudarevlist/pkg/enums"
)

func TestGetITemsByCategory(t *testing.T) {
	repo, ctx := InitRepo(context.Background())

	_, err := repo.GetItems(ctx, model.OptionsGetItems{
		PageIndex:       1,
		CountForOnePage: 20000,

		SearchQuery:         "",
		SimilarityThreshold: 0.05,

		Categories: []enums.Category{enums.CategoryAnime},
		Statuses:   []enums.ReleaseStatus{},
	})
	if err != nil {
		t.Fatal(err)
	}
	identifiers := make(map[string]string)
	/*
		for _, item := range items {
			m := animelayer.TryGetSomthingSemantizedFromNotes(item.Description)
			 		s := traverseMapNotesSemantized("Разрешение", m)
			   		if s == "" {
			   			s = traverseMapNotesSemantized("Видео", m)
			   		}
			s := traverseMapNotesSemantized("Субтитры", m)
			if s == "" {
				identifiers[item.Title] = item.CategoryPresentation
			}
		}
	*/

	s, err := json.Marshal(&identifiers)
	if err != nil {
		t.Fatal(err)
	}
	os.WriteFile("result.json", s, 0o644)
}

func traverseMapNotesSemantized(tag string, m *animelayer.NotesSematizied) string {
	for _, t := range m.Taged {
		if t.Tag == tag && len(t.Text) != 0 {
			return t.Text
		}
		if t.Childs != nil {
			s := traverseMapNotesSemantized(tag, t.Childs)
			if s != "" {
				return s
			}
		}
	}
	return ""
}

func TestGetFiltersByCategory(t *testing.T) {
	repo, ctx := InitRepo(context.Background())

	items, _ := repo.GetFilters(ctx, model.OptionsGetItems{
		PageIndex:       0,
		CountForOnePage: 20,

		SearchQuery:         "",
		SimilarityThreshold: 0.05,

		Categories: []enums.Category{
			enums.CategoryAnime,
			enums.CategoryAnimeHentai,
		},
	})

	println(items)
}
