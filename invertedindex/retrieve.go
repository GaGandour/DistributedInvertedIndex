package invertedindex

import "log"

func (ii InvertedIndex) Retrieve(token string) []int {
	log.Println("Retrieving token:", token)
	return []int{1, 2, 3, 4}
	// return ii.Token2docs[token]
}
