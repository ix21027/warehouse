package websocket

import (
	"admin_office_service/websocket/code"
)

func Route(reqData ReqData, ch chan string, s *Server) {
	switch reqData.Code {
	case code.GET_USER:
		GetUser(reqData.Data, ch, s)

	case code.GET_ALL_USERS:
		GetUsers(reqData.Data, ch, s)

	case code.BAN_USER:
		BanUser(reqData.Data, ch, s)

	case code.UNBAN_USER:
		UnBanUser(reqData.Data, ch, s)

	case code.CREATE_GOOD:
		CreateGood(reqData.Data, ch, s)

	case code.UPDATE_GOOD:
		UpdateGood(reqData.Data, ch, s)

	case code.DELETE_GOOD:
		DeleteGood(reqData.Data, ch, s)

	case code.FINISH_ORDER:
		FinishOrder(reqData.Data, ch, s)

	case code.GET_ORDERS_BY_USER:
		GetOrderByUser(reqData.Data, ch, s)

	case code.GET_ALL_ORDERS:
		GetAllOrders(reqData.Data, ch, s)

	default:
		ch <- "not supported route error"
	}
}
