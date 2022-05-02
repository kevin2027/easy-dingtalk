package utils

import "context"

func DingIdReduceBatch(dingIdReduceFn DingIdReduceFn, ctx context.Context, attr Attr, src ...string) (dest map[string]string) {
	if dingIdReduceFn != nil {
		return dingIdReduceFn(ctx, attr, src...)
	}
	dest = make(map[string]string)
	for _, s := range src {
		dest[s] = s
	}
	return
}

func DingIdReduce(dingIdReduceFn DingIdReduceFn, ctx context.Context, attr Attr, src string) (dest string) {
	if dingIdReduceFn != nil {
		dest1 := dingIdReduceFn(ctx, attr, src)
		if len(dest1) > 0 {
			return dest1[src]
		}
	}
	return src
}
