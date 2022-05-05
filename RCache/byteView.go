package RCache

//封装缓存值

//ByteView 一个只读数据结构,用来表示缓存值
type ByteView struct {
	b []byte
}

// Size 返回缓存长度
func (v ByteView) Size() int {
	return len(v.b)
}

//克隆缓存值
func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}

// ByteSlice 返回缓存值副本,防止被外界修改
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

//将缓存值转为字符串,有必要时进行备份
func (v ByteView) String() string {
	return string(v.b)
}

/*
注:
选择 byte 类型是为了能够支持任意的数据类型的存储，例如字符串、图片等
*/
