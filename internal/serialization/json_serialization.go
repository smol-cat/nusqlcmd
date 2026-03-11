package serialization

import (
	"database/sql"
	"encoding/json"

	"github.com/smol-cat/nusqlcmd/internal/core/mssql"
	"github.com/smol-cat/nusqlcmd/internal/core/sql_columns"
)

func scanRow(rows *sql.Rows, colTypes []*sql.ColumnType) ([]any, error) {
	sqlColumns := make([]sqlcolumns.SqlColumn, len(colTypes))
	mapTarget := make([]any, len(colTypes))

	for i, colType := range colTypes {
		typeName := colType.DatabaseTypeName()
		nullable, _ := colType.Nullable()
		sqlColumns[i] = mssql.MapTypeNameToSqlType(typeName, nullable)

		mapTarget[i] = sqlColumns[i].Value
	}

	if err := rows.Scan(mapTarget...); err != nil {
		return nil, err
	}

	resultCols := make([]any, len(colTypes))
	for i := range colTypes {
		resultCols[i] = sqlColumns[i].Scan(mapTarget[i])
	}

	return resultCols, nil
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
