package filesystem

import "testing"

func Test_Dashboard(t *testing.T) {
	
	fs := NewFS()

	fs.Dashboard([]string{"Name", "Age"}, [][]string{{"John", "20"}, {"Jane", "21"}})
}