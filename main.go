package main

import (
	"GoThas/model"
	"fmt"
)

func main() {
	g := &model.ArthasRequest{Command: "thread -n 1", Url: "http://localhost:8563/api"}
	g.New()
	fmt.Println(g.GetInfoByCom())

}
