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

	r := NewRouter(feedClient, postClient)
	r.Run(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
}
