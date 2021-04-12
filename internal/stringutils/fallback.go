package stringutils

// OrFallback returns the string if it's not empty, or the fallback.
func OrFallback(s, fallback string) string {
	if s != "" {
		return s
	}
	return fallback
}

// PointerOrFallback returns the string if it's not nil nor empty, or the fallback.
func PointerOrFallback(s *string, fallback string) string {
	if s == nil {
		return fallback
	}

	return OrFallback(*s, fallback)
}
