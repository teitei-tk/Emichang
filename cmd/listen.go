package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/teitei-tk/emichang"
)

func main() {
	http.HandleFunc("/callback", emichang.HandleFunc)

	listenPortNo := fmt.Sprintf(":%s", os.Getenv("EMICHANG_LISTEN_PORT_NO"))
	certFilePath := os.Getenv("EMICHANG_CERT_FILE_PATH")
	keyFilePath := os.Getenv("EMICHANG_KEY_FILE_PATH")

	if err := http.ListenAndServeTLS(listenPortNo, certFilePath, keyFilePath, nil); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
