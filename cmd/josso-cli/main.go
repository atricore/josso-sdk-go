package main

import (
	"flag"
	"fmt"
	"os"

	cli "github.com/atricore/josso-sdk-go"
)

func main() {
	os.Exit(mainWithCode())
}

func mainWithCode() int {

	s, err := cli.GetServerConfigFromEnv()
	if err != nil {
		fmt.Println(err)
		return 1
	}

	cli, err := cli.CreateClient(s)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	cli.Authn()

	flag.String("word", "foo", "a string")

	return 0

}
