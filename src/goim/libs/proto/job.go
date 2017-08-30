package proto

// TODO optimize struct after replace kafka
type KafkaMsg struct {
	OP       string   `json:"op"`
	RoomId   int32    `json:"roomid,omitempty"`
	ServerId int32    `json:"server,omitempty"`
	SubKeys  []string `json:"subkeys,omitempty"`
	Msg      []byte   `json:"msg"`
	Ensure   bool     `json:"ensure,omitempty"`
}

type KafkaUserMsg struct {
	OP       int32   `json:"op"`
	RoomId   int32    `json:"roomid,omitempty"`
	UserId   int64    `json:"userid,omitempty"`
	Key      string   `json:"key,omitempty"`
}