package bowling

import (
	"reflect"
	"testing"
)

func TestScore_Pins_LiveAbout(t *testing.T) {
	// These tests come from the following site:
	// https://www.liveabout.com/bowling-scoring-420895
	//
	// The comprehensive nature of the explanation with examples was very convenient.

	score := Score{
		Frame{Top: 10},                 // 1
		Frame{Top: 7, Btm: 3},          // 2
		Frame{Top: 7, Btm: 2},          // 3
		Frame{Top: 9, Btm: 1},          // 4
		Frame{Top: 10},                 // 5
		Frame{Top: 10},                 // 6
		Frame{Top: 10},                 // 7
		Frame{Top: 2, Btm: 3},          // 8
		Frame{Top: 6, Btm: 4},          // 9
		Frame{Top: 7, Btm: 3, Fill: 3}, // 10
	}

	type TestCase struct {
		name string
		s    Score
		want *int
	}
	tests := []TestCase{
		{"1", score[:1], nil},
		{"2", score[:2], nil},
		{"3", score[:3], intPtr(46)},
		{"4", score[:4], nil},
		{"5", score[:5], nil},
		{"6", score[:6], nil},
		{"7", score[:7], nil},
		{"8", score[:8], intPtr(138)},
		{"9", score[:9], nil},
		{"10", score[:10], intPtr(168)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Pins(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pins() = %v, want %v", got, tt.want)
				if got != nil {
					t.Errorf("dereferenced got = %v", *got)
				}
			}
		})
	}
}

func TestScore_FramePins_LiveAbout(t *testing.T) {
	// These tests come from the following site:
	// https://www.liveabout.com/bowling-scoring-420895
	//
	// The comprehensive nature of the explanation with examples was very convenient.

	score := Score{
		Frame{Top: 10},                 // 1
		Frame{Top: 7, Btm: 3},          // 2
		Frame{Top: 7, Btm: 2},          // 3
		Frame{Top: 9, Btm: 1},          // 4
		Frame{Top: 10},                 // 5
		Frame{Top: 10},                 // 6
		Frame{Top: 10},                 // 7
		Frame{Top: 2, Btm: 3},          // 8
		Frame{Top: 6, Btm: 4},          // 9
		Frame{Top: 7, Btm: 3, Fill: 3}, // 10
	}

	type TestCase struct {
		name       string
		frameIndex int
		want       *int
	}

	tests := []TestCase{
		{"1", 0, intPtr(20)},
		{"2", 1, intPtr(17)},
		{"3", 2, intPtr(9)},
		{"4", 3, intPtr(20)},
		{"5", 4, intPtr(30)},
		{"6", 5, intPtr(22)},
		{"7", 6, intPtr(15)},
		{"8", 7, intPtr(5)},
		{"9", 8, intPtr(17)},
		{"10", 9, intPtr(13)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := score.FramePins(tt.frameIndex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FramePins() = %v, want %v", got, tt.want)
			}
		})
	}
}

func intPtr(i int) *int {
	return &i
}
