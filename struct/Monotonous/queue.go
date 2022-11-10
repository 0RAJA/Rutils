package Monotonous

type Queue Assists

func (q *Queue) Push(v Comparable) {
	q.nums.PushBack(v)
	q.PushAssist(v)
}

func (q *Queue) Pop() interface{} {
	if q.nums.Len() <= 0 {
		return nil
	}
	t := q.nums.Front()
	q.PopAssist(t.Value.(Comparable))
	q.nums.Remove(t)
	return t.Value
}

func (q *Queue) Top() interface{} {
	if q.nums.Len() <= 0 {
		return nil
	}
	return q.nums.Front().Value
}

func (q *Queue) PopAssist(v Comparable) {
	if v.CompareTo(q.assistMin.Front().Value.(Comparable)) == 0 {
		q.assistMin.Remove(q.assistMin.Front())
	}
	// 移除最大
	if v.CompareTo(q.assistMax.Front().Value.(Comparable)) == 0 {
		q.assistMax.Remove(q.assistMax.Front())
	}
}

func (q *Queue) PushAssist(v Comparable) {
	// min: 从后往前剔除比当前值大的
	for q.assistMin.Len() > 0 && v.CompareTo(q.assistMin.Back().Value.(Comparable)) < 0 {
		q.assistMin.Remove(q.assistMin.Back())
	}
	q.assistMin.PushBack(v)
	// max：从后往前剔除比当前值小的
	for q.assistMax.Len() > 0 && v.CompareTo(q.assistMax.Back().Value.(Comparable)) > 0 {
		q.assistMax.Remove(q.assistMax.Back())
	}
	q.assistMax.PushBack(v)
}

func (q *Queue) Min() interface{} {
	return q.assistMin.Front().Value
}

func (q *Queue) Max() interface{} {
	return q.assistMax.Front().Value
}
