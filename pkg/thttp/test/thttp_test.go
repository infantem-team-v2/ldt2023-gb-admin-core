package test_http

import (
	"bank_api/internal/pkg/dependency"
	"bank_api/pkg/thttp"
	thttpHeaders "bank_api/pkg/thttp/headers"
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
	"time"
)

type TestCase struct {
	Public  string
	Private string
	Tracker string
}

// =======================REQUEST=======================//
func TestRequest(t *testing.T) {
	tdc := dependency.NewDependencyContainer().BuildDependencies().BuildContainer()

	t.Run("non query params", func(t *testing.T) {
		expectedResult := map[string]interface{}{
			"key": "value",
		}
		dest := map[string]interface{}{}
		hc := tdc.ContainerInstance().Get("httpClient").(*thttp.ThttpClient)
		result, statusCode, err := hc.Request("GET", "https://api.twinklepick.com/api/transaction/get", nil, nil, &dest, nil)
		if err != nil {
			t.Fatal(err)
		}
		if statusCode != 200 {
			t.Errorf("Request returned statusCode %d, expected 200", statusCode)
		}
		assert.Equal(t, expectedResult, result)
	})
}

// ==========================MULTI_LOAD=========================//
func TestMultiLoad(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Run(fmt.Sprintf("Testing highload: %d", i), TestLoad)
	}
}

// =========================LOAD==========================//
func TestLoad(b *testing.T) {
	tdc := dependency.NewDependencyContainer().BuildDependencies().BuildContainer()
	httpClient := tdc.ContainerInstance().Get("httpClient").(*thttp.ThttpClient)

	testCases := []TestCase{
		{
			Public:  "jm5r4eutgpk2hqgt5dkrqn3ihf75wtvsfdzou4wf7feh7unmmsr5sv53yazfxuv8",
			Private: "op9io5eyfne37xb2ad2ovyohb7oaqofjfp7yabuusz9tfbmm4z6hdkn7bx4gojfbq2fnuvxgecii5cocsrsu6x4ed9enui7z6t66hbfpdc4xcxmdca2xcb8t6tz2a8gz",
			Tracker: "64656703e9b11e97c821d998a2743a2a94b30fee36fe7380fd5591f2767ab13a145b02d1673dcb9a5e7077afd85f57b8118fbf12c8eca011c77b58fc5cb5d5bf",
		},
		{
			Public:  "kvwekcbmu9fkvjcixoqpisuxovrr5ubs5b96bbdiyftsbt7xf7inack3oeg6rxa8",
			Private: "hvhdqwqtpjwpxf6iskoobziibfjcqrxzj8qh2ausyurnk84zx3vxp9wbh237ywiryu2ra2vqu9wofzkdectzcp3oqqvkw6q4tsa3aufaxgbhgvscfqn76nw2wmzgwbuj",
			Tracker: "4af5962d41d225c4db2b1fb25633cc845f6cdbf5d0621324d377d5ff3b3a52fe4c5da4e91946de7516a51bb61fc9ce912a60a2a81e202f2f6c19b29af6d8c066",
		},
		{
			Public:  "w3ia3a72f8saddn7h8kmo8u9p7zqapsknpdye48oa4cig3vqdgr97tdh257nncm6",
			Private: "ddg9gnj2jeks7utw2y8xker8bndejxizfbbjsgbd5rze2kzcmq7o6qpcbqhdztjrkx6gntn4u26qdgtjzvr28g2kw7sukfffwo3i963xp8zf98ke9m8g769y85geqkod",
			Tracker: "4af5962d41d225c4db2b1fb25633cc845f6cdbf5d0621324d377d5ff3b3a52fe4c5da4e91946de7516a51bb61fc9ce912a60a2a81e202f2f6c19b29af6d8c066",
		},
		{
			Public:  "k9kq9icrsr24yphv958c8gnuzis77k6gwhqbpxcbf75kq2k59awjg55u8kz69yoj",
			Private: "uboq8qykxeowkdbsy6eksfbx69r7fh7bh8wm5ss25ufpc4m2ddk5dem5r9cqi9f89cj978txuyeyzjqzgg79tz45c9hr3gwrcg9kz73esc9vpn7ng44zuxoubrutshap",
			Tracker: "1b0af7add7747ad9815f241d178de7bdb4bc877faa05ed0150af18c955557de28b67100d6027a391145af76f0ff919369942772dbcfe73964fa791e59235f2f9",
		},
		{
			Public:  "vsdtbt6i5paff0i251lm4ls02teun38tu4zjudbkjaanqm4b2qkzqu1u6sh8uxzo",
			Private: "3c7va6ypnxr6216qls6p3vf971wbqvboekr8znyi974qigrh588wemsnz3zfxws8qnzkihoz5fu3govs8im9xkx8gkxcgfb29ttdeq8fb7qrflsk12ku8zopzvsa8nn1",
			Tracker: "86140afbb82903512e0592680debc94882d9b42a6f44b1048043d5582c8332b572cd4fce5e14ae81152e71a784f45770d9a88406ea1ff543173c6ffad131c5aa",
		},
	}

	timeSlice := make([]int64, 9, 100000)
	var total int64 = 0
	c := make(chan int64)
	for i := 0; i < 100000; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				body := map[string]string{"tracker_id": testCases[int(math.Abs(float64(i%5)))].Tracker}
				headers, _ := thttpHeaders.MakeAuthHeaders(body,
					testCases[i%5].Public,
					testCases[i%5].Private,
					thttp.POST,
				)
				start := time.Now().UnixNano()
				httpClient.Request(thttp.POST, "https://test.onecrypto.pro/api/transaction/get", headers, body, nil, nil)
				end := time.Now().UnixNano()
				timeSlice = append(timeSlice, end-start)
				c <- end - start
			}

		}()

	}
	total += <-c
	b.Logf("Avg time: %d", total/int64(len(timeSlice)))
}

// =======================MAKE_QUERY_STRING============================//
func TestMakeQueryString(t *testing.T) {
	hc := thttp.ThttpClient{}
	t.Run("non-empty query params", func(t *testing.T) {
		queryParams := map[string]interface{}{
			"key1": "value1",
			"key2": "value2",
		}

		expectedResult := "?key1=value1&key2=value2"

		result := hc.MakeQueryString(queryParams)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("empty query params", func(t *testing.T) {
		queryParams := map[string]interface{}{}

		expectedResult := ""

		result := hc.MakeQueryString(queryParams)

		assert.Equal(t, expectedResult, result)
	})
}
