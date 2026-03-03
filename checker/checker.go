package checker

const Weak = "СЛАБЫЙ ПАРОЛЬ"
const Medium = "СРЕДНИЙ ПАРОЛЬ"
const Strong = "СИЛЬНЫЙ ПАРОЛЬ"

func Analyze(password string) (string, []string) {
	runes := []rune(password)

	var hasDigit bool
	var hasUpper bool
	var hasLower bool
	var hasSpecial bool

	for _, r := range runes {
		switch {
		case r >= '0' && r <= '9':
			hasDigit = true
		case r >= 'A' && r <= 'Z':
			hasUpper = true
		case r >= 'a' && r <= 'z':
			hasLower = true
		default:
			hasSpecial = true
		}
	}

	length := len(runes)
	missing := make([]string, 0)

	if !hasDigit {
		missing = append(missing, "Добавьте цифры")
	}
	if !hasUpper {
		missing = append(missing, "Добавьте заглавные буквы")
	}
	if !hasLower {
		missing = append(missing, "Добавьте строчные буквы")
	}
	if !hasSpecial {
		missing = append(missing, "Добавьте спецсимволы")
	}
	if length < 8 {
		missing = append(missing, "Пароль слишком короткий")
	}

	if length >= 8 && hasDigit && hasUpper && hasLower && hasSpecial {
		return Strong, missing
	}

	if length >= 6 && hasDigit && (hasLower || hasUpper) {
		return Medium, missing
	}

	return Weak, missing
}
