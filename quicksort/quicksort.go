package mergesort

// Partition teilt das Array arr in zwei Teile, basierend auf dem ersten Element
// (dem Pivot). Alle Elemente kleiner oder gleich dem Pivot werden nach links,
// alle größeren nach rechts einsortiert. Gibt die neue Position des Pivotelements zurück.
func Partition(arr []int) int {
	// TODO
	i, j := 0, len(arr)-1
	pivot := arr[0]

	for i < j {
		for i < j && arr[j] > pivot {
			j--
		}
		for i < j && arr[i] < pivot {
			i++
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	return i
}

// QuickSort sortiert die übergebene Liste mittels des Quick-Sort-Algorithmus.
func QuickSort(arr []int) {
	// TODO
	if len(arr) < 2 {
		return
	}
}
