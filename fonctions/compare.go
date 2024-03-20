package fonctions

func Compare(stack_A, res_compare []int) bool {

	for i := 0; i < len(stack_A); i++ {
		if stack_A[i] != res_compare[i] {
			return false
		}
	}
	return true
}
