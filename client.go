package lecloud

import (
	"fmt"
	"sort"
	"time"
	"crypto/md5"
	"encoding/hex"
	"github.com/smartwalle/going/http"
)

const (
	LE_CLOUD_OPEN_API_URL = "http://api.letvcloud.com/open.php"
)

var (
	userUnique string
	secretKey  string
)

func UpdateKey(userUnique, secretKey string) {
	userUnique = userUnique
	secretKey  = secretKey
}

func Request(param ILeCloudParam) (results map[string]interface{}, err error) {
	results, err = RequestWithKey(secretKey, userUnique, param)
	return results, err
}

func RequestWithKey(secretKey, userUnique string, param ILeCloudParam) (results map[string]interface{}, err error) {
	var p = make(map[string]string)
	p["user_unique"] = userUnique
	p["timestamp"]   = fmt.Sprintf("%d", time.Now().Unix())
	p["format"]      = "json"
	p["ver"]         = "2.0"
	p["api"]         = param.APIName()

	var ps = param.Params()
	if ps != nil {
		for key, value := range ps {
			p[key] = value
		}
	}

	var c = http.NewClient()
	c.SetMethod("POST")
	c.SetURLString(LE_CLOUD_OPEN_API_URL)

	var keys = make([]string, len(p))
	for key, value := range p {
		c.SetParam(key, value)
		keys = append(keys, key)
	}
	sort.Strings(keys)

	c.SetParam("sign", sign(secretKey, keys, p))

	results, err = c.DoJsonRequest()
	return results, err
}

func sign(secretKey string, keys []string, param map[string]string) (s string) {
	for _, key := range keys {
		s = s + key + param[key]
	}
	s = fmt.Sprintf("%s%s", s, secretKey)

	var m = md5.New()
	m.Write([]byte(s))
	s = hex.EncodeToString(m.Sum(nil))
	return s
}