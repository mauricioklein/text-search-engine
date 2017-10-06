package report

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleReporter(t *testing.T) {
	sr := SimpleReporter{}
	buf := bytes.NewBuffer([]byte{})

	// Report a result
	sr.Report(buf, "file1.txt", 0.5432)

	actual, _ := buf.ReadString('\n')
	expected := "file1.txt: 54.32% match\n"
	assert.Equal(t, expected, actual)
}
