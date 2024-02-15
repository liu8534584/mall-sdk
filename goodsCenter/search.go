package goodsCenter

import (
	"context"
	"github.com/liu8534584/mall-sdk/goodsCenter/request"
	"github.com/liu8534584/mall-sdk/goodsCenter/response"
)

func (c *Client) Search(ctx context.Context, req request.SearchReq) (search []response.ItemDetail, err error) {
	headers := make(map[string]string)
	if requestId, ok := ctx.Value("X-Request-ID").(string); ok {
		headers["X-Request-ID"] = requestId
	}
	var res *response.SearchRes
	err = c.Post("/open/v1/search/search").
		SetHeaders(headers).
		SetBody(req).
		Do(ctx).
		Into(&res)
	if res != nil {
		search = res.Data
	}
	return
}

func (c *Client) ResolveTitle(ctx context.Context, req request.ResolveTitle) (search response.ItemDetail, err error) {
	var res *response.ResolveTitleResponse
	headers := make(map[string]string)
	if requestId, ok := ctx.Value("X-Request-ID").(string); ok {
		headers["X-Request-ID"] = requestId
	}
	err = c.Post("/open/v1/item/resolveTitle").
		SetBody(req).
		SetHeaders(headers).
		Do(ctx).
		Into(&res)
	if res != nil {
		search = res.Data
	}
	return
}

func (c *Client) ItemDetail(ctx context.Context, req request.ItemDetailParams) (search response.ItemDetail, err error) {
	var res *response.ItemDetailResp
	headers := make(map[string]string)
	if requestId, ok := ctx.Value("X-Request-ID").(string); ok {
		headers["X-Request-ID"] = requestId
	}
	err = c.Post("/open/v1/item/getItemInfo").
		SetBody(req).
		SetHeaders(headers).
		Do(ctx).
		Into(&res)
	if res != nil {
		search = res.Data
	}
	return
}

func (c *Client) BatchItemInfo(ctx context.Context, req request.BathItemDetailParams) (search []response.ItemDetail, err error) {
	var res *response.BathItemDetailResp
	headers := make(map[string]string)
	if requestId, ok := ctx.Value("X-Request-ID").(string); ok {
		headers["X-Request-ID"] = requestId
	}
	err = c.Post("/open/v1/item/batchItemInfo").
		SetBody(req).
		SetHeaders(headers).
		Do(ctx).
		Into(&res)
	if res != nil {
		search = res.Data
	}
	return
}

func (c *Client) BatchItemInfoV2(ctx context.Context, req request.BathItemDetailV2Params) (search map[string]response.ItemDetail, err error) {
	var res *response.BathItemDetailV2Resp
	err = c.Post("/open/v1/item/batchItemInfoV2").
		SetBody(req).
		Do(ctx).
		Into(&res)
	if res != nil {
		search = res.Data
	}
	return
}

func (c *Client) GetTbFullItemId(ctx context.Context, req request.TbFullItemIdReq) (search response.TbFullItemIdResp, err error) {
	var res *response.TbFullItemIdResp
	headers := make(map[string]string)
	if requestId, ok := ctx.Value("X-Request-ID").(string); ok {
		headers["X-Request-ID"] = requestId
	}
	err = c.Post("/open/v1/item/getTbFullItemId").
		SetBody(req).
		SetHeaders(headers).
		Do(ctx).
		Into(&res)
	if res != nil {
		search = *res
	}
	return
}

func (c *Client) ShortItemIds2Long(ctx context.Context, req request.ShortItemIds2LongReq) (tbFullItemId map[string]response.TbFullItemId, err error) {
	var res *response.ShortItemIds2LongResp
	headers := make(map[string]string)
	if requestId, ok := ctx.Value("X-Request-ID").(string); ok {
		headers["X-Request-ID"] = requestId
	}
	err = c.Post("/open/v1/item/ShortItemIds2Long").
		SetBody(req).
		SetHeaders(headers).
		Do(ctx).
		Into(&res)
	if res != nil {
		tbFullItemId = res.Data
	}
	return
}
