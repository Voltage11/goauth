package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func getEnvStringValue(paramName string) string {
	out := os.Getenv(paramName)

	if out == "" {
		log.Fatal(fmt.Printf("Отсутствует значение параметра окружения: %s", paramName))
	}

	return out
}

func getEnvStringValueWithDefault(paramName, valueDefault string) string {
	out := os.Getenv(paramName)

	if out == "" {
		return valueDefault
	}

	return out
}

func getEnvIntValue(paramName string) int {
	paramValueStr := getEnvStringValue(paramName)

	out, err := strconv.Atoi(paramValueStr)

	if err != nil {
		log.Fatal("Ошибка загрузки параметра окружения: ", paramName, "(", err, ")")
	}

	return out
}

func getEnvIntValueWithDefault(paramName string, valueDefault int) int {
	paramValueStr := getEnvStringValueWithDefault(paramName, "")

	if paramValueStr == "" {
		return valueDefault
	}

	out, err := strconv.Atoi(paramValueStr)

	if err != nil {
		return valueDefault
	}

	return out
}
