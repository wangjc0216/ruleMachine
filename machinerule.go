package rulemachine

import (
	"context"
	errlog "github.com/wangjc0216/errlog/v2"
)

func Execute(ctx context.Context) {
	errlog.Errorf(context.TODO(), "world is not good,but i am fine")
}
