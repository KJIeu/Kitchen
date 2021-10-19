package main

import "math/rand"

type CookList struct {
	cookList      []*Cook
	cookIdCounter int
}

func NewCookList() *CookList {
	ret := new(CookList)
	ret.cookIdCounter = 0
	for i := 0; i < chefN; i++ {
		randomCook := chefPersonas[rand.Intn(len(chefPersonas))]
		randomCook.id = ret.cookIdCounter
		ret.cookIdCounter++
		if i == 0 {
			randomCook.rank = 3
		}
		ret.cookList = append(ret.cookList, NewCook(&randomCook))
	}
	return ret
}

func (cl CookList) start() {
	for _, cook := range cl.cookList {
		go cook.startWorking()
	}
}

var chefPersonas = []Cook{{
	rank:        1,
	proficiency: 1,
	name:        "Jimmy Chef",
	catchPhrase: "Yes",
}, {
	rank:        2,
	proficiency: 2,
	name:        "Wolfgang Puck",
	catchPhrase: "Belissimo",
}, {
	rank:        1,
	proficiency: 3,
	name:        "Jamie Oliver",
	catchPhrase: "Impecable",
}, {
	rank:        3,
	proficiency: 2,
	name:        "James",
	catchPhrase: "This one is trash",
}, {
	rank:        3,
	proficiency: 3,
	name:        "Gordon Ramsay",
	catchPhrase: "WHERE IS THE LAMB SAUCE?",
}}
