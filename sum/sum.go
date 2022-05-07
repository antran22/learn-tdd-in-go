package sum

func Sum(array []int) int {
	sum := 0
	for _, element := range array {
		sum += element
	}
	return sum
}

func SumAll(inputs ...[]int) []int {
	result := make([]int, 0, len(inputs))
	for _, array := range inputs {
		result = append(result, Sum(array))
	}
	return result
}

func SumTails(inputs ...[]int) []int {
	result := make([]int, 0, len(inputs))
	for _, array := range inputs {
		result = append(result, Sum(Tails(array)))
	}
	return result
}

func Tails(array []int) []int {
	if len(array) == 0 {
		return []int{}
	}
	return array[1:]
}
