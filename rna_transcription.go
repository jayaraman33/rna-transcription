package strand

import "strings"

var dnaToRna = map[rune]rune{

	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

func ToRNA(dna string) string {
	toRNA := func(nucleotide rune) rune {
		return dnaToRna[nucleotide]
	}

	return strings.Map(toRNA, dna)
}
