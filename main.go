package main

import "blog/routers"

func main() {
	r := routers.Router()
	r.Spin()
}
