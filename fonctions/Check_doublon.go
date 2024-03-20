package fonctions

func Check_doublon(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				return false
			}
		}
	}
	return true
}
