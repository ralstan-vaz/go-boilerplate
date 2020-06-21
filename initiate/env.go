package initiate

import "os"

// Env Gets the enviroment variable from GO_ENV.
// By default the env will be development
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
