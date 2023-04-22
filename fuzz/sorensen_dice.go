package fuzz

func SorensenDice(source, target string) float64 {
	r1 := []rune(source)
	r2 := []rune(target)

	i := make(map[rune]bool)

	for _, r := range r1 {
		for _, r2 := range r2 {
			if r == r2 {
				i[r] = true
				break
			}
		}
	}

	s1 := len(r1)
	s2 := len(r2)
	sI := len(i)

	return float64(2*sI) / float64(s1+s2)
}
