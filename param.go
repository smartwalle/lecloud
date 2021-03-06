package lecloud

////////////////////////////////////////////////////////////////////////////////
type ILeCloudParam interface {
	// 用于提供访问的 API 接口名称
	APIName() string

	// 返回参数列表
	Params() map[string]string

	Method() string
}

////////////////////////////////////////////////////////////////////////////////
// 示例,别无它用
type LeCloudParam map[string]interface{}

func (this LeCloudParam) APIName() string {
	return ""
}

func (this LeCloudParam) Method() string {
	return "POST"
}
