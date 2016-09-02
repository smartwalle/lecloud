package lecloud

import "fmt"

////////////////////////////////////////////////////////////////////////////////
// GetPushTokenParam 直播活动推流Token获取接口
// http://help.lecloud.com/Wiki.jsp?page=activity.getPushToken
type GetPushTokenParam struct {
	ActivityId string // 是 直播活动ID
}

func (this GetPushTokenParam) APIName() string {
	return "lecloud.cloudlive.activity.getPushToken"
}

func (this GetPushTokenParam) Params() map[string]string {
	var m = make(map[string]string)
	m["activityId"] = this.ActivityId
	return m
}

func (this GetPushTokenParam) Method() string {
	return "GET"
}

////////////////////////////////////////////////////////////////////////////////
// GetPushURLParam 直播活动推流地址获取接口
// http://help.lecloud.com/Wiki.jsp?page=activity.getPushUrl
type GetPushURLParam struct {
	ActivityId string // 是 直播活动ID
}

func (this GetPushURLParam) APIName() string {
	return "lecloud.cloudlive.activity.getPushUrl"
}

func (this GetPushURLParam) Params() map[string]string {
	var m = make(map[string]string)
	m["activityId"] = this.ActivityId
	return m
}

func (this GetPushURLParam) Method() string {
	return "GET"
}

////////////////////////////////////////////////////////////////////////////////
// SearchActivityInfoParam 直播活动信息查询接口
// http://help.lecloud.com/Wiki.jsp?page=activity.vrsinfo.search

type SearchActivityInfoParam struct {
	ActivityId     string // 否 直播活动ID
	ActivityName   string // 否 直播活动名称
	ActivityStatus string // 否 直播活动状态。0：未开始 1：已开始 3：已结束
	OffSet         int    // 否 从第几条数据开始查询，默认0
	FetchSize      int    // 否 一次返回多少条数据，默认为10，最多不能超过100条
}

func (this SearchActivityInfoParam) APIName() string {
	return "lecloud.cloudlive.vrs.activity.vrsinfo.search"
}

func (this SearchActivityInfoParam) Params() map[string]string {
	var m = make(map[string]string)

	if len(this.ActivityId) > 0 {
		m["activityId"] = this.ActivityId
	}

	if len(this.ActivityName) > 0 {
		m["activityName"] = this.ActivityName
	}

	if len(this.ActivityStatus) > 0 {
		m["activityStatus"] = this.ActivityStatus
	}

	if this.OffSet > 0 {
		m["offset"] = fmt.Sprintf("%d", this.OffSet)
	}

	if this.FetchSize <= 0 {
		this.FetchSize = 10
	}
	m["fetchSize"] = fmt.Sprintf("%d", this.FetchSize)

	return m
}

func (this SearchActivityInfoParam) Method() string {
	return "GET"
}

////////////////////////////////////////////////////////////////////////////////
// CreateActivityParam 直播活动创建接口
// http://help.lecloud.com/Wiki.jsp?page=activity.create
//活动编码（云直播使用）
//001 发布会
//002 婚礼
//003 年会
//004 体育
//005 游戏
//006 旅游&户外
//007 财经
//008 演唱会
//009 烹饪
//010 宠物&动物
//011 访谈
//012 教育
//013 竞技
//014 剧场
//015 晚会
//016 电视节目
//017 秀场
//999 其他

type CreateActivityParam struct {
	ActivityName     string // 是 直播活动名称(200个字符以内)
	StartTime        string // 是 开始时间 格式yyyyMMddHHmmss
	EndTime          string // 是 结束时间 格式yyyyMMddHHmmss
	CoverImgUrl      string // 否 活动封面地址，如果为空，则系统会默认一张图片
	Description      string // 否 活动描述（1024个字符以内）
	LiveNum          int    // 是 机位数量，范围为：1,2,3,4. 默认为1
	CodeRateTypes    string // 是 流的码率类型，逗号分隔。由大到小排列。取值范围：10 流畅；13 标清；16 高清；19 超清；22 720P；25 1080P；99 原画
	NeedRecord       int    // 是 是否支持全程录制 0：否 1：是。默认为0
	NeedTimeShift    int    // 是 是否支持时移 0：否 1：是。默认为0
	NeedFullView     int    // 是 是否全景观看 0：否 1：是。默认为0
	ActivityCategory string // 是 活动分类，参见《活动编码》
	PlayMode         int    // 是 播放模式，0：实时直播。1：流畅直播
}

func (this CreateActivityParam) APIName() string {
	return "lecloud.cloudlive.activity.create"
}

