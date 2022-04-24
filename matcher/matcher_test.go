package matcher

import (
	"testing"
)

func TestMatcher(t *testing.T) {
	type args struct {
		prefixDataStore []string
		inputString     string
	}
	tests := []struct {
		name string
		args args
		want MatchResult
	}{
		{
			"valid",
			args{prefixDataStore: []string{""}, inputString: ""},
			MatchResult{Prefix: ""},
		},
		{
			"valid",
			args{prefixDataStore: []string{"abc", "def", "fgh"}, inputString: "abcmnekjnwldnewnejwkdn"},
			MatchResult{Prefix: "abc"},
		},
		{
			"invalid",
			args{prefixDataStore: []string{"abc", "def", "fgh", "aads"}, inputString: "bbcmnekjnwldnewnejwkdn"},
			MatchResult{Prefix: ""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Matcher(tt.args.prefixDataStore, tt.args.inputString); got.Prefix != tt.want.Prefix {
				t.Errorf("Matcher() = %v, want %v", got, tt.want)
			}
		})
	}
}
