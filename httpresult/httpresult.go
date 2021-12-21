package httpresult

import (
	"encoding/json"
	"github.com/debudda/option"
	"io"
	"net/http"
)

type response[T any] struct {
	StatusCode int
	Body       T
	Header     http.Header
}

type SimpleResponse struct {
	response[[]byte]
}

type SimpleJSONResponse[T any] struct {
	response[T]
}

func Response(res *http.Response, err error) option.Result[*SimpleResponse] {
	if err != nil {
		return option.Err[*SimpleResponse](err)
	}
	raw, err := io.ReadAll(res.Body)
	if err != nil {
		return option.Err[*SimpleResponse](err)
	}
	defer res.Body.Close()
	return option.Ok(&SimpleResponse{response: response[[]byte]{
		StatusCode: res.StatusCode,
		Body:       raw,
		Header:     res.Header.Clone(),
	}})
}

func ResponseJSON[T any](res *http.Response, err error) option.Result[*SimpleJSONResponse[T]] {
	if err != nil {
		return option.Err[*SimpleJSONResponse[T]](err)
	}
	raw, err := io.ReadAll(res.Body)
	if err != nil {
		return option.Err[*SimpleJSONResponse[T]](err)
	}
	defer res.Body.Close()
	var v T
	if err := json.Unmarshal(raw, &v); err != nil {
		return option.Err[*SimpleJSONResponse[T]](err)
	}
	return option.Ok(&SimpleJSONResponse[T]{response: response[T]{
		StatusCode: res.StatusCode,
		Body:       v,
		Header:     res.Header.Clone(),
	}})
}

func Body(res *http.Response, err error) option.Result[[]byte] {
	if err != nil {
		return option.Err[[]byte](err)
	}
	raw, err := io.ReadAll(res.Body)
	if err != nil {
		return option.Err[[]byte](err)
	}
	res.Body.Close()
	return option.Ok(raw)
}

func BodyCode(res *http.Response, err error) (option.Result[[]byte], int) {
	if err != nil {
		return option.Err[[]byte](err), 0
	}
	return Body(res, err), res.StatusCode
}

func JSON[T any](res *http.Response, err error) option.Result[T] {
	if err != nil {
		return option.Err[T](err)
	}
	raw, err := io.ReadAll(res.Body)
	if err != nil {
		return option.Err[T](err)
	}
	defer res.Body.Close()
	var v T
	if err := json.Unmarshal(raw, &v); err != nil {
		return option.Err[T](err)
	}
	return option.Ok(v)
}

func JSONCode[T any](res *http.Response, err error) (option.Result[T], int) {
	if err != nil {
		return option.Err[T](err), 0
	}
	return JSON[T](res, err), res.StatusCode
}
