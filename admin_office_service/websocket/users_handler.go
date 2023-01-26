package websocket

func GetUser(data map[string]string, n *Server) {
	n.SendReqToUserSvcGetByID([]byte(data["id"]))
}

func GetUsersByStatus(data map[string]string, n *Server) {
	n.SendReqToUserSvcGetByStatus([]byte(data["status"]))
}

//func BanUser(data map[string]string, n *Server) {
//	ch <- "hello from BanUser"
//}

func GetUsersByLogin(data map[string]string, n *Server) {
	n.SendReqToUserSvcGetByLogin([]byte(data["login"]))
}
