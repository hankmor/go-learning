package main

// Hub 是一个集成器，维护了活跃的 Client 和广播到 Client 消息
type Hub struct {
	// 已经注册的Client
	clients map[*Client]bool

	// 入站的消息
	broadcast chan []byte

	// 客户端的注册请求通道
	register chan *Client

	// 客户端的注销通道
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

// 启动集成器, 获取注册/注销，广播消息给 Client
func (h *Hub) run() {
	for {
        // 检测注册注销
		select {
		case client := <-h.register:
			// 如果收到注册请求，则注册
			h.clients[client] = true
		case client := <-h.unregister:
			// 如果收到注销请求，则注销
			if _, ok := h.clients[client]; ok {
			 	delete(h.clients, client)
				close(client.send) // 关闭消息通道，不再接收消息
			} 
		case msg := <-h.broadcast:
			// 如果收到了广播消息，则广播到所有客户端 
			for client := range h.clients {
				select {
				case client.send <- msg:
					// 成功发送消息
				default:
					// 无法发送消息，则关闭该客户端
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
