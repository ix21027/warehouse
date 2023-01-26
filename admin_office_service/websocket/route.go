package websocket

import (
	"admin_office_service/websocket/code"
)

func Route(reqData ReqData, n *Server) {
	switch reqData.Code {
	case code.GET_USER_BY_ID:
		GetUser(reqData.Data, n)

	case code.GET_USERS_BY_STATUS:
		GetUsersByStatus(reqData.Data, n)

	case code.GET_USERS_BY_LOGIN:
		GetUsersByLogin(reqData.Data, n)

	//case code.BAN_USER:
	//	BanUser(reqData.Data, n)

	case code.CREATE_GOOD:
		CreateGood(reqData.Data, n)

	case code.UPDATE_GOOD:
		UpdateGood(reqData.Data, n)

	case code.DELETE_GOOD:
		DeleteGood(reqData.Data, n)

	case code.FINISH_ORDER:
		FinishOrder(reqData.Data, n)

	case code.GET_ORDERS_BY_USER:
		GetOrderByUser(reqData.Data, n)

	case code.GET_ALL_ORDERS:
		GetAllOrders(reqData.Data, n)

	default:
		ch <- "not supported route error"
	}
}
