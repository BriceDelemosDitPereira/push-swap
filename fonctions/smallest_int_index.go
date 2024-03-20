package fonctions

func Smallest_int_index(stack_A []int) int {

	smallestIndex := 0          // Variable pour stocker l'index du plus petit élément
	smallestValue := stack_A[0] // Variable pour stocker la valeur du plus petit élément

	for i := 1; i < len(stack_A); i++ {
		if stack_A[i] < smallestValue {
			smallestValue = stack_A[i]
			smallestIndex = i
		}
	}
	return smallestIndex
}
