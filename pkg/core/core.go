package core

import (
	"github.com/joho/godotenv"
)

func main() {
	
}

func loadEnviromentVariables(envFilePath string) {
	godotenv.Load(envFilePath)	
}