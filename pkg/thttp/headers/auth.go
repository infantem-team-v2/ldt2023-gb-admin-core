package thttpHeaders

import (
	"encoding/json"
	"gb-admin-core/pkg/tsecure"
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

	bodyBytes, mErr := json.Marshal(body)
	if mErr != nil {
		return nil, mErr
	}

	signature := tsecure.CalcSignature(privateKey, string(bodyBytes), tsecure.SHA512)

	headers[ContentTypeKey] = TypeApplicationJSON.String()
	headers["T-Public-Key"] = publicKey
	headers["TimeStamp"] = strconv.Itoa(int(timestamp))
	headers["T-512-Signature"] = signature

	return headers, nil
}
