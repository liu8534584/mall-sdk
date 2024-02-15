package goodsCenter

import (
	"context"
	"errors"
	"github.com/liu8534584/mall-sdk/goodsCenter/request"
	"github.com/liu8534584/mall-sdk/goodsCenter/response"
)

func (c *Client) TransLinkReq(ctx context.Context, req request.ItemTransLinkReq) (search response.ItemTransLink, err error) {
	var res *response.TransLinkRes
	headers := make(map[string]string)
	if requestId, ok := ctx.Value("X-Request-ID").(string); ok {
		headers["X-Request-ID"] = requestId
	}
	err = c.Post("/open/v1/transLink/itemTransLink").
		SetBody(req).
		SetHeaders(headers).
		Do(ctx).
		Into(&res)
	if err != nil || res == nil {
		return
	}
	search = res.Data
	if res.Code != 10001 {
		return search, errors.New(res.Message)
	}
	return
}

func (c *Client) TransLinkByLeagueAccountReq(ctx context.Context, req request.ItemTransLinkReq) (search response.ItemTransLink, err error) {
	var res *response.TransLinkRes
	headers := make(map[string]string)
	if requestId, ok := ctx.Value("X-Request-ID").(string); ok {
		headers["X-Request-ID"] = requestId
	}
	err = c.Post("/open/v1/transLink/itemTransLinkByLeagueAccount").
		SetBody(req).
		SetHeaders(headers).
		Do(ctx).
		Into(&res)
	if res != nil {
		search = res.Data
	}
	return
}

func (c *Client) BathTransLinkReq(ctx context.Context, req request.BathTransLinkReq) (search response.BathTransLink, err error) {
	var res *response.BathTransLinkRes
	headers := make(map[string]string)
	if requestId, ok := ctx.Value("X-Request-ID").(string); ok {
		headers["X-Request-ID"] = requestId
	}
	err = c.Post("/open/v1/transLink/bathTransLink").
		SetBody(req).
		SetHeaders(headers).
		Do(ctx).
		Into(&res)
	if res != nil {
		search = res.Data
	}
	return
}

func (c *Client) BathTransLinkV2(ctx context.Context, req request.BathTransLinkV2Req) (search response.BathTransLink, err error) {
	var res *response.BathTransLinkRes
	headers := make(map[string]string)
	if requestId, ok := ctx.Value("X-Request-ID").(string); ok {
		headers["X-Request-ID"] = requestId
	}
	err = c.Post("/open/v1/transLink/bathTransLinkV2").
		SetBody(req).
		SetHeaders(headers).
		Do(ctx).
		Into(&res)
	if res != nil {
		search = res.Data
	}
	return
}

func (c *Client) ShareMiniCodeReq(ctx context.Context, req request.ShareMiniCodeReq) (search response.ShareMiniCode, err error) {
	var res *response.ShareMiniCodeRes
	headers := make(map[string]string)
	if requestId, ok := ctx.Value("X-Request-ID").(string); ok {
		headers["X-Request-ID"] = requestId
	}
	err = c.Post("/open/v1/transLink/shareMiniCode").
		SetBody(req).
		SetHeaders(headers).
		Do(ctx).
		Into(&res)
	if res != nil {
		search = res.Data
	}
	return
}

func (c *Client) BatchRedirectUrl(ctx context.Context, req []request.RedirectUrlReq) (search []response.RedirectUrlRes, err error) {
	var res *response.BathRedirectUrlRes
	headers := make(map[string]string)
	if requestId, ok := ctx.Value("X-Request-ID").(string); ok {
		headers["X-Request-ID"] = requestId
	}
	err = c.Post("/open/v1/transLink/batchRedirectUrl").
		SetBody(req).
		SetHeaders(headers).
		Do(ctx).
		Into(&res)
	if res != nil {
		search = res.Data
	}
	return
}

func (c *Client) BatchRedirectUrlV2(ctx context.Context, req request.RedirectUrlV2Req) (search []response.RedirectUrlRes, err error) {
	var res *response.BathRedirectUrlRes
	headers := make(map[string]string)
	if requestId, ok := ctx.Value("X-Request-ID").(string); ok {
		headers["X-Request-ID"] = requestId
	}
	err = c.Post("/open/v1/transLink/batchRedirectUrlV2").
		SetBody(req).
		SetHeaders(headers).
		Do(ctx).
		Into(&res)
	if res != nil {
		search = res.Data
	}
	return
}

func (c *Client) BatchRedirectUrlV3(ctx context.Context, req request.RedirectUrlV2Req) (search []response.RedirectUrlRes, err error) {
	var res *response.BathRedirectUrlRes
	headers := make(map[string]string)
	if requestId, ok := ctx.Value("X-Request-ID").(string); ok {
		headers["X-Request-ID"] = requestId
	}
	err = c.Post("/open/v1/transLink/batchRedirectUrlV3").
		SetBody(req).
		SetHeaders(headers).
		Do(ctx).
		Into(&res)
	if res != nil {
		search = res.Data
	}
	return
}

func (c *Client) BatchRedirectUrlV4(ctx context.Context, req request.RedirectUrlV4Req) (search []response.RedirectUrlV4Res, err error) {
	var res *response.BathRedirectUrlV4Res
	headers := make(map[string]string)
	if requestId, ok := ctx.Value("X-Request-ID").(string); ok {
		headers["X-Request-ID"] = requestId
	}
	err = c.Post("/open/v1/transLink/batchRedirectUrlV4").
		SetBody(req).
		SetHeaders(headers).
		Do(ctx).
		Into(&res)
	if res != nil {
		search = res.Data
	}
	return
}

func (c *Client) ItemShareInfo(ctx context.Context, req request.ShareInfoReq) (search response.ItemShareInfo, err error) {
	var res *response.ItemShareInfoRes
	headers := make(map[string]string)
	if requestId, ok := ctx.Value("X-Request-ID").(string); ok {
		headers["X-Request-ID"] = requestId
	}
	err = c.Post("/open/v1/item/getShareInfo").
		SetBody(req).
		SetHeaders(headers).
		Do(ctx).
		Into(&res)
	if res != nil {
		search = res.Data
	}
	return
}
