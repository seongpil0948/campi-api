package app

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
)

type FireApp struct {
	Ctx  context.Context
	Inst *firebase.App
}

func (f FireApp) ToString() string {
	_id, err := f.Inst.InstanceID(f.Ctx)
	return fmt.Sprintf("FireApp: %v, %v", _id, err)
}

var instance *FireApp

func newApp() *FireApp {
	appInst := new(FireApp)
	appInst.Ctx = context.Background()
	app, err := firebase.NewApp(appInst.Ctx, nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	appInst.Inst = app
	return appInst
}

func GetFireInstance() *FireApp {
	if instance == nil {
		instance = newApp()
	}
	return instance
}
