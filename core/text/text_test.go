package text

import "testing"

func TestCountCharacters(t *testing.T) {
	total, err := CountCharacters("testdata/data.txt")
	if err != nil {
		t.Error("Unexpected error:", err)
	}
	if total != 21 {
		t.Error("Expected 21, got", total)
	}
	_, err = CountCharacters("testdata/no_file.txt")
	if err == nil {
		t.Error("Expected an error")
	}
}
