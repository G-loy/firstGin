package main

import (
	"firstGin/api/user"
	"firstGin/databases"
	"firstGin/routers"
	"fmt"
)

func main() {
	databases.Init()
	routers.Include(user.Routers)
	r := routers.Init()
	if err := r.Run(); err != nil {
		fmt.Printf("startup service failed, err:%v\n\n", err)
	}
}
