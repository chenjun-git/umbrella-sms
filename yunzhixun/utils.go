package yunzhixun

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"strings"
)

//control over the request lifecycle
var httpClient = &http.Client{
	Timeout: 2000 * time.Millisecond,
}

func SetClientTimeout(duration time.Duration) {
	httpClient.Timeout = duration
}

func apiURL(apiType YUNZHIXUN_API_TYPE) string {
	switch apiType {
	case SMS_SEND_SINGLE:
		return YUNZHIXUN_API_HOST + "/ol/sms/sendsms"
	case SMS_SEND_GROUP:
		return YUNZHIXUN_API_HOST + "/ol/sms/sendsms_batch"
	case SMS_SEND_STATUS:
		return YUNZHIXUN_API_HOST + "/"
	case SMS_TEMPLATE_ADD:
		return YUNZHIXUN_API_HOST + "/ol/sms/addsmstemplate"
	case SMS_TEMPLATE_QUERY:
		return YUNZHIXUN_API_HOST + "/ol/sms/getsmstemplate"
	case SMS_TEMPLATE_EDIT:
		return YUNZHIXUN_API_HOST + "/ol/sms/editsmstemplate"
	case SMS_TEMPLATE_DELETE:
		return YUNZHIXUN_API_HOST + "/ol/sms/deletegetsmstemplate"
	default:
		panic(fmt.Errorf("invalid apiType %v", apiType))
	}
}

func httpReqWithParams(method, url string, v interface{}) ([]byte, error) {
	parmas, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, url, strings.NewReader(string(parmas)))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf(resp.Status)
	}

	return ioutil.ReadAll(resp.Body)
}
