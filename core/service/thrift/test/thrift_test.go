/**
 * Copyright 2015 @ at3.net.
 * name : thrift_test.go
 * author : jarryliu
 * date : 2016-11-13 13:11
 * description :
 * history :
 */
package test

import (
	"errors"
	"github.com/jsix/gof/log"
	"go2o/core/infrastructure/domain"
	"go2o/gen-code/thrift/define"
	"strings"
	"testing"
)

func TestLogin(t *testing.T) {
	cli, err := MemberServeClient()
	if err != nil {
		t.Error(err)
		return
	}

	defer trans.Close()
	t.Logf("连接开启：%v", cli.Transport.IsOpen())
	pwd := domain.MemberSha1Pwd("123456")
	r, _ := cli.CheckLogin("jarry6", pwd, false)
	t.Logf("登录(1)结果：\n MemberId:%d\n Result:%v", r.ID, r.Result_)

	pwd = domain.MemberSha1Pwd("123000")
	r, _ = cli.CheckLogin("jarry6", pwd, false)
	t.Logf("登录(2)结果：\n MemberId:%d\n Result:%v", r.ID, r.Result_)

	arr, _ := cli.InviterArray(16893, 5)
	t.Log("邀请人：", arr)
}

func TestSSORegister(t *testing.T) {
	cli, err := FoundationServeClient()
	if err == nil {
		defer trans.Close()
		sa := &define.SsoApp{
			ID:     1,
			Name:   "gp",
			ApiUrl: "http://localhost:14281/member/sync_m.p",
		}
		s, _ := cli.RegisterApp(sa)
		arr := strings.Split(s, ":")
		if arr[0] != "1" {
			t.Error(errors.New("注册SSO-APP出错：" +
				s + "; api-url:" + sa.ApiUrl))
		} else {
			log.Println("得到的token为：", arr[1])
		}
	} else {
		t.Log("连接失败：", err.Error())
	}
}
