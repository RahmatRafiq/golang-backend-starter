package helpers

import "fmt"

import "os"

func GetEnv(key string, defaultValue any) string {
	value := os.Getenv(key)
	if len(value) == 0 && defaultValue != nil {
		switch defaultValue.(type) {
		case string:
			value = defaultValue.(string)
		case int:
			value = fmt.Sprint(defaultValue.(int))
		}
	}

	return value
}
