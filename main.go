package main

import (
	core_api "github.com/Redmoon-2333/innospark-idl-personal/kitex_gen/core_api/coreapi"
	"log"
)

func main() {
	svr := core_api.NewServer(new(CoreApiImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
