package serialization

import (
	"database/sql"
	"encoding/json"

	"github.com/smol-cat/nusqlcmd/internal/config"
	"github.com/smol-cat/nusqlcmd/internal/core"
)

func getSqlColumns(colTypes []*sql.ColumnType, driver string) ([]core.SqlColumn, error) {
	sqlColumns := make([]core.SqlColumn, len(colTypes))
	mapType, err := getMapper(driver)
	if err != nil {
		return nil, err
	}

	for i, colType := range colTypes {
		typeName := colType.DatabaseTypeName()
		nullable, _ := colType.Nullable()
		sqlColumns[i] = mapType(typeName, nullable)
	}

	return sqlColumns, nil
}

func scanRow(rows *sql.Rows, sqlColumns []core.SqlColumn) ([]any, error) {
	mapTarget := make([]any, len(sqlColumns))

	for i, col := range sqlColumns {
		mapTarget[i] = col.AllocateValue()
	}

	if err := rows.Scan(mapTarget...); err != nil {
		return nil, err
	}

	resultCols := make([]any, len(sqlColumns))
	for i := range sqlColumns {
		resultCols[i] = sqlColumns[i].Scan(mapTarget[i])
	}

	return resultCols, nil
}

func SerializeToJson(rows *sql.Rows, runtimeConfig config.RuntimeConfig) (string, error) {
	colNames, err := rows.Columns()
	if err != nil {
		return "", err
	}

	colTypes, err := rows.ColumnTypes()
	if err != nil {
		return "", err
	}

	sqlColumns, err := getSqlColumns(colTypes, runtimeConfig.Driver)
	if err != nil {
		return "", err
	}

	result := []map[string]any{}
	for rows.Next() {
		cols, err := scanRow(rows, sqlColumns)
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
