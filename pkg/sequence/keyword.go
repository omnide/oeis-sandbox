package sequence

// Keywords are the keywords used in the OEIS.
type Keyword uint64

//go:generate enumer -type=Keyword -json -text -yaml -transform=lower -trimprefix=Keyword
const (
	// base: dependent on base used for sequence
	KeywordBase Keyword = iota

	// bref: sequence is too short to do any analysis with
	KeywordBref

	// changed: sequence changed in last two weeks (set automatically)
	KeywordChanged

	// cofr: a continued fraction expansion of a number
	KeywordCofr

	// cons: a decimal expansion of a number
	KeywordCons

	// core: an important sequence
	KeywordCore

	// dead: an erroneous sequence
	KeywordDead

	// dumb: an unimportant sequence
	KeywordDumb

	// dupe: duplicate of another sequence
	KeywordDupe

	// easy: it is very easy to produce terms of sequence
	KeywordEasy

	// eigen: an eigensequence: a fixed sequence for some transformation - see the files transforms and transforms (2) for further information.
	KeywordEigen

	// fini: a finite sequence
	KeywordFini

	// frac: numerators or denominators of sequence of rationals
	KeywordFrac

	// full: the full sequence is given
	KeywordFull

	// hard: next term not known, may be hard to find. Would someone please extend this sequence?
	KeywordHard

	// hear: worth listening to
	KeywordHear

	// look: just look at this sequence, interesting graph
	KeywordLook

	// less: reluctantly accepted
	KeywordLess

	// more: more terms are needed! would someone please extend this sequence?
	KeywordMore

	// mult: Multiplicative: a(mn)=a(m)a(n) if g.c.d.(m,n)=1
	KeywordMult

	// new: New (added within last two weeks, roughly)
	KeywordNew

	// nice: an exceptionally nice sequence
	KeywordNice

	// nonn: a sequence of nonnegative numbers
	KeywordNonn

	// obsc: obscure, better description needed
	KeywordObsc

	// probation: included on a provisional basis, but may be deleted later at the discretion of the editor.
	KeywordProbation

	// sign: sequence contains negative numbers
	KeywordSign

	// tabf: An irregular (or funny-shaped) array of numbers made into a sequence by reading it row by row
	KeywordTabf

	// tabl: typically a triangle of numbers, such as Pascal's triangle, made into a sequence by reading it row by row.
	KeywordTabl

	// uned: Not edited, so may contain basic errors
	KeywordUned

	// walk: counts walks (or self-avoiding paths)
	KeywordWalk

	// word: depends on words for the sequence in some language
	KeywordWord
)

// // Function that parses a string into a Keyword.
// func ParseKeyword(s string) Keyword {
// 	return Keyword(s)
// }

// // Funciton that returns a string representation of a Keyword.
// func (k Keyword) String() string {
// 	return string(k)
// }
