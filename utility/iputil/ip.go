package iputil

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/yimuysl001/gtoolboxs/utility/logger"
	"io/ioutil"
	"net"
	"net/http"
	"os/exec"
	"strings"
)

const (
	whoisApi = "https://whois.pconline.com.cn/ipJson.jsp?json=true&ip="
	dyndns   = "http://members.3322.org/dyndns/getip"
)

type IpLocationData struct {
	Ip           string `json:"ip"`
	Country      string `json:"country"`
	Region       string `json:"region"`
	Province     string `json:"province"`
	ProvinceCode int64  `json:"province_code"`
	City         string `json:"city"`
	CityCode     int64  `json:"city_code"`
	Area         string `json:"area"`
	AreaCode     int64  `json:"area_code"`
}

type WhoisRegionData struct {
	Ip         string `json:"ip"`
	Pro        string `json:"pro" `
	ProCode    string `json:"proCode" `
	City       string `json:"city" `
	CityCode   string `json:"cityCode"`
	Region     string `json:"region"`
	RegionCode string `json:"regionCode"`
	Addr       string `json:"addr"`
	Err        string `json:"err"`
}

// GetClientIp 获取客户端IP
func GetClientIp(r *ghttp.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = r.GetClientIp()
	}

	// 如果存在多个，默认取第一个
	if gstr.Contains(ip, ",") {
		ip = gstr.StrTillEx(ip, ",")
	}

	if gstr.Contains(ip, ", ") {
		ip = gstr.StrTillEx(ip, ", ")
	}

	return ip
}

// GetLocalIP 获取服务器内网IP
func GetLocalIP() (ip string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, addr := range addrs {
		ipAddr, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.IsLoopback() {
			continue
		}
		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}
		return ipAddr.IP.String(), nil
	}
	return
}

func GetPublicIP2() (ip string, err error) {
	response, err := http.Get(dyndns)
	if err != nil {
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	ip = strings.ReplaceAll(string(body), "\n", "")
	return
}

func GetLocalIPv4() (ipv4Address string, err error) {
	defer func() {
		r := recover()
		logger.Logger.IfError(r)

	}()

	ipv4Address = ""
	cmd := exec.Command("ipconfig")
	// 执行命令，并返回结果
	output, err := cmd.Output()
	if err != nil {
		return
	}
	// 因为结果是字节数组，需要转换成string
	ipv4Address = string(output)
	ipv4Address = strings.ReplaceAll(ipv4Address, "\r\n", ":")
	split := strings.Split(ipv4Address, ":")
	for _, s := range split {
		s = strings.TrimSpace(s)
		ip := net.ParseIP(s)
		if ip != nil {
			if ip.To4() != nil {
				ipv4Address = s
				return
			}
		}
	}

	return "", nil
}
