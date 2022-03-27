package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type PaymentStatus struct {
	Result bool `json:"result"`
}

func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/pp/rest/v1/charge", MakeCharge)
	http.HandleFunc("/pp/rest/health", HealthCheck)
	http.ListenAndServe(":9000", nil)
}

func MakeCharge(w http.ResponseWriter, r *http.Request) {

	var paymentStatus PaymentStatus
	if rand.Int()%2 == 0 {
		paymentStatus = PaymentStatus{Result: true}
	} else {
		paymentStatus = PaymentStatus{Result: false}
	}

	js, err := json.Marshal(paymentStatus)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "ok")
}
