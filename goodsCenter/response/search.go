package response

import "github.com/liu8534584/mall-sdk/goodsCenter/response/base"

type SearchRes struct {
	Code     base.Code     `json:"code"`
	Message  string        `json:"msg"`
	Demotion base.Demotion `json:"demotion"`
	Data     []ItemDetail  `json:"data"`
}

type ItemDetail struct {
	Item              Item               `json:"item"`
	Coupon            Coupon             `json:"coupon"`
	HasCoupon         bool               `json:"hasCoupon"`
	Extra             []Extra            `json:"extra"`
	HsRedPacket       HsRedPacket        `json:"hsRedPacket"`
	Shop              Shop               `json:"shop"`
	ShopTotalScore    ShopTotalScore     `json:"shopTotalScore"` //乐买买店铺平飞
	LinkParams        string             `json:"linkParams"`     //各渠道自己处理
	GzhParams         string             `json:"gzhParams"`
	IsShowGzh         bool               `json:"isShowGzh"`
	GzhUrl            string             `json:"gzhUrl"`
	IsPinGou          int64              `json:"isPinGou"`
	Pid               string             `json:"pid"`
	Rid               int64              `json:"rid"`
	FavoriteStatus    bool               `json:"favoriteStatus"`
	Tip               Tip                `json:"tip"`
	Ad                Ad                 `json:"ad"`
	IsFeizhuProduct   bool               `json:"isFeizhuProduct"`
	IsBiJia           bool               `json:"isBiJia"`
	IsShowPddSuperRed bool               `json:"isShowPddSuperRed"`
	PddSearchId       string             `json:"pddSearchId"`
	ZsDuoId           string             `json:"zsDuoId"`
	MaterialInfo      []ItemMaterialInfo `json:"materialInfo"`
	Cursor            int                `json:"cursor"`
	IsFromCloudSend   string             `json:"isFromCloudSend"`
}

type ItemMaterialInfo struct {
	ImageList       []string `json:"imageList"`
	VideoCoverImage string   `json:"videoCoverImage"`
	Video           string   `json:"video"`
	Content         string   `json:"content"`
	ItemId          string   `json:"itemId"`
	Title           string   `json:"title"`
	GoodType        uint8    `json:"goodType"`
	Id              uint64   `json:"id"`
}

type Ad struct {
	Type        string `json:"type"`
	Link        string `json:"link"`
	ImageUrl    string `json:"imageUrl"`
	AspectRatio int64  `json:"aspectRatio"`
}
type Tip struct {
	Image  string `json:"image"`
	Link   string `json:"link"`
	IsShow int64  `json:"isShow"`
}
type HsRedPacket struct {
	ShowText   string `json:"showText"`
	Money      string `json:"money"`
	Status     int64  `json:"status"`
	ChooseText string `json:"chooseText"`
}
type Coupon struct {
	ActivityId              string `json:"activityId"`
	ActivityUrl             string `json:"activityUrl"`
	Amount                  string `json:"amount"`
	Title                   string `json:"title"`
	CouponDate              string `json:"couponDate"`
	IsShopCoupon            int8   `json:"isShopCoupon"`
	CouponLimitPrice        string `json:"couponLimitPrice"`
	CouponPrice             int64  `json:"couponPrice"`
	GiftMoney               int64  `json:"giftMoney"`
	CouponEndTime           uint64 `json:"couponEndTime"`        // 结束时间
	CouponStartTime         uint64 `json:"couponStartTime"`      // 开始时间
	CouponRemainQuantity    uint64 `json:"couponRemainQuantity"` //优惠券库存
	CouponTotalQuantity     uint64 `json:"couponTotalQuantity"`  //优惠券总库存
	FailureTime             int64  `json:"failureTime"`
	IsCouponGetByActivityId bool   `json:"isCouponGetByActivityId"`
}

