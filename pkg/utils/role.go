package utils

const (
	ADMIN    = "ADMIN"
	STANDARD = "STANDARD"
)

func CheckRole(role string) bool {
	roleArr := [...]string{ADMIN, STANDARD}

	for v := range roleArr {
		if roleArr[v] == role {
			return true
		}
	}
	return false
}
