package common

import (
	"bytes"
	"crypto/tls"
	"errors"
	"github.com/valyala/fasthttp"
	"io"
	"mime/multipart"
	"time"
)

var httpClient = &fasthttp.Client{}

func Get(url string, headers map[string]string, handler RespHandler) error {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}()
	req.SetRequestURI(url)
	req.Header.SetMethod(fasthttp.MethodGet)
	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}
	errDo := httpClient.Do(req, resp)
	if errDo != nil {
		return errDo
	}
	body := resp.Body()
	if body != nil {
		return handler(body)
	} else {
		return errors.New("empty")
	}
}

func PostXml(path string, xmlString string, tlsConfig *tls.Config, timeout time.Duration, handler RespHandler) error {

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}()
	req.SetRequestURI(path)
	req.Header.SetMethod(fasthttp.MethodPost)
	req.Header.SetContentType("application/xml")
	req.SetBodyString(xmlString)
	httpClient.WriteTimeout = timeout
	httpClient.ReadTimeout = timeout
	if tlsConfig != nil {
		httpClient.TLSConfig = tlsConfig
	}
	errDo := httpClient.Do(req, resp)
	if errDo != nil {
		return errDo
	}
	body := resp.Body()
	if body != nil {
		return handler(body)
	} else {
		return errors.New("empty")
	}
}
func PostFile(url string, data map[string]string, headers map[string]string, nameField, fileName string, file io.Reader, handler RespHandler) error {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}()
	req.SetRequestURI(url)
	bodyBufer := new(bytes.Buffer)
	writer := multipart.NewWriter(bodyBufer)
	formFile, err := writer.CreateFormFile(nameField, fileName)
	if err != nil {
		return err
	}
	_, err = io.Copy(formFile, file)
	if err != nil {
		return err
	}
	for key, val := range data {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return err
	}
	req.Header.SetMethod(fasthttp.MethodPost)
	req.Header.SetContentType(writer.FormDataContentType())
	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}
	req.SetBody(bodyBufer.Bytes())
	errDo := httpClient.Do(req, resp)
	if errDo != nil {
		return errDo
	}
	body := resp.Body()
	if body != nil {
		return handler(body)
	} else {
		return errors.New("empty")
	}
}
func PostJson(path string, jsonData string, headers map[string]string, handler RespHandler) error {

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}()
	req.SetRequestURI(path)
	req.Header.SetMethod(fasthttp.MethodPost)
	req.Header.SetContentType("application/json")
	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}
	req.SetBodyString(jsonData)
	errDo := httpClient.Do(req, resp)
	if errDo != nil {
		return errDo
	}
	body := resp.Body()
	if body != nil {
		return handler(body)
	} else {
		return errors.New("empty")
	}
}

func PostForm(url string, data map[string]string, headers map[string]string, handler RespHandler) error {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}()
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
	errDo := httpClient.Do(req, resp)
	if errDo != nil {
		return errDo
	}
	body := resp.Body()
	if body != nil {
		return handler(body)
	} else {
		return errors.New("empty")
	}
}
