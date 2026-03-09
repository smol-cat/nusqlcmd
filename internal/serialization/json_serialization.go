package serialization

import (
	"database/sql"
	"encoding/json"
)

func scanRow(rows *sql.Rows, colCount int) ([]*string, error) {
	var rawCols = make([]any, colCount)

	for i := range rawCols {
		var alloc *string
		rawCols[i] = &alloc
	}

	if err := rows.Scan(rawCols...); err != nil {
		return nil, err
	}

	var cols = make([]*string, colCount)
	for i := range rawCols {
		cols[i] = *(rawCols[i].(**string))
	}

	return cols, nil
}

func SerializeToJson(rows *sql.Rows) (string, error) {
	colNames, err := rows.Columns()
	if err != nil {
		return "", err
	}

	result := []map[string]any{}
	for rows.Next() {
		cols, err := scanRow(rows, len(colNames))
		if err != nil {
			return "", err
		}

		colsMap := map[string]any{}
		for i := range colNames {
			colsMap[colNames[i]] = cols[i]
		}

		result = append(result, colsMap)
	}

	resultString, err := json.MarshalIndent(result, "", "    ")
	return string(resultString), err
}
