package repository_pgx

import (
	pgx_sqlc "collector/pkg/recollection/repository/pgx/sqlc"
	"context"
	"errors"
	"time"

	"github.com/dmji/go-animelayer-parser"
	"github.com/jackc/pgx/v5"
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

func (repo *repository) InsertItem(ctx context.Context, item *animelayer.Item, category animelayer.Category) error {

	tx, err := repo.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	r := repo.query.WithTx(tx)

	categoryPgx, err := categoryToPgxCategory(category)
	if err != nil {
		return err
	}

	now := time.Now()
	lastCheckedDate, err := timeToPgTimestamp(&now)
	if err != nil {
		return err
	}

	createdDate, err := timeToPgTimestamp(item.Updated.CreatedDate)
	if err != nil {
		return err
	}

	updatedDate, err := timeToPgTimestamp(item.Updated.UpdatedDate)
	if err != nil {
		return err
	}

	itemId, err := r.InsertItem(ctx,
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

	if err != nil {
		return err
	}

	err = r.InsertUpdateNote(ctx, pgx_sqlc.InsertUpdateNoteParams{
		ItemID:      itemId,
		UpdateDate:  lastCheckedDate,
		UpdateTitle: "New",
	})

	//
	// Commit
	//
	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}