type Shop struct {
	ShopIcon string  `json:"shopIcon"`
	ShopId   string  `json:"shopId"`
	ShopName string  `json:"shopName"`
	IsSelf   bool    `json:"isSelf"` //是否自营
	Distance float64 `json:"distance"`
}
type Extra struct {
	Icon     int64  `json:"icon"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Link     string `json:"link"`
	Type     string `json:"type"`
}

type Item struct {
	GoodsSource          int64    `json:"goodsSource"`
	ItemUrl              string   `json:"itemUrl"`
	OriginTbItemId       string   `json:"originTbItemId"`
	BizSceneId           string   `json:"bizSceneId"`
	ItemId               string   `json:"itemId"`
	HsItemId             string   `json:"hsItemId"`
	ShopId               string   `json:"shopId"`
	GoodsId              uint64   `json:"goodsId"` //Pdd 用
	Image                string   `json:"image"`
	AuctionImages        []string `json:"auctionImages"` //轮播图
	DetailImages         []string `json:"detailImages"`  //详情图
	Recommend            string   `json:"recommend"`
	OriginPrice          string   `json:"oPrice"`   //原价
	Price                string   `json:"price"`    //券后价
	Discount             string   `json:"discount"` //折扣信息
	OriginPriceValue     int64    `json:"OriginPriceValue"`
	PriceValue           int64    `json:"priceValue"`
	Sales                int64    `json:"sale"`       //销量
	SalesTip             string   `json:"monthSales"` //销量tip
	Tabs                 []string `json:"tabs"`
	Title                string   `json:"title"`
	Tags                 []TbTags `json:"tags"`
	TabTemplateType      int64    `json:"tabTemplateType"`
	HasSubsidy           bool     `json:"hasSubsidy"`
	SubsidyAmount        int64    `json:"subsidyAmount"`
	Fee                  string   `json:"fee"`
	FeeValue             int64    `json:"feeValue"`
	RateInfo             string   `json:"rateInfo"`
	TkRate               int64    `json:"tkRate"`       //佣金率 1550表示15.5%
	ActivityTags         []uint64 `json:"activityTags"` //商品活动标记数组，例：[4,7]，4-秒杀 7-百亿补贴等
	PddSubsidyActivityId int      `json:"pddSubsidyActivityId"`
	KolNum               int64    `json:"kolNum"`    //带货达人
	KolNumTip            string   `json:"kolNumTip"` //带货达人tip
	KolAvatars           []string `json:"kolAvatars"`

	PreSell      PreSell `json:"preSell"` //预售
	PromotionTip string  `json:"promotionTip"`

	AdCode   string `json:"adCode"`
	DealType int64  `json:"dealType"`

	CashGiftAmount uint64 `json:"cash_gift_amount"`
}

type PreSell struct {
	PresaleDeposit         int64  `json:"presaleDeposit"`         //预售商品-定金（元）
	PresaleDiscountFeeText string `json:"presaleDiscountFeeText"` //售商品-商品优惠信息
	PresaleStartTime       int64  `json:"presaleStartTime"`       //预售开始时间
	PresaleEndTime         int64  `json:"presaleEndTime"`         //预售结束时间
}

type TbTags struct {
	Text  string `json:"text"`
	Color string `json:"color"`
}

type Bijia struct {
	IsBiJia        bool   `json:"isBijia"`
	BiJiaLink      string `json:"bijiaLink"`
	PddBiJiaNotice string `json:"pddBijiaNotice"`
}

type ShopTotalScore struct {
	// 商家体验分
	ShopScore ShopScore `json:"shopScore"`
	// 商品体验分
	ProductScore ProductScore `json:"productScore"`
	// 物流体验分
	LogisticsScore LogisticsScore `json:"logisticsScore"`
	// 商家服务分
	ServiceScore ServiceScore `json:"serviceScore"`
}

type OriginShopTotalScore struct {
	LogisticsScore LogisticsScore `json:"logistics_score"`
	ProductScore   ProductScore   `json:"product_score"`
	ServiceScore   ServiceScore   `json:"service_score"`
	ShopScore      ShopScore      `json:"shop_score"`
}

type ServiceScore struct {
	// 文本
	Text string `json:"text"`
	// 得分
	Score string `json:"score"`
	// 等级（1:高 2:中 3:低）
	Level int16 `json:"level"`
}

type ShopScore struct {
	// 文本
	Text string `json:"text"`
	// 得分
	Score string `json:"score"`
	// 等级（1:高 2:中 3:低）
	Level int16 `json:"level"`
}
type ProductScore struct {
	// 文本
	Text string `json:"text"`
	// 得分
	Score string `json:"score"`
	// 等级（1:高 2:中 3:低）
	Level int16 `json:"level"`
}
type LogisticsScore struct {
	// 文本
	Text string `json:"text"`
	// 得分
	Score string `json:"score"`
	// 等级（1:高 2:中 3:低）
	Level int16 `json:"level"`
}

type ItemSearchParam struct {
	Keyword     string   `json:"keyword"`
	ItemId      string   `json:"itemId"`
	ActivityId  string   `json:"activityId"`
	Page        int      `json:"page"`
	PageIndex   string   `json:"pageIndex"`
	PageSize    int      `json:"pageSize"`
	Cid         int      `json:"cid"`
	SecondCid   int      `json:"secondCid"`
	Sort        int      `json:"sort"`
	ItemIdList  []string `json:"itemIdList"`
	ExtraParams string   `json:"extraParams"`
}

type ItemSearchResponse struct {
	List      []ItemDetail `json:"list"`
	IsEnd     bool         `json:"isEnd"`
	PageIndex string       `json:"pageIndex"`
}

type TbFullItemIdResp struct {
	FullItemId   string `json:"fullItemId"`
	UniqueItemId string `json:"uniqueItemId"`
}
