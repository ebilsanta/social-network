package server

import (
	"fmt"
	"os"

	services "github.com/ebilsanta/social-network/backend/complex_services/feed_complex_service/services"
)

func Init() {
	feedClient, feedConn := services.InitFeedService()
	defer feedConn.Close()

	postClient, postConn := services.InitPostService()
	defer postConn.Close()

	userClient, userConn := services.InitUserService()
	defer userConn.Close()

	r := NewRouter(feedClient, postClient, userClient)
	r.Run(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
}
