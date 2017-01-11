package lecloud

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/smartwalle/going/http"
	"sort"
	"net/url"
	"time"
)

const (
	LE_CLOUD_OPEN_API_URL      = "http://api.letvcloud.com/open.php"
	LE_CLOUD_OPEN_API_LIVE_URL = "http://api.open.letvcloud.com/live/execute"
)

var (
	UserId     string
	UserUnique string
	SecretKey  string
)

func UpdateKey(userId, userUnique, secretKey string) {
	UserId = userId
	UserUnique = userUnique
	SecretKey = secretKey
}

func VideoRequest(param ILeCloudParam) (result map[string]interface{}, err error) {
	result, err = VideoRequestWithKey(SecretKey, UserUnique, param)
	return result, err
}

func LiveRequest(param ILeCloudParam) (result map[string]interface{}, err error) {
	result, err = LiveRequestWithKey(SecretKey, UserId, param)
	return result, err
}

func VideoRequestWithKey(secretKey, userUnique string, param ILeCloudParam) (result map[string]interface{}, err error) {
	var p = make(map[string]string)
	p["user_unique"] = userUnique
	p["ver"] = "2.0"
	p["api"] = param.APIName()
	p["format"] = "json"
	result, err = requestWithKey(secretKey, userUnique, LE_CLOUD_OPEN_API_URL, p, param)
	return result, err
}

func LiveRequestWithKey(secretKey, userId string, param ILeCloudParam) (result map[string]interface{}, err error) {
	var p = make(map[string]string)
	p["userid"] = userId
	p["ver"] = "3.1"
	p["method"] = param.APIName()
	result, err = requestWithKey(secretKey, userId, LE_CLOUD_OPEN_API_LIVE_URL, p, param)
	return result, err
}

func requestWithKey(secretKey, userUnique, domain string, p map[string]string, param ILeCloudParam) (result map[string]interface{}, err error) {
	p["timestamp"] = fmt.Sprintf("%d000", time.Now().Unix())

	var ps = param.Params()
	if ps != nil {
		for key, value := range ps {
			p[key] = value
		}
	}

	var pp = url.Values{}
	var keys []string
	for key, value := range p {
		pp.Add(key, value)
		keys = append(keys, key)
	}
	sort.Strings(keys)
	pp.Set("sign", sign(secretKey, keys, p))

	result, err = http.JSONRequest(param.Method(), domain, pp)


	//var c = http.NewClient()
	//c.SetMethod(param.Method())
	//c.SetURLString(domain)
	//c.SetHeader("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	//
	//var keys []string
	//for key, value := range p {
	//	c.SetParam(key, value)
	//	keys = append(keys, key)
	//}
	//sort.Strings(keys)
	//
	//c.SetParam("sign", sign(secretKey, keys, p))
	//
	//result, err = c.DoJsonRequest()
	return result, err
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
