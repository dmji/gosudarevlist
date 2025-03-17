package service

import (
	"context"

	"github.com/dmji/go-myanimelist/mal/maltype"
)

func (s *service) GetItem(ctx context.Context, id int) (*maltype.Anime, error) {
	opts := s.client.Anime.DetailsOptions
	t, _, err := s.client.Anime.Details(ctx, id,
		opts.Fields(
			opts.AnimeFields.ID(),
			opts.AnimeFields.Title(),
			opts.AnimeFields.MainPicture(),
			opts.AnimeFields.AlternativeTitles(),
			opts.AnimeFields.StartDate(),
			opts.AnimeFields.EndDate(),
			opts.AnimeFields.Synopsis(),
			opts.AnimeFields.Mean(),
			opts.AnimeFields.Rank(),
			opts.AnimeFields.Popularity(),
			opts.AnimeFields.NumListUsers(),
			opts.AnimeFields.NumScoringUsers(),
			opts.AnimeFields.NSFW(),
			opts.AnimeFields.CreatedAt(),
			opts.AnimeFields.UpdatedAt(),
			opts.AnimeFields.MediaType(),
			opts.AnimeFields.Status(),
			opts.AnimeFields.Genres(),
			opts.AnimeFields.MyListStatus(),
			opts.AnimeFields.NumEpisodes(),
			opts.AnimeFields.StartSeason(),
			opts.AnimeFields.Broadcast(),
			opts.AnimeFields.Source(),
			opts.AnimeFields.AverageEpisodeDuration(),
			opts.AnimeFields.Rating(),
			opts.AnimeFields.Pictures(),
			opts.AnimeFields.Background(),
			opts.AnimeFields.RelatedAnime(),
			opts.AnimeFields.RelatedManga(),
			opts.AnimeFields.Recommendations(),
			opts.AnimeFields.Studios(),
			opts.AnimeFields.Statistics(),
		))
	if err != nil {
		return nil, err
	}

	return t, nil
}
