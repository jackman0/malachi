package bowling

import (
	"strconv"
	"strings"
)

type Frame struct {
	Top  int
	Btm  int
	Fill int
}

func (f *Frame) IsSpare() bool {
	if f.Top == 10 {
		return false
	}

	// This may look kind of funny.
	// I've gotten in the habit of making sure that a series of checks like this
	// always be negative or positive for the sake of being able to easily scan
	// the list and spot problems.
	if !(f.Top+f.Btm == 10) {
		return false
	}

	return true
}
func (f *Frame) Sum() int {
	return f.Btm + f.Top + f.Fill
}

func (f *Frame) IsStrike() bool {
	if f.Top == 10 {
		return true
	}
	return false
}

func ParseFrame(s string) *Frame {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)

	if s == "x" || s == "10"{
		// This should be the only case (a strike) in which no commas are present.
		return &Frame{
			Top:  10,
			Btm:  0,
			Fill: 0,
		}
	}

	chunks := strings.Split(s, ",")
	if len(chunks) < 2 || len(chunks) > 3 {
		return nil
	}

	var (
		top  int
		btm  int
		fill int
		err  error
	)

	// TOP
	if chunks[0] == "x" {
		// The top should never be a spare.
		top = 10
	} else {
		top, err = strconv.Atoi(chunks[0])
		if err != nil {
			return nil
		}
	}
	if top < 0 || top > 10 {
		return nil
	}

	// BTM
	if chunks[1] == "/" {
		btm = 10 - top
	} else if chunks[1] == "x" {
		// I'm not sure if this is necessary. I just want to be robust.
		btm = 10
	} else {
		btm, err = strconv.Atoi(chunks[1])
		if err != nil {
			return nil
		}
	}
	if btm < 0 || btm > 10 {
		return nil
	}

	// FILL
	if len(chunks) == 3 {
		if chunks[2] == "/" {
			fill = 10 - btm
		} else if chunks[2] == "x" {
			// I'm not sure if this is necessary. I just want to be robust.
			fill = 10
		} else {
			fill, err = strconv.Atoi(chunks[2])
			if err != nil {
				return nil
			}
		}
		if fill < 0 || fill > 10 {
			return nil
		}
	}

	return &Frame{
		Top:  top,
		Btm:  btm,
		Fill: fill,
	}
}
