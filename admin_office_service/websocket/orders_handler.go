package websocket

func FinishOrder(data any, n *Server) {
	ch <- "hello from FinishOrder"
}

func GetOrderByUser(data any, n *Server) {
	ch <- "hello from GetOrderByUser"
}

func GetAllOrders(data any, n *Server) {
	ch <- "hello from GetAllOrders"
}
