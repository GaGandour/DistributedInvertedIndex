package invertedindex

func Intersect(orderedList1 []int, orderedList2 []int) []int {
	length1 := len(orderedList1)
	length2 := len(orderedList2)

	result := []int{}

	i1 := 0
	i2 := 0
	for i1 < length1 && i2 < length2 {
		if orderedList1[i1] == orderedList2[i2] {
			result = append(result, orderedList1[i1])
			i1++
			i2++
		} else if orderedList1[i1] < orderedList2[i2] {
			i1++
		} else {
			i2++
		}
	}
	return result
}
