package selector

import (
	"context"
	"github.com/signmem/falcon-plus/common/redisdb"
	"github.com/signmem/raftleader/g"
	"github.com/signmem/redislock"
	"time"
)

var (
	Role string
	lock *redislock.Lock
)


func Start() {

	config := g.Config().Redis
	lock_key := config.LockKey
	lock_time := config.LockTime
	server := config.Server + ":" + config.Port
	maxActive := config.MaxActive
	maxIdle := config.MaxIdle
	ask_Lock_Time := config.AskLockTime
	Role = "slave"

	redisClient, err := redisdb.RedisClient(maxIdle, maxActive, server)

	if err != nil {
		g.Logger.Errorf("connect redis error:%s", err)
		return
	}

	ctx := context.Background()
	defer redisClient.Close()
	locker := redislock.New(redisClient)

	for {

		if Role == "slave" {

			lock, err = redisdb.LockRedis(locker, ctx, lock_key, lock_time)

			if err != nil {
				time.Sleep(time.Duration(ask_Lock_Time) * time.Second)
				Role = "slave"
				continue
			}
			Role = "master"

		} else {
			err = lock.Refresh(ctx, time.Duration(lock_time) * time.Second, nil)
			if err != nil {
				g.Logger.Errorf("reflesh lock error:%s", err)
				Role = "slave"
				continue
			}
			time.Sleep( 1 * time.Second)
		}
	}
}

func RoleCheck() {
	roleNow := "slave"

	for {

		if roleNow != Role {
			g.Logger.Debugf("role change, now is: %s", Role)

		} else {
			t := time.Now().Unix()
			if t % 120 == 0 {
				g.Logger.Debugf("role now: %s", Role)
			}
		}
		roleNow = Role
		time.Sleep( 1 * time.Second)
	}

}