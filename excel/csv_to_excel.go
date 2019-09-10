package excel

import (
	"errors"
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

const MAX_Column = 26

var (
	ColumnPrefix = [MAX_Column]rune{}
)

func init() {

	var (
		start = int('A')
	)

	for i := range ColumnPrefix {
		ColumnPrefix[i] = rune(start + i)
	}

}

func New() *ExcelFile {
	e := &ExcelFile{
		Sheet: "Sheet1",
		File:  excelize.NewFile(),
		Row:   1,
	}
	return e
}

type ExcelFile struct {
	Sheet string
	File  *excelize.File
	Row   int
}

func (e *ExcelFile) ChangeSheet(sheet string) {
	e.Sheet = sheet
}

func (e *ExcelFile) AddRow(row []interface{}) error {

	if len(row) > MAX_Column {
		return errors.New(fmt.Sprintf("要插入的列长度%d超过目前代码设定的最大数量%d", len(row), MAX_Column))
	}

	for i, v := range row {
		e.File.SetCellValue(e.Sheet, GetAxisID(i, e.Row), v)
	}

	e.Row++

	return nil
}

func GetAxisID(column int, row int) string {
	return fmt.Sprintf("%c%d", ColumnPrefix[column], row)
}

func (e *ExcelFile) SaveAs(path string) error {
	return e.File.SaveAs(path)
}