func (this CreateActivityParam) Params() map[string]string {
	var m = make(map[string]string)
	m["activityName"] = this.ActivityName
	m["startTime"] = this.StartTime
	m["endTime"] = this.EndTime
	if len(this.CoverImgUrl) > 0 {
		m["coverImgUrl"] = this.CoverImgUrl
	}
	if len(this.Description) > 0 {
		m["description"] = this.Description
	}
	if this.LiveNum < 1 || this.LiveNum > 4 {
		this.LiveNum = 1
	}
	m["liveNum"] = fmt.Sprintf("%d", this.LiveNum)
	m["codeRateTypes"] = this.CodeRateTypes
	m["needRecord"] = fmt.Sprintf("%d", this.NeedRecord)
	m["needTimeShift"] = fmt.Sprintf("%d", this.NeedTimeShift)
	m["needFullView"] = fmt.Sprintf("%d", this.NeedFullView)
	m["activityCategory"] = this.ActivityCategory
	m["playMode"] = fmt.Sprintf("%d", this.PlayMode)
	return m
}

func (this CreateActivityParam) Method() string {
	return "POST"
}

////////////////////////////////////////////////////////////////////////////////
// UpdateActivityVrsInfoParam 直播活动媒资信息修改接口
// http://help.lecloud.com/Wiki.jsp?page=activity.vrsinfo.modify
type ModifyActivityVrsInfoParam struct {
	ActivityId       string // 是 直播活动ID
	ActivityName     string // 否 直播活动名称(200个字符以内)
	StartTime        string // 否 开始时间 格式yyyyMMddHHmmss
	EndTime          string // 否 结束时间 格式yyyyMMddHHmmss
	CoverImgUrl      string // 否 活动封面地址，如果为空，则系统会默认一张图片
	Description      string // 否 活动描述（1024个字符以内）
	ActivityCategory string // 否 活动分类参见如下《活动编码》，无二级编码时直接填写一级编码，参见如下《扩展字段参数列表》
	Extensions       string // 扩展字段，活动分类修改时，需要修改分类的扩展字段，也可在活动分类不变的情况下单独修改扩展字段，按如下方式传参数： 参数名为要修改的扩展字段如： {"host":"主队名称","guest":"客队名称"}
}

func (this ModifyActivityVrsInfoParam) APIName() string {
	return "lecloud.cloudlive.vrs.activity.vrsinfo.modify"
}

func (this ModifyActivityVrsInfoParam) Params() map[string]string {
	var m = make(map[string]string)
	m["activityId"] = this.ActivityId

	if len(this.ActivityName) > 0 {
		m["activityName"] = this.ActivityName
	}

	if len(this.StartTime) > 0 {
		m["startTime"] = this.StartTime
	}
	if len(this.EndTime) > 0 {
		m["endTime"] = this.EndTime
	}
	if len(this.CoverImgUrl) > 0 {
		m["coverImgUrl"] = this.CoverImgUrl
	}
	if len(this.Description) > 0 {
		m["description"] = this.Description
	}
	if len(this.ActivityCategory) > 0 {
		m["activityCategory"] = this.ActivityCategory
	}
	if len(this.Extensions) > 0 {
		m["extensions"] = this.Extensions
	}
	return m
}

func (this ModifyActivityVrsInfoParam) Method() string {
	return "POST"
}

////////////////////////////////////////////////////////////////////////////////
// SearchActivityStreamInfo 直播活动流信息查询接口
// http://help.lecloud.com/Wiki.jsp?page=streaminfo.search
type SearchActivityStreamInfo struct {
	ActivityId string // 是 直播活动ID
}

func (this SearchActivityStreamInfo) APIName() string {
	return "lecloud.cloudlive.vrs.activity.streaminfo.search"
}

func (this SearchActivityStreamInfo) Params() map[string]string {
	var m = make(map[string]string)
	m["activityId"] = this.ActivityId
	return m
}

func (this SearchActivityStreamInfo) Method() string {
	return "GET"
}

////////////////////////////////////////////////////////////////////////////////
// StopActivityParam 直播活动结束接口
// http://help.lecloud.com/Wiki.jsp?page=activity.stop
type StopActivityParam struct {
	ActivityId string // 是 直播活动ID
}

func (this StopActivityParam) APIName() string {
	return "lecloud.cloudlive.activity.stop"
}

func (this StopActivityParam) Params() map[string]string {
	var m = make(map[string]string)
	m["activityId"] = this.ActivityId
	return m
}

func (this StopActivityParam) Method() string {
	return "POST"
}

////////////////////////////////////////////////////////////////////////////////
// GetActivityMachineStateParam 直播活动机位状态查询接口
// http://help.lecloud.com/Wiki.jsp?page=getActivityMachineState
type GetActivityMachineStateParam struct {
	ActivityId string // 是 直播活动ID
}

func (this GetActivityMachineStateParam) APIName() string {
	return "letv.cloudlive.activity.getActivityMachineState"
}

func (this GetActivityMachineStateParam) Params() map[string]string {
	var m = make(map[string]string)
	m["activityId"] = this.ActivityId
	return m
}

func (this GetActivityMachineStateParam) Method() string {
	return "GET"
}

////////////////////////////////////////////////////////////////////////////////
// ModifyActivityCoverImgParam 直播活动封面上传接口
// http://help.lecloud.com/Wiki.jsp?page=modifyCoverImgnew
type ModifyActivityCoverImgParam struct {
	ActivityId string // 是 直播活动ID
	File       string // 是 要上传的封面图片
}

