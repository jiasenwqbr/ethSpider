package levelDB

import (
	"ethSpider/logger"
	logging "github.com/ipfs/go-log/v2"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"time"
)

type ORM interface {
	Register(email string, pwd string) error
	Login(email string, pwd string) error
	ResetPwd(email string, pwd string, repeatPwd string) error
	SetSession(email string, token string, duration time.Duration) error
}

type LvDB struct {
	Db              *leveldb.DB
	SessionDuration time.Duration
	Logger          logging.StandardLogger
}

func NewDB(dbUri string, logName string) (*LvDB, error) {
	db, err := leveldb.OpenFile(dbUri, nil)
	if err != nil {
		return nil, err
	}

	return &LvDB{
		Db:              db,
		SessionDuration: 60 * time.Minute,
		Logger:          logger.SetupLog(logName),
	}, nil
}

func (l *LvDB) put(key, value []byte, wo *opt.WriteOptions) error {
	err := l.Db.Put(key, value, wo)
	if err != nil {
		l.Logger.Errorf("Put data error:%v", err)
		return err
	}
	return nil
}

func (l *LvDB) get(key []byte, ro *opt.ReadOptions) ([]byte, error) {
	data, err := l.Db.Get(key, ro)
	if err != nil {
		l.Logger.Errorf("Get data error:%v", err)
		return data, err
	}
	return data, nil
}

func (l *LvDB) delete(key []byte, wo *opt.WriteOptions) error {
	err := l.Db.Delete(key, nil)
	if err != nil {
		l.Logger.Errorf("Delete key error:%v", err)
		return err
	}

	return nil
}
