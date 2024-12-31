package main

type Stat int

const (
	Str Stat = iota
	Dex
	Con
	Int
	Wis
	Cha
)

type Item struct {
	Name string
	Desc string
	Stat Stat
}
