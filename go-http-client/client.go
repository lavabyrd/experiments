package main

import (
	"log"
	"io/ioutil"
	"net/http"
	"os"

  "github.com/joho/godotenv"
)

func main() {

  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  DATAHUB_AVA := os.Getenv("DH_AVA_API_KEY")

	url := "https://avalanche--mainnet--indexer.datahub.figment.io/apikey/" + DATAHUB_AVA + "/"
	urlupdated := url + os.Args[1]
	resp, err := http.Get(urlupdated)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.Status == "404 Not Found" {
		log.Println("Failed url endpoint is " + os.Args[1])
	} else {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			log.Println("Failed url endpoint is " + os.Args[1])
		}
		//Convert the body to type string
		sb := string(body)
		log.Printf(sb)
	}

}
