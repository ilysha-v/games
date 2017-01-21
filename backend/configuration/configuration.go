package configuration

import "os"

// GetDatabaseHost - get db host
func GetDatabaseHost() string {
	value := os.Getenv("MONGOHOST")
	if value == "" {
		return "localhost"
	}

	return value
}
