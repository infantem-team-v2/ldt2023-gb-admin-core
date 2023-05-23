package etc

import "strconv"

func MustParseToInt(a string) (i int) {
	i, _ = strconv.Atoi(a)
	return i
}
