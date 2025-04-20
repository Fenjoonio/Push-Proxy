package main

import (
	"io"
	"log"
	"net/http"
)

const expoAPI = "https://exp.host/--/api/v2/push/send"

func handler(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}

	// Create a new request to Expo's API
	req, err := http.NewRequest(r.Method, expoAPI, r.Body)
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	// Copy headers from the original request
	for name, values := range r.Header {
		for _, value := range values {
			req.Header.Add(name, value)
		}
	}

	// Make the request to Expo API
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to contact Expo API", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	// Copy the response from Expo to the client
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func main() {
	http.HandleFunc("/push", handler)
	log.Println("Proxy server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
