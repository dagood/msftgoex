package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func sumHandler(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query().Get("v")

	s := sha256.New()
	_, err := s.Write([]byte(r.RequestURI))
	if err != nil {
		log.Fatalln(err)
	}
	sum := s.Sum(nil)
	sumStr := hex.EncodeToString(sum)

	data, err := json.MarshalIndent(struct {
		Request string
		Type    string
		Value   string
		Sum     string
	}{
		r.RequestURI,
		fmt.Sprintf("%T", s),
		v,
		sumStr,
	}, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}
	if _, err = w.Write(data); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	http.HandleFunc("/sum", sumHandler)
	http.ListenAndServe(":8080", nil)
}
