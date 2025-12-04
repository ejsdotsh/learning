// Package protein translates RNA sequences into proteins.
package protein

import "errors"

var (
	ErrInvalidBase = errors.New("invalid BASE")
	ErrStop        = errors.New("STOP")
)

// FromRNA reurns a list of proteins and an error for an RNA string.
func FromRNA(rna string) ([]string, error) {
	n := 3
	r := []rune(rna)
	var protein []string

	for i := 0; i < len(r); i += n {
		end := i + n
		end = min(end, len(r))
		p, err := FromCodon(string(r[i:end]))
		if err != nil {
			if err == ErrStop {
				return protein, nil
			}
			return protein, err
		}
		protein = append(protein, p)
	}

	return protein, nil
}

// FromCodon returns the protein and an error for a codon.
func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop

	default:
		return "", ErrInvalidBase
	}

}
