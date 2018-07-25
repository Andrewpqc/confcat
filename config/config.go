package config

import (
	"io/ioutil"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
	"log"
	"errors"
	"strconv"
)

type Config struct {
	filename         string
	latestReloadTime int64
	conf             map[string]string
	rwLock           sync.RWMutex
}

func init() {
	log.SetFlags(log.Ldate | log.Lshortfile)
}

func (c *Config) parseFile(filename string) error {
	c.rwLock.Lock()
	defer c.rwLock.Unlock()

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("err %v\n", err)
		return err
	}

	for _, line := range strings.Split(string(content), "\n") {
		if line != "" && !strings.HasPrefix(line, "#") {
			kv := strings.Split(line, "=")
			c.conf[strings.TrimSpace(string(kv[0]))] = strings.TrimSpace(string(kv[1]))
		}
	}
	c.latestReloadTime = time.Now().Unix()
	return nil
}

func (c *Config) reload() {
	ticker := time.NewTicker(3 * time.Second)
	for _ = range ticker.C {
		fileInfo, err := os.Stat(c.filename)
		if err != nil {
			log.Fatalf("reload error")
		}
		modTime := fileInfo.ModTime().Unix()
		if modTime > c.latestReloadTime {
			fmt.Printf("detect config file:%s changed,"+
				"reloading...\n", c.filename)
			c.parseFile(c.filename)
			fmt.Printf("reload successed!\n")
		}
	}
}

func (c *Config) GetString(key string, defaultVal string) (string) {
	if c.conf[key] == "" {
		return defaultVal
	}
	return c.conf[key]
}

func (c *Config) GetInt(key string, defaultVal int) (int) {
	if c.conf[key] == "" {
		return defaultVal
	}

	result, err := strconv.Atoi(c.conf[key])
	if err != nil {
		log.Println(err)
	}
	return result
}

func (c *Config) GetFloat(key string, defaultVal float64) (float64) {
	if c.conf[key] == "" {
		return defaultVal
	}
	ft, err := strconv.ParseFloat(c.conf[key], 64)
	if err != nil {
		log.Println(err)
	}
	return ft
}

func NewConfig(filename string) (Config, error) {
	var c Config
	c.conf = make(map[string]string, 512)
	c.filename = filename
	if err := c.parseFile(c.filename); err != nil {
		return c, errors.New("new config error")
	}
	go c.reload()
	return c, nil
}
