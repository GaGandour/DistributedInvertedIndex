package invertedindex

func (ii InvertedIndex) retrieve(token string) []int {
	return ii.token2docs[token]
}
