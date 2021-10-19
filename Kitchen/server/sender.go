package main

import "time"

type Sender struct {
	OrderId        int          `json:"order_id"`
	TableId        int          `json:"table_id"`
	Items          []int        `json:"items"`
	Priority       int          `json:"priority"`
	MaxWait        int          `json:"max_wait"`
	PickUpTime     int64        `json:"pick_up_time"`
	ChefingTime    int          `json:"chefing_time"`
	ChefingDetails []MealSender `json:"chefing_details"`
}

func newSender(order *Order) *Sender {
	ret := new(Sender)
	ret.OrderId = order.id
	ret.TableId = order.tableId
	ret.Items = order.items
	ret.Priority = order.priority
	ret.MaxWait = order.maxWait
	ret.PickUpTime = order.pickUpTime
	ret.ChefingTime = int(time.Now().Unix() - order.pickUpTime)
	var chefingDetails []MealSender
	for _, meal := range order.mealList {
		chefingDetails = append(chefingDetails, MealSender{meal.foodId, meal.chefId})
	}
	ret.ChefingDetails = chefingDetails
	return ret
}
