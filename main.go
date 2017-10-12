package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
)

type Config struct {
	Port string `yaml:"port"`
}

type PaymentStatus []string

var (
	port   = os.Getenv("PORT")
	config Config
)

func init() {
	yamlFile, err := ioutil.ReadFile("./config.yaml")

	if err != nil {
		log.Println("Error while reading config file ", err)
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Println("Error while unmarshalling config ", err)
	}

	if port == "" {
		port = config.Port
	}
}

func main() {
	http.HandleFunc("/payments", paymentCheck)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}

func paymentCheck(w http.ResponseWriter, r *http.Request) {
	if uid := r.PostFormValue("id"); uid != "" {
		statuses := PaymentStatus{
			"OK",
			"LOW MONEY",
			"INCORRECT",
		}

		rand.Seed(42)

		fmt.Fprintln(w, statuses[rand.Intn(len(statuses))])
	}
}
