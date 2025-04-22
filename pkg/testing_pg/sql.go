package testing_pg

import (
	"bytes"
	"fmt"
	"slices"
	"strings"
	"text/tabwriter"
)

type ResultRows struct {
	Headers []string
	Rows    []map[string]any
}

func (r *ResultRows) TabbedString(pickHeaders ...string) string {
	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 0, 2, ' ', 0)

	for i, header := range r.Headers {
		if len(pickHeaders) > 0 && !slices.Contains(pickHeaders, header) {
			continue
		}
		if i > 0 {
			fmt.Fprint(w, "\t")
		}
		fmt.Fprint(w, header)
	}
	fmt.Fprintln(w)

	for i, header := range r.Headers {
		if len(pickHeaders) > 0 && !slices.Contains(pickHeaders, header) {
			continue
		}
		if i > 0 {
			fmt.Fprint(w, "\t")
		}
		fmt.Fprint(w, strings.Repeat("-", len(header)))
	}
	fmt.Fprintln(w)

	for _, row := range r.Rows {
		for i, header := range r.Headers {
			if len(pickHeaders) > 0 && !slices.Contains(pickHeaders, header) {
				continue
			}
			if i > 0 {
				fmt.Fprint(w, "\t")
			}
			val := row[header]
			fmt.Fprintf(w, "%v", val)
		}
		fmt.Fprintln(w)
	}

	w.Flush()
	return buf.String()
}
