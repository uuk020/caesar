package global

import (
	"caesar/config"
	"github.com/go-redis/redis"

	"github.com/jmoiron/sqlx"
)

var (
	Setting config.Server
	DB      *sqlx.DB
	Redis   *redis.Client
)
