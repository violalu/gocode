package utils

import (
 "gopkg.in/mgo.v2/bson"
)

const (
 DB_URL := "mongodb://localhost"
)
func getSession() *mgo.Session{
 s, err:= mgo.Dial(DB_URL)
 if err != nil {
  panic(err)
 }

 return s
}
