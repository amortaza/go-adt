package adt

import (
	"container/list"
)

type CallbacksByID struct {
	cbsByID map[string] *list.List
}

func NewCallbacksByID() *CallbacksByID {
	ans := &CallbacksByID{}

	ans.cbsByID = make(map[string] *list.List)

	return ans
}

func (c *CallbacksByID) Clear(id string) {
	delete(c.cbsByID, id)
}

func (c *CallbacksByID) ClearAll(id string) {
	c.cbsByID = make(map[string] *list.List)
}

func (c *CallbacksByID) HasId(id string) bool {
	_, ok := c.cbsByID[id]

	return ok
}

func (c *CallbacksByID) CallAll(id string, param interface{}) {
	cbs, ok := c.cbsByID[id]

	if ok {
		for e := cbs.Front(); e != nil; e = e.Next() {
		    	cb := e.Value.(func(interface{}))

			cb(param)
		}
	}
}

func (c *CallbacksByID) Add(id string, cb func(interface{})) {

	cbs, ok := c.cbsByID[id]

	if !ok {
		cbs = list.New()

		c.cbsByID[id] = cbs
	}

	cbs.PushBack(cb)
}


