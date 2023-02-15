package httptools

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"tools-go/httptools/constant"
)

// SendGet GET HTTP
func SendGet(targetUrl string, payload url.Values) (string, error) {

	uri, err := url.ParseRequestURI(targetUrl)
	if err != nil {
		return "", err
	}
	uri.RawQuery = payload.Encode()
	resp, err := http.Get(uri.String())
	if err != nil {
		return "", err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(respBody), nil
}

// SendGetWithHeader GET HTTP WITH HEADER
func SendGetWithHeader(targetUrl string, payload url.Values, header map[string]string, client *http.Client) (
	string, error) {

	uri, err := url.ParseRequestURI(targetUrl)
	if err != nil {
		return "", err
	}
	uri.RawQuery = payload.Encode()
	return sendPostWithHeader(http.MethodGet, uri.String(), nil, header, client)
}

// SendPostForm form post
func SendPostForm(targetUrl string, payload url.Values) (string, error) {

	resp, err := http.PostForm(targetUrl, payload)
	if err != nil {
		return "", err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(resp.Body)

	all, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return string(all), nil
}

// SendPostStringWithHeader string with header
func SendPostStringWithHeader(targetUrl string, payload string, header map[string]string,
	client *http.Client) (string, error) {

	header["Content-Type"] = constant.Form
	body := strings.NewReader(payload)
	return sendPostWithHeader(http.MethodPost, targetUrl, body, header, client)
}

// SendPostFormWithHeader  form urlValues with header
func SendPostFormWithHeader(targetUrl string, payload url.Values, header map[string]string,
	client *http.Client) (string, error) {

	header["Content-Type"] = constant.Form
	body := strings.NewReader(payload.Encode())
	return sendPostWithHeader(http.MethodPost, targetUrl, body, header, client)
}

// SendPostJsonWithHeader json with header
func SendPostJsonWithHeader(targetUrl string, payload map[string]interface{}, header map[string]string,
	client *http.Client) (string, error) {

	header["Content-Type"] = constant.Json
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	body := bytes.NewReader(payloadBytes)
	return sendPostWithHeader(http.MethodPost, targetUrl, body, header, client)
}

// sendPostWithHeader send post request with header
func sendPostWithHeader(method string, targetUrl string, body io.Reader, header map[string]string,
	client *http.Client) (string, error) {

	// build request
	// req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, "http://", bodyReader)
	req, err := http.NewRequest(method, targetUrl, body)
	if err != nil {
		return "", err
	}

	// set header
	for key, value := range header {
		req.Header.Set(key, value)
	}

	var resp *http.Response
	if client == nil {
		// httptools := &http.Client{Timeout: 5 * time.Second}
		client = http.DefaultClient
	}

	// start
	resp, err = client.Do(req)
	if err != nil {
		return "", err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(resp.Body)

	all, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return string(all), nil
}

//// SendPostSimpleParamString  param build example: "key=value&id=123"
//func SendPostSimpleParamString(targetUrl string, payload string, contentType string) (string, error) {
//	return sendPostSimple(targetUrl, strings.NewReader(payload), contentType)
//}
//
//// SendPostSimpleParamUrlValues param build example: payload = url.Values{"name": {"han"}, "age": {"18"}}
//func SendPostSimpleParamUrlValues(targetUrl string, payload url.Values, contentType string) (string, error) {
//	reqData := payload.Encode()
//	return sendPostSimple(targetUrl, strings.NewReader(reqData), contentType)
//}
//
//// SendPostSimpleParamJson param application/json
//func SendPostSimpleParamJson(targetUrl string, payload map[string]interface{}) (string, error) {
//	payloadBytes, err := json.Marshal(payload)
//	if err != nil {
//		return "", err
//	}
//	body := bytes.NewReader(payloadBytes)
//	return sendPostSimple(targetUrl, body, constant.Json)
//}
//
//// sendPostSimple simple post
//func sendPostSimple(targetUrl string, reqBody io.Reader, contentType string) (string, error) {
//	if contentType == "" {
//		contentType = constant.Form
//	}
//
//	resp, err := http.Post(targetUrl, contentType, reqBody)
//	if err != nil {
//		return "", err
//	}
//
//	defer func(Body io.ReadCloser) {
//		err := Body.Close()
//		if err != nil {
//
//		}
//	}(resp.Body)
//
//	respBody, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		return "", err
//	}
//
//	return string(respBody), nil
//}
