package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  string `json:"result"`
}

func main() {
	responseObject := Response{}
	argsWithProg := os.Args[1]
	url := "https://api-testnet.polygonscan.com/api?module=stats&action=tokensupply&contractaddress=" + argsWithProg
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	jsonErr := json.Unmarshal(body, &responseObject)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println("Number of tokens:", responseObject.Result)
}

/* Due to lack of information regarding the api and where ot get the information's regarding the api and how to access specific information described in the test i just did this little part
i could do more if i had the endpoints or the api documentation. */

/* I have experience with almost every technology listed on the test description, except things involving blockchain and NATS */