func (this ModifyActivityCoverImgParam) APIName() string {
	return "lecloud.cloudlive.activity.modifyCoverImg"
}

func (this ModifyActivityCoverImgParam) Params() map[string]string {
	var m = make(map[string]string)
	m["activityId"] = this.ActivityId
	return m
}

func (this ModifyActivityCoverImgParam) Method() string {
	return "POST"
}

////////////////////////////////////////////////////////////////////////////////
// GetActivityPlayerPageURLParam 直播活动播放页地址获取
// http://help.lecloud.com/Wiki.jsp?page=playerpage.getUrl
type GetActivityPlayerPageURLParam struct {
	ActivityId string // 是 直播活动ID
}

func (this GetActivityPlayerPageURLParam) APIName() string {
	return "lecloud.cloudlive.activity.playerpage.getUrl"
}

func (this GetActivityPlayerPageURLParam) Params() map[string]string {
	var m = make(map[string]string)
	m["activityId"] = this.ActivityId
	return m
}

func (this GetActivityPlayerPageURLParam) Method() string {
	return "POST"
}

////////////////////////////////////////////////////////////////////////////////
// CreateRecTaskParam 直播创建打点录制任务接口
// http://help.lecloud.com/Wiki.jsp?page=createRecTasknew
type CreateRecTaskParam struct {
	LiveId    string // 是 直播ID，直播id查询参考《活动流信息查询接口》文档
	StartTime string // 是 开始时间 格式yyyyMMddHHmmss
	EndTime   string // 是 结束时间 格式yyyyMMddHHmmss
}

func (this CreateRecTaskParam) APIName() string {
	return "lecloud.cloudlive.rec.createRecTask"
}

func (this CreateRecTaskParam) Params() map[string]string {
	var m = make(map[string]string)
	m["liveId"] = this.LiveId
	m["startTime"] = this.StartTime
	m["endTime"] = this.EndTime
	return m
}

func (this CreateRecTaskParam) Method() string {
	return "POST"
}

////////////////////////////////////////////////////////////////////////////////
// SearchRecResult 直播打点录制结果查询接口
// http://help.lecloud.com/Wiki.jsp?page=SearchResultnew
type SearchRecResult struct {
	LiveId    string // 否 直播ID，直播id查询参考《活动流信息查询接口》文档
	TaskId    string // 否 任务id
	Offset    int    // 否 开始行数
	Size      int    // 否 每页记录数
	StartTime string // 否 开始时间 格式yyyyMMdd
	EndTime   string // 否 结束时间 格式yyyyMMdd
}

func (this SearchRecResult) APIName() string {
	return "lecloud.cloudlive.rec.searchResult"
}

func (this SearchRecResult) Params() map[string]string {
	var m = make(map[string]string)
	m["liveId"] = this.LiveId
	return m
}

func (this SearchRecResult) Method() string {
	return "GET"
}

////////////////////////////////////////////////////////////////////////////////
// SecurityConfig 直播活动安全信息设置接口
// http://www.lecloud.com/zh-cn/help/2016/07/27/121.html?LeftMenu=api_zb_trans
type SecurityConfig struct {
	ActivityId                string // 是 直播活动ID
	NeededPushAuth            int    // 是 是否启用推流鉴权: 0、否;1、是
	PushUrlValidTime          int    // 否 推流地址有效时长,单位s,启用推流鉴权时有效
	LiveKey                   string // 否 直播安全码,计算推流地址时用到的安全码,如果为空的话,则使用客户的安全码
	NeedIpWhiteList           int    // 是 是否启用IP推流白名单: 0、否;1、是
	PushIpWhiteList           string // 否 推流IP白名单。多个IP时,用英文半角逗号分隔,IP最多配置10个。
	NeedPlayerDomainWhiteList int    // 是 是否启用域名白名单: 0、否;1、是
	PlayerDomainWhiteList     string // 否 域名白名单。多个域名时,用英文半角逗号分隔,最多配置10个。
}

func (this SecurityConfig) APIName() string {
	return "lecloud.cloudlive.activity.sercurity.config"
}

func (this SecurityConfig) Params() map[string]string {
	var m = make(map[string]string)
	m["activityId"] = this.ActivityId
	m["neededPushAuth"] = fmt.Sprintf("%d", this.NeededPushAuth)
	if this.NeededPushAuth == 1 {
		m["pushUrlValidTime"] = fmt.Sprintf("%d", this.PushUrlValidTime)
	}
	if this.LiveKey != "" {
		m["liveKey"] = this.LiveKey
	}
	m["needIpWhiteList"] = fmt.Sprintf("%d", this.NeedIpWhiteList)
	if this.NeedIpWhiteList == 1 {
		m["pushIpWhiteList"] = this.PushIpWhiteList
	}
	m["needPlayerDomainWhiteList"] = fmt.Sprintf("%d", this.NeedPlayerDomainWhiteList)
	if this.NeedPlayerDomainWhiteList == 1 {
		m["playerDomainWhiteList"] = this.PlayerDomainWhiteList
	}

	return m
}

func (this SecurityConfig) Method() string {
	return "POST"
}
