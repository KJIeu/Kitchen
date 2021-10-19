package main

import "time"

type Kitchen struct {
	kitchenWeb KitchenWeb
	orderList  *OrderList
	ovens      *ApparatList
	stoves     *ApparatList
	chefList   *CookList
	connected  bool
}

func (k *Kitchen) start() {
	k.chefList = NewCookList()
	k.orderList = NewOrderList()
	k.ovens = newApparat(ovenN)
	k.stoves = newApparat(stoveN)

	go k.tryConnectDiningHall()
	k.kitchenWeb.start()
}

func (k *Kitchen) tryConnectDiningHall() {
	k.connected = false
	for k.connected {
		if k.kitchenWeb.establishConnection() {
			k.connectionSuccessful()
			break
		} else {
			time.Sleep(time.Second)
		}
	}
}

func (k *Kitchen) deliver(delivery *Sender) {
	k.kitchenWeb.deliver(delivery)
}

func (k *Kitchen) connectionSuccessful() {
	if k.connected {
		return
	}
	k.connected = true
	k.chefList.start()
}

func (k *Kitchen) getStatus() string {
	ret := ""
	for _, chef := range k.chefList.cookList {
		ret += HTMLWriter(chef.getStatus())
	}
	ret += "\tOvens:\n"
	ret += k.ovens.getStatus()
	ret += "\tStoves in use: \n"
	ret += k.stoves.getStatus()
	ret += "\tOrderList:\n"
	ret += k.orderList.getStatus()

	return ret
}
