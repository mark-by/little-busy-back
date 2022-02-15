package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/mark-by/little-busy-back/auth/internal/domain/entity"
	"github.com/mark-by/little-busy-back/auth/internal/domain/repository"
	"github.com/pkg/errors"
	"log"
	"time"
)

type Options struct {
	Host string
	Port string
	User string
}

type Session struct {
	connPool *redis.Pool
}

func NewSession(options *Options) *Session {
	return &Session{
		connPool: &redis.Pool{
			Dial: func() (redis.Conn, error) {
				conn, err := redis.DialURL(fmt.Sprintf("redis://%s:%s:%s",
					options.User,
					options.Host,
					options.Port))
				if err != nil {
					log.Fatal("fail to connect to redis: ", err)
				}
				return conn, nil
			},
			MaxIdle:   80,
			MaxActive: 12000,
		},
	}
}

func (s Session) Get(sessionID string) (*entity.Session, error) {
	conn := s.connPool.Get()
	defer conn.Close()

	userID, err := redis.Int(conn.Do("GET", s.sessionKey(sessionID)))
	if err != nil {
		if errors.Is(redis.ErrNil, err) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "fail to get authorization from redis")
	}

	return &entity.Session{
		ID:     sessionID,
		UserID: userID,
	}, nil
}

func (s Session) Delete(sessionID string) error {
	conn := s.connPool.Get()
	defer conn.Close()

	_, err := redis.Int(conn.Do("DEL", s.sessionKey(sessionID)))
	if err != nil {
		return err
	}
	return nil
}

func (s Session) Create(session *entity.Session) error {
	conn := s.connPool.Get()
	defer conn.Close()

	expirationTime := session.Expiration.Sub(time.Now()).Seconds()
	if expirationTime <= 0 {
		return errors.New("wrong expiration time")
	}

	reply, err := redis.String(conn.Do("SET", s.sessionKey(session.ID), session.UserID, "EX", expirationTime))
	if err != nil {
		return errors.Wrap(err, "fail to set authorization in redis")
	}
	if reply != "OK" {
		return errors.New("result is not ok")
	}
	return nil
}

var _ repository.Session = &Session{}
