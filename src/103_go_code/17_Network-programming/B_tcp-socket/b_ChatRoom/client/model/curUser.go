package model

import (
	"awesomeProject/src/103_go_code/17_Network-programming/B_tcp-socket/b_ChatRoom/common/message"
	"net"
)

//因为在客户端，我们很多地方会使用到curUser,我们将其作为一个全局
type CurUser struct {
	Conn net.Conn
	message.User
}
