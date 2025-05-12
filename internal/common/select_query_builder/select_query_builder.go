package select_query_builder

import (
	"strings"
)

func New() *SelectQueryBuilder {
	return &SelectQueryBuilder{
		Ctes:          []string{},
		AndConditions: []string{},
		Query:         "",
		Cursor:        "",
		Order:         "",
		GroupBy:       "",
		Having:        "",
		InnerJoins:    []string{},
		LeftJoins:     []string{},
		Limit:         "",
	}
}

type SelectQueryBuilder struct {
	Ctes          []string
	AndConditions []string
	Query         string
	Cursor        string
	Order         string
	GroupBy       string
	Having        string
	InnerJoins    []string
	LeftJoins     []string
	Limit         string
}

func (s *SelectQueryBuilder) And(condition string) {
	s.AndConditions = append(s.AndConditions, condition)
}

func (s *SelectQueryBuilder) Build() string {
	q := s.Query

	ctesStr := ""
	if len(s.Ctes) > 0 {
		ctesStr = "WITH " + strings.Join(s.Ctes, ",\n")
	}

	if ctesStr != "" {
		q = ctesStr + "\n" + q
	}

	conditionsStr := " WHERE 1=1 \n"

	if len(s.AndConditions) > 0 {
		conditionsStr += " AND " + strings.Join(s.AndConditions, " AND\n")
	}

	if s.Cursor != "" {
		conditionsStr = conditionsStr + " AND " + s.Cursor
	}

	if len(s.InnerJoins) > 0 {
		q += "INNER JOIN " + strings.Join(s.InnerJoins, "\nINNER JOIN ")
		q += "\n"
	}

	if len(s.LeftJoins) > 0 {
		q += "LEFT JOIN " + strings.Join(s.LeftJoins, "\nLEFT JOIN ")
		q += "\n"
	}

	q += conditionsStr

	if s.GroupBy != "" {
		q += " GROUP BY " + s.GroupBy
	}

	if s.Having != "" {
		q += " HAVING " + s.Having
	}

	if s.Order != "" {
		q += " ORDER BY " + s.Order
	}

	if s.Limit != "" {
		q += " LIMIT " + s.Limit
	}

	return q
}
