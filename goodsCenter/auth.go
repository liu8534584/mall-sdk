package goodsCenter

import (
	"context"
	"github.com/liu8534584/mall-sdk/goodsCenter/request"
	"github.com/liu8534584/mall-sdk/goodsCenter/response"
)

func (c *Client) AuthStatus(ctx context.Context, req request.AuthStatusReq) (search response.AuthStatusRes, err error) {
	var res *response.AuthStatusResp
	headers := make(map[string]string)
	if requestId, ok := ctx.Value("X-Request-ID").(string); ok {
		headers["X-Request-ID"] = requestId
	}
	err = c.Post("/open/v1/auth/status").
		SetHeaders(headers).
		SetBody(req).
		Do(ctx).
		Into(&res)
	if res != nil {
		search = res.Data
	}
	return
}
