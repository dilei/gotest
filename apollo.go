package main

import (
	"bufio"
	"encoding/json"

	// "encoding/json"
	"fmt"
	"github.com/philchia/agollo"
	"os"
	"strings"
)

type APP struct {
	ImgType                  string          `json:"const_img_type"`
	ImgIdc                   int             `json:"const_img_idc"`
	SetUpTest                int             `json:"const_set_up_test"`
	FlashSalesProductType    []int           `json:"const_flash_sales_product_type"`
	BindProductType          int             `json:"const_bind_product_type"`
	StockAlarmNum            int             `json:"const_stock_alarm_num"`
	GlobalMallStock          []string        `json:"const_global_mall_stock"`
	GlobalMallWarehouse      []string        `json:"const_global_mall_warehouse"`
	GlobalPublishStock       []string        `json:"const_global_publish_stock"`
	GiftCardCategoryPathLow  int             `json:"const_gift_card_category_path_low"`
	GiftCardCategoryPathHigh int             `json:"const_gift_card_category_path_high"`
	ShopDDVipDiscount        map[int]float64 `json:"shop_ddvip_discount"`
	PartnerShop              []int64         `json:"partner_shop"`
	StockDealPromoType       []string        `json:"stock_deal_promo_type"`
	DDOwnerShopId            []int           `json:"ddowner_shop_id"`
}

func main() {
	err := agollo.StartWithConfFile("app.properties")
	if err != nil {
		fmt.Println(err)
		return
	}
	val := agollo.GetStringValue("partner_shop", "unknow")
	fmt.Println(val)
	res := agollo.GetNameSpaceContent("app.json", "unknow")
	fmt.Println(res)
	var app APP
	err = json.Unmarshal([]byte(res), &app)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(app)

	// 监听
	go func() {
		for {
			events := agollo.WatchUpdate()
			changeEvent := <-events
			bytes, _ := json.Marshal(changeEvent)
			fmt.Println("event:", string(bytes))
			res := agollo.GetNameSpaceContent("app.json", "unknow")
			fmt.Println(res)
		}
	}()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("is stop?(y/n)")
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		if strings.Compare("y", text) == 0 {
			fmt.Println("stop")
			break
		}
	}

	// 非kv形式的配置
	// res := agollo.GetNameSpaceContent("go-prodapi.yaml", "unknow")
	// fmt.Println(res)

	// res2 := agollo.GetAllKeys("application")
	// fmt.Println(res2)

	// 没这个方法了？
	// agollo.SubscribeToNamespaces("newNamespace1", "newNamespace2")
}

type Conf struct {
	Db struct {
		Product string `yaml:"product"`
		Price   string `yaml:"price"`
		Prod    string `yaml:"prod"`
		Promo   string `yaml:"promo"`
	}
	Memcache string `yaml:"memcache"`
	Url      struct {
		ImgVersionApi                 string `yaml:"img_version_api"`
		AttribImgPath                 string `yaml:"attrib_imgpath"`
		PriceQueryApi                 string `yaml:"price_query_api"`
		PromoFindProducts             string `yaml:"promo_find_products"`
		PromoFindPromotionsShop       string `yaml:"promo_find_promotions_shop"`
		PromoFindPromotionsProduct    string `yaml:"promo_find_promotions_product"`
		PromoFindPromotionsCollection string `yaml:"promo_find_promotions_collection"`
		PromoShopProductExcept        string `yaml:"promo_shop_product_except"`
		PromoLimitBuy                 string `yaml:"promo_limit_buy"`
		ShopApi                       string `yaml:"shop_api"`
		ShippingFeeApi                string `yaml:"shipping_fee_api"`
		QueryProdStockApi             string `yaml:"query_prod_stock_api"`
		WarehouseConfigApi            string `yaml:"warehouse_config_api"`
	}
	Constant struct {
		ImgType                  string   `yaml:"img_type"`
		ImgIdc                   int      `yaml:"img_idc"`
		SetUpTest                int      `yaml:"set_up_test"`
		FlashSalesProductType    []int    `yaml:"flash_sales_product_type"`
		BindProductType          int      `yaml:"bind_product_type"`
		StockAlarmNum            int      `yaml:"stock_alarm_num"`
		GlobalMallStock          []string `yaml:"global_mall_stock"`
		GlobalMallWarehouse      []string `yaml:"global_mall_warehouse"`
		GlobalPublishStock       []string `yaml:"global_publish_stock"`
		GiftCardCategoryPathLow  int      `yaml:"gift_card_category_path_low"`
		GiftCardCategoryPathHigh int      `yaml:"gift_card_category_path_high"`
	}
	ShopDDVipDiscount  map[int]float64 `yaml:"shopDDVipDiscount"`
	PartnerShop        []int64         `yaml:"partnerShop"`
	StockDealPromoType []string        `yaml:"stockDealPromoType"`
}
