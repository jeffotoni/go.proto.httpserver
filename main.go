package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/golang/protobuf/proto"
	protoc "github.com/jeffotoni/go.protobuffers/customer"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/customer/proto", Customer)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func Customer(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		log.Println("Error ioutil:", err)
		return
	}

	var pCustomer = &protoc.Customer{}
	err = proto.Unmarshal(b, pCustomer)
	if err != nil {
		w.WriteHeader(400)
		log.Fatal("unmarshaling error: ", err)
		return
	}

	fmt.Println(pCustomer)
	w.WriteHeader(200)
	w.Write([]byte("ok"))
}
