package initiate

import "os"

// Env Initializes the enviroment variable
// The default is set to develop
func Env() (string, error) {

	env := os.Getenv("GO_ENV")

	if env == "" {
		err := os.Setenv("GO_ENV", "development")
		if err != nil {
			return "", err
		}
		env = os.Getenv("GO_ENV")
	}

	return env, nil
}
