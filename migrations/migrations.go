package migrations

import (
	"context"
	"embed"
	"fmt"

	"github.com/dmji/gosudarevlist/pkg/logger"
	"github.com/pressly/goose/v3"
)

//go:embed */*.sql
var embedMigrations embed.FS

func Up(ctx context.Context, driver, dbstring, dialect, dir string) error {
	goose.SetLogger(&loggerGoose{ctx: ctx})
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect(dialect); err != nil {
		return err
	}

	db, err := goose.OpenDBWithDriver(driver, dbstring)
	if err != nil {
		return err
	}

	if err := goose.UpContext(ctx, db, dir); err != nil {
		return err
	}

	return nil
}

type loggerGoose struct {
	ctx context.Context
}

func (l *loggerGoose) Fatalf(format string, v ...interface{}) {
	logger.Fatalw(l.ctx, fmt.Sprintf(format, v...))
}

func (l *loggerGoose) Printf(format string, v ...interface{}) {
	logger.Infow(l.ctx, fmt.Sprintf(format, v...))
}
