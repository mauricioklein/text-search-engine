package report

import (
	"io"
)

// Reporter defines the interface for all the
// reporter implementations of the system
type Reporter interface {
	Report(out io.Writer, filename string, rank float64)
}
