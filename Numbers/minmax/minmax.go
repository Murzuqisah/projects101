package minmax

func MaxMin(a, b int) (int, int) { 
	if a > b {
		return a, b
	} else {
		return b, a
	}
	

}