package utils

import (
	"context"
	"strconv"
	"time"

	"github.com/simplecolding/douyin/hertz-server/biz/redis"
)

// return 鉴权,用户名,uid
func Auth(ctx context.Context, tk string) (bool, string, int64, string) {
	redisCtx, cancel := context.WithTimeout(ctx, 500*time.Hour)
	defer cancel()
	username := redis.RD.Get(redisCtx, tk).Val()
	uidString := redis.RD.Get(redisCtx, username).Val()
	uid, _ := strconv.ParseInt(uidString, 10, 64)
	if len(username) == 0 {
		println(username)
		return false, "", 0, ""
	}
	return true, username, uid, uidString
}
