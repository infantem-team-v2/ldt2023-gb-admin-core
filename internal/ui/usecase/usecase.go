package usecase

import (
	"encoding/json"
	"fmt"
	authInterface "gb-auth-gate/internal/auth/interface"
	"gb-auth-gate/internal/pkg/common"
	"gb-auth-gate/internal/ui/model"
	"gb-auth-gate/pkg/terrors"
	"gb-auth-gate/pkg/thttp"
	thttpHeaders "gb-auth-gate/pkg/thttp/headers"
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

	headers, err := thttpHeaders.MakeAuthHeaders(nil, service.PublicKey, service.PrivateKey, "GET")
	rawResponse, statusCode, err := uuc.HttpClient.Request(
		thttp.GET,
		fmt.Sprintf("%s/calc/element/active", service.URL),
		headers,
		nil,
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
	}
	return response, statusCode, nil
}
