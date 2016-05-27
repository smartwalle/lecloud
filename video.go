package lecloud

import "fmt"

////////////////////////////////////////////////////////////////////////////////
// GetVideoInfoParam 获取单个视频信息
// http://help.lecloud.com/Wiki.jsp?page=VideoInfo
type GetVideoInfoParam struct {
	VideoId string  // 必须 视频ID
}

func (this GetVideoInfoParam) APIName() string {
	return "video.get"
}

func (this GetVideoInfoParam) Params() map[string]string {
	var m = make(map[string]string)
	m["video_id"] = this.VideoId
	return m
}

func (this GetVideoInfoParam) Method() string {
	return "POST"
}

////////////////////////////////////////////////////////////////////////////////
// GetVideoListParam 获取视频列表信息
// http://help.lecloud.com/Wiki.jsp?page=VideoInfo
type GetVideoListParam struct {
	VideoName string // 可选    根据视频名称模糊搜索
	Index     int    // 可选    开始页索引,默认值为1
	Size      int    // 可选    分页大小, 默认值为10, 最大值为100
	Status    int    // 可选    视频状态：0表示全部；10表示可以正常播放；20表示转码失败；21表示审核失败；22表示片源错误；23表示发布失败；24表示上传失败；30表示处理中；31表示审核中；32表示无视频源；33表示上传初始化；34表示视频上传中；40表示停用；默认值为0
}

func (this GetVideoListParam) APIName() string {
	return "video.list"
}

func (this GetVideoListParam) Params() map[string]string {
	var m = make(map[string]string)
	if len(this.VideoName) > 0 {
		m["video_name"] = this.VideoName
	}
	if this.Index < 1 {
		this.Index = 1
	}
	m["index"] = fmt.Sprintf("%d", this.Index)

	if this.Size < 10 || this.Size > 100 {
		this.Size = 10
	}
	m["size"] = fmt.Sprintf("%d", this.Size)

	return m
}

func (this GetVideoListParam) Method() string {
	return "POST"
}

