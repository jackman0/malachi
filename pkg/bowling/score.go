package bowling

type Score []Frame

func NewScore() Score {
	return make(Score, 0, 10)
}

func (s Score) AddFrame(frame Frame) Score {
	if len(s) == 10 {
		// You can't add a frame to full score card.
		return s
	}

	if len(s) != 9 && (frame.Fill != 0 || frame.Sum()>10) {
		// Don't permit a fill frame on anything but the tenth frame.
		return s
	}

	return append(s, frame)
}

func (s Score) Pins() *int {
	// The return value is nullable because it's possible to *not* be able to
	// print a score, when the last from contains a strike or spare.

	if len(s) == 0 {
		return nil
	}

	var sum int
	for i := range s {
		pins := s.FramePins(i)
		if pins == nil {
			return nil
		}
		sum += *pins
	}

	return &sum
}

func (s Score) FramePins(i int) *int {
	// The return value is nullable because it's possible to *not* be able to
	// print a score.

	// z is the last index of the score.
	z := len(s) - 1

	f0 := s[i]

	// The most basic case, a non-strike, non-spare frame
	if !f0.IsSpare() && !f0.IsStrike() {
		pins := f0.Top + f0.Btm
		return &pins
	}

	// The spare.
	if !f0.IsStrike() {
		pins := 10

		// What's the next roll?
		// If this is the final frame, we should have a fill.

		if i==9{
			pins += f0.Fill
		}else{
			// We need the next frame to get the pins.
			if i == z {
				// If there is no next frame, return nil (not calculable).
				return nil
			}
			f1 := s[i+1]
			pins += f1.Top
		}

		return &pins
	}

	// Final case, the strike.
	pins := 10

	// We need the next two rolls.
	// If we're in the tenth frame, we will have a Btm and a Fill.
	if i == 9{
		pins += f0.Btm + f0.Fill
	}else{
		// Just so you know, I'm not proud of this nesting.
		// I'm sure we could write some helper functions to get this done more cleanly, but it's getting late.

		// We need at least the next frame to get the pins.

		if i == z {
			// If there is no next frame, return nil (not calculable).
			return nil
		}
		f1 := s[i+1]
		if f1.IsStrike() {
			pins += 10

			// There was only one throw in the next frame.
			// Get the second next frame.

			if i+1 == z {
				// There is no second next frame.
				return nil
			}

			f2 := s[i+2]

			pins += f2.Top
		} else {
			// An open frame or a spare both contain two throws.
			// Use whatever is in the next frame.
			pins += f1.Top + f1.Btm
		}

	}


	return &pins
}
