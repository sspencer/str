package str

import (
	"math"
	"strconv"
	"strings"
)

// ChunkString chops string into as many equal size parts as possible,
// starting from left for positive (n), or right for negative (n).
//
// e.g. ChunkString("1234", 3)  -> ["123", "4"]
// e.g. Chunkstring("1234", -3) -> ["1", "234"]
func ChunkString(buf string, lim int) []string {
	if lim == 0 {
		return []string{buf}
	}

	sn := len(buf)
	var m, b, e int
	var parts []string
	if lim < 0 {
		lim = 0 - lim
		m = sn % lim
	}

	if sn%lim == 0 {
		parts = make([]string, sn/lim)
	} else {
		parts = make([]string, sn/lim+1)
	}

	for i := range parts {
		if m > 0 {
			e = m
			m = 0
		} else {
			b = e
			e = b + lim
		}

		if e > sn {
			e = sn
		}

		parts[i] = buf[b:e]
	}

	return parts
}

// Comma creates a human readable integer by adding commas
// for thousands separators.
//
// e.g. Comma(1234567) -> 1,234,567
func Comma(n int64) string {
	sign := ""

	// only special case
	if n == math.MinInt64 {
		return "-9,223,372,036,854,775,808"
	}

	if n < 0 {
		sign = "-"
		n = 0 - n
	}

	parts := ChunkString(strconv.FormatInt(n, 10), -3)
	return sign + strings.Join(parts, ",")
}
