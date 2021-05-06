package regions

import (
	"testing"
)

func TestRemoveMetaData(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"Nyíregyháza", "Nyíregyháza"},
		{"Rehiyon ng Calabarzon", "Calabarzon"},
		{"London, City of", "London"},
		{"Il-lvant tal-Ingilterra", "Il-lvant tal-Ingilterra"},
		{"Walloon region (Belgium)", "Walloon"},
		{"Kemerovskaja oblast'", "Kemerovskaja"},
		{"Virgin Islands, U.S.", "Virgin Islands"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			actual := removeMetaData(tt.input)
			if actual != tt.expected {
				t.Errorf("got %q, want %q", actual, tt.expected)
			}
		})
	}
}
