package main

import (
	"fmt"
	"io"
)

// Reporter defines the interface for all the
// reporter implementations of the system
type Reporter interface {
	Report(out io.Writer, filename string, rank float64)
}

// SimpleReporter defines the simple results
// reporter for the system
type SimpleReporter struct{}

// Report reports the rank result for a given file
func (sr SimpleReporter) Report(out io.Writer, filename string, rank float64) {
	out.Write([]byte(
		fmt.Sprintf("%s: %.2f%% match\n", filename, rank*100.0),
	))
}
