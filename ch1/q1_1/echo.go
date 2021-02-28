package q1_1

func Echo(input []string) string {
	var s, sep string
	for i := 0; i < len(input); i++ {
		s += sep + input[i]
		sep = " "
	}

	return s
}
