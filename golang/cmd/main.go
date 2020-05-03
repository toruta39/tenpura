package main

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ok")
	})

	logrus.Info("Listening at 8080")
	logrus.Fatal(http.ListenAndServe(":8080", nil))
}
