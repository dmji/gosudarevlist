package animelayer_client

import (
	"context"
	"time"

	"github.com/dmji/go-animelayer-parser"
	"github.com/dmji/gosudarevlist/internal/updater/model"
	"github.com/dmji/gosudarevlist/pkg/enums"
)

func New(i animelayer.ItemProvider) *client {
	return &client{
		client: i,
	}
}

type client struct {
	client animelayer.ItemProvider
}

func (c *client) GetItemByIdentifier(ctx context.Context, identifier string) (*model.AnimelayerItem, error) {
	item, err := c.client.GetItemByIdentifier(ctx, identifier)
	if err != nil {
		return nil, err
	}

	return conv(item), nil
}

func (c *client) GetItemsFromCategoryPages(ctx context.Context, category enums.Category, iPage int) ([]*model.AnimelayerItem, error) {
	items, err := c.client.GetItemsFromCategoryPages(ctx, modelCategoryToAnimelayerCategory(category), iPage)
	if err != nil {
		return nil, err
	}

	itemsModel := make([]*model.AnimelayerItem, 0, len(items))
	for i := range items {
		itemsModel = append(itemsModel, conv(&items[i]))
	}
	return itemsModel, nil
}

func conv(item *animelayer.Item) *model.AnimelayerItem {
	lastCheckedDate := time.Now()

	releaseStatus := enums.ReleaseStatusOnAir
	if item.IsCompleted {
		releaseStatus = enums.ReleaseStatusCompleted
	} else {
		data := item.Updated.UpdatedDate
		if data == nil {
			data = item.Updated.CreatedDate
		}
		if data != nil {
			yearAfterUpdate := data.Add(time.Hour * 8760 /*year*/)
			if lastCheckedDate.After(yearAfterUpdate) {
				releaseStatus = enums.ReleaseStatusIncompleted
			}
		}
	}

	return &model.AnimelayerItem{
		Identifier:       item.Identifier,
		Title:            item.Title,
		ReleaseStatus:    releaseStatus,
		LastCheckedDate:  &lastCheckedDate,
		CreatedDate:      item.Updated.CreatedDate,
		UpdatedDate:      item.Updated.UpdatedDate,
		RefImageCover:    item.RefImageCover,
		RefImagePreview:  item.RefImagePreview,
		BlobImageCover:   "",
		BlobImagePreview: "",
		TorrentFilesSize: item.Metrics.FilesSize,
		Notes:            item.Notes,
		Category:         animelayerCategoryToModelCategory(item.Category),
	}
}

func animelayerCategoryToModelCategory(category animelayer.Category) enums.Category {
	switch category {
	case animelayer.CategoryAnime:
		return enums.CategoryAnime
	case animelayer.CategoryAnimeHentai:
		return enums.CategoryAnimeHentai
	case animelayer.CategoryManga:
		return enums.CategoryManga
	case animelayer.CategoryMangaHentai:
		return enums.CategoryMangaHentai
	case animelayer.CategoryMusic:
		return enums.CategoryMusic
	case animelayer.CategoryDorama:
		return enums.CategoryDorama
	default:
		return enums.CategoryAnime
	}
}

func modelCategoryToAnimelayerCategory(category enums.Category) animelayer.Category {
	switch category {
	case enums.CategoryAnime:
		return animelayer.CategoryAnime
	case enums.CategoryAnimeHentai:
		return animelayer.CategoryAnimeHentai
	case enums.CategoryManga:
		return animelayer.CategoryManga
	case enums.CategoryMangaHentai:
		return animelayer.CategoryMangaHentai
	case enums.CategoryMusic:
		return animelayer.CategoryMusic
	case enums.CategoryDorama:
		return animelayer.CategoryDorama
	case enums.CategoryAll:
		return animelayer.CategoryAll
	default:
		return animelayer.CategoryAnime
	}
}
