package websocket

func CreateGood(data any, n *Server) {
	n.SendNatsMsg("sub", "createGood")
	ch <- "hello from CreateGood"
}

func UpdateGood(data any, s *Server) {
	ch <- "hello from UpdateGood"
}

func DeleteGood(data any, s *Server) {
	ch <- "hello from DeleteGood"
}
