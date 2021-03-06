package main

import (
    "github.com/micro/go-log"
    "net/http"
    "github.com/micro/go-web"
    "github.com/julienschmidt/httprouter"
    _ "IhomeWeb/models"
)

func main() {
	// create new web service
	//创建1个新的web服务
    service := web.NewService(
            web.Name("go.micro.web.IhomeWeb"),
            web.Version("latest"),
            web.Address(":10086"),
    )

	// initialise service
	//服务初始化
    if err := service.Init(); err != nil {
            log.Fatal(err)
    }


    //使用路由中间件来映射页面
    rou := httprouter.New()

    rou.NotFound = http.FileServer(http.Dir("html"))


	// register html handler
	//映射前端页面
	//service.Handle("/", http.FileServer(http.Dir("html")))
	service.Handle("/", rou)



	// run service
    if err := service.Run(); err != nil {
            log.Fatal(err)
    }
}
