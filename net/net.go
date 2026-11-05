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

		if resp.StatusCode == http.StatusTooManyRequests || (resp.StatusCode == http.StatusForbidden && resp.Header.Get("X-RateLimit-Remaining") == "0") {
			// Backoff defensivo para Primary/Secondary Rate Limit ou 429
			backoff := 30 * time.Second
			if retryAfter := resp.Header.Get("Retry-After"); retryAfter != "" {
				if seconds, err := time.ParseDuration(retryAfter + "s"); err == nil && seconds > 0 {
					backoff = seconds
				}
			}
			time.Sleep(backoff)
			snippet := string(bodyText)
			if len(snippet) > 200 {
				snippet = snippet[:200] + "..."
			}
			return []byte{}, fmt.Errorf("github api rate limit exceeded (HTTP %d): %s", resp.StatusCode, snippet)
		}

		if resp.StatusCode >= 500 {
			// Erros transitórios de infraestrutura da GitHub (502/503/504)
			time.Sleep(10 * time.Second)
			snippet := string(bodyText)
			if len(snippet) > 200 {
				snippet = snippet[:200] + "..."
			}
			return []byte{}, fmt.Errorf("github server error (HTTP %d): %s", resp.StatusCode, snippet)
		}

		if resp.StatusCode != http.StatusOK {
			snippet := string(bodyText)
			if len(snippet) > 200 {
				snippet = snippet[:200] + "..."
			}
			return []byte{}, fmt.Errorf("unexpected http status %d: %s", resp.StatusCode, snippet)
		}

		return bodyText, nil
	}
}

