package piscine

func alpha(a rune) bool {
	if (a >= 'A' && a <= 'Z') || (a >= 'a' && a <= 'z') || (a >= '0' && a <= '9') {
		return true
	}
	return false
}

func Capitalize(s string) string {
	slice := []rune(s)
	vide := true
	for i := 0; i < len(s); i++ {
		if alpha(slice[i]) && vide {
			if slice[i] >= 'a' && slice[i] <= 'z' {
				slice[i] -= 32
			}
			vide = false
		} else if slice[i] >= 'A' && slice[i] <= 'Z' {
			slice[i] += 32
		} else if !alpha(slice[i]) {
			vide = true
		}
	}
	return string(slice)
}
