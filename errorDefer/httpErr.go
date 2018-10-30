package errorDefer

import (
	"net/http"
	"os"
	"io/ioutil"
	"log"
)

func showFailContent(writer http.ResponseWriter, request *http.Request) error {
	fileName := request.URL.Path[len("/list/"):]
	fHead, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer fHead.Close()
	buff, err := ioutil.ReadAll(fHead)
	if err != nil {
		return err
	}
	writer.Write(buff)
	return nil
}

type errDeferHeader func(writer http.ResponseWriter, request *http.Request) error

func weepErr(p errDeferHeader) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := p(writer, request)
		if err != nil {
			log.Printf("error: %v", err)
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}

func HttpErrDemo() {
	http.HandleFunc("/list/", weepErr(showFailContent))
	err := http.ListenAndServe(":8889", nil)
	if err != nil {
		log.Fatalf("err: %v", err)
	}
}
