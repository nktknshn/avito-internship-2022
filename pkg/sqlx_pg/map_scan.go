package sqlx_pg

import (
	"context"

	"github.com/jmoiron/sqlx"
)

func NamedSelectMapScan(
	ctx context.Context,
	conn *sqlx.DB,
	dest *[]map[string]any,
	query string,
	arg any,
) error {

	queryStr, params, err := conn.BindNamed(query, arg)

	if err != nil {
		return err
	}

	rows, err := conn.QueryxContext(ctx, queryStr, params...)

	if err != nil {
		return err
	}

	defer rows.Close()

	if rowsErr := rows.Err(); rowsErr != nil {
		return rowsErr
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
