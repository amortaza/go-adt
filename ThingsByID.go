package adt

import (
	"container/list"
)

type ThingsByID struct {
	thingsById map[string] *list.List
}

func NewThingsByID() *ThingsByID {
	ans := &ThingsByID{}

	ans.thingsById = make(map[string] *list.List)

	return ans
}

func (c *ThingsByID) Clear(id string) {
	delete(c.thingsById, id)
}

func (c *ThingsByID) ClearAll(id string) {
	c.thingsById = make(map[string] *list.List)
}

func (c *ThingsByID) Has(id string) bool {
	_, ok := c.thingsById[id]

	return ok
}

func (c *ThingsByID) Get(id string) (*list.List, bool) {

	a,b := c.thingsById[id]

	return a,b
}

func (c *ThingsByID) Add(id string, thing interface{}) {

	things, ok := c.thingsById[id]

	if !ok {
		things = list.New()

		c.thingsById[id] = things
	}

	things.PushBack(thing)
}


