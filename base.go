package hbgo

type IRoom interface {
	Id() uint32
	AsyncCall(id interface{}, args ...interface{})
	GetStartTime()
}

type IPlayer interface {
}
