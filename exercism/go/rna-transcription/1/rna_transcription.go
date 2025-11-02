package strand

import "strings"

func ToRNA(dna string) string {
	rna := []string{}
	for _, n := range dna {
		switch n {
		case 'G':
			rna = append(rna, "C")
		case 'C':
			rna = append(rna, "G")
		case 'T':
			rna = append(rna, "A")
		case 'A':
			rna = append(rna, "U")

		}
	}
	return strings.Join(rna, "")
}
