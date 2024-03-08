package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {

	http.HandleFunc("/", getRoot)
	http.HandleFunc("/luhncheck", getLuhnCheck)

	log.Fatal(http.ListenAndServe(":7770", nil))

	// var card string
	// _, err := fmt.Scanln(&card)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// passedLuhnCheck := luhnCheck(card)
	// fmt.Printf("Did %s pass Luhn Check: %t", card, passedLuhnCheck)
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request :/")
	io.WriteString(w, "Hello, Skulls!")
}

func getLuhnCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request :/luhncheck")
	cardno := r.URL.Query().Get("card")
	isValid := luhnCheck(cardno)
	data := map[string]bool{
		"passedLuhnCheck": isValid,
	}
	w.Header().Set("Content-Type", "application/json")
	// io.WriteString(w, strconv.FormatBool(isValid))
	json.NewEncoder(w).Encode(data)
}

func luhnCheck(cardNo string) (passedLuhnCheck bool) {

	alternate := false
	sum := 0

	for i := len(cardNo) - 1; i >= 0; i-- {
		adder, err := strconv.Atoi(string(cardNo[i]))

		if err != nil {
			return false
		}
		if alternate {
			twice := 2 * adder
			adder = (twice / 10) + (twice % 10)
		}
		sum += adder
		alternate = !alternate
	}
	passedLuhnCheck = (sum%10 == 0)
	return
}
