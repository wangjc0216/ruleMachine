package rulemachine

import (
	"context"
	errlog "github.com/wangjc0216/errlog"
)

func Execute(ctx context.Context) {
	errlog.Errorf("world is not good,but i am fine")
}
