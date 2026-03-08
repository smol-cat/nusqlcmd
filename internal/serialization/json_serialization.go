package serialization

import (
	"database/sql"
	"encoding/json"

	"github.com/smol-cat/nusqlcmd/internal/common"
)

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
