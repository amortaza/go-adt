package adt

import (
	"container/list"
)

type CallbacksByID struct {
	cbsById map[string] *list.List
}

func NewCallbacksByID() *CallbacksByID {
	ans := &CallbacksByID{}

	ans.cbsById = make(map[string] *list.List)

	return ans
}

func (c *CallbacksByID) Clear(id string) {
	delete(c.cbsById, id)
}

func (c *CallbacksByID) ClearAll() {
	c.cbsById = make(map[string] *list.List)
}

func (c *CallbacksByID) HasId(id string) bool {
	_, ok := c.cbsById[id]

	return ok
}

func (c *CallbacksByID) CallAll(id string, param interface{}) {
	cbs, ok := c.cbsById[id]

	if ok {
		for e := cbs.Front(); e != nil; e = e.Next() {
		    	cb := e.Value.(func(interface{}))

			cb(param)
		}
	}
}

func (c *CallbacksByID) Add(id string, cb func(interface{})) {

	cbs, ok := c.cbsById[id]

	if !ok {
		cbs = list.New()

		c.cbsById[id] = cbs
	}

	cbs.PushBack(cb)
}


