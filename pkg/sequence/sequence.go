// Package sequence provides types and a parser for the OEIS sequence database.
//
// The OEIS sequence database is a collection of integer sequences. The internal data format is
// textual and described at https://oeis.org/eishelp1.html.
package sequence

import (
	"bufio"
	"bytes"
	"fmt"
	"math/big"
)

type Sequence interface {
	// Aliases returns alternative ids of sequences like M0082 or
	GetAliases() []string

	// Concatenate STU terms as a string of ascii integers
	TermsString() string

	// Compose the entire record in the original format
	RecordString() string

	// Get the address of an internal []string type field for the given code
	GetStringListField(code byte) *[]string
}

// SequenceRecord represents the textual format of an OEIS sequence record.
//
// A quick summary of the fields is given below.
//
// %I A000001 Identification line (required)
// %S A000001 First line of sequence (required)
// %T A000001 2nd line of sequence.
// %U A000001 3rd line of sequence.
// %N A000001 Name (required)
// %D A000001 Detailed reference line.
// %D A000001 Detailed references (2).
// %H A000001 Link to other site.
// %H A000001 Link to other site (2).
// %F A000001 Formula.
// %F A000001 Formula (2).
// %Y A000001 Cross-references to other sequences.
// %A A000001 Author (required)
// %O A000001 Offset (required)
// %E A000001 Extensions, errors, etc.
// %e A000001 examples to illustrate initial terms.
// %p A000001 Maple program.
// %t A000001 Mathematica program.
// %o A000001 Program in another language.
// %K A000001 Keywords (required)
// %C A000001 Comments.
type SequenceRecord struct {

	// Keep track of parsed line counts for validation purposes
	LineCounts map[byte]int
	FieldOrder []byte
	STUCounts  [3]int

	// Required fields

	// %I, A000001
	Identity string

	// Additional identity line content is separate, this includes aliases, the latest revision
	// mumber and the associated timestamp
	IdentityPlus string

	// %N, Brief descriptive name for the sequence
	Name string

	// %A, Author(s) of the sequence
	Author string

	// %O, Offset of the sequence
	Offset SequenceOffset

	// %S, %T, %U, Comma separated list of sequence values, e.g. Concat(%S, %T, %U)
	Sequence []big.Int

	// %K, Comma separated list of keywords
	Keywords []Keyword

	// Optional fields

	// %D, Detailed reference line
	DetailedReferences []string

	// %H, Links to other sites
	Links []string

	// %F, Formula (if not included in the Name)
	Formulae []string

	// %Y, Cross-references to other sequences
	CrossReferences []string

	// %E, Extensions, errors, etc.
	Errata []string

	// %e, Examples to illustrate initial terms.
	Examples []string

	// %p, Maple program.
	MapleProgram []string

	// %t, Mathematica program.
	MathematicaProgram []string

	// %0, Program in another language.
	OtherProgram []string

	// %C, Comments.
	Comments []string
}

// SequenceOffset represents offset characteristics of a sequence.
type SequenceOffset struct {

	// subscript of first term (the first valid input to the sequence)
	// for example 0 for sequences covering all non-negative integers,
	// 1 for sequences covering all positive integers, and -1 for
	InitialValue big.Int

	// position of first entry greater than or equal to 2 in magnitude
	// (or 1 if no such entry exists)
	FirstGreaterThanOne big.Int
}

// SequenceData is a modernized representation of a sequence record with greater parity with
// common tree-like data structures.
type SequenceData struct {

	// ID is the OEIS sequence id, e.g. A000001
	ID int

	// Aliases are alternative ids of sequences, Mnnnn or Nnnnn from other publications
	Aliases []string

	// Revision is the revision number of the sequence
	Revision int

	// Created is the timestamp of the creation of the sequence
	CreatedAt string

	// LastModified is the timestamp of the last modification of the sequence
	LastModifiedAt string

	// Name is the brief descriptive name for the sequence
	Name string
}

// Author is the author of a sequence definition
type Author struct {
	// Name is the name of the author
	Name string
}

type Comment struct {
}

type CrossReference struct {
}

func (seq *SequenceRecord) KeywordsString() string {
	buf := new(bytes.Buffer)
	out := bufio.NewWriter(buf)

	for i, value := range seq.Keywords {
		if i == len(seq.Keywords)-1 {
			fmt.Fprintf(out, "%s", value.String())
		} else {
			fmt.Fprintf(out, "%s,", value.String())
		}
	}
	out.Flush()
	return buf.String()
}

func (seq *SequenceRecord) TermsString() string {
	buf := new(bytes.Buffer)
	out := bufio.NewWriter(buf)

	for i, value := range seq.Sequence {
		if i == len(seq.Sequence)-1 {
			fmt.Fprintf(out, "%s", value.String())
		} else {
			fmt.Fprintf(out, "%s,", value.String())
		}
	}
	out.Flush()
	return buf.String()
}

func (seq *SequenceRecord) GetStringListField(code byte) *[]string {
	switch code {
	case 'D':
		return &seq.DetailedReferences
	case 'H':
		return &seq.Links
	case 'F':
		return &seq.Formulae
	case 'Y':
		return &seq.CrossReferences
	case 'E':
		return &seq.Errata
	case 'e':
		return &seq.Examples
	case 'p':
		return &seq.MapleProgram
	case 't':
		return &seq.MathematicaProgram
	case 'o':
		return &seq.OtherProgram
	case 'C':
		return &seq.Comments
	}
	return nil
}
