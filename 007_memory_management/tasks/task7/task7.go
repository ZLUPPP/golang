package task7

import (
	"bytes"
	"strconv"
	"sync"
)

var bufPool = sync.Pool{
	New: func() any { b := new(bytes.Buffer); b.Grow(4096); return b },
}

func RenderPool(ids []int) int {
	buf := bufPool.Get().(*bytes.Buffer)
	buf.Reset()

	for _, id := range ids {
		buf.WriteString("id=")
		buf.WriteString(strconv.Itoa(id))
		buf.WriteByte('\n')
	}

	n := buf.Len()
	bufPool.Put(buf)
	return n
}

func RenderNew(ids []int) int {
	buf := new(bytes.Buffer)
	buf.Grow(4096)

	for _, id := range ids {
		buf.WriteString("id=")
		buf.WriteString(strconv.Itoa(id))
		buf.WriteByte('\n')
	}

	return buf.Len()
}
