package resty

import(
	"fmt"
	"sync"
	"encoding/json"

	"application/logger"
	resty "github.com/go-resty/resty/v2"
)

var client RestyClient = nil
var lock *sync.Mutex = &sync.Mutex{}

type RestyClient interface {
	GetClient() *resty.Client
	Send(req *resty.Request, url string, reqType requestType) (*resty.Response, error) 
	CheckResponse(restyRes *resty.Response, err error, expectedStatusCode int, serviceName serviceProvider) ([]byte, error)
}

type restyClient struct {
	client *resty.Client
}

var CreateRestyClient = func() {
	if client == nil {
		lock.Lock()
		defer lock.Unlock()
		if client == nil {
			client = &restyClient{ client : resty.New() }
		}
	}
}

var GetRestyClient = func() RestyClient {
	return client
}

func (rc *restyClient) GetClient() *resty.Client {
	return rc.client
}

func (rc *restyClient) Send(req *resty.Request, url string, reqType requestType) (*resty.Response, error) {
	var res *resty.Response
	var err error
	switch reqType {
	case GET :
		res, err = req.Get(url)
	case PUT:
		res, err = req.Put(url)
	case POST:
		res, err = req.Post(url)
	case PATCH:
		res, err = req.Patch(url)
	case DELETE:
		res, err = req.Delete(url)
	default:
		logger.ThrowDebugLog("Request type is not supported")
	}

	return res, err
} 

func (rc *restyClient) CheckResponse(restyRes *resty.Response, err error, expectedStatusCode int, serviceName serviceProvider) ([]byte, error) {
	if err != nil {
		logger.ThrowErrorLog(err)
		return nil, err
	}

	if restyRes.StatusCode() != expectedStatusCode {
		logger.ThrowDebugLog("Did not find the expected status code!")
		var error interface{}
		if err := json.Unmarshal(restyRes.Body(), &error); err != nil {
			logger.ThrowDebugLog("error in unmarshalling the response")
			return nil, err
		}
		return nil, fmt.Errorf("The expected status code was not seen.")
	}

	return restyRes.Body(), nil
}