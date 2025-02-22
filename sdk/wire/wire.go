//go:build wireinject
// +build wireinject

package main

import (
	"fmt"

	"github.com/google/wire"
)

type (
	LevelName    string
	MonitorName  string
	MonitorBlood int
	MapSize      int
)

// 关卡
type Level struct {
	Name    LevelName
	Monster Monster
	Map     Map
}

func (l *Level) Start() {
	fmt.Printf("关卡: %s, 地图大小: %d, 野怪: %s 血量为: %d \n", l.Name, l.Map.Size, l.Monster.Name, l.Monster.Blood)
}

func NewLevel(name LevelName, monster Monster, m Map) Level {
	return Level{
		Name:    name,
		Monster: monster,
		Map:     m,
	}
}

// 野怪
type Monster struct {
	Name  MonitorName
	Blood MonitorBlood
}

func NewMonster(name MonitorName, blood MonitorBlood) Monster {
	return Monster{Name: name, Blood: blood}
}

// 地图
type Map struct {
	Size MapSize
}

func NewMap(size MapSize) Map {
	return Map{
		Size: size,
	}
}

func main() {
	// 关卡需要由野怪与地图构成

	// 不使用wire时，需要组装这几个对象关系
	m := NewMap(100)
	monster := NewMonster("史莱姆", 10)
	level := NewLevel("关卡1", monster, m)
	level.Start()
}

// 使用 wire进行初始化关卡
// 它是根据类型进行匹配注入参数的，如果两个入参类型一致，那么两个函数都会注入该参数
func InitLevel(name MonitorName, blood MonitorBlood, size MapSize, levelName LevelName) Level {
	wire.Build(NewMonster, NewMap, NewLevel)
	return Level{}
}
