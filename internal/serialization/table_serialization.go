package serialization

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/smol-cat/nusqlcmd/internal/common"
)

const MAX_COL_LENGTH = 30
const MIN_COL_LENGTH = 10

func getColLengths(colTypes []*sql.ColumnType) []int {
	lengths := []int{}

	for i := range colTypes {
		length, _ := colTypes[i].Length()
		length = min(MAX_COL_LENGTH, length)
		length = max(MIN_COL_LENGTH, length)
		lengths = append(lengths, int(length))
	}

	return lengths
}

func writeRow(table *strings.Builder, cols []string, lengths []int) {
	row := strings.Builder{}
	row.WriteString("|")

	for i := range cols {
		column := fmt.Sprintf("%-"+strconv.Itoa(lengths[i])+"s|", cols[i])
		row.WriteString(column)
	}

	table.WriteString(row.String() + "\n")
}

func writeTopRow(table *strings.Builder, colNames []string, lengths []int) {
	topRow := strings.Builder{}
	topRow.WriteString("|")

	cols := []string{}
	lenSum := 0
	for i := range colNames {
		lenSum += lengths[i]
		cols = append(cols, colNames[i])
	}

	var line = strings.Repeat("-", lenSum + len(colNames) + 1) + "\n"
	table.WriteString(line)
	writeRow(table, cols, lengths)
	table.WriteString(line)
}

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

func SerializeRowsToTable(rows *sql.Rows) string {
	colNames, err := rows.Columns()
	common.PanicOnErr(err)
	colTypes, err := rows.ColumnTypes()
	common.PanicOnErr(err)

	lengths := getColLengths(colTypes)

	table := strings.Builder{}
	writeTopRow(&table, colNames, lengths)

	for rows.Next() {
		cols := scanRow(rows, len(colTypes))
		writeRow(&table, cols, lengths)
	}

	return table.String()
}
