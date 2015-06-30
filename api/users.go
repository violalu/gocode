package api

import (
 "./utils"
 "./model"
)

type UserController struct {
  session *mgo.Session
}

func NewUserController(session *mgo.Session) *UserController {
  return &UserController{session}
}

func (uc *UserController) CreateUser(response *http.ResponseWriter, request *http.Request, p httprouter.Params) {


}
