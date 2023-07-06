package utils

func GetHeros() []string {
	return []string{"ironman", "capamerica"}
}

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
