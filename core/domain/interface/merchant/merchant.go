/**
 * Copyright 2014 @ z3q.net.
 * name :
 * author : jarryliu
 * date : 2013-12-12 16:53
 * description :
 * history :
 */

package merchant

import (
	"go2o/core/domain/interface/merchant/mss"
	"go2o/core/domain/interface/merchant/user"
)

const (
	RegisterModeClosed         = 0 // 关闭注册
	RegisterModeNormal         = 1 // 正常注册
	RegisterModeMustInvitation = 2 // 必须邀请注册
	RegisterModeMustRedirect   = 3 // 必须直接注册
)

type (
	IMerchant interface {
		GetAggregateRootId() int

		GetValue() Merchant

		SetValue(*Merchant) error

		// 保存
		Save() (int, error)

		// 获取商户的域名
		GetMajorHost() string

		// 获取销售配置
		GetSaleConf() SaleConf

		// 保存销售配置
		SaveSaleConf(*SaleConf) error

		// 获取站点配置
		GetSiteConf() ShopSiteConf

		// 保存站点配置
		SaveSiteConf(*ShopSiteConf) error

		// 注册权限验证,如果没有权限注册,返回错误
		RegisterPerm(isInvitation bool) error

		// 新建商店
		CreateShop(*Shop) IShop

		// 获取所有商店
		GetShops() []IShop

		// 获取营业中的商店
		GetBusinessInShops() []IShop

		// 获取商店
		GetShop(int) IShop

		// 删除门店
		DeleteShop(shopId int) error

		// 返回用户服务
		UserManager() user.IUserManager

		// 返回设置服务
		ConfManager() IConfManager

		// 获取会员等级服务
		LevelManager() ILevelManager

		// 获取键值管理器
		KvManager() IKvManager

		// 企业资料服务
		ProfileManager() IProfileManager

		// API服务
		ApiManager() IApiManager

		// 获取会员键值管理器
		MemberKvManager() IKvManager

		// 消息系统管理器
		MssManager() mss.IMssManager
	}

	//合作商
	Merchant struct {
		Id   int    `db:"id" pk:"yes" auto:"yes"`
		Usr  string `db:"usr"`
		Pwd  string `db:"pwd"`
		Name string `db:"name"`
		Logo string `db:"logo"`
		// 省
		Province int `db:"province"`
		// 市
		City int `db:"city"`
		// 区
		District      int   `db:"district"`
		ExpiresTime   int64 `db:"expires_time"`
		JoinTime      int64 `db:"join_time"`
		UpdateTime    int64 `db:"update_time"`
		LoginTime     int64 `db:"login_time"`
		LastLoginTime int64 `db:"last_login_time"`
	}
	SaleConf struct {
		MerchantId              int     `db:"merchant_id" auto:"no" pk:"yes"` // 合作商编号
		CashBackPercent         float32 `db:"cb_percent"`                     // 返现比例,0则不返现
		CashBackTg1Percent      float32 `db:"cb_tg1_percent"`                 // 一级比例
		CashBackTg2Percent      float32 `db:"cb_tg2_percent"`                 // 二级比例
		CashBackMemberPercent   float32 `db:"cb_member_percent"`              // 会员比例
		IntegralBackNum         int     `db:"ib_num"`                         // 每一元返多少积分
		IntegralBackExtra       int     `db:"ib_extra"`                       // 每单额外赠送
		AutoSetupOrder          int     `db:"oa_open"`                        // 自动设置订单
		OrderTimeOutMinute      int     `db:"oa_timeout_minute"`              // 订单超时分钟数
		OrderConfirmAfterMinute int     `db:"oa_confirm_minute"`              // 订单自动确认时间
		OrderTimeOutReceiveHour int     `db:"oa_receive_hour"`                // 订单超时自动收货
		RegisterMode            int     `db:"register_mode"`                  // 必须注册模式
		ApplyCsn                float32 `db:"apply_csn"`                      // 提现手续费费率
		TransCsn                float32 `db:"trans_csn"`                      // 转账手续费费率
		FlowConvertCsn          float32 `db:"flow_convert_csn"`               // 活动账户转为赠送可提现奖金手续费费率
		PresentConvertCsn       float32 `db:"present_convert_csn"`            // 赠送账户转换手续费费率
	}
)