package bowling_test

import (
	"reflect"
	"testing"

	"malachi/pkg/bowling"
)

func TestParseFrame(t *testing.T) {
	type TestCase struct {
		name string
		arg  string
		want *bowling.Frame
	}

	tests := []TestCase{
		{
			name: "strike-X",
			arg:  "X",
			want: &bowling.Frame{
				Top: 10,
				Btm: 0,
			},
		},
		{
			name: "strike-10",
			arg:  "10",
			want: &bowling.Frame{
				Top: 10,
				Btm: 0,
			},
		},
		{
			name: "7-spare",
			arg:  "7,/",
			want: &bowling.Frame{
				Top: 7,
				Btm: 3,
			},
		},
		{
			name: "5-6, too big",
			arg:  "5,6",
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bowling.ParseFrame(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseFrame() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFrame_IsSpare(t *testing.T) {
	type TestCase struct {
		name  string
		frame bowling.Frame
		want  bool
	}
	tests := []TestCase{
		{
			name:  "strike",
			frame: bowling.Frame{Top: 10},
			want:  false,
		},
		{
			name:  "spare",
			frame: bowling.Frame{Top: 3, Btm: 7},
			want:  true,
		},
		{
			name:  "not-spare",
			frame: bowling.Frame{Top: 3, Btm: 3},
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.frame.IsSpare(); got != tt.want {
				t.Errorf("IsSpare() = %v, want %v", got, tt.want)
			}
		})
	}
}
