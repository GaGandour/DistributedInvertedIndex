package invertedindex

func (ii InvertedIndex) Retrieve(token string) []int {
	return ii.Token2docs[token]
}
