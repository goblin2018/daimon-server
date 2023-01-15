package websocket

import (
	"daimon/api"
	"daimon/pkg/log"
	"sync"
)

// 初始化 wsManager 管理器
var wsManager = WebsocketManager{
	clients:          make(map[string]map[string]*Client),
	registerChan:     make(chan *Client, 128),
	unRegisterChan:   make(chan *Client, 128),
	groupMessageChan: make(chan *api.Message, 128),
	messageChan:      make(chan *api.Message, 128),
	broadcastChan:    make(chan *api.Message, 128),
	groupCount:       0,
	clientCount:      0,
}

type WebsocketManager struct {
	clients          map[string]map[string]*Client
	groupCount       int
	clientCount      int
	Lock             sync.Mutex
	registerChan     chan *Client
	unRegisterChan   chan *Client
	messageChan      chan *api.Message
	groupMessageChan chan *api.Message
	broadcastChan    chan *api.Message
}

func (m *WebsocketManager) start() {
	log.L.Infof("websocket manager start")
	for {
		select {
		// 注册
		case client := <-m.registerChan:
			log.L.Infof("clinet [%s] connect", client.Id)
			log.L.Infof("register client [%s] to group [%s]", client.Id, client.Group)

			m.Lock.Lock()
			if m.clients[client.Group] == nil {
				m.clients[client.Group] = make(map[string]*Client)
				m.groupCount += 1
			}
			m.clients[client.Group][client.Id] = client
			m.clientCount += 1
			m.Lock.Unlock()

		case client := <-m.unRegisterChan:
			log.L.Infof("unregister client [%s] from group [%s]", client.Id, client.Group)
			m.Lock.Lock()
			if _, ok := m.clients[client.Group]; ok {
				if _, ok := m.clients[client.Group][client.Id]; ok {
					close(client.Message)
					delete(m.clients[client.Group], client.Id)
					m.clientCount -= 1
					if len(m.clients[client.Group]) == 0 {
						delete(m.clients, client.Group)
						m.groupCount -= 2
					}
				}
			}

			m.Lock.Unlock()
		}
	}
}

// 处理单个client发送数据
func (m *WebsocketManager) toClientService() {
	for {
		select {
		case data := <-m.messageChan:
			if groupMap, ok := m.clients[data.Group]; ok {
				if conn, ok := groupMap[data.Id]; ok {
					conn.Message <- data.Message
				}
			}
		}
	}
}

// 处理 group 广播数据
func (m *WebsocketManager) toGroupService() {
	for {
		select {
		case data := <-m.groupMessageChan:
			if groupMap, ok := m.clients[data.Group]; ok {
				for _, conn := range groupMap {
					conn.Message <- data.Message
				}
			}
		}
	}
}

// 处理广播数据
func (m *WebsocketManager) broadcastService() {
	for {
		select {
		case data := <-m.broadcastChan:
			for _, v := range m.clients {
				for _, conn := range v {
					conn.Message <- data.Message
				}
			}

		}
	}
}

// 向指定client发送数据
func (m *WebsocketManager) Send(id string, group string, message []byte) {
	data := &api.Message{
		Id:      id,
		Group:   group,
		Message: message,
	}
	m.messageChan <- data
}

// 向指定 group 广播
func (m *WebsocketManager) SendToGroup(group string, message []byte) {
	data := &api.Message{
		Group:   group,
		Message: message,
	}
	m.groupMessageChan <- data
}

// 广播
func (m *WebsocketManager) SendToAll(message []byte) {
	data := &api.Message{
		Message: message,
	}
	m.broadcastChan <- data
}

// 注册
func (m *WebsocketManager) RegisterClient(client *Client) {
	m.registerChan <- client
}

// 注销
func (m *WebsocketManager) UnRegisterClient(client *Client) {
	m.unRegisterChan <- client
}

// 当前组个数
func (m *WebsocketManager) LenGroup() uint {
	return uint(m.groupCount)
}

// 当前连接个数
func (m *WebsocketManager) LenClient() uint {
	return uint(m.clientCount)
}

// 获取管理器信息

func (m *WebsocketManager) Info() map[string]interface{} {
	info := make(map[string]interface{})
	info["groupLen"] = m.LenGroup()
	info["clientLen"] = m.LenClient()

	return info

}
