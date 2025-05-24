package main

type IRegistry interface {
	Add(svc Service)
	Remove(serviceID string)
	Get(name string) []Service
}
