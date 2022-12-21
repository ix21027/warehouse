package service

type IGood interface {
	GetByID(any)
	GetAll(any)
}

type IOrder interface {
	CRUDer
}
type IUser interface {
	CRUDer
}

type CRUDer interface {
	Create(any)
	GetByID(any)
	Update(any)
	Delete(any)
}

type INatsServer interface {
	SendMsg(subj, msg string)
}
