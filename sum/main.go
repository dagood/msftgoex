package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"
)

func sumHandler(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query().Get("v")

	s := sha256.New()
	if _, err := s.Write([]byte(r.RequestURI)); err != nil {
		log.Fatalln(err)
	}
	sumStr := hex.EncodeToString(s.Sum(nil))

	t := time.Now().UTC().Format(time.RFC3339Nano)
	if _, err := s.Write([]byte(t)); err != nil {
		log.Fatalln(err)
	}
	sumStr2 := hex.EncodeToString(s.Sum(nil))

	data, err := json.MarshalIndent(struct {
		Request string
		Version string
		Type    string
		Value   string
		Sum     string
		Value2  string
		Sum2    string
	}{
		r.RequestURI,
		runtime.Version(),
		fmt.Sprintf("%T", s),
		v,
		sumStr,
		v + t,
		sumStr2,
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
