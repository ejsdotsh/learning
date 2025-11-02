package secret

func Handshake(code uint) []string {
	var res []string

	if code&1 == 1 {
		res = append(res, "wink")
	}

	if code&2 == 2 {
		res = append(res, "double blink")
	}

	if code&4 == 4 {
		res = append(res, "close your eyes")
	}

	if code&8 == 8 {
		res = append(res, "jump")
	}

	if code&16 == 16 {
		for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
			res[i], res[j] = res[j], res[i]
		}
	}

	return res
}
