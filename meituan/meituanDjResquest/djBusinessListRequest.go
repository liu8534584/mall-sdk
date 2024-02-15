package meituanDjResquest

import (
	"encoding/json"
	"fmt"
	"github.com/liu8534584/mall-sdk/meituan"
	"github.com/liu8534584/mall-sdk/meituan/meituanDjResponse"
	"time"
)

type DjBussinessListRequest struct {
	param *DjBussinessListParams
	meituan.DjBaseRequest
}

func New(config *meituan.MeituanDjConfig) *DjBussinessListRequest {
	request := DjBussinessListRequest{param: &DjBussinessListParams{
		Token:     config.AccessToken,
		MediaSrc1: config.MediaSrc1,
		MediaSrc2: config.MediaSrc2,
		MediaPvId: fmt.Sprintf("HS_%d", time.Now().UnixNano()),
	}}
	request.SetConfig(config)
	request.SetClient(meituan.DefaultMeituanDjClient)
	return &request
}

func (m *DjBussinessListRequest) GetApiParamsObject() interface{} {

	return m.param
}

func (m *DjBussinessListRequest) GetParams() *DjBussinessListParams {
	return m.param
}

func (m *DjBussinessListRequest) GetApiMethod() string {
	return "POST"
}

func (m *DjBussinessListRequest) GetApiPath() string {
	return meituan.DjBaseUrl + "/act/cpsapi/get_list"
}

func (m *DjBussinessListRequest) Execute() (*meituanDjResponse.DjBusinessResponse, error) {
	execute, err := m.GetClient().Execute(m)
	if err != nil {
		return nil, err
	}
	var res meituanDjResponse.DjBusinessResponse
	_ = json.Unmarshal([]byte(execute), &res)

	return &res, err
}

type DjBussinessListParams struct {
	Token       string       `json:"token"`
	MediaSrc1   string       `json:"mediaSrc1"`
	MediaSrc2   string       `json:"mediaSrc2"`
	MediaPvId   string       `json:"mediaPvId"`
	Lat         string       `json:"lat"`         // 纬度-用户 LBS 位置坐标(保留小数点 6 位)
	Lon         string       `json:"lon"`         // 经度-用户 LBS 位置坐标(保留小数点 6 位)
	Deviceid    DeviceIdInfo `json:"deviceid"`    // 否 用户设备 id
	Unionid     string       `json:"unionid"`     // 否 微信小程序用户 id
	Phone       string       `json:"phone"`       // 否 手机号
	MediaUserid string       `json:"mediaUserid"` // 否 媒体用户标识，
	/**
	 * //  否（首页请求 不传，后续分 页要传）分页追踪标识，首次请求不携带，由美团侧在首
	 * 次请求时返回填充此值，在同一 mediaPvid 的
	 * 后续分页请求中需回传 pagePvId。
	 */
	PagePvId string `json:"pagePvId"`
}

type DeviceIdInfo struct {
	Id           string `json:"id"`             // 经度-用户 LBS 位置坐标(保留小数点 6 位)
	DeviceIdType string `json:"device_id_type"` //  // 必填，设备 id 类型 IMEI = 1、IDFA = 2、OAID = 3、ANDROID_ID = 4、MAC = 5、CAID = 6
}
