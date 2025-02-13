package main

import (
	"context"
	"errors"
)

var ErrNotFound = errors.New("not found")

func main() {

}

func Get(ctx context.Context, address, key string) (string, error) {
	return "", nil
}

func GetWithCancel(ctx context.Context, key, address string) (string, error) {
	replicas := []string{"127.0.0.1", "127.0.0.2", "127.0.0.3"}
	ch := make(chan string)
	c, cancel := context.WithCancel(ctx)

	for _, r := range replicas {
		go func() {
			result, err := Get(ctx, r, key)
			if err != nil {
				//
			} else {
				ch <- result
				cancel()
			}
		}()
	}

	select {
	case <-c.Done():
		return "", c.Err()
	case val := <-ch:
		return val, nil
		// default:
		// 	return "", nil
	}

}
