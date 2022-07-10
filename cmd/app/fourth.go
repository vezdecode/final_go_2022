package main

import "vk/internal/http"

func main() {
	app := http.GetRoute()
	_ = app.Run("0.0.0.0:8000")
}
