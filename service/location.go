package service

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
)

// 百度lbs相关
const (
	BaiduAk    = "EcaqNxUoy0LHtsVKXIshPOqHZZAHN7sj"          // BaiduAk
	BaiduSk    = "jh9cjRcM8CpBkuB4Sm5mnBMgAboqcUz3"          // BaiduSk
	BaiduIPUrl = "https://api.map.baidu.com/location/ip?" // 百度ip获取位置信息url
)

// ip相关header
const (
	XForwardedFor = "X-Forwarded-For"
	XRealIP       = "X-Real-IP"
)

/*GetLocation 根据req获取当前请求的地理位置信息,
如header中有longitude(经度)、latitude(纬度)字段,
则根据经纬度逆地理编码结构化为详细地址,如果没有上面两个字段,
则通过获取调用方ip获取详细地址
*/
func GetLocation(req *http.Request) (string, string, error) {
	// 获取经度
	longitude := req.Header.Get("longitude")

	// 获取纬度
	latitude := req.Header.Get("latitude")

	// 获取远程ip
	ip := remoteIP(req)

	// 如果获取到为内网ip则不传给百度
	if !isPublicIP(net.ParseIP(ip)) {
		ip = ""
	}

	// 如代理ip数组不为空, 则取第一个ip作为调用客户端ip
	if longitude == "" || latitude == "" {
		resp, err := http.Get(createBaiduReqURL(ip))
		if err != nil {
			return "0", "0", err
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "0", "0", err
		}

		m := make(map[string]interface{})
		err = json.Unmarshal(body, &m)
		if err != nil {
			return "0", "0", err
		}

		if v, ok := m["status"].(int); !ok {
			if v != 0 {
				return "0", "0", errors.New("百度返回错误")
			}
		}

		if vc, ok := m["content"].(map[string]interface{}); ok {
			if vp, ok := vc["point"].(map[string]interface{}); ok {
				var xok bool
				var yok bool

				// 经纬度
				longitude, xok = vp["x"].(string)
				latitude, yok = vp["y"].(string)
				if xok && yok {
					fmt.Println("longitude:" + longitude)
					fmt.Println("latitude:" + latitude)
				}
			} else {
				return "0", "0", errors.New("point字段未取到")
			}
		} else {
			return "0", "0", errors.New("content字段未取到")
		}
	}

	// 返回纬经度
	return latitude, longitude, nil
}

/**
  组装百度LBS请求的URL
*/
func createBaiduReqURL(realIP string) (url string) {
	paramsMap := make(map[string]string)
	if realIP != "" {
		paramsMap["ip"] = realIP
	}
	paramsMap["ak"] = BaiduAk
	paramsMap["coor"] = "bd09ll"
	paramsStr := toQueryString(paramsMap)
	sn := createBaiduLbsSn(paramsStr)
	url = BaiduIPUrl + paramsStr + "&sn=" + sn
	return
}

/**
  对Map内所有的value做uft8编码, 拼接返回结果
*/
func toQueryString(paramsMap map[string]string) string {
	if len(paramsMap) == 0 {
		return ""
	}

	var queryString string
	for k, v := range paramsMap {
		queryString = queryString + k + "=" + url.QueryEscape(v) + "&"
	}

	return queryString[0 : len(queryString)-1]
}

/**
  生成百度lbssn字符串
*/
func createBaiduLbsSn(paramsStr string) (sn string) {
	wholeStr := "/location/ip?" + paramsStr + BaiduSk

	// 百度lbs sn
	sn = fmt.Sprintf("%x", md5.Sum([]byte(url.QueryEscape(wholeStr))))
	fmt.Print("wholeStr:", wholeStr)
    //fmt.Print("sn:", sn)
	return
}

// remoteIP 返回远程客户端的 IP，如 192.168.1.1
func remoteIP(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get(XRealIP); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get(XForwardedFor); ip != "" {
		arr := strings.Split(ip, ",")
		if len(arr) > 1 {
			remoteAddr = arr[0]
		} else {
			remoteAddr = ip
		}
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}

	return remoteAddr
}

/**
   是否为公网IP
 */
func isPublicIP(IP net.IP) bool {
	if IP.IsLoopback() || IP.IsLinkLocalMulticast() || IP.IsLinkLocalUnicast() {
		return false
	}
	if ip4 := IP.To4(); ip4 != nil {
		switch true {
		case ip4[0] == 10:
			return false
		case ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31:
			return false
		case ip4[0] == 192 && ip4[1] == 168:
			return false
		default:
			return true
		}
	}
	return false
}
