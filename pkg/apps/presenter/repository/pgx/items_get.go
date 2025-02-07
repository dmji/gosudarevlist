package repository_pgx

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"github.com/dmji/go-animelayer-parser"
	"github.com/dmji/gosudarevlist/pkg/apps/presenter/model"
	pgx_sqlc "github.com/dmji/gosudarevlist/pkg/apps/presenter/repository/pgx/sqlc"
	"github.com/dmji/gosudarevlist/pkg/enums"
	"github.com/dmji/gosudarevlist/pkg/logger"
	"github.com/dmji/gosudarevlist/pkg/pgx_utils"
	"github.com/dmji/gosudarevlist/pkg/time_formater.go"
)

func categoryPresentation(ctx context.Context, s enums.Category, bShow bool) string {
	if bShow {
		return s.Presentation(ctx)
	}

	return ""
}

func (r *repository) GetItems(ctx context.Context, opt model.OptionsGetItems) ([]model.ItemCartData, error) {
	startID := (opt.PageIndex - 1) * opt.CountForOnePage

	items, err := r.query.GetItems(ctx, pgx_sqlc.GetItemsParams{
		Count:               opt.CountForOnePage,
		OffsetCount:         startID,
		SimilarityThreshold: opt.SimilarityThreshold,

		SearchQuery:   opt.SearchQuery,
		CategoryArray: categoriesToAnimelayerCategories(opt.Categories, true),
		StatusArray:   releaseStatusAnimelayerArrToPgxReleaseStatusAnimelayerArr(ctx, opt.Statuses, true),
	})
	if err != nil {
		logger.Errorw(ctx, "Pgx repo error | GetItems", "error", err)
		return nil, err
	}

	cardItems := make([]model.ItemCartData, 0, len(items))
	for _, item := range items {
		if item.Category == pgx_sqlc.CategoryAnimelayerAnime {
			if i := slices.IndexFunc(cardItems, func(e model.ItemCartData) bool { return e.Title == item.Title }); i != -1 {
				cardItems[i].AnimeLayerRefs = append(cardItems[i].AnimeLayerRefs, &model.ItemCartHrefData{
					Href:  fmt.Sprintf("https://animelayer.ru/torrent/%s/", item.Identifier),
					Text:  itemNotesToHrefText(len(cardItems[i].AnimeLayerRefs)+1, &item),
					Image: item.RefImageCover,
				})
				continue
			}
		}

		cardItems = append(cardItems, model.ItemCartData{
			Title:         item.Title,
			CreatedDate:   time_formater.Format(ctx, pgx_utils.TimeFromPgTimestamp(item.CreatedDate)),
			UpdatedDate:   time_formater.Format(ctx, pgx_utils.TimeFromPgTimestamp(item.UpdatedDate)),
			TorrentWeight: item.TorrentFilesSize,
			AnimeLayerRefs: []*model.ItemCartHrefData{
				{
					Href:        fmt.Sprintf("https://animelayer.ru/torrent/%s/", item.Identifier),
					Text:        itemNotesToHrefText(1, &item),
					Image:       item.RefImageCover,
					Description: itemNotesToDescriptions(&item),
				},
			},
			CategoryPresentation: categoryPresentation(ctx, pgxCategoriesToCategory(item.Category), len(opt.Categories) != 1),
			ReleaseStatus:        pgxReleaseStatusAnimelayerToReleaseStatusAnimelayer(ctx, item.ReleaseStatus),
		})
	}

	return cardItems, nil
}

func itemNotesToHrefText(i int, item *pgx_sqlc.AnimelayerItem) []string {
	baseText := fmt.Sprintf("Torrent №%d", i)
	if item.Category != pgx_sqlc.CategoryAnimelayerAnime {
		return []string{baseText}
	}

	m := animelayer.TryGetSomthingSemantizedFromNotes(item.Notes)

	resolution := traverseMapNotesSemantized("Разрешение", m)
	if resolution == "" {
		resolution = traverseMapNotesSemantized("Видео", m)
	}
	resolution += " " + item.TorrentFilesSize
	subs := traverseMapNotesSemantized("Субтитры", m)

	baseText = strings.Join([]string{baseText, resolution}, ": ")
	return []string{baseText, "Субтитры: " + subs}
}

func itemNotesToDescriptions(item *pgx_sqlc.AnimelayerItem) model.ItemCartDescriptions {
	m := animelayer.TryGetSomthingSemantizedFromNotes(item.Notes)
	return model.ItemCartDescriptions{
		Type:            traverseMapNotesSemantized("Тип", m),
		Genres:          splitGenres(traverseMapNotesSemantized("Жанр", m)),
		Year:            traverseMapNotesSemantized("Год выхода", m),
		EpisodeCount:    traverseMapNotesSemantized("Кол серий", m),
		EpisodeDuration: traverseMapNotesSemantized("Продолжительность", m),
		UpdateReaseon:   traverseMapNotesSemantized("Торрент был обновлен", m),
	}
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

func splitGenres(genreString string) []string {
	genres := strings.Split(genreString, ",")
	for gi := range genres {
		for {
			if len(genres[gi]) == 0 {
				break
			}
			if genres[gi][0] == ' ' || genres[gi][0] == ':' {
				genres[gi] = genres[gi][1:]
				continue
			}

			l := len(genres[gi]) - 1
			if genres[gi][l] == ' ' || genres[gi][l] == ':' {
				genres[gi] = genres[gi][:l-1]
				continue
			}
			break
		}
	}
	return slices.DeleteFunc(genres, func(e string) bool { return len(e) == 0 })
}
