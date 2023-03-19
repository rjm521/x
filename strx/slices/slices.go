package slices

// Equal reports whether two slices are equal: the same length and all
// elements equal. If the lengths are different, Equal returns false.
// Otherwise, the elements are compared in index order, and the
// comparison stops at the first unequal pair.
func Equal(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, n := range s1 {
		if n != s2[i] {
			return false
		}
	}
	return true
}

// Filter appends to d each element e of s for which keep(e) returns true.
// It returns the modified d. d may be s[:0], in which case the kept
// elements will be stored in the same slice.
// if the slices overlap in some other way, the results are unspecified.
// To create a new slice with the filtered results, pass nil for d.
func Filter(d, s []string, keep func(string) bool) []string {
	for _, v := range s {
		if keep(v) {
			d = append(d, v)
		}
	}
	return d
}

// Contains reports whether str is present in strs.
func Contains(strs []string, str string) bool {
	return Index(strs, str) >= 0
}

// Index returns the index of the first occurrence of str in strs, or -1 if
// not present
func Index(strs []string, str string) int {
	for i, v := range strs {
		if v == str {
			return i
		}
	}
	return -1
}

// Clone returns a new clone of strs.
func Clone(strs []string) []string {
	if strs == nil {
		return nil
	}
	c := make([]string, len(strs))
	copy(c, strs)
	return c
}
