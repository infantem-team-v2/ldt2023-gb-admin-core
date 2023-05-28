package usecase

import (
	"encoding/json"
	"fmt"
	authInterface "gb-admin-core/internal/auth/interface"
	"gb-admin-core/internal/pkg/common"
	"gb-admin-core/internal/ui/model"
	"gb-admin-core/pkg/terrors"
	"gb-admin-core/pkg/thttp"
	thttpHeaders "gb-admin-core/pkg/thttp/headers"
	"github.com/sarulabs/di"
)

type UiUseCase struct {
	AuthUC     authInterface.UseCase `di:"authUC"`
	HttpClient *thttp.ThttpClient    `di:"httpClient"`
}

func BuildUiUseCase(ctn di.Container) (interface{}, error) {
	return &UiUseCase{
		AuthUC:     ctn.Get("authUC").(authInterface.UseCase),
		HttpClient: ctn.Get("httpClient").(*thttp.ThttpClient),
	}, nil
}

func (uuc *UiUseCase) GetCalcActiveElements() (interface{}, uint16, error) {
	service, err := uuc.AuthUC.GetAuthServiceByName(model.UiService)
	if err != nil {
		return nil, 0, err
	}

	var response model.GetCalcActiveElementsResponse

	headers, err := thttpHeaders.MakeAuthHeaders("", service.PublicKey, service.PrivateKey, "GET")
	fmt.Printf("\n%+v\n", headers)
	rawResponse, statusCode, err := uuc.HttpClient.Request(
		thttp.GET,
		fmt.Sprintf("%s/calc/element/active", service.URL),
		headers,
		"",
		&response,
		map[string]interface{}{
			"source": "admin",
		},
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

func (uuc *UiUseCase) SetActiveForElements(params *model.SetActiveForElementRequest) (interface{}, uint16, error) {
	service, err := uuc.AuthUC.GetAuthServiceByName(model.UiService)
	if err != nil {
		return nil, 0, err
	}

	var response common.Response

	headers, err := thttpHeaders.MakeAuthHeaders(params, service.PublicKey, service.PrivateKey, "PATCH")
	fmt.Printf("\n%+v\n", headers)
	rawResponse, statusCode, err := uuc.HttpClient.Request(
		thttp.PATCH,
		fmt.Sprintf("%s/calc/element/active", service.URL),
		headers,
		params,
		&response,
		nil,
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
