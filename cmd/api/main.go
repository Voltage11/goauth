package main

import (
	"fmt"
	"goauth/internal/config"
)

func main() {
	// Загрузка конфигурации
	cfg := config.GetConfig()
	fmt.Println(cfg)
}
