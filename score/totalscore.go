package score

// Total returns the total score computed over the set of scores provided.
// The total is a grade in the range 0-100.
func Total(scores []*Score) uint8 {
	totalWeight := float32(0)
	var max, score, weight []float32
	for _, ts := range scores {
		totalWeight += float32(ts.Weight)
		weight = append(weight, float32(ts.Weight))
		score = append(score, float32(ts.Score))
		max = append(max, float32(ts.MaxScore))
	}

	total := float32(0)
	for i := 0; i < len(scores); i++ {
		if score[i] > max[i] {
			score[i] = max[i]
		}
		total += ((score[i] / max[i]) * (weight[i] / totalWeight))
	}

	return uint8(total * 100)
}
