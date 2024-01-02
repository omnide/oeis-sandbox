package sequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeyword_CodeGen(t *testing.T) {

	// Ensures that the code generator is working correctly, currently 31 keywords
	assert.Equal(t, 31, len(KeywordValues()))

	// Strings are lowercase lowercase
	assert.Equal(t, KeywordBase.String(), "base")

	// Lookup by string matches value
	{
		val, err := KeywordString("base")
		assert.NoError(t, err)
		assert.Equal(t, val, KeywordBase)
	}

	// Lookup is case insensitive
	{
		val, err := KeywordString("BAse")
		assert.NoError(t, err)
		assert.Equal(t, val, KeywordBase)
	}
}
