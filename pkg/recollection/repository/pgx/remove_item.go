package repository_pgx

import "context"

func (r *repository) RemoveItem(ctx context.Context, identifier string) error {

	return r.query.RemoveItem(ctx, identifier)

}
