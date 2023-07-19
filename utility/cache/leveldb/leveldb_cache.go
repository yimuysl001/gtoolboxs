package leveldb_cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/yimuysl001/gtoolboxs/utility/logger"
	"time"
)

type AdapterLeveldb struct {
	db *leveldb.DB
}
type fileContent struct {
	Duration int64       `json:"duration"`
	Data     interface{} `json:"data,omitempty"`
}

func (a *AdapterLeveldb) Set(ctx context.Context, key interface{}, value interface{}, duration time.Duration) error {
	if value == nil || duration < 0 {
		return a.db.Delete(gconv.Bytes(key), nil)
	}

	return a.Save(key, value, duration)

}

func (a *AdapterLeveldb) SetMap(ctx context.Context, data map[interface{}]interface{}, duration time.Duration) error {
	//TODO implement me
	panic("implement me")
}

func (a *AdapterLeveldb) SetIfNotExist(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (ok bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (a *AdapterLeveldb) SetIfNotExistFunc(ctx context.Context, key interface{}, f gcache.Func, duration time.Duration) (ok bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (a *AdapterLeveldb) SetIfNotExistFuncLock(ctx context.Context, key interface{}, f gcache.Func, duration time.Duration) (ok bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (a *AdapterLeveldb) Get(ctx context.Context, key interface{}) (*gvar.Var, error) {
	fetch, err := a.read(gconv.Bytes(key))
	if err != nil {
		return nil, err
	}
	return gvar.New(fetch.Data), nil
}

func (a *AdapterLeveldb) GetOrSet(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (result *gvar.Var, err error) {
	//TODO implement me
	panic("implement me")
}

func (a *AdapterLeveldb) GetOrSetFunc(ctx context.Context, key interface{}, f gcache.Func, duration time.Duration) (result *gvar.Var, err error) {
	//TODO implement me
	panic("implement me")
}

func (a *AdapterLeveldb) GetOrSetFuncLock(ctx context.Context, key interface{}, f gcache.Func, duration time.Duration) (result *gvar.Var, err error) {
	//TODO implement me
	panic("implement me")
}

func (a *AdapterLeveldb) Contains(ctx context.Context, key interface{}) (bool, error) {
	_, err := a.read(gconv.Bytes(key))
	return err == nil, err

}

func (a *AdapterLeveldb) Size(ctx context.Context) (size int, err error) {
	//TODO implement me
	panic("implement me")
}

func (a *AdapterLeveldb) Data(ctx context.Context) (data map[interface{}]interface{}, err error) {
	//TODO implement me
	panic("implement me")
}

func (a *AdapterLeveldb) Keys(ctx context.Context) (keys []interface{}, err error) {
	//TODO implement me
	panic("implement me")
}

func (a *AdapterLeveldb) Values(ctx context.Context) (values []interface{}, err error) {
	//TODO implement me
	panic("implement me")
}

func (a *AdapterLeveldb) Update(ctx context.Context, key interface{}, value interface{}) (oldValue *gvar.Var, exist bool, err error) {

	exist, err = a.Contains(ctx, key)

	oldValue, err = a.Get(ctx, gconv.Bytes(key))

	return
}

func (a *AdapterLeveldb) UpdateExpire(ctx context.Context, key interface{}, duration time.Duration) (oldDuration time.Duration, err error) {
	var (
		v       *gvar.Var
		oldTTL  int64
		fileKey = gconv.Bytes(key)
	)
	// TTL.
	expire, err := a.GetExpire(ctx, fileKey)
	if err != nil {
		return
	}
	oldTTL = int64(expire)
	if oldTTL == -2 {
		// It does not exist.
		oldTTL = -1
		return
	}
	oldDuration = time.Duration(oldTTL) * time.Second
	// DEL.
	if duration < 0 {
		err = a.db.Delete(fileKey, nil)
		return
	}
	v, err = a.Get(ctx, fileKey)
	if err != nil {
		return
	}
	err = a.Set(ctx, fileKey, v.Val(), duration)

	return
}

func (a *AdapterLeveldb) GetExpire(ctx context.Context, key interface{}) (time.Duration, error) {
	content, err := a.read(gconv.Bytes(key))
	if err != nil {
		return -1, nil
	}

	if content.Duration <= time.Now().Unix() {
		return -1, nil
	}

	return time.Duration(time.Now().Unix()-content.Duration) * time.Second, nil
}

func (a *AdapterLeveldb) Remove(ctx context.Context, keys ...interface{}) (lastValue *gvar.Var, err error) {
	if len(keys) == 0 {
		return nil, nil
	}
	// Retrieves the last key value.
	if lastValue, err = a.Get(ctx, gconv.Bytes(keys[len(keys)-1])); err != nil {
		return nil, err
	}

	for key := range keys {
		err = a.db.Delete(gconv.Bytes(key), nil)
		if err != nil {
			break
		}
	}

	return

}

func (a *AdapterLeveldb) Clear(ctx context.Context) error {
	//TODO implement me
	iter := a.db.NewIterator(nil, nil)
	defer iter.Release()
	for iter.Next() {
		key := iter.Key()
		// 删除数据
		err := a.db.Delete(key, nil)
		if err != nil {
			//log.Printf("删除数据失败：%v", err)
			logger.Logger.Error(ctx, fmt.Sprintf("删除数据失败：%v", err))
			continue
		}
		logger.Logger.TraceCtx(ctx, fmt.Sprintf("已删除数据：%s\n", key))
	}

	return iter.Error()

}

func (a *AdapterLeveldb) Close(ctx context.Context) error {
	return a.db.Close()
}

// NewAdapterFile creates and returns a new memory cache object.
func NewAdapterFile(dir string) gcache.Adapter {
	db, err := leveldb.OpenFile(dir, nil)
	logger.Logger.PanicErrorCtx(context.Background(), err)
	return &AdapterLeveldb{
		db: db,
	}
}

// Save a value in File storage by key
func (a *AdapterLeveldb) Save(key interface{}, value interface{}, lifeTime time.Duration) error {
	duration := int64(0)

	if lifeTime > 0 {
		duration = time.Now().Unix() + int64(lifeTime.Seconds())
	}

	content := &fileContent{duration, value}

	data, err := json.Marshal(content)
	if err != nil {
		return err
	}

	return a.db.Put(gconv.Bytes(key), data, nil)
}

func (a *AdapterLeveldb) read(key []byte) (*fileContent, error) {
	value, err := a.db.Get(key, nil)
	if err != nil {
		return nil, err
	}

	content := &fileContent{}
	if err := json.Unmarshal(value, content); err != nil {
		return nil, err
	}

	if content.Duration == 0 {
		return content, nil
	}

	if content.Duration <= time.Now().Unix() {
		_ = a.db.Delete(key, nil) // c.Delete(key)
		//return nil, errors.New("cache expired")
		return nil, errors.New("cache expired")
	}

	return content, nil
}
