package helpers

import "os"

// GetEnv retorna o valor da variável de ambiente ou um valor padrão se não estiver definida
func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
