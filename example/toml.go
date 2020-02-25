package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

// Config config of pipelone
type Config struct {
	Input     ConfigItem            `toml:"input"`
	Processor map[string]ConfigItem `toml:"processor"`
	Output    map[string]ConfigItem `toml:"output"`
}

// ConfigItem config of each layer
type ConfigItem struct {
	Name   string         `toml:"type"`
	Config toml.Primitive `toml:"config"`
}

type HttpConfig struct {
	Route     string `toml:"route"`
	QueueSize int    `toml:"queueSize"`
}

func DecodeConfig(md toml.MetaData, primValue toml.Primitive) (c interface{}, err error) {
	c = make(map[string]HttpConfig)
	if err = md.PrimitiveDecode(primValue, c); err != nil {
		return nil, err
	}
	// 类型转换
	m := c.(map[string]HttpConfig)
	return mapToSlice(m), nil
}

func mapToSlice(m map[string]HttpConfig) []HttpConfig {
	s := make([]HttpConfig, 0, len(m))
	for _, v := range m {
		s = append(s, v)
	}
	return s
}

func main() {
	content := `
	[input]
	type="http"
	[input.config.1]
	route="/realtime"
	queueSize=100
	[input.config.2]
	route="/unrealtime"
	queueSize=200
	`
	c := new(Config)
	md, err := toml.Decode(content, c)
	if err != nil {
		fmt.Println("decode error")
	}

	hc, err := DecodeConfig(md, c.Input.Config)
	if err != nil {
		fmt.Println("decode http error")
	}

	fmt.Println(hc)
}
