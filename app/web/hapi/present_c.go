package hapi

import (
	"fmt"
	"github.com/jsix/goex/echox"
	"github.com/jsix/gof/crypto"
	"github.com/jsix/gof/storage"
	"github.com/jsix/gof/util"
	"go2o/core/domain/interface/ad"
	"go2o/core/service/rsi"
	"net/http"
	"strconv"
	"strings"
)

// 用于支持界面呈现的控制器
type presentC struct {
}

func (p *presentC) getMd5(s string) string {
	return crypto.Md5([]byte(s))[8:16]
}

// 广告
func (p *presentC) AdApi(c *echox.Context) error {
	callback := c.QueryParam("callback")
	namesParams := strings.TrimSpace(c.QueryParam("keys"))
	names := strings.Split(namesParams, "|")
	userId, _ := util.I32Err(strconv.Atoi(c.QueryParam("user_id")))
	as := rsi.AdService
	result := make(map[string]*ad.AdDto, len(names))
	key := fmt.Sprintf("go2o:repo:ad:%d:front:%s", userId,
		p.getMd5(namesParams))
	rds := c.App.Storage().(storage.IRedisStorage)
	if rds == nil {
		panic("storage need redis support")
	}
	var seconds int64 = 0
	rds.RWJson(key, &result, func() interface{} {
		//从缓存中读取
		for _, n := range names {
			//分别绑定广告
			dto := as.GetAdAndDataByKey(userId, n)
			if dto == nil {
				result[n] = nil
				continue
			}
			result[n] = dto
		}
		seconds = rsi.FoundationService.GetRegistry().CacheAdMaxAge
		return result
	}, seconds)
	return c.JSONP(http.StatusOK, callback, result)
}
