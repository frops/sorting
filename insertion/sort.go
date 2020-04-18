package insertion

func sort(sliceToSort []int64) []int64 {
	for i := 1; i < len(sliceToSort); i++ {
		stone := sliceToSort[i]
		j := i - 1
		for j >= 0 && sliceToSort[j] > stone {
			sliceToSort[j+1] = sliceToSort[j]
			j--
		}
		sliceToSort[j+1] = stone
	}
	return sliceToSort
}
