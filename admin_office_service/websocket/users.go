package websocket

func GetUser(data any, ch chan string, s *Server) {
	s.SendNatsMsg("sub", "GetUser")
	ch <- "hello from GetUser"
}

func GetUsers(data any, ch chan string, s *Server) {
	ch <- "hello from GetUsers"
}

func BanUser(data any, ch chan string, s *Server) {
	ch <- "hello from BanUser"
}

func UnBanUser(data any, ch chan string, s *Server) {
	ch <- "hello from UnBanUser"
}
