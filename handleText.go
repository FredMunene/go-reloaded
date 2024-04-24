package reloaded

import "os"

func ReadTextFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func WriteTextFile(filename string, content string) error {
	err := os.WriteFile(filename, []byte(content), 0o644)
	return err
}
