package adt

import (
	"container/list"
)

type CallbacksById struct {
	cbsById map[string] *list.List
}

func NewCallbacksById() *CallbacksById {

	ans := &CallbacksById{}

	ans.cbsById = make(map[string] *list.List)

	return ans
}

func (c *CallbacksById) ClearId(id string) {

	delete(c.cbsById, id)
}

func (c *CallbacksById) ClearAll() {

	c.cbsById = make(map[string] *list.List)
}

func (c *CallbacksById) HasId(id string) bool {

	_, ok := c.cbsById[id]

	return ok
}

func (c *CallbacksById) CallAll(id string, param interface{}) {

	cbs, ok := c.cbsById[id]

	if ok {
		for e := cbs.Front(); e != nil; e = e.Next() {
	    	cb := e.Value.(func(interface{}))

			cb(param)
		}
	}
}

func (c *CallbacksById) Add(id string, cb func(interface{})) {

	cbs, ok := c.cbsById[id]

	if !ok {
		cbs = list.New()

		c.cbsById[id] = cbs
	}

	cbs.PushBack(cb)
}


