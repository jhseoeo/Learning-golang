package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

func logic(ctx context.Context, info string) (string, error) {
	return "", nil
}

func Middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// wrap the context with stuff
		r = r.WithContext(ctx)
		handler.ServeHTTP(w, r)
	})
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	data := r.FormValue("data")
	result, err := logic(ctx, data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(result))
}

// ----------------------------------------------------------------------------------

func processResponse(rc io.ReadCloser) (string, error) {
	return "", nil
}

type ServiceCaller struct {
	client *http.Client
}

func (sc ServiceCaller) callAnotherService(ctx context.Context, data string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, "http://example.com?data="+data, nil)
	if err != nil {
		return "", err
	}
	req = req.WithContext(ctx)

	resp, err := sc.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}

	// processing response
	id, err := processResponse(resp.Body)
	return id, err
}

// ----------------------------------------------------------------------------------

func main() {
	ctx := context.Background()
	result, err := logic(ctx, "a string")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
