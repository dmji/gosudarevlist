package env

import (
	"github.com/joho/godotenv"
)

func LoadEnv(deep int, canPanic bool) {
	path := ".env"
	for i := range deep {
		if i != 0 {
			path = "../" + path
		}
		err := godotenv.Load(path)
		if err == nil {
			return
		}
	}

	if canPanic {
		panic(".env not found")
	}
}
