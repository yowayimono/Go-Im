package server

import (
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"im/config"
	"im/models"
	"log"
	"time"

	mq "github.com/Ambition6666/coderzh.github.io"
	re "github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Rediscc *re.Client
	DB      *gorm.DB
	MQ      *mq.MyRabbitMQ
	Etcd    *clientv3.Client
)

// 注册数据库引擎
func InitServer() {
	InitRedis()
	InitMysql()
	InitMQ()
	InitEtcd()
}

// --------------------------------注册服务----------------------------------
// 注册redis
func InitRedis() {
	addr := config.RedisHost + ":" + config.RedisPort
	fmt.Println(addr)
	Rediscc = re.NewClient(&re.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func InitEtcd() {
	addr := config.EtcdHost + ":" + config.EtcdPort

	// 创建etcd客户端
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{addr},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	if err != nil {
		log.Fatal(err)
	}
}

// 注册mysql
func InitMysql() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.MysqlUser, config.MysqlPwd, config.MysqlHost, config.MysqlPort, config.MysqlDbName)
	fmt.Println(dsn)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		fmt.Println(err)
	}
}

// 注册MQ
func InitMQ() {
	mq.InitMqConfig(config.RabbitMQUser, config.RabbitMQPwd, config.RabbitMQHost, config.RabbitMQPort)
	MQ = mq.NewRabbitMQSimple("cache")
}

// --------------------------------获取操作服务的句柄----------------------------------
// 得到redis句柄
func GetRedis() *re.Client {
	return Rediscc
}

// 得到mysql的句柄
func GetMysqlDB() *gorm.DB {
	return DB
}

// 得到mq句柄
func GetMQ() *mq.MyRabbitMQ {
	return MQ
}

func GetEtcd() *clientv3.Client {
	return Etcd
}

// --------------------------------初始化表----------------------------------
func RForm() {
	DB.AutoMigrate(&models.Friend_application_list{})
	DB.AutoMigrate(&models.Hail_fellow{})
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Message{})
}
