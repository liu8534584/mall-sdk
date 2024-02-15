package meituanCard

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/forgoer/openssl"
	"github.com/liu8534584/mall-sdk/utils"
	"github.com/spf13/cast"
	"io"
	"net/http"
	"sort"
	"strings"
	"time"
)

// MeituanCardClient
type MeituanCardClient struct {
}

func NewMeituanCardClient() *MeituanCardClient {
	return &MeituanCardClient{}
}

func (m *MeituanCardClient) Execute(params MeituanCardRequest) (string, error) {
	contentMd5 := ""

	marshal, _ := json.Marshal(params.GetApiParamsObject())
	method := "GET"
	var url string
	if params.IsPost() {
		method = "POST"
		contentMd5 = m.base64AndMD5(marshal)
		url = CardBaseUrl + params.GetApiPath()
	} else {
		data := map[string]interface{}{}
		_ = json.Unmarshal(marshal, &data)
		param := map[string]string{}
		for k, v := range data {
			param[k] = cast.ToString(v)
		}
		url = m.setQuery(CardBaseUrl+params.GetApiPath(), param)
	}

	headers := map[string]string{
		"Content-Type":   "application/json; charset=UTF-8",
		"S-Ca-App":       params.GetConfig().AppKey,
		"S-Ca-Timestamp": cast.ToString(time.Now().UnixMilli()),
		"Content-MD5":    contentMd5,
	}
	signHeaderKey, signHeader := m.createSignHeader(headers)
	headers["S-Ca-Signature"] = m.createSign(params.GetConfig().AppSecret, method, contentMd5, params.GetApiPath(), signHeader)
	if signHeaderKey != "" {
		headers["S-Ca-Signature-Headers"] = signHeaderKey
	}

	if params.IsPost() {
		return m.HttpPostJson(url, string(marshal), headers)
	}

	return utils.HttpGet(url, time.Second*2)
}

func (m *MeituanCardClient) base64AndMD5(bytes []byte) string {
	if bytes == nil {
		return ""
	}
	md5hash := md5.New()
	md5hash.Write(bytes)
	md5Result := md5hash.Sum(nil)

	base64Result := base64.StdEncoding.EncodeToString(md5Result)
	return base64Result
}

func (m *MeituanCardClient) HttpPostJson(url string, data string, headers map[string]string) (string, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil && req == nil {
		return "", err
	}

	if req == nil {
		return "", err
	}

	client := &http.Client{
		Timeout: 2 * time.Second,
	}

	if len(headers) > 0 {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}
	//发送请求
	resp, err := client.Do(req)
	if resp == nil {
		return "", err
	}
	//关闭请求
	defer resp.Body.Close()

	if err != nil {
		return "", err
	}

	// 有gzip压缩时,需要解压缩读取返回内容
	//if resp.Header.Get("Content-Encoding") == "gzip" {
	//	reader, _ := gzip.NewReader(resp.Body) // gzip解压缩
	//	defer reader.Close()
	//	io.Copy(os.Stdout, reader)
	//}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (m *MeituanCardClient) createSignHeader(headers map[string]string) (string, string) {
	headerStr := ""
	headerKeyStr := ""
	if len(headers) > 0 {
		var keys []string
		for k := range headers {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, v := range keys {
			if v != "S-Ca-Signature" && v != "S-Ca-Signature-Headers" && v != "Content-MD5" {
				headerKeyStr += fmt.Sprintf("%s,", v)
				headerStr += fmt.Sprintf("%s:%s\n", v, headers[v])
			}
		}
	}
	return strings.TrimRight(headerKeyStr, ","), headerStr
}

func (m *MeituanCardClient) createSign(secret, method, contentMd5, url, hearderStr string) string {
	stringToSign := fmt.Sprintf("%s\n%s\n%s%s", method, contentMd5, hearderStr, url)

	keyBytes := []byte(secret)
	mac := hmac.New(sha256.New, keyBytes)
	mac.Write([]byte(stringToSign))
	signatureBytes := mac.Sum(nil)

	return base64.StdEncoding.EncodeToString(signatureBytes)
}

func (m *MeituanCardClient) setQuery(link string, params map[string]string) string {
	if params == nil {
		return link
	}

	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var ss = ""
	for _, v := range keys {
		if ss == "" {
			ss += v + "=" + params[v]
		} else {
			ss += "&" + v + "=" + params[v]
		}
	}

	return link + "?" + ss
}

func (m *MeituanCardClient) GetUtmMedium(appKey, pid string) string {
	src := []byte(pid)
	key := []byte(appKey)
	dst, _ := openssl.AesECBEncrypt(src, key, openssl.PKCS7_PADDING)
	return hex.EncodeToString(dst)
}

var DefaultMeituanCardClient *MeituanCardClient = NewMeituanCardClient()
