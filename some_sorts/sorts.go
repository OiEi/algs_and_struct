package some_sorts

func QuickSort(input []int) []int {
	if len(input) < 2 {
		return input
	}

	startElem := input[0]

	var leftArr []int
	var rightArr []int

	for _, k := range input[1:] {
		if startElem > k {
			leftArr = append(leftArr, k)
		} else {
			rightArr = append(rightArr, k)
		}
	}

	input = append(QuickSort(leftArr), startElem)
	input = append(input, QuickSort(rightArr)...)

	return input
}

func BubleSort(input []int) []int {

	inputLen := len(input)

	for i := inputLen; i > 1; i-- {
		for k, _ := range input[:i] {
			if k != inputLen-1 && input[k] > input[k+1] {
				input[k], input[k+1] = input[k+1], input[k]
			}
		}
	}

	return input
}
