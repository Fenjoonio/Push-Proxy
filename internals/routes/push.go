package routes

import (
	"net/http"

	"github.com/freakingeek/phoxy/internals/handlers"
)

func SetupRoutes() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/push/send", handlers.PushHandler)
	http.HandleFunc("/push/getExpoPushToken", handlers.GetTokenHandler)
	http.HandleFunc("/push/updateDeviceToken", handlers.UpdateTokenHandler)
}
