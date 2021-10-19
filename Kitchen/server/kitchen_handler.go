package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type KitchenHandler struct {
	packetsReceived   int32
	postReceived      int32
	latestOrder       OrderProp
	latestOrderString string
}

func (oh *KitchenHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		{
			response := "OK"

			latestOrder := new(OrderProp)
			var requestBody = make([]byte, r.ContentLength)
			r.Body.Read(requestBody)
			json.Unmarshal(requestBody, latestOrder)
			kitchen.orderList.addOrder(parseOrder(latestOrder))

			fmt.Fprint(w, response)
		}
	case http.MethodGet:
		{
			fmt.Fprintln(w, "<head><meta http-equiv=\"refresh\" content=\"1\" /></head>")
			if kitchen.connected {
			} else {
				fmt.Fprintln(w, HTMLWriter("connection enstablished: "+diningHallHost+diningHallPort))
				err := kitchen.kitchenWeb.connectionError
				if err != nil {
					fmt.Fprintln(w, HTMLWriter("Connection error: "+err.Error()))
				}
			}
			fmt.Fprintln(w, kitchen.getStatus())
		}
	case http.MethodConnect:
		{
			kitchen.connectionSuccessful()
			fmt.Fprint(w, "OK")
		}
	default:
		{
			fmt.Fprintln(w, "UNSUPPORTED METHOD")
		}
	}

}
