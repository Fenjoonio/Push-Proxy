package routes

import (
	"net/http"

	"github.com/freakingeek/phoxy/internals/handlers"
)

func SetupRoutes() {
	http.HandleFunc("/push/send", handlers.SendHandler)
	http.HandleFunc("/push/getExpoPushToken", handlers.GetTokenHandler)
	http.HandleFunc("/push/updateDeviceToken", handlers.UpdateTokenHandler)
}
