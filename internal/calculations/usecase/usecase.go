package usecase

import (
	"encoding/json"
	"fmt"
	accountModel "gb-auth-gate/internal/account/model"
	authInterface "gb-auth-gate/internal/auth/interface"
	"gb-auth-gate/internal/calculations/model"
	"gb-auth-gate/internal/pkg/common"
	"gb-auth-gate/pkg/terrors"
	"gb-auth-gate/pkg/thttp"
	thttpHeaders "gb-auth-gate/pkg/thttp/headers"
	"gb-auth-gate/pkg/tutils/etc"
	"github.com/sarulabs/di"
)

type CalculationsUseCase struct {
	httpClient *thttp.ThttpClient    `di:"httpClient"`
	authUC     authInterface.UseCase `di:"authUC"`
}

func BuildCalculationsUseCase(ctn di.Container) (interface{}, error) {
	return &CalculationsUseCase{
		httpClient: ctn.Get("httpClient").(*thttp.ThttpClient),
		authUC:     ctn.Get("authUC").(authInterface.UseCase),
	}, nil
}

func (cu *CalculationsUseCase) sendRequestToAPI(uri, method string, params, response interface{}, queryParams map[string]interface{}) (interface{}, uint16, error) {
	service, err := cu.authUC.GetAuthServiceByName(accountModel.PdsCalcService)
	if err != nil {
		return nil, 0, err
	}

	headers, err := thttpHeaders.MakeAuthHeaders(params, service.PublicKey, service.PrivateKey, method)
	fmt.Printf("\n%+v\n", headers)
	rawResponse, statusCode, err := cu.httpClient.Request(
		method,
		fmt.Sprintf("%s%s", service.URL, uri),
		headers,
		params,
		&response,
		queryParams,
	)
	if err != nil {
		return nil, 0, terrors.Raise(err, 200005)
	}
	if statusCode != 200 {
		var commonResponse common.Response
		if err := json.Unmarshal(rawResponse, &commonResponse); err != nil {
			return nil, 0, terrors.Raise(err, 200005)
		}
		return commonResponse, statusCode, nil
	}
	return response, statusCode, nil
}

func (cu *CalculationsUseCase) BaseCalculate(params *model.BaseCalculateRequest, userId *int) (*model.BaseCalculateResponse, error) {

	var response model.BaseCalculateResponse
	_, statusCode, err := cu.sendRequestToAPI("/calc/create",
		thttp.POST,
		model.MakeCalcRequestLogic{
			UserID:  userId,
			Company: params.BasicCompanyLogic,
		},
		&response,
		nil)
	if err != nil {
		return nil, err
	}
	if etc.GetCodeGroup(int(statusCode)) != etc.Successful {
		return nil, terrors.Raise(nil, 200006)
	}
	return &response, nil

}

func (cu *CalculationsUseCase) ImprovedCalculate(params *model.BaseCalculateRequest, userId *int) (*model.ImprovedCalculateResponse, error) {
	baseResponse, err := cu.BaseCalculate(params, userId)
	if err != nil {
		return nil, err
	}
	return &model.ImprovedCalculateResponse{
		BaseCalculateResponse: *baseResponse,
		Link:                  "https://cdn.ldt2023.infantem.tech/trahni_psa.pdf",
	}, nil
}

func (cu *CalculationsUseCase) GetResult(trackerId string) (*model.BaseCalculateResponse, error) {
	var response model.BaseCalculateResponse
	_, statusCode, err := cu.sendRequestToAPI("/calc/info",
		thttp.GET,
		nil,
		&response,
		map[string]interface{}{
			"tracker_id": trackerId,
		})
	if err != nil {
		return nil, err
	}
	if etc.GetCodeGroup(int(statusCode)) != etc.Successful {
		return nil, terrors.Raise(nil, 200006)
	}

	return &response, nil
}
