package websocket

import (
	"admin_office_service/websocket/codes"
)

func Route(reqData ReqData, ch chan string, s *Server) {
	switch reqData.Code {
	case codes.GET_USER:
		GetUser(reqData.Data, ch, s)

	case codes.GET_ALL_USERS:
		GetUsers(reqData.Data, ch, s)

	case codes.BAN_USER:
		BanUser(reqData.Data, ch, s)

	case codes.UNBAN_USER:
		UnBanUser(reqData.Data, ch, s)

	case codes.CREATE_GOOD:
		CreateGood(reqData.Data, ch, s)

	case codes.UPDATE_GOOD:
		UpdateGood(reqData.Data, ch, s)

	case codes.DELETE_GOOD:
		DeleteGood(reqData.Data, ch, s)

	case codes.FINISH_ORDER:
		FinishOrder(reqData.Data, ch, s)

	case codes.GET_ORDERS_BY_USER:
		GetOrderByUser(reqData.Data, ch, s)

	case codes.GET_ALL_ORDERS:
		GetAllOrders(reqData.Data, ch, s)

	default:
		ch <- "not supported route error"
	}
}
