package handlers

import (
	"fmt"
	"io"
	"net/http"
)

const (
	expoSendAPI         = "https://exp.host/--/api/v2/push/send"
	expoGetPushTokenAPI = "https://exp.host/--/api/v2/push/getExpoPushToken"
	expoUpdateTokenAPI  = "https://exp.host/--/api/v2/push/updateDeviceToken"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Home Page!")
}

func PushHandler(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}

	fmt.Print("im in push handler")

	req, err := http.NewRequest(r.Method, expoSendAPI, r.Body)
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	for name, values := range r.Header {
		for _, value := range values {
			req.Header.Add(name, value)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to contact API", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()
	fmt.Print("im after request")

	for name, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(name, value)
		}
	}

	w.WriteHeader(resp.StatusCode)
	fmt.Print("THE END1...")

	io.Copy(w, resp.Body)
	fmt.Print("THE END2...")

	fmt.Fprintf(w, "Push recieved")

	fmt.Print("THE END3...")

}

func createProxyHandler(targetURL string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		client := &http.Client{}

		req, err := http.NewRequest(r.Method, targetURL, r.Body)
		if err != nil {
			http.Error(w, "Failed to create request", http.StatusInternalServerError)
			return
		}

		for name, values := range r.Header {
			for _, value := range values {
				req.Header.Add(name, value)
			}
		}

		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, "Failed to contact API", http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		for name, values := range resp.Header {
			for _, value := range values {
				w.Header().Add(name, value)
			}
		}

		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)
	}
}

var SendHandler = createProxyHandler(expoSendAPI)
var GetTokenHandler = createProxyHandler(expoGetPushTokenAPI)
var UpdateTokenHandler = createProxyHandler(expoUpdateTokenAPI)
