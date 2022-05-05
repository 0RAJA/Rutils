package treenums

type TreeNum struct {
	nums []int
}

// NewTreeNum 初始化一个存储n个元素的树状数组
func NewTreeNum(n int) *TreeNum {
	return &TreeNum{nums: make([]int, n+1)}
}

//计算x二进制的最低位1及之后的数。101000 -> 001000
func (t *TreeNum) lowbit(x int) int {
	return x & (-x)
}

// Add x点值+d;x = [0,n-1]
func (t *TreeNum) Add(x, d int) {
	x++
	for ; x < len(t.nums); x += t.lowbit(x) { //因为动态维护一个前缀和数组，所以依次都要+d
		t.nums[x] += d
	}
}

// Ask 查询x下标处的前缀和;x = [0,n-1]
func (t *TreeNum) Ask(x int) (ret int) {
	x++
	for ; x >= 1 && x < len(t.nums); x -= t.lowbit(x) {
		ret += t.nums[x]
	}
	return
}

/*
单点修改，区间查询
	[x]+d == add(x,d)
	[l,r] == ask(r)-ask(l-1)
区间修改，单点查询(使用树状数组维护一个查分数组b，记录查分数组的前缀和)
	[l,r]+d == add(l,d);add(r+1,-d) //[l]+d使得后面的都+d，[r+1]-d 保证区间外的值不受影响
	a[x] == ask(x)+a[x] // 查分数组记录了区间的值得变化情况
*/
