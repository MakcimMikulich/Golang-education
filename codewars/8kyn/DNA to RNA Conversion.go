package main

func DNAtoRNA(dna string) string {

	rna := []rune(dna)

	for i, value := range dna {
		if value == 'T' {
			rna[i] = 'U'
		}
	}
	return string(rna)
}
