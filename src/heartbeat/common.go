package heartbeat

const (
	CMsgTypeSync = iota
	CMsgTypeRsp
)

const (
	CMsgIdOffset = 0
	CMsgMagicOffset = 8
	CMsgBacklogOffset = 16
	CMsgLenOffset = 20
	CMsgTypeOffset = 22
	CMsgLen = 24
)

type HeartbeatMsg struct {
	Id      uint64 // 标记用户
	Magic   uint64 // 动态验证用户有效性
	Backlog uint32 // 待处理事务数量, response only
	Len     uint16 // 消息的长度, Max 1500
	Type    uint8  // 消息类型
}
