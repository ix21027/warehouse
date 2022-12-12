package codes

const (
	USERS  = 1
	GOODS  = 20
	ORDERS = 40
)

const (
	GET_USER = USERS + iota
	GET_ALL_USERS
	BAN_USER
	UNBAN_USER
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
