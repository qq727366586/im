package timer

import (
	"sync"
	"time"
)

const (
	infiniteDuration = time.Duration(1<<63 - 1)
)

type TimersData struct {
	key    string      // channel 对应的唯一key
	expire time.Time   // 过期时间
	fn     func()      // 回调函数
	index  int         // 堆的索引
	next   *TimersData // 下一个节点
}

type HeapTimer struct {
	lock   sync.Mutex
	free   *TimersData   // 空闲链表
	timers []*TimersData // 初始化的容量
	signal *time.Timer   // 信号
	num    int           // 数量
}

// 初始化
func NewHeapTimer(num int) (h *HeapTimer) {
	h = new(HeapTimer)
	h.init(num)
	return h
}

// 初始化
func (h *HeapTimer) Init(num int) {
	h.init(num)
}

func (h *HeapTimer) init(num int) {
	h.timers = make([]*TimersData, 0, num)
	h.signal = time.NewTimer(infiniteDuration)
	h.num = num
	go h.start()
	h.grow()
}

// 等待计数器到期
func (h *HeapTimer) start() {
	for {
		h.expire()
		<-h.signal.C
	}
}

func (h *HeapTimer) expire() {
	var (
		d  time.Duration
		td *TimersData
		fn func()
	)
	h.lock.Lock()
	defer h.lock.Unlock()
	for {
		if len(h.timers) == 0 {
			d = infiniteDuration
			break
		}
		// 看第一个元素是否过期
		td = h.timers[0]
		if td.Delay() > 0 {
			break
		}
		// 否则 执行计数器 回调函数
		fn = td.fn
		// 让使用者选择 put 操作
		h.del(td)
		h.lock.Unlock()
		// 需要判断 nil
		if fn != nil {
			fn()
		}
		h.lock.Lock()
	}
	h.signal.Reset(d)
}

func (h *HeapTimer) grow() {
	// 预申请空闲指针变量
	tds := make([]TimersData, h.num)
	h.free = &tds[0]

	td := h.free
	for i := 1; i < h.num; i++ {
		td.next = &tds[i]
		td = td.next
	}
	td.next = nil
}

// 添加计时器 log(n)
func (h *HeapTimer) Add(expire time.Duration, fn func()) (td *TimersData) {
	h.lock.Lock()
	defer h.lock.Unlock()
	td = h.get()
	td.expire = time.Now().Add(expire)
	td.fn = fn
	h.add(td)
	return
}

// 删除计时器 log(n)
func (h *HeapTimer) Del(td *TimersData) {
	h.lock.Lock()
	defer h.lock.Unlock()
	h.del(td)
	h.put(td)
}

// 更新计时器 log(n)
func (h *HeapTimer) Set(td *TimersData, expire time.Duration) {
	h.lock.Lock()
	defer h.lock.Unlock()
	h.del(td)
	td.expire = time.Now().Add(expire)
	h.add(td)
}

func (h *HeapTimer) get() (td *TimersData) {
	if td = h.free; td == nil {
		h.grow()
		td = h.free
	}
	h.free = td.next
	return
}

func (h *HeapTimer) put(td *TimersData) {
	td.fn = nil
	// 放入头部
	td.next = h.free
	h.free = td
}

func (h *HeapTimer) del(td *TimersData) {
	i := td.index
	last := len(h.timers) - 1
	// 先判断是否可能已经过期 早就被移除了
	if i < 0 || i > last || h.timers[i] != td {
		return
	}
	// 否则与最后一个交换 在进行 Heap Insert/Down 操作
	if i != last {
		h.swap(i, last)
		h.down(i, last)
		h.up(i)
	}
	h.timers[last].index = -1
	h.timers = h.timers[:last]
}

func (h *HeapTimer) add(td *TimersData) {
	td.index = len(h.timers)
	h.timers = append(h.timers, td)
	h.up(td.index)
	// 如果是第一个, 需要重置定时器
	if td.index == 0 {
		delay := td.Delay()
		h.signal.Reset(delay)
	}
}

// t.Sub(timer.Now()) 的快捷方式
func (td *TimersData) Delay() time.Duration {
	return time.Until(td.expire)
}

// Heap Insert 过程
func (h *HeapTimer) up(j int) {
	for {
		i := (j - 1) / 2
		if i >= j || !h.less(j, i) {
			break
		}
		h.swap(i, j)
		j = i
	}
}

// Heap Down 过程
func (h *HeapTimer) down(i, n int) {
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 {
			break
		}
		j := j1
		if j2 := j1 + 1; j2 < n && !h.less(j1, j2) {
			j = j2
		}
		if !h.less(j, i) {
			break
		}
		h.swap(i, j)
		i = j
	}
}

// 过期时间对比 构建最小堆
func (h *HeapTimer) less(i, j int) bool {
	return h.timers[i].expire.Before(h.timers[j].expire)
}

// 交换
func (h *HeapTimer) swap(i, j int) {
	h.timers[i], h.timers[j] = h.timers[j], h.timers[i]
	h.timers[i].index = i
	h.timers[j].index = j
}
