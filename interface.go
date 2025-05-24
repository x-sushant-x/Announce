package main

type IRegistry interface {
	Add(svc Service)
	Remove(name, id string)
	Get(name string) []Service
}
