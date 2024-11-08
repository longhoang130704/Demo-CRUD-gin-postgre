package main

import (
	"GIN-TUTORIAL/db"
	"GIN-TUTORIAL/routes"
)

func main() {
    router := routes.InitRoutes()
    router.SetTrustedProxies([]string{"192.168.1.1", "10.0.0.0/24"})
    db.ConnectDatabase()
    db.MigrateDatabase(db.DB)
    router.Run(":8080")
}



