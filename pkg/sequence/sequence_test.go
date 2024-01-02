package sequence

import (
	"testing"

	testdata "github.com/peteole/testdata-loader"
	"github.com/stretchr/testify/assert"
)

func TestSequenceRecord_String(t *testing.T) {
	bytes := testdata.GetTestFile("pkg/sequence/sequence_testdata/A007318.seq")
	seq := SequenceRecord{}
	err := seq.UnmarshalText(bytes)
	assert.NoError(t, err)
	assert.Equal(t, 78, len(seq.Sequence), "A007318 STU lines should have 78 terms")

	// We must be able to restore the exact record text as it was parsed from a sequence file
	{
		text, err := seq.MarshalText()
		assert.NoError(t, err)
		assert.Equal(t, string(bytes), string(text))
	}
}
