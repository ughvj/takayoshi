package types

import (
	"testing"
	"database/sql"
)

func Test_GenerateBulkSentence(t *testing.T) {
	// Given.
	data := QuestionOptions{
		{
			1,
			sql.NullBool{Valid: true, Bool: true},
			sql.NullInt32{Valid: false},
			1,
			1,
		},
		{
			1,
			sql.NullBool{Valid: true, Bool: false},
			sql.NullInt32{Valid: false},
			1,
			2,
		},
		{
			1,
			sql.NullBool{Valid: true, Bool: false},
			sql.NullInt32{Valid: false},
			1,
			3,
		},
		{
			1,
			sql.NullBool{Valid: true, Bool: false},
			sql.NullInt32{Valid: false},
			1,
			4,
		},
	}

	expected := "(true,null,1,1),(false,null,1,2),(false,null,1,3),(false,null,1,4)"

	// When.
	actual := data.GenerateBulkSentence()

	// Then.
	if expected != actual {
		t.Errorf("Strings do not match. Expected: %s, Got: %s", expected, actual)
	}

}
