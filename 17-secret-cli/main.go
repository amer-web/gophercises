package main

import (
	"flag"
	"fmt"
	"log"
	"secret-cli/secret"
)

func main() {
	flag.Parse()
	nameFlag := flag.Arg(0)
	nameKey := flag.Arg(1)
	if nameFlag == "" || nameKey == "" {
		fmt.Println("u should enter valid")
		return
	}
	fileVault := secret.Vault{}
	switch nameFlag {
	case "set":
		nameValue := flag.Arg(2)
		if nameValue == "" {
			fmt.Println("u should enter value of key")
			return
		}
		fileVault.Set(nameKey, nameValue)
		fmt.Println("set value successfully")
	case "get":
		value, err := fileVault.Get(nameKey)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(value)
	default:
		fmt.Println("u should enter set or get")
	}
}
