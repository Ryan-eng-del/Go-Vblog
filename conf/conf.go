package conf

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	gormmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"time"
	// 导入MySQL驱动
	_ "github.com/go-sql-driver/mysql"
)

func NewDefaultConfig() *Config {
	return &Config{
		App:   newDefaultApp(),
		MySQL: newDefautMySQL(),
	}
}

type Config struct {
	App   *app   `toml:"app" json:"app"`
	MySQL *mysql `toml:"mysql" json:"mysql"`
}

func (c *Config) String() string {
	jd, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	return string(jd)
}

func newDefaultApp() *app {
	return &app{
		Name: "vblog",
		HTTP: newDefaultHttp(),
	}
}

type app struct {
	// 应用名称
	Name string `toml:"name" env:"APP_NAME" json:"name"`
	HTTP *http  `toml:"http" json:"http"`
}

func newDefaultHttp() *http {
	return &http{
		Host: "localhost",
		Port: "7070",
	}
}

type http struct {
	Host string `toml:"host" json:"host" env:"HTTP_HOST"`
	Port string `toml:"port" json:"port" env:"HTTP_PORT"`
}

func (h *http) Addr() string {
	return fmt.Sprintf("%s:%s", h.Host, h.Port)
}

func newDefautMySQL() *mysql {
	return &mysql{
		Host:     "localhost",
		Port:     "3306",
		Database: "vblog",
		Username: "root",
		Password: "123456",
	}
}

type mysql struct {
	Host     string `toml:"host" json:"host" env:"MYSQL_HOST"`
	Port     string `toml:"port" json:"port" env:"MYSQL_PORT"`
	Database string `toml:"database" json:"database" env:"MYSQL_DATABASE"`
	Username string `toml:"username" json:"username" env:"MYSQL_USERNAME"`
	Password string `toml:"password" json:"password" env:"MYSQL_PASSWORD"`

	// 连接池设置
	// 最大连接数
	MaxOpenConn int `toml:"max_open_conn" json:"max_open_conn" env:"MYSQL_MAX_OPEN_CONN"`
	// 最大的最大闲置连接数
	MaxIdleConn int `toml:"max_idel_conn" json:"max_idel_conn" env:"MYSQL_MAX_IDLE_CONN"`
	// 连接的有效时间, 小于服务端的设置时间
	MaxLifeTime int `toml:"max_life_time" json:"max_life_time" env:"MYSQL_MAX_LIFE_TIME"`
	// 一个闲置的连接多久没用会被释放
	MaxIdleTime int `toml:"max_idel_time" json:"max_idel_time" env:"MYSQL_MAX_IDEL_TIME"`

	lock   sync.Mutex
	dbconn *sql.DB
	orm    *gorm.DB
}

func (m *mysql) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&multiStatements=true",
		m.Username,
		m.Password,
		m.Host,
		m.Port,
		m.Database,
	)
}

func (m *mysql) GetORMDB() *gorm.DB {
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.orm == nil {
		db, err := gorm.Open(gormmysql.Open(m.Dsn()))
		if err != nil {
			panic(err)
		}
		m.orm = db
	}

	return m.orm
}

// GetDB 数据连接 需要单列模式
func (m *mysql) GetDB() *sql.DB {
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.dbconn == nil {
		conn, err := m.getDB()
		if err != nil {
			panic(err)
		}
		m.dbconn = conn
	}

	return m.dbconn
}

// 通过MySQL配置获取一个连接池
func (m *mysql) getDB() (*sql.DB, error) {
	var err error

	db, err := sql.Open("mysql", m.Dsn())
	if err != nil {
		return nil, fmt.Errorf("connect to mysql<%s> error, %s", m.Dsn(), err.Error())
	}

	// 设置连接池参数
	db.SetMaxOpenConns(m.MaxOpenConn)
	db.SetMaxIdleConns(m.MaxIdleConn)
	if m.MaxLifeTime != 0 {
		db.SetConnMaxLifetime(time.Second * time.Duration(m.MaxLifeTime))
	}
	if m.MaxIdleConn != 0 {
		db.SetConnMaxIdleTime(time.Second * time.Duration(m.MaxIdleTime))
	}

	//通过Ping来测试当前MySQL服务是否可达
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping mysql<%s> error, %s", m.Dsn(), err.Error())
	}
	return db, nil
}
