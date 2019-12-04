package apollotest

import (
	"fmt"
	"gopkg.in/yaml.v2"
)

//profile variables
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

var (
	config     *Conf
	configPath string
)

func NewConf() *Conf {
	if config != nil {
		return config
	}

	var (
		err      error
		yamlFile []byte
	)

	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		fmt.Println(err.Error())
	}

	return config
}
