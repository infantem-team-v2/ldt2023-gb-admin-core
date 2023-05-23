package thttp

import (
	"encoding/json"
	"fmt"
	"gb-auth-gate/config"
	"gb-auth-gate/pkg/tlogger"
	"gb-auth-gate/pkg/tutils/etc"
	"github.com/sarulabs/di"
	http "github.com/valyala/fasthttp"
	"time"
)

type ThttpClient struct {
	httpClient *http.Client
	Config     *config.Config  `di:"config"`
	Logger     tlogger.ILogger `di:"logger"`
}

func BuildHttpClient(ctn di.Container) (interface{}, error) {
	cfg := ctn.Get("config").(*config.Config)
	logger := ctn.Get("logger").(tlogger.ILogger)
	httpClient := http.Client{
		Name:         "bank-client",
		ReadTimeout:  time.Duration(cfg.HttpConfig.TimeOut) * time.Second,
		WriteTimeout: time.Duration(cfg.HttpConfig.TimeOut) * time.Second,
	}

	return &ThttpClient{
		Config:     cfg,
		Logger:     logger,
		httpClient: &httpClient,
	}, nil
}

func (hc *ThttpClient) MakeQueryString(
	queryParams map[string]interface{},
) (query string) {
	var queryCount int

	for k, v := range queryParams {
		switch queryCount {
		case 0:
			query += fmt.Sprintf("?%s=%s", k, v)
		default:
			query += fmt.Sprintf("&%s=%s", k, v)
		}
		queryCount++
	}

	return query
}

// Request Basic function to send request w/ params
func (hc *ThttpClient) Request(
	method, reqUrl string,
	reqHeaders map[string]string,
	reqParams, destStruct interface{},
	queryParams map[string]interface{},
) (rawResponse interface{}, statusCode uint16, err error) {
	req := http.AcquireRequest()
	var queryString string
	if queryParams != nil {
		queryString = hc.MakeQueryString(queryParams)
	}
	req.SetRequestURI(fmt.Sprintf("%s%s", reqUrl, queryString))
	req.Header.SetMethod(method)
	for k, v := range reqHeaders {
		req.Header.Set(k, v)
	}

	reqParamsData, err := json.Marshal(reqParams)
	if err != nil {
		return nil, 0, err
	}
	req.SetBody(reqParamsData)
	var resp http.Response
	//----------------TEST---------------->//
	httpClient := http.Client{Name: "bank-client", ReadTimeout: time.Duration(10) * time.Second, WriteTimeout: time.Duration(10) * time.Second}
	err = httpClient.Do(req, &resp)
	//<----------------TEST----------------//
	//err = hc.httpClient.Do(req, &resp) //PROD
	if err != nil {
		return nil, 0, err
	}
	err = json.Unmarshal(resp.Body(), destStruct)
	if err != nil {
		return nil, 0, err
	}
	hc.logRequest(req, &resp)
	return destStruct, uint16(resp.StatusCode()), nil
}

func (hc *ThttpClient) logRequest(req *http.Request, res *http.Response) {
	logMessage := fmt.Sprintf("REQUEST:\n %s\nRESPONSE:\n %s", req.Body(), res.Body())
	statusGroup := etc.GetCodeGroup(res.StatusCode())
	switch statusGroup {
	case etc.ClientError:
		hc.Logger.Warnf(logMessage)
		break
	case etc.ServerError:
		hc.Logger.Errorf(logMessage)
		break
	default:
		hc.Logger.Infof(logMessage)
		break
	}
}
