package main

type Order struct {
	id          int
	tableId     int
	items       []int
	mealCounter int32
	priority    int
	pickUpTime  int64
	maxWait     int
	mealList    []*Meal
}

func parseOrder(postOrder *OrderProp) *Order {
	ret := new(Order)
	ret.id = postOrder.Id
	ret.tableId = postOrder.TableId
	ret.items = postOrder.Items
	ret.mealCounter = 0
	ret.priority = postOrder.Priority
	ret.pickUpTime = postOrder.PickUpTime
	ret.maxWait = postOrder.MaxWait
	for _, id := range postOrder.Items {
		ret.mealCounter += 1
		meal := newMeal(ret, id)
		ret.mealList = append(ret.mealList, meal)
	}
	return ret
}

func (order Order) isReady() bool {
	if order.mealCounter != 0 {
		return false
	}
	return true
}
