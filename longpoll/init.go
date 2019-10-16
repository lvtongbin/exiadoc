package longpoll

import (
	"fmt"
	"trism/golongpoll"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

var (
	manager *golongpoll.LongpollManager
)

func init() {
	var err error
	manager, err = golongpoll.StartLongpoll(golongpoll.Options{
		LoggingEnabled: false,
	})
	if err != nil {
		fmt.Println("start golongpoll failed")
	}
	beego.Get("/v1/longpoll/receiveMessage", receiveMessage(manager))
}

func receiveMessage(manager *golongpoll.LongpollManager) func(ctx *context.Context) {
	return func(ctx *context.Context) {
		manager.SubscriptionHandler(ctx.ResponseWriter, ctx.Request)
	}
}

// Publish ...
func Publish(category string, data interface{}) error {
	return manager.Publish(category, data)
}
