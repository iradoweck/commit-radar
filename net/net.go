package net

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

const DefaultRequestTimeout = 30 * time.Second

type Requester func(req *http.Request) ([]byte, error)

type Wrapper func(Requester) Requester

func Compose(wrappers ...Wrapper) Wrapper {
	return func(r Requester) Requester {
		for _, wrapper := range wrappers {
			r = wrapper(r)
		}
		return r
	}
}

func TokenAuth(token string) Wrapper {
	return func(r Requester) Requester {
		return func(req *http.Request) ([]byte, error) {
			req.Header.Add("Authorization", fmt.Sprintf("bearer %s", token))
			return r(req)
		}
	}
}

func MakeRequester(client *http.Client) Requester {
	return func(req *http.Request) ([]byte, error) {
		ctx, cancel := context.WithTimeout(req.Context(), DefaultRequestTimeout)
		defer cancel()

		reqWithCtx := req.WithContext(ctx)
		resp, err := client.Do(reqWithCtx)
		if err != nil {
			return []byte{}, err
		}
		defer resp.Body.Close()

		bodyText, err := io.ReadAll(resp.Body)
		if err != nil {
			return []byte{}, err
		}

		return bodyText, nil
	}
}
