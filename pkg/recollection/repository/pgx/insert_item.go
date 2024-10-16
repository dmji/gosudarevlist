package repository_pgx

import (
	pgx_sqlc "collector/pkg/recollection/repository/pgx/sqlc"
	"context"
	"errors"
	"time"

	"github.com/dmji/go-animelayer-parser"
	"github.com/jackc/pgx/v5/pgtype"
)

func categoryToPgxCategory(cat animelayer.Category) (pgx_sqlc.CategoryAnimelayer, error) {
	switch cat {
	case animelayer.Categories.Anime():
		return pgx_sqlc.CategoryAnimelayerAnime, nil
	case animelayer.Categories.AnimeHentai():
		return pgx_sqlc.CategoryAnimelayerAnimeHentai, nil
	case animelayer.Categories.Manga():
		return pgx_sqlc.CategoryAnimelayerManga, nil
	case animelayer.Categories.MangaHentai():
		return pgx_sqlc.CategoryAnimelayerMangaHentai, nil
	case animelayer.Categories.Dorama():
		return pgx_sqlc.CategoryAnimelayerDorama, nil
	case animelayer.Categories.Music():
		return pgx_sqlc.CategoryAnimelayerMusic, nil
	}

	return pgx_sqlc.CategoryAnimelayerAnime, errors.New("undefined category")
}

func (r *repository) InsertItem(ctx context.Context, item *animelayer.ItemDetailed, category animelayer.Category) error {

	categoryPgx, err := categoryToPgxCategory(category)
	if err != nil {
		return err
	}

	lastCheckedDate := pgtype.Timestamp{}
	if err := lastCheckedDate.Scan(time.Now()); err != nil {
		return err
	}

	createdDate := pgtype.Timestamp{}
	if item.Updated.CreatedDate != nil {
		if err := createdDate.Scan(*item.Updated.CreatedDate); err != nil {
			return err
		}
	}

	updatedDate := pgtype.Timestamp{}
	if item.Updated.UpdatedDate != nil {
		if err := updatedDate.Scan(*&item.Updated.UpdatedDate); err != nil {
			return err
		}
	}

	return r.query.InsertItem(ctx,
		pgx_sqlc.InsertItemParams{
			Identifier:       item.Identifier,
			Title:            item.Title,
			IsCompleted:      item.IsCompleted,
			LastCheckedDate:  lastCheckedDate,
			CreatedDate:      createdDate,
			UpdatedDate:      updatedDate,
			RefImageCover:    item.RefImageCover,
			RefImagePreview:  item.RefImagePreview,
			BlobImageCover:   "",
			BlobImagePreview: "",
			TorrentFilesSize: item.Metrics.FilesSize,
			Notes:            item.Notes,
			Category:         categoryPgx,
		},
	)

}
