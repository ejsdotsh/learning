// Package rotationalcipher implements a basic rotation cipher.
package rotationalcipher

func RotationalCipher(plain string, shiftKey int) string {
	// if there is no shift, return the string
	if shiftKey == 0 {
		return plain
	}

	result := make([]byte, len(plain))

	for i := 0; i < len(plain); i++ {
		ch := plain[i]

		switch {

		// uppercase letters
		case 'A' <= ch && ch <= 'Z':
			result[i] = 'A' + byte((int(ch-'A')+shiftKey)%26)

		// lowercase letters
		case 'a' <= ch && ch <= 'z':
			result[i] = 'a' + byte((int(ch-'a')+shiftKey)%26)

		// non-letters
		default:
			result[i] = ch
		}
	}

	return string(result)
}
