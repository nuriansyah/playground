package sortKM

// TODO: answer here
func bubbleSort(array []int) []int {
	return array
}

//concrete strategy implementation
type AscendingSort struct{}

func (as *AscendingSort) Sort(array []int) {
	//choose any sort algo you want
	// TODO: answer here
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}

}
