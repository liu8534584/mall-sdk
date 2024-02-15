package goodsCenter

import (
	"context"
	"github.com/liu8534584/mall-sdk/goodsCenter/request"
	"github.com/liu8534584/mall-sdk/goodsCenter/response"
)

func (c *Client) MeituanCategory(ctx context.Context, req request.MeituanCategoryReq) (search response.MeituanCategoryRes, err error) {
	var res *response.MeituanCategoryResp
	headers := make(map[string]string)
	if requestId, ok := ctx.Value("X-Request-ID").(string); ok {
		headers["X-Request-ID"] = requestId
	}
	params := map[string]interface{}{
		"cityId": req.CityId,
		"lng":    req.Lng,
		"lat":    req.Lat,
	}
	err = c.Get("/open/v1/meituan/category").
		SetQueryParamsAnyType(params).
		SetHeaders(headers).
		Do(ctx).
		Into(&res)
	if res != nil {
		search = res.Data
	}
	return
}
