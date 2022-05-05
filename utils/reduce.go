package utils

import "context"

const (
	AttrUserid  = "userid"
	AttrUnionId = "unionId"
	AttDeptId   = "deptId"
)

type DingIdReduceFn func(ctx context.Context, attr string, src ...string) (dest map[string]string)

type DingIdReduceAble interface {
	SetReduceFn(fn DingIdReduceFn)
}

type DingIdReduceStruct struct {
	DingIdReduceFn DingIdReduceFn
}

func (d *DingIdReduceStruct) SetReduceFn(fn DingIdReduceFn) {
	d.DingIdReduceFn = fn
}

func (d *DingIdReduceStruct) ReduceBatch(ctx context.Context, attr string, src ...string) (dest map[string]string) {
	if d.DingIdReduceFn != nil {
		return d.DingIdReduceFn(ctx, attr, src...)
	}
	dest = make(map[string]string)
	for _, s := range src {
		dest[s] = s
	}
	return
}

func (d *DingIdReduceStruct) Reduce(ctx context.Context, attr string, src string) (dest string) {
	if d.DingIdReduceFn != nil {
		dest1 := d.DingIdReduceFn(ctx, attr, src)
		if len(dest1) > 0 {
			return dest1[src]
		}
	}
	return src
}
