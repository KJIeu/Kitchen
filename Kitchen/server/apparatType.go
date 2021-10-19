package main

import (
	"math"
	"strconv"
	"sync"
)

type ApparatList struct {
	numOfApparat int
	list         []*Apparat
	listMutex    sync.Mutex
}

func (al *ApparatList) getApparatAndWait(now int64) (*Apparat, int) {

	al.listMutex.Lock()
	appa := al.list[0].getValues()
	minWait := math.MaxInt32

	for _, loopAppa := range al.list {
		timeLeft := loopAppa.getTimeLeft(now)
		if timeLeft == 0 {
			minWait = 0
			appa = loopAppa.getValues()
			break
		}
		if minWait > timeLeft {
			minWait = timeLeft
			appa = loopAppa.getValues()
		}
	}
	al.listMutex.Unlock()

	return appa, minWait
}

func newApparat(numOfApparat int) *ApparatList {
	ret := new(ApparatList)
	ret.numOfApparat = numOfApparat
	for i := 0; i < numOfApparat; i++ {
		ret.list = append(ret.list, new(Apparat))
	}
	return ret
}

func (al *ApparatList) getTimeLeft(now int64) int {
	minWait := math.MaxInt32
	for i, _ := range al.list {
		timeLeft := al.list[i].getTimeLeft(now)
		if timeLeft == 0 {
			return 0
		}
		if minWait > timeLeft {
			minWait = timeLeft
		}
	}
	return minWait
}

func (al *ApparatList) getStatus() string {
	ret := ""
	for i, Apparat := range al.list {
		identification := "Id:" + strconv.Itoa(i)
		if Apparat.busy == 1 {
			identification += " Used by chef id:"
			if Apparat.chef != nil {
				identification += strconv.Itoa(Apparat.chef.id)
			}
		} else {
			identification += " Free"
		}
		ret += HTMLWriter(identification)
	}
	return ret
}
