package lecloud

import "fmt"

////////////////////////////////////////////////////////////////////////////////
// GetPushTokenParam 直播活动推流Token获取接口
// http://help.lecloud.com/Wiki.jsp?page=activity.getPushToken
type GetPushTokenParam struct {
	ActivityId  string // 是 直播活动ID
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
// GetActivityInfoSearchParam 直播活动信息查询接口
// http://help.lecloud.com/Wiki.jsp?page=activity.vrsinfo.search

type GetActivityInfoSearchParam struct {
	ActivityId     string // 否 直播活动ID
	ActivityName   string // 否 直播活动名称
	ActivityStatus string // 否 直播活动状态。0：未开始 1：已开始 3：已结束
	OffSet         int    // 否 从第几条数据开始查询，默认0
	FetchSize      int    // 否 一次返回多少条数据，默认为10，最多不能超过100条
}

func (this GetActivityInfoSearchParam) APIName() string {
	return "lecloud.cloudlive.vrs.activity.vrsinfo.search"
}

func (this GetActivityInfoSearchParam) Params() map[string]string {
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

	if this.FetchSize <=0 {
		this.FetchSize = 10
	}
	m["fetchSize"] = fmt.Sprintf("%d", this.FetchSize)

	return m
}

func (this GetActivityInfoSearchParam) Method() string {
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
	ActivityName     string  // 是 直播活动名称(200个字符以内)
	StartTime        string  // 是 开始时间 格式yyyyMMddHHmmss
	EndTime          string  // 是 结束时间 格式yyyyMMddHHmmss
	CoverImgUrl      string  // 否 活动封面地址，如果为空，则系统会默认一张图片
	Description      string  // 否 活动描述（1024个字符以内）
	LiveNum          int     // 是 机位数量，范围为：1,2,3,4. 默认为1
	CodeRateTypes    string  // 是 流的码率类型，逗号分隔。由大到小排列。取值范围：10 流畅；13 标清；16 高清；19 超清；22 720P；25 1080P；99 原画
	NeedRecord       int     // 是 是否支持全程录制 0：否 1：是。默认为0
	NeedTimeShift    int     // 是 是否支持时移 0：否 1：是。默认为0
	NeedFullView     int     // 是 是否全景观看 0：否 1：是。默认为0
	ActivityCategory string  // 是 活动分类，参见《活动编码》
	PlayMode         int     // 是 播放模式，0：实时直播。1：流畅直播
}

func (this CreateActivityParam) APIName() string {
	return "lecloud.cloudlive.activity.create"
}

func (this CreateActivityParam) Params() map[string]string {
	var m = make(map[string]string)
	m["activityName"] = this.ActivityName
	m["startTime"]    = this.StartTime
	m["endTime"]      = this.EndTime
	if len(this.CoverImgUrl) > 0 {
		m["coverImgUrl"]  = this.CoverImgUrl
	}
	if len(this.Description) > 0 {
		m["description"]  = this.Description
	}
	if this.LiveNum < 1 || this.LiveNum > 4 {
		this.LiveNum = 1
	}
	m["liveNum"]       = fmt.Sprintf("%d", this.LiveNum)
	m["codeRateTypes"] = this.CodeRateTypes
	m["needRecord"]    = fmt.Sprintf("%d", this.NeedRecord)
	m["needTimeShift"] = fmt.Sprintf("%d", this.NeedTimeShift)
	m["needFullView"]  = fmt.Sprintf("%d", this.NeedFullView)
	m["activityCategory"] = this.ActivityCategory
	m["playMode"]         = fmt.Sprintf("%d", this.PlayMode)
	return m
}

func (this CreateActivityParam) Method() string {
	return "POST"
}