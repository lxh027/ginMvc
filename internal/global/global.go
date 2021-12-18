package global

import (
	"gorm.io/gorm"
	"mvc/pkg/config"
	"mvc/pkg/db"
	"mvc/pkg/redis"
	"path"
)

var DatabaseCfg config.Database
var LogCfg config.Log
var RedisCfg config.Redis
var ServerCfg config.Server
var SessionCfg config.Session

var MysqlClient *gorm.DB
var RedisClient *redis.Client

func Init(env string) {
	loadCfg(env)
	initMysql(&DatabaseCfg)
	initRedis(&RedisCfg)
}

func loadCfg(env string) {
	basePath := path.Join("./config", env)

	databasePath := path.Join(basePath, "database.yml")
	DatabaseCfg = config.NewDatabaseConfig(databasePath)

	logPath := path.Join(basePath, "log.yml")
	LogCfg = config.NewLogConfig(logPath)

	redisPath := path.Join(basePath, "redis.yml")
	RedisCfg = config.NewRedisConfig(redisPath)

	serverPath := path.Join(basePath, "server.yml")
	ServerCfg = config.NewServerConfig(serverPath)

	sessionPath := path.Join(basePath, "session.yml")
	SessionCfg = config.NewSessionConfig(sessionPath)
}

func initMysql(cfg *config.Database) {
	MysqlClient = db.NewMysqlConn(cfg)
}

func initRedis(cfg *config.Redis) {
	RedisClient = redis.NewRedisClient(cfg)
}
