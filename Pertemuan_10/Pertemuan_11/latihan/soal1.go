package latihan

func IsAllowWork(gender string, age int) bool {
	result := false
	if gender == "Pria" {
		if age > 17 && age < 60 {
			result = true
		}
	} else if gender == "Wanita" {
		if age > 19 && age < 60 {
			result = true
		}
	}

	return result
}
