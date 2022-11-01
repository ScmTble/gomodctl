package parse

import (
	"testing"
)

func TestMustParse(t *testing.T) {
	MustParse()
	if ModFile == nil {
		t.Fail()
	}
}
