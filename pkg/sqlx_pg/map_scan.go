package sqlx_pg

import (
	"context"

	"github.com/jmoiron/sqlx"
)

func NamedSelectMapScan(ctx context.Context, conn *sqlx.DB, dest *[]map[string]any, query string, arg interface{}) error {
	const action = "PostgresAdapter NamedSelectMapScan "
	const method = "NamedSelectMapScan"

	queryStr, params, err := conn.BindNamed(query, arg)

	if err != nil {
		return err
	}

	rows, err := conn.QueryxContext(ctx, queryStr, params...)

	if err != nil {
		return err
	}

	for rows.Next() {
		row := map[string]any{}
		err = rows.MapScan(row)
		if err != nil {
			return err
		}
		*dest = append(*dest, row)
	}

	return err
}
