// Code generated by "stringer -type=ServerMessageType"; DO NOT EDIT.

package vnc

import "fmt"

const _ServerMessageType_name = "FramebufferUpdateMsgTypeSetColorMapEntriesMsgTypeBellMsgTypeServerCutTextMsgType"

var _ServerMessageType_index = [...]uint8{0, 24, 49, 60, 80}

func (i ServerMessageType) String() string {
	if i >= ServerMessageType(len(_ServerMessageType_index)-1) {
		return fmt.Sprintf("ServerMessageType(%d)", i)
	}
	return _ServerMessageType_name[_ServerMessageType_index[i]:_ServerMessageType_index[i+1]]
}
