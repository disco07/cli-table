package cli_table

type Table struct {
	header Rows
	rows   []Rows
	border Border
}

func (t *Table) New() *Table {
	return &Table{}
}

func (t *Table) addRow(row []Rows) *Table {
	for _, r := range row {
		t.rows = append(t.rows, r)
	}

	return t
}
