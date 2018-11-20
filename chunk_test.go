package str_test

import (
	"testing"

	"github.com/sspencer/str"
)

func TestChunks(t *testing.T) {
	tt := []struct {
		limit  int
		input  string
		output []string
	}{
		{0, "1234", []string{"1234"}},
		{1, "123", []string{"1", "2", "3"}},
		{-1, "123", []string{"1", "2", "3"}},
		{2, "12345", []string{"12", "34", "5"}},
		{-2, "12345", []string{"1", "23", "45"}},
		{3, "", []string{}},
		{3, "1", []string{"1"}},
		{3, "1234", []string{"123", "4"}},
		{3, "12345678", []string{"123", "456", "78"}},
		{-3, "", []string{}},
		{-3, "1", []string{"1"}},
		{-3, "1234", []string{"1", "234"}},
		{-3, "12345678", []string{"12", "345", "678"}},
	}

	for _, tc := range tt {
		t.Run(tc.input, func(t *testing.T) {
			r := str.ChunkString(tc.input, tc.limit)
			if len(r) != len(tc.output) {
				t.Errorf("For %s (limit=%d), expected %v, not %v", tc.input, tc.limit, tc.output, r)
				return
			}

			for i, c := range r {
				if c != tc.output[i] {
					t.Errorf("For index %d of %s (limit=%d), expected %s, not %s", i, tc.input, tc.limit, tc.output[i], c)
					return
				}
			}
		})
	}
}
