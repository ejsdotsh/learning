// bob returns some limited responses that could be attributable to a teenager.
package bob

import (
	"strings"
	"unicode"
)

// Remark is used as a convenience type for matching what kind of remark it is.
type Remark struct {
	remark string
}

// isSilence returns true if the remark is empty.
func (r Remark) isSilence() bool {
	return r.remark == ""
}

// isYelling returns true only if there are letters and they are all uppercase.
func (r Remark) isYelling() bool {
	letters := strings.IndexFunc(r.remark, unicode.IsLetter) >= 0
	upper := strings.ToUpper(r.remark) == r.remark

	return letters && upper
}

// isExasperated returns true if isYelling and isQuestion are both true.
func (r Remark) isExasperated() bool {
	return r.isYelling() && r.isQuestion()
}

// isQuestion returns true if the remark ends with a '?'.
func (r Remark) isQuestion() bool {
	return strings.HasSuffix(r.remark, "?")
}

// newRemark returns a Remark with all leading and trailing whitespace removed.
func newRemark(remark string) Remark {
	return Remark{strings.TrimSpace(remark)}
}

// Hey returns Bob's limited responses.
func Hey(remark string) string {
	r := newRemark(remark)
	switch {
	case r.isExasperated():
		return "Calm down, I know what I'm doing!"
	case r.isYelling():
		return "Whoa, chill out!"
	case r.isQuestion():
		return "Sure."
	case r.isSilence():
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}
