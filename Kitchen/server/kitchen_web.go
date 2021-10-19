package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type KitchenWeb struct {
	kitchenServer   http.Server
	kitchenHandler  KitchenHandler
	kitchenClient   http.Client
	connectionError error
}

func (kw *KitchenWeb) start() {
	kw.kitchenServer.Addr = kitchenServerPort
	kw.kitchenServer.Handler = &kw.kitchenHandler

	fmt.Println(time.Now())
	fmt.Println("Kitchen is listening and serving on port" + kitchenServerPort)
	if err := kw.kitchenServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func (kw *KitchenWeb) deliver(delivery *Sender) {

	requestBody, marshallErr := json.Marshal(delivery)
	if marshallErr != nil {
		fmt.Println("Marshalling error:", marshallErr)
	}

	request, newRequestError := http.NewRequest(http.MethodPost, diningHallHost+diningHallPort+"/delivery", bytes.NewBuffer(requestBody))
	if newRequestError != nil {
		fmt.Println("Could not create new request. Error:", newRequestError)
	} else {
		_, doError := kw.kitchenClient.Do(request)
		if doError != nil {
			fmt.Println("ERROR Sending request. ERR:", doError)
			return
		}
	}
}

func (kw *KitchenWeb) establishConnection() bool {
	if kitchen.connected == true {
		return false
	}
	request, _ := http.NewRequest(http.MethodConnect, diningHallHost+diningHallPort+"/", bytes.NewBuffer([]byte{}))
	response, err := kw.kitchenClient.Do(request)
	if err != nil {
		kw.connectionError = err
		return false
	}
	var responseBody = make([]byte, response.ContentLength)
	response.Body.Read(responseBody)
	if string(responseBody) != "OK" {
		return false
	}
	return true
}
