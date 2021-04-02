package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

const PORT = 8888
const DEFAULT_SIZE = 128

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	log.SetOutput(os.Stdout)
}

func parseSizeQueryParam(req *http.Request) (int, error) {
	mb, exists := req.URL.Query()["mb"]
	if exists && len(mb[0]) > 0 {
		size, err := strconv.Atoi(mb[0])
		if err != nil {
			return 0, err
		}

		if size < 0 {
			err := fmt.Errorf("query param mb < 0")
			return 0, err
		}

		return size * 1024 * 1024, nil
	}

	return DEFAULT_SIZE, nil
}

func createZeroContent(size int) []byte {
	return make([]byte, size)
}

func handleHttp(w http.ResponseWriter, req *http.Request) {
	size, err := parseSizeQueryParam(req)
	if err != nil {
		errmsg := fmt.Sprintf("%d %s: %s", http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err)
		log.Printf("%s %s %s %s", req.RemoteAddr, req.Method, req.RequestURI, errmsg)

		http.Error(w, errmsg, http.StatusBadRequest)
		return
	}

	log.Printf("%s %s %s %d", req.RemoteAddr, req.Method, req.RequestURI, size)

	zeroContent := createZeroContent(size)

	w.WriteHeader(http.StatusOK)
	w.Write(zeroContent)
}

func main() {
	log.Printf("Start listening on port %d ...", PORT)

	http.HandleFunc("/", handleHttp)

	addr := fmt.Sprintf(":%d", PORT)
	log.Fatal(http.ListenAndServe(addr, nil))
}
