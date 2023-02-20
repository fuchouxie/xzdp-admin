package main

import (
	"xzdp-admin/app/router"
	"xzdp-admin/system/core/myGin"
)

func main() {
	myEngine := myGin.InitGin()
	router.InitRouter(myEngine)
	myGin.Start()
}
