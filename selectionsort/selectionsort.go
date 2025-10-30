package selectionsort

// SmallestElementIndex gibt den Index des kleinsten Elements im Array arr ab Index start zurück.
func SmallestElementIndex(arr []int, start int) int {
	// TODO
	minIndex := start
	for i := start + 1; i < len(arr); i++ {
		if arr[i] < arr[minIndex] {
			minIndex = i
		}
	}
	return minIndex
}

// SelectionSort sortiert die übergebene Liste mittels des Selection-Sort-Algorithmus.
func SelectionSort(arr []int) {
	// TODO
	for i := 0; i < len(arr)-1; i++ {
		minIndex := SmallestElementIndex(arr, i)
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}

}
