package websocket

import (
	"daimon/api"
	"daimon/pkg/ctx"
	"daimon/pkg/e"
	"daimon/pkg/log"
	"net/http"

	uuid "github.com/satori/go.uuid"

	"github.com/gorilla/websocket"
)

type WsService struct {
}

func NewService() *WsService {
	return &WsService{}
}

func init() {
	go wsManager.start()
	go wsManager.toClientService()
	go wsManager.toGroupService()
	go wsManager.broadcastService()
}

// gin 处理 websocket handler
// todo: specify a group for connection
func (s *WsService) Connect(ctx *ctx.Context) (res *api.Message, err e.Error) {
	res = &api.Message{}
	upGrader := websocket.Upgrader{
		// cross origin domain
		CheckOrigin: func(r *http.Request) bool { return false },
		//  sec-websocket-protocl header
		Subprotocols: []string{ctx.GetHeader("Sec-WebSocket-Protocol")},
	}

	conn, er := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if er != nil {
		err = e.WsConnectError
		log.L.Warnf("websocket connect error: %s", ctx.Param("channel"))
		return
	}

	client := &Client{
		Id: uuid.NewV4().String(),
		// Group:   ctx.Param("channel"),
		// use a test group for test
		Group:   "test",
		Socket:  conn,
		Message: make(chan []byte, 1024),
	}
	wsManager.RegisterClient(client)
	res.Id = client.Id
	res.Group = client.Group
	return
}

func (s *WsService) SendToMe(ctx *ctx.Context, req *api.Message) {
	wsManager.Send(req.Id, req.Group, []byte("this is test message"))

}
