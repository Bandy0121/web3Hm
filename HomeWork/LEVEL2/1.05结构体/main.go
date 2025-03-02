package main

import "sync"

type Person struct {
	Name  string
	Age   int
	Call  func() byte
	Map   map[string]string
	Ch    chan string
	Arr   [32]uint8
	Slice []interface{}
	Ptr   *int
	once  sync.Once
}

type Other struct{}

type Person1 struct {
	Name  string            `json:"name" gorm:"column:<name>"`
	Age   int               `json:"age" gorm:"column:<name>"`
	Call  func()            `json:"-" gorm:"column:<name>"`
	Map   map[string]string `json:"map" gorm:"column:<name>"`
	Ch    chan string       `json:"-" gorm:"column:<name>"`
	Arr   [32]uint8         `json:"arr" gorm:"column:<name>"`
	Slice []interface{}     `json:"slice" gorm:"column:<name>"`
	Ptr   *int              `json:"-"`
	O     Other             `json:"-"`
}
type Custom struct {
	int
	string
	Other string
}
type A struct {
	a string
}

type B struct {
	A
	b string
}

type C struct {
	A
	B
	a string
	b string
	c string
}
