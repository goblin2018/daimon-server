package websocket

import (
	"daimon/api"
	"daimon/pkg/ctx"
	"daimon/pkg/log"
	"encoding/json"
	"net/http"
	"time"

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
func (s *WsService) Connect(ctx *ctx.Context) {
	upGrader := websocket.Upgrader{
		// cross origin domain
		CheckOrigin: func(r *http.Request) bool { return true },
		//  sec-websocket-protocl header
		Subprotocols: []string{ctx.GetHeader("Sec-WebSocket-Protocol")},
	}

	conn, er := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if er != nil {
		log.L.Warnf("websocket connect error: %s", er.Error())
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

	go client.Read()
	go client.Write()

	// todo 这里是一个测试消息，用于测试连接成功之后，向客户端发送消息
	go func() {
		time.Sleep(time.Second * 1)
		// 连接成功之后，发送 连接的id 和 group
		type msg struct {
			Id    string `json:"id" form:"id"`
			Group string `json:"group" form:"group"`
		}

		message := msg{
			Id:    client.Id,
			Group: client.Group,
		}
		b, _ := json.Marshal(message)
		wsManager.Send(client.Id, client.Group, b)
	}()

}

// 向指定的客户端发送消息
// todo 按照需求修改消息内容
func (s *WsService) SendToMe(ctx *ctx.Context, req *api.Message) {
	wsManager.Send(req.Id, req.Group, []byte("this is test message"))

}
