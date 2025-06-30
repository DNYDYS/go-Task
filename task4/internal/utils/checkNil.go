package util

func IsNilOrEmptyString(s *string) bool {
	return s == nil || *s == ""
}
