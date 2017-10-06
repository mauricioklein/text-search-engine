package ranking

// Algorithm defines the interface for the ranking algorithms
type Algorithm interface {
	Calculate(text string, sentence string) float64
}
