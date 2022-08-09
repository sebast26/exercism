package protein

import "errors"

var (
	translations = map[string]string{
		"AUG": "Methionine",
		"UUU": "Phenylalanine",
		"UUC": "Phenylalanine",
		"UUA": "Leucine",
		"UUG": "Leucine",
		"UCU": "Serine",
		"UCC": "Serine",
		"UCA": "Serine",
		"UCG": "Serine",
		"UAU": "Tyrosine",
		"UAC": "Tyrosine",
		"UGU": "Cysteine",
		"UGC": "Cysteine",
		"UGG": "Tryptophan",
	}
	stops = map[string]struct{}{
		"UAA": {},
		"UAG": {},
		"UGA": {},
	}
	ErrStop        = errors.New("stop")
	ErrInvalidBase = errors.New("invalid base")
)

func FromRNA(rna string) ([]string, error) {
	if len(rna) % 3 != 0 {
		return nil, ErrInvalidBase
	}

	chunks := chunks(rna, 3)
	var proteins []string
	for _, chank := range chunks {
		p, err := FromCodon(chank)
		if err != nil {
			if err == ErrStop {
				return proteins, nil
			}
			return proteins, err
		}
		proteins = append(proteins, p)
	}
	return proteins, nil
}

func FromCodon(codon string) (string, error) {
	if _, ok := stops[codon]; ok {
		return "",ErrStop
	}

	protein, ok := translations[codon]
	if !ok {
		return "", ErrInvalidBase
	}
	return protein, nil
}

func chunks(s string, chunkSize int) []string {
    if len(s) == 0 {
        return nil
    }
    if chunkSize >= len(s) {
        return []string{s}
    }
    var chunks = make([]string, 0, (len(s)-1)/chunkSize+1)
    currentLen := 0
    currentStart := 0
    for i := range s {
        if currentLen == chunkSize {
            chunks = append(chunks, s[currentStart:i])
            currentLen = 0
            currentStart = i
        }
        currentLen++
    }
    chunks = append(chunks, s[currentStart:])
    return chunks
}