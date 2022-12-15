package websocket

func CreateGood(data any, ch chan string, s *Server) {
	s.SendNatsMsg("sub", "createGood")
	ch <- "hello from CreateGood"
}

func UpdateGood(data any, ch chan string, s *Server) {
	ch <- "hello from UpdateGood"
}

func DeleteGood(data any, ch chan string, s *Server) {
	ch <- "hello from DeleteGood"
}
