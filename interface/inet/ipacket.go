/*
* @Author: liws
* @Date:   2021/11/20 14:53
* @Desc:
**/
package inet

type Packet interface {
	Unpack(binaryData []byte) (IMessage, error)
	Pack(msg IMessage) ([]byte, error)
	GetHeadLen() uint32
}
