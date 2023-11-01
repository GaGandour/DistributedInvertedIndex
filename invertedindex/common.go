package invertedindex

// InvertedIndex is the data structure that will be used for searching documents
type InvertedIndex struct {
	token2docs map[string][]int
	docs       []int
}

// Task is the exposed struct of the Framework that the calling code should initialize
// with the specific implementation of the operation.
type Task struct {
	// Fundamental argument
	query string

	// Inverted Index functions
	Intersect IntersectFunc
	Fetch     FetchFunc

	// Number of partitions in the query
	NumPartitions int

	// Channels for data
	InputChan  chan []string
	OutputChan chan []int
}

type (
	IntersectFunc func([]int) []int
	FetchFunc     func([]string) []int
)
