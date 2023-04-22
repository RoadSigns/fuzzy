package fuzzy

// Levenshtein allows us to calculate the edit distance between the source and the target.
// Additional information for this algorithm can be found https://en.wikipedia.org/wiki/Levenshtein_distance
func Levenshtein(source, target string) int {
	s := len(source)
	t := len(target)

	dist := make([][]int, s+1)
	for i := range dist {
		dist[i] = make([]int, t+1)
	}

	for i := 0; i <= s; i++ {
		dist[i][0] = i
	}
	for j := 0; j <= t; j++ {
		dist[0][j] = j
	}

	for j := 1; j <= t; j++ {
		for i := 1; i <= s; i++ {
			if source[i-1] == target[j-1] {
				dist[i][j] = dist[i-1][j-1]
			} else {
				dist[i][j] = min(dist[i-1][j]+1, min(dist[i][j-1]+1, dist[i-1][j-1]+1))
			}
		}
	}

	return dist[s][t]
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
