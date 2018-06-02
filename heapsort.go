package heapsort

func heapify(arr []int, n, i int) {
	largest := i
	l := i*2 + 1
	r := i*2 + 2

	if l < n && arr[i] < arr[l] {
		largest = l
	}

	if r < n && arr[largest] < arr[r] {
		largest = r
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}

func heapsort(arr []int) {
	n := len(arr)
	for i := n; i > -1; i-- {
		heapify(arr, n, i)
	}

	for i := n - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		heapify(arr, i, 0)
	}
}

type Item struct {
	Name  string
	Value int
}

func heapsortItem(arr []Item) {
	n := len(arr)
	for i := n; i > -1; i-- {
		heapifyItem(arr, n, i)
	}

	for i := n - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		heapifyItem(arr, i, 0)
	}
}

func heapifyItem(arr []Item, n, i int) {
	largest := i
	l := i*2 + 1
	r := i*2 + 2

	if l < n && arr[i].Value < arr[l].Value {
		largest = l
	}

	if r < n && arr[largest].Value < arr[r].Value {
		largest = r
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapifyItem(arr, n, largest)
	}
}
