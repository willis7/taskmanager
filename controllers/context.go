package controllers

import (
	"gopkg.in/mgo.v2"
	"github.com/willis7/taskmanager/common"
)

// Struct used for maintaining Http Request Context
type Context struct {
	MongoSession *mgo.Session
}

// Close mgo.Session
func (c *Context) Close() {
	c.MongoSession.Close()
}

// Return mgo.collection for the given name
func (c *Context) DbCollection(name string) *mgo.Collection {
	return c.MongoSession.DB(common.AppConfig.Database).C(name)
}

// Create a new Context obect for each Http request
func NewContext() *Context {
	session := common.GetSession().Copy()
	context := &Context{MongoSession:session}
	return context
}
