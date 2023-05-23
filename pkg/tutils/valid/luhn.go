package valid

import "errors"

// Luhn is an implementation of Luhn algorithm.
// For more information see https://en.wikipedia.org/wiki/Luhn_algorithm
func Luhn(s string) error {
	if len(s) == 0 {
		return errors.New("empty string")
	}

	var sum int
	var alter bool
	for i := len(s) - 1; i >= 0; i-- {
		d := int(s[i] - '0')
		if alter {
			d *= 2
		}
		sum += d / 10
		sum += d % 10
		alter = !alter
	}

	if sum%10 != 0 {
		return errors.New("invalid checksum")
	}
	return nil
}
