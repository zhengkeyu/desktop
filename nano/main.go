package main

import (
	"time"

	"github.com/lonng/nano"
	"github.com/lonng/nano/component"
	"github.com/lonng/nano/serialize/json"
	"github.com/lonng/nano/session"
)

type HallManager struct {
	component.Base
	Group *nano.Group
}

//可以重新实现
//Init()
//AfterInit()
//BeforeShutdown()
//Shutdown()
func (m *HallManager) AfterInit() {}

var Hall = &HallManager{Group: nano.NewGroup("hall")}

func main() {

	comps := &component.Components{}
	comps.Register(Hall, component.WithName("HallManager"))
	nano.Listen(":8700",
		nano.WithHeartbeatInterval(10*time.Second),
		//nano.WithLogger(log.WithField("component", "TPDZGame")),
		nano.WithSerializer(json.NewSerializer()),
		nano.WithComponents(comps),
		nano.WithIsWebsocket(true),
	)
}

type LoginToGameServerRequest struct {
	Version string `json:"version"` //客户端版本号
	Token   string `json:"token"`   //上传的token
	Uid     int64  `json:"uid"`     //这个Uid是大厅服务器传给玩家的
	//后加
	Coins int //机器人登录的钱
}

func (m *HallManager) Login(s *session.Session, req *LoginToGameServerRequest) error {
	return nil
}
