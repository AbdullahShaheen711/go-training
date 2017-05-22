package exercises

import (
	"strings"
)

type DNA struct {
	Strand string
	Length int
}

func CalcHamming(s1 *DNA, s2 *DNA) (int, bool) {
	hammingDiff := 0
	s1.Strand = strings.ToUpper(s1.Strand)
	s2.Strand = strings.ToUpper(s2.Strand)
	s1Chars := []rune(s1.Strand)
	s2Chars := []rune(s2.Strand)
	if s1.Length != s2.Length {
		return hammingDiff, false
	} else {
		for i, c1 := range s1Chars {
			if c1 != s2Chars[i] {
				hammingDiff++
			}
		}
		return hammingDiff, true
	}
}

//func main() {
//	var dna1, dna2 *exercises.DNA = &exercises.DNA{"GGACGGATTCTG", len("GGACGGATTCTG")}, &exercises.DNA{"AGGACGGATTCT", len("AGGACGGATTCT")}
//	hammingDiff, isValid := exercises.CalcHamming(dna1, dna2)
//	if !isValid {
//		fmt.Println("Disallowed")
//	} else {
//		fmt.Println("Hamming Difference: ", hammingDiff)
//	}
//}
