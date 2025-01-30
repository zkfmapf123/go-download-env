package filesystem

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

type fs struct {

}

func NewFS() *fs {
	return &fs{}
}

func (f fs) Dashboard(header []string, data [][]string) {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)

	for _, v := range data {
		table.Append(v)
	}

	table.Render()
}
