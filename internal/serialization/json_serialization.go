package serialization

import (
	"database/sql"
	"encoding/json"

	"github.com/smol-cat/nusqlcmd/internal/common"
)

func scanRow(rows *sql.Rows, colCount int) []string {
	var rawCols = make([]any, colCount)

	for i := range rawCols {
		var alloc string
		rawCols[i] = &alloc
	}

	err := rows.Scan(rawCols...)
	common.PanicOnErr(err)

	var cols = make([]string, colCount)
	for i := range rawCols {
		cols[i] = *(rawCols[i].(*string))
	}

	return cols
}

func SerializeToJson(rows *sql.Rows) string {
	colNames, err := rows.Columns()
	common.PanicOnErr(err)

	result := []map[string]any{}
	for rows.Next() {
		cols := scanRow(rows, len(colNames))

		colsMap := map[string]any{}
		for i := range colNames {
			colsMap[colNames[i]] = cols[i]
		}

		result = append(result, colsMap)
	}

	resultString, err := json.MarshalIndent(result, "", "    ")
	common.PanicOnErr(err)

	return string(resultString)
}
