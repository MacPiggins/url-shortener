package Base63

import "testing"

func TestConvertNumber(t *testing.T) {

	tests := []struct {
		name   string
		number uint8
		want   string
	}{
		{
			name:   "number = 0",
			number: 0,
			want:   "a",
		},
		{
			name:   "number = 1",
			number: 1,
			want:   "b",
		},
		{
			name:   "number = 2",
			number: 2,
			want:   "c",
		},
		{
			name:   "number = 60",
			number: 60,
			want:   "8",
		},
		{
			name:   "number = 61",
			number: 61,
			want:   "9",
		},
		{
			name:   "number = 62",
			number: 62,
			want:   "_",
		},
		{
			name:   "number = 124",
			number: 124,
			want:   "b9",
		},
		{
			name:   "number = 255",
			number: 255,
			want:   "ed",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if res := convertNumber(test.number); res != test.want {
				t.Errorf("convertNumber() = %v, want %v", res, test.want)
			}
		})
	}

}

func TestConvertToBase63(t *testing.T) {

	tests := []struct {
		name  string
		input []byte
		want  string
	}{
		{
			name:  "input = []byte{0}",
			input: []byte{0},
			want:  "a",
		},
		{
			name:  "input = []byte{1}",
			input: []byte{1},
			want:  "b",
		},
		{
			name:  "input = []byte{2}",
			input: []byte{2},
			want:  "c",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if res := ConvertToBase63(test.input); res != test.want {
				t.Errorf("ConvertToBase63() = %v, want %v", res, test.want)
			}
		})
	}
}
