package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/1Password/connect-sdk-go/connect"
	"github.com/1Password/connect-sdk-go/onepassword"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print(".env file does not exist, will get the variables from the environment")
	}
}

func main() {
	var err error

	var argHelp string = "Please run following command: go run main.go -vault=\"vault_name\" -item=\"item_name\" -label=\"label_name\""

	vaultARG := flag.String("vault", "", argHelp)
	itemARG := flag.String("item", "", argHelp)
	labelARG := flag.String("label", "", argHelp)
	flag.Parse()

	if *vaultARG == "" || *itemARG == "" || *labelARG == "" {
		fmt.Println(argHelp)
		return
	}

	var onepassURL string = ""
	onepassURLValue, onepassURLPresent := os.LookupEnv("1PASS_URL")
	if onepassURLPresent {
		onepassURL = onepassURLValue
	} else {
		panic("Missing ENV Variable 1PASS_URL")
	}

	var onepassToken string = ""
	onepassTokenValue, onepassTokenPresent := os.LookupEnv("1PASS_TOKEN")
	if onepassTokenPresent {
		onepassToken = onepassTokenValue
	} else {
		panic("Missing ENV Variable 1PASS_TOKEN")
	}

	client := connect.NewClient(onepassURL, onepassToken)

	vaults, err := client.GetVaults()
	if err != nil {
		panic(err)
	}

	var vault string = ""
	for _, v := range vaults {
		if v.Name == *vaultARG {
			vault = v.ID
			break
		}
	}

	if vault == "" {
		panic("Vault not found")
	}

	var item *onepassword.Item
	item, err = client.GetItemByTitle(*itemARG, vault)
	if err != nil {
		panic(err)
	}

	for _, f := range item.Fields {
		if f.Label == *labelARG {
			fmt.Println(f.Value)
		}
	}
}
