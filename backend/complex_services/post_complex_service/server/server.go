package server

import (
	"fmt"
	"os"

	api "github.com/ebilsanta/social-network/backend/complex_services/post_service/services"
)

func Init() {
	postClient, conn := api.InitPostService()
	defer conn.Close()

	r := NewRouter(postClient)

	r.Run(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
}
