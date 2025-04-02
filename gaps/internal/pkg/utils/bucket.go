package utils

type Bucket struct {
	cap  int
	set  map[int]struct{}
	list []int
}

func NewBucket(bucket []int, cap int) *Bucket {
	b := &Bucket{
		cap:  cap,
		set:  make(map[int]struct{}),
		list: bucket,
	}
	for _, element := range bucket {
		b.set[element] = struct{}{}
	}
	return b
}

func (b *Bucket) Add(newElement int) bool {
	if _, ok := b.set[newElement]; ok {
		return true
	}
	if len(b.set) >= b.cap {
		return false
	}
	if b.cap == 3 && len(b.list) == 2 {
		x := b.list[0]
		y := b.list[1]
		z := newElement
		if !haveThreeSumCondition(x, y, z) {
			return false
		}
	}
	b.set[newElement] = struct{}{}
	b.list = append(b.list, newElement)
	return true
}

func haveThreeSumCondition(x, y, z int) bool {
	return x+y == z || x+z == y || z+y == x
}
