package goutils

type Range_t struct {
    Start int64
    End   int64
}

type RangeNode_t struct {
    Range_t
    Next *RangeNode_t
}

type RangeLink_t struct {
    length int
    header RangeNode_t
    tail *RangeNode_t
}

/**
 * start: 4
 * length: 3
 * end: start + length = 7
 * content: [4, 5, 6]
 */

func line(foo_s int64, foo_e int64, bar_s int64, bar_e int64) (vec int, start int64, end int64) {
    // foo在左边
    if foo_e < bar_s {
        return -1, -1, -1
    }
    // foo在右边
    if bar_e < foo_s {
        return 1, -1, -1
    }
    tar_s := foo_s
    // foo的起点在bar中间
    if bar_s < foo_s {
        tar_s = bar_s
    }
    tar_e := foo_e
    if foo_e < bar_e {
        tar_e = bar_e
    }
    return 0, tar_s, tar_e
}

func _NewRangeLink() *RangeLink_t {
    this := &RangeLink_t{
        length: 0,
        header: RangeNode_t {
            Next: nil,
        },
        tail: nil,
    }

    this.tail = &this.header

    return this
}

func NewRangeLink(arr []Range_t) *RangeLink_t {
    this := _NewRangeLink()
    if nil == arr {
        return this
    }

    this.length = len(arr)
    p := &(this.header)

    for i := 0; i < this.length; i++ {
        newNode := &RangeNode_t {
            Next: nil,
        }
        newNode.Start = arr[i].Start
        newNode.End = arr[i].End
        p.Next = newNode
        p = newNode
    }
    this.tail = p

    return this
}

// 头查看
func (this *RangeLink_t) Front() *Range_t {
    firstNode := this.header.Next
    if nil == firstNode {
        return nil
    }
    return &firstNode.Range_t
}

// 头删除
func (this *RangeLink_t) Pop() {
    p := this.header.Next
    if nil == p {
        return
    }
    this.header.Next = p.Next
    if nil == p.Next {
        this.tail = &this.header
    }
    this.length--
}

// 头添加
func (this *RangeLink_t) Push(newNode *RangeNode_t) {
    newNode.Next = this.header.Next
    this.header.Next = newNode
    this.length++
}

// 尾添加
func (this *RangeLink_t) Append(newNode *RangeNode_t) {
    this.tail.Next = newNode
    this.tail = newNode
    this.length++
}

// 挂载
func (this *RangeLink_t) Mount(start int64, end int64) {
    newNode := &RangeNode_t {
        Range_t: Range_t {
            Start: start,
            End: end,
        },
        Next: nil,
    }

    p := &(this.header)

    for {
        // 尽头
        if nil == p.Next {
            p.Next = newNode
            this.length++
            this.tail = newNode
            return
        }

        curNode := p.Next
        vec, lineStart, lineEnd := line(start, end, curNode.Start, curNode.End)

        // 头插入
        if -1 == vec {
            newNode.Next = curNode
            p.Next = newNode
            curNode = newNode
            this.length++
            return
        }

        // 节点扩大
        if 0 == vec {
            curNode.Start = lineStart
            curNode.End = lineEnd
            nextNode := curNode.Next

            // 与后序节点连接
            if nil != nextNode && nextNode.Start <= curNode.End {
                curNode.End = nextNode.End
                curNode.Next = nextNode.Next
                this.length--
                // 最后一个节点
                if nextNode == this.tail {
                    this.tail = curNode
                }
            }

            // 如果前序不是头节点，与前序节点连接
            if p != &(this.header) && curNode.Start <= p.End {
                p.End = curNode.End
                p.Next = curNode.Next
                this.length--
                // 最后一个节点
                if curNode == this.tail {
                    this.tail = p
                }
            }
            return
        }

        // 下一个节点
        p = p.Next
    }
}

// 转成数组
func (this *RangeLink_t) ToArray() []Range_t {
    ret := make([]Range_t, this.length)
    p := this.header.Next

    for i := 0; i < this.length; i++ {
        ret[i].Start = p.Start
        ret[i].End = p.End
        p = p.Next
    }
    return ret
}

// 取反
func (this *RangeLink_t) Converse(start int64, end int64) *RangeLink_t {
    if end <= start {
        return nil
    }
    ret := _NewRangeLink()
    _start := start
    _end := end
    p := this.header.Next
    for ; nil != p; p = p.Next {
        if start < p.Start {
            _start = start
        }
        _end = p.Start
        if end < p.Start {
            _end = end
        }
        if start <= p.End {
            start = p.End
        }
        if _end <= _start {
            continue
        }
        ret.Append(&RangeNode_t {
            Range_t: Range_t { Start: _start, End: _end },
            Next: nil,
        })
    }
    // now, start is lastNode.End
    if start < end {
        ret.Append(&RangeNode_t {
            Range_t: Range_t { Start: start, End: end },
            Next: nil,
        })
    }
    return ret
}
