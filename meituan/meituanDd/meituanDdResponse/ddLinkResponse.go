package meituanDdResponse

type DdLinkResponse struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
	Data LinkVO `json:"data"`
}

type LinkVO struct {
	Activity             string `json:"activity"`             //活动ID
	CommonLink           string `json:"commonLink"`           //活动长链接
	ShortLink            string `json:"shortLink"`            //活动短链接
	CommonQrCode         string `json:"commonQrCode"`         //活动二维码
	MiniProgramPath      string `json:"miniProgramPath"`      //活动小程序路径
	MiniProgramQrCode    string `json:"miniProgramQrCode"`    //活动小程序码图片链接
	MaterialDownloadLink string `json:"materialDownloadLink"` //活动物料文件链接
	BeginTime            int    `json:"beginTime"`            //活动开始时间（时间戳，单位秒）
	EndTime              int    `json:"endTime"`              //活动结束时间（时间戳，单位秒）
	RuleInfo             string `json:"ruleInfo"`             //活动推广规则详情
}
