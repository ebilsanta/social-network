package server

import (
	"fmt"
	"os"

	api "github.com/ebilsanta/social-network/backend/complex_services/user_service/services"
)

func Init() {
	userClient, conn := api.InitUserService()
	defer conn.Close()

	r := NewRouter(userClient)

	r.Run(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
}
