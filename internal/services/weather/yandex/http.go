package yandex

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getRequest(ctx context.Context, url string, result any, f ...func(req *http.Request)) error {
	return request(ctx, http.MethodGet, url, nil, result, f...)
}

func request(ctx context.Context, method, url string, reqBody, result any, f ...func(req *http.Request)) error {
	b, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	r := bytes.NewReader(b)
	req, err := http.NewRequestWithContext(ctx, method, url, r)

	for _, fun := range f {
		fun(req)
	}

	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("status code: %v", res.StatusCode)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if result == nil {
		return nil
	}

	return json.Unmarshal(body, result)
}
