package proto

type ConnArg struct {
	Token  string
	Server int32
}

type ConnReply struct {
	Key    string
	RoomId int32
}

type DisconnArg struct {
	Key    string
	RoomId int32
}

type DisconnReply struct {
	Has bool
}

type MarkReadArg struct{
	Key string
	MarkKey string
}

type MarkReadReply struct{
	Msg string
}

type UnreadNumArg struct {
	Key string
}

type UnreadNumReply struct{
	Result []byte
}