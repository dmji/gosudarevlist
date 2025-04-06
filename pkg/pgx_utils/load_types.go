package pgx_utils

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func loadType(ctx context.Context, conn *pgx.Conn, typeName string, bArray bool) error {
	dt, err := conn.LoadType(ctx, typeName)
	if err != nil {
		return err
	}

	conn.TypeMap().RegisterType(dt)

	if bArray {
		dt, err = conn.LoadType(ctx, "_"+typeName)
		if err != nil {
			return err
		}

		conn.TypeMap().RegisterType(dt)
	}

	return nil
}

func AnimelayerPostgresAfterConnectFunction() func(ctx context.Context, conn *pgx.Conn) error {
	return func(ctx context.Context, conn *pgx.Conn) error {
		err := loadType(ctx, conn, "RELEASE_STATUS_ANIMELAYER", true)
		if err != nil {
			return err
		}

		return nil
	}
}
