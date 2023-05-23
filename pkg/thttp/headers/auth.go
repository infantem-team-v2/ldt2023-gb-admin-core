package thttpHeaders

import (
	"encoding/json"
	"gb-auth-gate/pkg/thttp"
	"gb-auth-gate/pkg/tsecure"
	"strconv"
	"strings"
	"time"
)

// MakeAuthHeaders Function that build default auth headers to work with other services
func MakeAuthHeaders(body interface{}, publicKey, privateKey string, method string) (headers map[string]string, err error) {
	headers = make(map[string]string, 0)
	timestamp := time.Now().UTC().Unix()
	strBuilder := strings.Builder{}
	_, err = strBuilder.WriteString(strconv.Itoa(int(timestamp)))
	if err != nil {
		return nil, err
	}
	if method == thttp.GET {

		_, mErr := strBuilder.WriteString(publicKey)
		if mErr != nil {
			return nil, mErr
		}
	} else {
		bodyBytes, mErr := json.Marshal(body)
		if mErr != nil {
			return nil, mErr
		}
		_, mErr = strBuilder.WriteString(string(bodyBytes))
		if mErr != nil {
			return nil, mErr
		}
	}

	signature := tsecure.CalcSignature(privateKey, strBuilder.String(), tsecure.SHA512)

	headers[ContentTypeKey] = TypeApplicationJSON.String()
	headers["ApiPublic"] = publicKey
	headers["TimeStamp"] = strconv.Itoa(int(timestamp))
	headers["Signature"] = signature

	return headers, nil
}
