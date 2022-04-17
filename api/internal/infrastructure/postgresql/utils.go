package postgresql

import (
	"github.com/jackc/pgconn"
	"github.com/mark-by/little-busy-back/api/internal/domain/entity"
	"github.com/pkg/errors"
)

func convertPgxError(err error) error {
	var pgErr *pgconn.PgError
	if !errors.As(err, &pgErr) {
		return err
	}

	switch pgErr.Code {
	case "23505":
		return entity.DuplicateError
	default:
		return err
	}
}
