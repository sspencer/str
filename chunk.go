package str

// ChunkString chunks string into equal size parts,
// starting from left for positive (n),
// or right for negative (n).
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
