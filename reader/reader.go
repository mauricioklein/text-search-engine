package reader

// Reader defines an interface that implements the
// methods to read files from some resource and translate
// it to File instances
type Reader interface {
	Read(path string) ([]File, error)
}
