package utils

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"ligang/sysinit"
	"time"

	"github.com/astaxie/beego/cache"
	// redis driver
	_ "github.com/astaxie/beego/cache/redis"
)

var cc cache.Cache

// InitCache InitCache
func InitCache() {
	fmt.Println("Init Redis")
	settings := sysinit.InitDefaultSettings()
	// host := beego.AppConfig.String("cache::redis_host")
	host := settings.RedisHost
	//passWord := beego.AppConfig.String("cache::redis_password")
	var err error
	defer func() {
		if r := recover(); r != nil {
			cc = nil
		}
	}()
	//cc, err = cache.NewCache("redis", `{"conn":"`+host+`","password":"`+passWord+`"}`)
	cc, err = cache.NewCache("redis", `{"conn":"`+host+`"}`)
	if err != nil {
		LogError("Connect to the redis host " + host + " failed")
		LogError(err)
	}
}

// SetCache SetCache
func SetCache(key string, value interface{}, timeout int) error {
	data, err := Encode(value)
	if err != nil {
		return err
	}
	if cc == nil {
		return errors.New("cc is nil")
	}

	defer func() {
		if r := recover(); r != nil {
			LogError(r)
			cc = nil
		}
	}()
	timeouts := time.Duration(timeout) * time.Second
	err = cc.Put(key, data, timeouts)
	if err != nil {
		LogError(err)
		LogError("SetCache失敗，key:" + key)
	}
	return err
}

// GetCache GetCache
func GetCache(key string, to interface{}) error {
	if cc == nil {
		return errors.New("cc is nil")
	}

	defer func() {
		if r := recover(); r != nil {
			LogError(r)
			cc = nil
		}
	}()

	data := cc.Get(key)
	if data == nil {
		return errors.New("Cache不存在")
	}

	err := Decode(data.([]byte), to)
	if err != nil {
		LogError(err)
		LogError("GetCache失敗，key:" + key)
	}

	return err
}

// DelCache DelCache
func DelCache(key string) error {
	if cc == nil {
		return errors.New("cc is nil")
	}
	defer func() {
		if r := recover(); r != nil {
			//fmt.Println("get cache error caught: %v\n", r)
			cc = nil
		}
	}()
	err := cc.Delete(key)
	if err != nil {
		return errors.New("Cache刪除失敗" + err.Error())
	}
	return nil
}

// Encode Encode
func Encode(data interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Decode Decode
func Decode(data []byte, to interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(to)
}
