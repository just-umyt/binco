package start

func (det *Det) CheckProcess(ind int, ball string) bool {
	if det.CorrectBalls[0][ind] == ball {
		if det.CorrectBalls[1][ind] == ball {
			det.CorrectBalls[1][ind] = ""
			det.EndBalls[1][ind] = ball
			return true
		}
		det.CorrectBalls[0][ind] = ""
		det.EndBalls[0][ind] = ball
		return true
	}

	return false
}
