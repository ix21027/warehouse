package code

const (
	USERS  = 1
	GOODS  = 20
	ORDERS = 40
)

const (
	GET_USER_BY_ID = USERS + iota
	GET_USERS_BY_STATUS
	GET_USERS_BY_LOGIN
	BAN_USER
)

const (
	CREATE_GOOD = GOODS + iota
	UPDATE_GOOD
	DELETE_GOOD
)

const (
	FINISH_ORDER = ORDERS + iota
	GET_ORDERS_BY_USER
	GET_ALL_ORDERS
)
