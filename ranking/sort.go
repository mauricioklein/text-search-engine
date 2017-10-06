package ranking

// ByScoreAndName defines the sorting algorithm of
// result by ranking and filename
type ByScoreAndName []RankResult

func (bsan ByScoreAndName) Len() int {
	return len(bsan)
}

func (bsan ByScoreAndName) Swap(i, j int) {
	bsan[i], bsan[j] = bsan[j], bsan[i]
}

func (bsan ByScoreAndName) Less(i, j int) bool {
	if bsan[i].Rank != bsan[j].Rank {
		// Sort by rank
		return bsan[i].Rank < bsan[j].Rank
	}

	/*
		Sort by filename:

		Here we use some tricky approach:

		If we sort the names ASCENDINGLY, we'll be displaying the files with
		same rank in the reverse order, because the matches are displayed
		reversely (from the biggest to the smallest match):
			file1.txt: 75%
		   	file3.txt: 50%
			file2.txt: 50%

		Sorting the names DESCENDINGLY causes the files with same rank to be
		displayed in the "reverse-reverse order", such as:
			file1.txt: 75%
			file2.txt: 50%
			file3.txt: 50%
	*/
	return bsan[i].File.Name() > bsan[j].File.Name()
}
