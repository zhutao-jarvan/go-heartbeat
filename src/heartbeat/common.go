package heartbeat

const (
	cMsgTypeSync = iota
	cMsgTypeRspNone
	cMsgTypeRsp
)

type HeartbeatHead struct {
	Id    uint64 // 标记用户
	Magic uint64 // 动态验证用户有效性, addr + random Data
	Len   uint16 // 消息的长度
	Type  uint8  // 消息类型
}
