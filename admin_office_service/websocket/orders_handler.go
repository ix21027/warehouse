package websocket

func FinishOrder(data any, ch chan string, s *Server) {
	ch <- "hello from FinishOrder"
}

func GetOrderByUser(data any, ch chan string, s *Server) {
	ch <- "hello from GetOrderByUser"
}

func GetAllOrders(data any, ch chan string, s *Server) {
	ch <- "hello from GetAllOrders"
}
