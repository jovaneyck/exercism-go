package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("strands must be of equal length")
	}

	numberDifferences := 0
	for i := range a {
		if a[i] != b[i] {
			numberDifferences++
		}
	}
	return numberDifferences, nil
}
