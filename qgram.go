package fuzzy

func QGram(source string, target string, q int) float64 {
	if q <= 0 {
		return 0.0
	}

	qGrams1 := generateQGrams(source, q)
	qGrams2 := generateQGrams(target, q)

	i := 0

	for _, q1 := range qGrams1 {
		for j, q2 := range qGrams2 {
			if q1 == q2 {
				i++
				qGrams2 = append(qGrams2[:j], qGrams2[j+1:]...)
				break
			}
		}
	}

	return float64(i) / float64(len(qGrams1)+len(qGrams2))
}

func generateQGrams(s string, q int) []string {
	var qGrams []string

	for i := 0; i < len(s)-q+1; i++ {
		qGrams = append(qGrams, s[i:i+q])
	}

	return qGrams
}
