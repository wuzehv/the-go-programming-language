package q2_5

func PopCount(x uint64) (num int) {
	for ; x > 0; x &= x - 1 {
		num++
	}

	return
}
