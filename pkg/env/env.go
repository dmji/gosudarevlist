package env

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv(name string, deep int) error {
	path := name
	for i := range deep {
		if i != 0 {
			path = "../" + path
		}
		err := godotenv.Load(path)
		if err == nil {
			return nil
		}
	}

	return errors.New("specified file not found")
}

func Check(names ...string) func() ([]bool, []string) {
	checkNames := append([]string{}, names...)
	return func() ([]bool, []string) {
		plist := make([]bool, 0)
		nlist := make([]string, 0)
		for _, name := range checkNames {
			_, bOk := os.LookupEnv(name)
			plist = append(plist, bOk)
			if !bOk {
				nlist = append(nlist, name)
			}
		}
		return plist, nlist
	}
}
