package report

import (
	"fmt"
	"io"
)

// SimpleReporter defines the simple results
// reporter for the system
type SimpleReporter struct{}

// Report reports the rank result for a given file
func (sr SimpleReporter) Report(out io.Writer, filename string, rank float64) {
	out.Write([]byte(
		fmt.Sprintf("%s: %.2f%% match\n", filename, rank*100.0),
	))
}
