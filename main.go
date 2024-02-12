package main

import "github.com/tabed23/user-boilerplate-auth/routes"



func main() {
	routes.Routes()
	routes.Run("0.0.0.0:8080")
}
