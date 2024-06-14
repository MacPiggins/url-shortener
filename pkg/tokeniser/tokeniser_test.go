package tokeniser

import (	
	"testing"
)

func TestGenerateToken(t *testing.T) {

	tests := []struct {
		name string
		link string
		want string
	}{
		{
			name: "deterministic",
			link: "abcdefg",
			want: GenerateToken("abcdefg"),
		},		
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if res := GenerateToken(test.link); res != test.want {
				t.Errorf("GenerateToken() = %v, want %v", res, test.want)
			}
		})
	}
}

func TestGenerateTokenLength(t *testing.T) {

	tests := []struct {
		name string
		link string
		want int
	}{
		{
			name: "length",
			link: "abcdefg",
			want: 10,
		},		
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if res := GenerateToken(test.link); len(res) != test.want {
				t.Errorf("GenerateToken() = %v, want %v", len(res), test.want)
			}
		})
	}
}