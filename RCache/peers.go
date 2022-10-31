package RCache

import (
	pb "github.com/0RAJA/Rutils/RCache/geecachepb"
)

// 分布式调用接口

// PeerPicker 对应 HttpPool,其 PickPeer() 方法用于根据传入的 key 选择相应节点 PeerGetter。
type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

// PeerGetter 接口对应 HTTP 客户端 PeerGetter,其 Get() 方法用于从对应 group 查找缓存值
type PeerGetter interface {
	Get(in *pb.Request, out *pb.Response) error
}
