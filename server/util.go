package server

import (
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"time"
)

type Expirable interface {
	CreatedAt() int64
	ExpiresIn() int
}

type Cache interface {
	Set(exp Expirable) error
	Get(exp Expirable) error
}

type FileCache struct {
	Path string
}

func NewFileCache(path string) *FileCache {
	return &FileCache{
		Path: path,
	}
}

func (c *FileCache) Set(exp Expirable) error {
	data, err := json.Marshal(exp)
	if err != nil {
		return err
	}
	if err = ioutil.WriteFile(c.Path, data, 0644); err != nil {
		return err
	}
	return nil
}

func (c *FileCache) Get(exp Expirable) error {
	data, err := ioutil.ReadFile(c.Path)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(data, exp); err != nil {
		return err
	}
	created := exp.CreatedAt()
	expires := exp.ExpiresIn()
	if time.Now().Unix() > created+int64(expires-60) {
		return errors.New("Data is already expired")
	}
	return nil
}

type InMemoryCache struct {
	data []byte
}

func NewInMemoryCache() *InMemoryCache {
	return &InMemoryCache{}
}

func (c *InMemoryCache) Set(exp Expirable) error {
	data, err := json.Marshal(exp)
	if err != nil {
		return err
	}
	c.data = data
	return nil
}

func (c *InMemoryCache) Get(exp Expirable) error {
	if err := json.Unmarshal(c.data, exp); err != nil {
		return err
	}
	created := exp.CreatedAt()
	expires := exp.ExpiresIn()
	if time.Now().Unix() > created+int64(expires-60) {
		return errors.New("Data is already expired")
	}
	return nil
}

func sha1Sign(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}
