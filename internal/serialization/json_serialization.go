package serialization

import (
	"database/sql"
	"encoding/json"

	"github.com/smol-cat/nusqlcmd/internal/core/mssql"
)

func scanRow(rows *sql.Rows, colTypes []*sql.ColumnType) ([]any, error) {
	rawCols := make([]any, len(colTypes))

	for i, colType := range colTypes {
		typeName := colType.DatabaseTypeName()
		nullable, _ := colType.Nullable()
		rawCols[i] = mssql.MapTypeNameToGoType(typeName, nullable)
	}

	if err := rows.Scan(rawCols...); err != nil {
		return nil, err
	}

	cols := make([]any, len(colTypes))
	for i := range rawCols {
		cols[i] = mssql.GetValueFromScanned(rawCols[i])
	}

	return cols, nil
}

func SerializeToJson(rows *sql.Rows) (string, error) {
	colNames, err := rows.Columns()
	if err != nil {
		return "", err
	}

	colTypes, err := rows.ColumnTypes()
	if err != nil {
		return "", err
	}

	result := []map[string]any{}
	for rows.Next() {
		cols, err := scanRow(rows, colTypes)
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
