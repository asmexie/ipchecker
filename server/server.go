package main

import (
	"fmt"
	"net/http"

	"github.com/asmexie/go-logger/logger"
	"github.com/gin-gonic/gin"
)

//localhost:8080/api/ipapp/taobao?ip=127.0.0.1
//http://ip.taobao.com/service/getIpInfo.php?ip=192.34.23.23
//response: {"code":0,"data":{"ip":"192.34.23.23","country":"美国","area":"","region":"加利福尼亚","city":"洛杉矶","county":"XX","isp":"XX","country_id":"US","area_id":"","region_id":"US_104","city_id":"US_1018","county_id":"xx","isp_id":"xx"}}
func iptaobao(c *gin.Context) {
	ip := c.Query("ip")

	//do requset from taobao
	url := fmt.Sprintf("http://ip.taobao.com/service/getIpInfo.php?ip=%s", ip)
	buf, err := doRequest(url)
	if err != nil {
		logger.Error(err.Error())
	}
	c.String(http.StatusOK, string(buf))
}
