package main

import "sync"

type Apparat struct {
	busy         int32
	meal         *Meal
	chef         *Cook
	queueWait    int
	prepareMutex sync.Mutex
	valueMutex   sync.Mutex
}

func (a *Apparat) getTimeLeft(now int64) int {
	a.valueMutex.Lock()
	defer a.valueMutex.Unlock()
	if a.busy == 0 {
		return 0
	}
	return a.meal.getTimeLeft(now) + a.queueWait
}
func (a *Apparat) setValues(busy int32, chef *Cook, meal *Meal) {
	a.valueMutex.Lock()
	a.busy = busy
	a.meal = meal
	a.chef = chef
	a.valueMutex.Unlock()
}

func (a *Apparat) addQueueWait(value int) {
	a.valueMutex.Lock()
	a.queueWait += value
	a.valueMutex.Unlock()
}

func (a *Apparat) getValues() *Apparat {
	a.valueMutex.Lock()
	defer a.valueMutex.Unlock()
	return a
}

func (a *Apparat) useApparat(chef *Cook, meal *Meal, now int64) {
	timeForCurrentMeal := meal.getTimeLeft(now)
	a.addQueueWait(timeForCurrentMeal)
	a.prepareMutex.Lock()
	a.addQueueWait(-timeForCurrentMeal)
	a.setValues(1, chef, meal)

	meal.prepare(chef, now)

	a.setValues(0, nil, nil)
	a.prepareMutex.Unlock()
}
