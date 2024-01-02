package sequence

import (
	"bufio"
	"bytes"
	"fmt"
	"math/big"
	"strings"
)

// Parse a .seq formatted records string into a new SequenceRecord.
// All valid lines have a format of '%{code} {seq_id} {content}'
// with {code} being a single character code, {seq_id} being the
// identity of the sequence and {content} being a single line of text.
func (rec *SequenceRecord) UnmarshalText(text []byte) error {
	// first character of actual content begins after '%C A000000 '
	const pre = 11

	// for each line in the string, assign the appropriate fields
	// to the SequenceRecord
	s := string(text)
	for _, line := range strings.Split(s, "\n") {

		// Skip over lines with no actual content
		if len(line) <= pre || line[0] != '%' {
			continue
		}

		// Use the space positions as sentinels
		if line[2] != ' ' || line[10] != ' ' {
			return fmt.Errorf("format: %s", line)
		}

		// Assign lines based on the code
		code := line[1]
		switch code {
		case 'I':
			// %I A000001
			rec.Identity = line[3 : pre-1]

			// Aliases and other info
			rec.IdentityPlus = line[pre:]
		case 'N':
			rec.Name = line[pre:]
		case 'A':
			rec.Author = line[pre:]
		case 'O':
			values := strings.Split(line[pre:], ",")
			if len(values) != 2 {
				return fmt.Errorf("invalid offset: %s", line[pre:])
			}
			if _, ok := (&rec.Offset.InitialValue).SetString(values[0], 10); !ok {
				return fmt.Errorf("invalid initial value: %s", values[0])
			}
			if _, ok := (&rec.Offset.FirstGreaterThanOne).SetString(values[1], 10); !ok {
				return fmt.Errorf("invalid first greater than one: %s", values[1])
			}
		case 'S', 'T', 'U':
			values := strings.Split(strings.TrimRight(line[pre:], ","), ",")
			bigs := make([]big.Int, len(values))
			for i, v := range values {
				if _, ok := (&bigs[i]).SetString(v, 10); !ok {
					return fmt.Errorf("invalid sequence value: %s", v)
				}
			}
			rec.Sequence = append(rec.Sequence, bigs...)
			idx := line[1] - 'S'
			rec.STUCounts[idx] = len(bigs)
		case 'K':
			values := strings.Split(line[pre:], ",")
			keywords := make([]Keyword, len(values))
			for i, value := range values {
				if keyword, err := KeywordString(value); err != nil {
					return fmt.Errorf("invalid keyword: %s", value)
				} else {
					keywords[i] = keyword
				}
			}
			rec.Keywords = append(rec.Keywords, keywords...)
		case 'D', 'H', 'F', 'Y', 'E', 'e', 'p', 't', 'o', 'C':
			field := rec.GetStringListField(code)
			if field != nil {
				*field = append(*field, line[pre:])
			} else {
				return fmt.Errorf("invalid field code: %c", code)
			}
		}

		// Some bookkeeping for validation and stringification
		if len(rec.FieldOrder) == 0 {
			rec.FieldOrder = append(rec.FieldOrder, line[1])
		} else if rec.FieldOrder[len(rec.FieldOrder)-1] != line[1] {
			rec.FieldOrder = append(rec.FieldOrder, line[1])
		}
		if rec.LineCounts == nil {
			rec.LineCounts = make(map[byte]int)
		}
		rec.LineCounts[line[1]]++
	}

	return nil
}
func (rec *SequenceRecord) MarshalText() ([]byte, error) {
	buf := new(bytes.Buffer)
	out := bufio.NewWriter(buf)

	for _, code := range rec.FieldOrder {
		switch code {
		case 'I':
			fmt.Fprintf(out, "%%I %s", rec.Identity)
			if rec.IdentityPlus != "" {
				fmt.Fprintf(out, " %s", rec.IdentityPlus)
			}
			out.WriteByte('\n')
		case 'N':
			fmt.Fprintf(out, "%%N %s %s\n", rec.Identity, rec.Name)
		case 'A':
			fmt.Fprintf(out, "%%A %s %s\n", rec.Identity, rec.Author)
		case 'O':
			fmt.Fprintf(out, "%%O %s %s,%s\n", rec.Identity, rec.Offset.InitialValue.String(), rec.Offset.FirstGreaterThanOne.String())
		case 'S', 'T', 'U':
			line_offset := int(code - 'S')
			begin := 0
			end := rec.STUCounts[line_offset]
			for i := 0; i < line_offset; i++ {
				begin += rec.STUCounts[i]
				end += rec.STUCounts[i]
			}
			fmt.Fprintf(out, "%%%c %s ", code, rec.Identity)
			for i, value := range rec.Sequence[begin:end] {
				if code == 'U' && i == rec.STUCounts[line_offset]-1 {
					fmt.Fprintf(out, "%s", value.String())
				} else {
					fmt.Fprintf(out, "%s,", value.String())
				}
			}
			fmt.Fprintln(out)
		case 'K':
			fmt.Fprintf(out, "%%K %s %s\n", rec.Identity, rec.KeywordsString())
		case 'D', 'H', 'F', 'Y', 'E', 'e', 'p', 't', 'o', 'C':
			field := rec.GetStringListField(code)
			if field != nil {
				for _, value := range *field {
					fmt.Fprintf(out, "%%%c %s %s\n", code, rec.Identity, value)
				}
			} else {
				return nil, fmt.Errorf("invalid field code: %c", code)
			}
		}
	}

	out.Flush()
	return buf.Bytes(), nil
}
