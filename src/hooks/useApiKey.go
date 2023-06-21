Go does not have an equivalent to React hooks or local storage, but you can achieve similar functionality using environment variables and a package like "github.com/joho/godotenv". Here's a Go version of the code:

```go
package main

import (
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}
}

func useApiKey() string {
	return os.Getenv("OPENAI_KEY")
}

func main() {
	apiKey := useApiKey()
}
```

This code assumes that you have a ".env" file in your project directory with the following content:

```
OPENAI_KEY=your_api_key_here
```

The "useApiKey" function retrieves the value of the "OPENAI_KEY" environment variable.