package main

import (
	"math"
	"strconv"
	"sync"
	"time"
)

func removeFromArr(arr *[]*Meal, ptr *Meal) {
	index := -1
	for i, meal := range *arr {
		if meal == ptr {
			index = i
			break
		}
	}
	if index != -1 {
		*arr = append((*arr)[:index], (*arr)[index+1:]...)
	}
}

type OrderList struct {
	deliveryMutex sync.Mutex
	mealMutex     sync.Mutex
	ovenList      []*Meal
	stoveList     []*Meal
	nilList       []*Meal
	orderArr      []*Order
}

func NewOrderList() *OrderList {
	ret := new(OrderList)
	ret.deliveryMutex = sync.Mutex{}
	ret.mealMutex = sync.Mutex{}
	ret.ovenList = []*Meal{}
	ret.stoveList = []*Meal{}
	ret.nilList = []*Meal{}
	ret.orderArr = []*Order{}
	return ret
}

func (orderList *OrderList) addOrder(order *Order) {
	orderList.orderArr = append(orderList.orderArr, order)
	for _, meal := range order.mealList {
		apparatId := meal.Apparat
		switch apparatId {
		case 0:
			orderList.nilList = append(orderList.nilList, meal)
		case 1:
			orderList.ovenList = append(orderList.ovenList, meal)
		case 2:
			orderList.stoveList = append(orderList.stoveList, meal)
		}
	}
}

func (orderList *OrderList) getSender() *Sender {
	orderList.deliveryMutex.Lock()
	defer orderList.deliveryMutex.Unlock()

	for i, order := range orderList.orderArr {
		if order.isReady() {
			for _, meal := range order.mealList {
				apparatId := meal.Apparat
				switch apparatId {
				case 0:
					removeFromArr(&orderList.nilList, meal)
				case 1:
					removeFromArr(&orderList.ovenList, meal)
				case 2:
					removeFromArr(&orderList.stoveList, meal)
				}
			}
			orderList.orderArr = append(orderList.orderArr[:i], orderList.orderArr[i+1:]...)
			return newSender(order)
		}
	}
	return nil
}

func (orderList *OrderList) getMeal(chef *Cook) *Meal {

	orderList.mealMutex.Lock()
	defer orderList.mealMutex.Unlock()

	now := time.Now().Unix()
	overallMin := math.MaxInt64
	var ret *Meal
	ovenTimeLeft := kitchen.ovens.getTimeLeft(now)
	for _, meal := range orderList.ovenList {
		readMeal := meal.get()
		if readMeal.prepared == 0 && readMeal.busy == 0 && readMeal.complexity <= chef.rank {
			timeLeft := readMeal.getTimeLeft(now) + ovenTimeLeft
			if overallMin > timeLeft {
				overallMin = timeLeft
				ret = readMeal
			}
		}
	}
	stoveTimeLeft := kitchen.stoves.getTimeLeft(now)
	for _, meal := range orderList.stoveList {
		readMeal := meal.get()
		if readMeal.prepared == 0 && readMeal.busy == 0 && readMeal.complexity <= chef.rank {
			timeLeft := readMeal.getTimeLeft(now) + stoveTimeLeft
			if overallMin > timeLeft {
				overallMin = timeLeft
				ret = readMeal
			}
		}
	}
	for _, meal := range orderList.nilList {
		readMeal := meal.get()
		if readMeal.prepared == 0 && readMeal.busy == 0 && readMeal.complexity <= chef.rank {
			timeLeft := readMeal.getTimeLeft(now)
			if overallMin > timeLeft {
				overallMin = timeLeft
				ret = readMeal
			}
		}
	}

	if ret != nil {
		return ret.get()
	}

	return ret
}

func (orderList *OrderList) getStatus() string {
	var ret string

	now := time.Now().Unix()
	for _, order := range orderList.orderArr {
		ret += HTMLWriter("Order " + strconv.Itoa(order.id) + "\nItems to be done: " + strconv.Itoa(int(order.mealCounter)) +
			" Time passed " + strconv.Itoa(int(now-order.pickUpTime)) + " Time limit: " + strconv.Itoa(order.maxWait))
	}
	return ret
}
