package bytes

import "sync"

type Buffer struct {
	buf  []byte
	next *Buffer
}

// 获取缓存字节
func (b *Buffer) Bytes() []byte {
	return b.buf
}

type BytePool struct {
	lock sync.Mutex
	free *Buffer
	max  int
	num  int
	size int
}

func NewBytePool(num, size int) (bp *BytePool) {
	bp = new(BytePool)
	bp.init(num, size)
	return
}

func (bp *BytePool) Init(num, size int) {
	bp.init(num, size)
}

func (bp *BytePool) init(num, size int) {
	bp.num = num
	bp.size = size
	bp.max = num * size
	bp.grow()
}

func (bp *BytePool) grow() {
	var (
		buf        []byte
		bufferList []Buffer
		buffer     *Buffer
		i          int
	)
	// 预申请一段连续的内存空间
	buf = make([]byte, bp.max)
	// 预申请连续的buf空间
	bufferList = make([]Buffer, bp.num)
	bp.free = &bufferList[0]
	buffer = bp.free
	for i = 1; i < bp.num; i++ {
		buffer.buf = buf[(i-1)*bp.size : i*bp.size]
		buffer.next = &bufferList[i]
		buffer = buffer.next
	}
	buffer.buf = buf[(i-1)*bp.size : i*bp.size]
	buffer.next = nil
}

// 从空闲栈获取
func (bp *BytePool) Get() (b *Buffer) {
	bp.lock.Lock()
	defer bp.lock.Unlock()
	if b = bp.free; b == nil {
		bp.grow()
		b = bp.free
	}
	bp.free = b.next
	return
}

// 放回空闲栈
func (bp *BytePool) Put(b *Buffer) {
	bp.lock.Lock()
	defer bp.lock.Unlock()
	b.next = bp.free
	bp.free = b
}
