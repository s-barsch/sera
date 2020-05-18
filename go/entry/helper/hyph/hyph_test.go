package hyph

import (
	"testing"
)

func TestHyphenateWord(t *testing.T) {
	h, err := LoadPattern("./hyph-de.dic")	
	if err != nil {
		t.Error(err)
	}
	str := h.HyphenateWord("Schonheit")
	t.Log(str)
	str = h.HyphenateWord("Schönheit")
	t.Log(str)
}
