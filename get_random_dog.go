package main


import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http" 
	"os"
)

func main() {
	response, err := http.Get("https://dog.ceo/api/breeds/image/random")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(responseData);

}