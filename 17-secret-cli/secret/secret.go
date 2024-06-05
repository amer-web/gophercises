package secret

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"secret-cli/encrypt"
	"strings"
)

var keyCode = "amer"

type Vault struct {
	encodingKey string
	filePath    string
}

func FileVault() *Vault {
	return &Vault{}
}
func (v *Vault) Set(keyName, keyValue string) error {
	data := keyName + "::" + keyValue
	hexData, err := encrypt.Encrypt(keyCode, data)
	if err != nil {
		fmt.Println("can't encrypt your data")
		return err
	}
	if err := createAndSetData(hexData); err != nil {
		return err
	}
	return nil
}
func (v *Vault) Get(keyName string) (string, error) {
	keysValues, _ := readFile()
	if value, ok := keysValues[keyName]; ok {
		return value, nil
	}
	return "", errors.New("your key is not found")
}

func createAndSetData(data string) error {
	file, err := os.OpenFile("encrypt.enc", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(data + "\n")
	if err != nil {
		return err
	}
	return writer.Flush()
}
func readFile() (map[string]string, error) {
	file, err := os.Open("encrypt.enc")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var keysValues = map[string]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		output, _ := encrypt.Decrypt(keyCode, scanner.Text())
		parts := strings.Split(output, "::")
		if len(parts) == 2 {
			keysValues[parts[0]] = parts[1]
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return keysValues, nil
}
