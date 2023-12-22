package cmd

import (
	"context"

	"github.com/gogf/gf/v2/os/gcron"
)

func UserOrderDefaultComments(ctx context.Context) (err error) {
	//每分钟执行一次
	_, err = gcron.Add(ctx, "0 */1 * * * *", func(ctx context.Context) {

		return
	}, "UserOrderDefaultComments")
	return err
}
