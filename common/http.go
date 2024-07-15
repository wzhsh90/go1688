package common

import (
	"bytes"
	"crypto/tls"
	"github.com/valyala/fasthttp"
	"io"
	"mime/multipart"
	"time"
)

var httpClient = &fasthttp.Client{}

func Get(url string, headers map[string]string) (respDa []byte, err error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(url)
	req.Header.SetMethod(fasthttp.MethodGet)
	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	errDo := httpClient.Do(req, resp)
	if errDo != nil {
		return nil, errDo
	} else {
		body := resp.Body()
		if body != nil {
			return body, nil
		} else {
			return []byte(""), nil
		}
	}
}

func PostXml(path string, xmlString string, tlsConfig *tls.Config, timeout time.Duration) ([]byte, error) {

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(path)
	req.Header.SetMethod(fasthttp.MethodPost)
	req.Header.SetContentType("application/xml")
	req.SetBodyString(xmlString)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	httpClient.WriteTimeout = timeout
	httpClient.ReadTimeout = timeout
	if tlsConfig != nil {
		httpClient.TLSConfig = tlsConfig
	}
	errDo := httpClient.Do(req, resp)
	if errDo != nil {
		return nil, errDo
	} else {
		body := resp.Body()
		if body != nil {
			return body, nil
		} else {
			return []byte(""), nil
		}
	}
}
func PostFile(url string, data map[string]string, headers map[string]string, nameField, fileName string, file io.Reader) (respDa []byte, err error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(url)
	bodyBufer := new(bytes.Buffer)
	writer := multipart.NewWriter(bodyBufer)
	formFile, err := writer.CreateFormFile(nameField, fileName)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(formFile, file)
	if err != nil {
		return nil, err
	}
	for key, val := range data {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}
	req.Header.SetMethod(fasthttp.MethodPost)
	req.Header.SetContentType(writer.FormDataContentType())
	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}
	req.SetBody(bodyBufer.Bytes())
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	errDo := httpClient.Do(req, resp)
	if errDo != nil {
		return nil, errDo
	} else {
		body := resp.Body()
		if body != nil {
			return body, nil
		} else {
			return []byte(""), nil
		}
	}
}
func PostJson(path string, jsonData string, headers map[string]string) ([]byte, error) {

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(path)
	req.Header.SetMethod(fasthttp.MethodPost)
	req.Header.SetContentType("application/json")
	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}
	req.SetBodyString(jsonData)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	errDo := httpClient.Do(req, resp)
	if errDo != nil {
		return nil, errDo
	} else {
		body := resp.Body()
		if body != nil {
			return body, nil
		} else {
			return []byte(""), nil
		}
	}
}

func PostForm(url string, data map[string]string, headers map[string]string) (respDa []byte, err error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(url)
	args := req.PostArgs()
	for k, v := range data {
		args.Add(k, v)
	}
	req.Header.SetMethod(fasthttp.MethodPost)
	req.Header.SetContentType("application/x-www-form-urlencoded")
	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	errDo := httpClient.Do(req, resp)
	if errDo != nil {
		return nil, errDo
	} else {
		body := resp.Body()
		if body != nil {
			return body, nil
		} else {
			return []byte(""), nil
		}
	}
}
