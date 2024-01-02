package sequence

import (
	"fmt"
	"regexp"
	"slices"
)

type SequenceRecordValidator interface {
	Validate() error

	CheckRequiredFields() error
}

func (s *SequenceRecord) Validate() error {
	if err := s.CheckRequiredFields(); err != nil {
		return err
	}
	return nil
}

func (s *SequenceRecord) CheckRequiredFields() error {
	const required = "INAOSTUK"

	for _, c := range required {
		if !slices.Contains(s.FieldOrder, byte(c)) {
			return fmt.Errorf("missing required field: %c", c)
		}
	}

	// Check identity
	{
		re := regexp.MustCompile(`^A\d{6}$`)
		if !re.MatchString(s.Identity) {
			return fmt.Errorf("invalid identity: %s", s.Identity)
		}
	}

	// Check name
	{
		// Could be any textual description, but must be present
		if s.Name == "" {
			return fmt.Errorf("missing name")
		}
	}

	// Check author
	{
		// Who do we thank?
		// TODO: make some rules around this empirically if possible
		if s.Author == "" {
			return fmt.Errorf("missing author")
		}
	}

	return nil
}
