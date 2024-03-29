package main

import (
	"github.com/gofiber/fiber/v2"

)

type Proxy interface{
	Accept(key string) bool
	Proxy (c *fiber.Ctx) error
}

var Procies = []Proxy{
	
}



func ProxyHandler(c *fiber.Ctx)error {
	for _,v:= range Procies{
		if v.Accept(c.Params("key")){
			return v.Proxy(c)
		}
	}
	c.Response().SetStatusCode(404)
	return nil

}