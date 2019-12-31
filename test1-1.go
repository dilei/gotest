/*
this a doc for go
*/
package main

import (
	"strings"
)

// this a main func
func main() {
	/*
			values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

			resultChan := make(chan int, 2)
			go sum(values[:len(values)/2], resultChan)
			go sum(values[len(values)/2:], resultChan)
			sum1, sum2 := <-resultChan, <-resultChan

			fmt.Println("result:", sum1, sum2, sum1+sum2)

			url := fmt.Sprintf("%s", values)
			fmt.Println(url)

			strs := []string{"product_id", "product_name", "display_status", "is_direct_buy", "main_product_id", "standard_id", "is_move_store_disabled", "shop_id", "product_medium", "product_type", "category_id", "is_small_store_bh", "is_share_product", "is_catalog_product", "is_vip_discount", "is_shop_vip", "has_ebook", "ebook_product_id", "paper_book_product_id", "is_have_c2c", "book_id", "selected_template", "selected_product_id", "all_product_ids", "share_product_id", "shop_product_code", "describe_map", "package_size", "subname", "have_image", "num_images", "vendor_address", "vendor_tel", "device_id", "device_type", "is_have_infor", "sup_id", "share_display_priority", "input_date", "first_input_date", "special_sale", "icon_flag_life_five_star", "icon_flag_bang", "icon_flag_sole", "icon_flag_mall_new", "icon_flag_mall_hot", "icon_flag_character", "is_has_subsidy", "subsidy_amt", "is_has_limit_buy", "limit_buy_num", "service_provider_id", "service_provider_name", "cover_area_id", "cover_area", "is_has_360image", "is_integral_conversion", "is_monthly_payment", "one_ebook_subtype", "two_ebook_subtype", "device_type_name", "is_sold_out", "is_self_deliver", "is_single_package", "product_version", "relation_type", "support_seven_days", "activity_type"}
			strs := []string{
				"manufacturer", "aliases", "publisher", "author_id", "author_name", "set_id", "product_size", "series_id",
				"series_name", "print_copy", "paper_quality", "publish_date", "number_of_pages", "number_of_words",
				"version_num", "special_instruction", "mag_chinese_num", "mag_yearly_price", "mag_publish_num",
				"mag_total_publish_num", "mag_holding_company", "mag_publishing_cycle", "print_date", "author_format",
				"singer", "conductor", "director", "actors", "description_language", "prod_length", "effective_period",
				"subtitle_language", "audio_format", "video_format", "area", "material", "accessory", "screen_ratio",
				"binding", "packing_style", "originalcompany", "distributor", "producer", "award", "volumes", "feature",
				"import_num", "wyjz", "gqyz", "copyright", "singer_format", "director_format", "actor_format",
				"manufacture_format", "topical_word", "lyrics", "composer", "area_code", "main_book_name", "volume_name",
				"volume_number", "reign_title", "cip_number", "author_country", "editor", "publishing_brands", "is_suit",
				"barcode_img", "edition", "products_copyright_image", "painter", "translator", "copies_number",
				"version_type", "is_pinyin",
			}
		strs := []string{
			"is_small_store_bh", "is_cod_disabled", "is_cod_small_warehouse", "cod_small_warehouse_id", "shippingfee",
			"refund_expire_days", "confirm_refund_expire_days", "exchange_expire_days", "confirm_exchange_expire_days",
			"volume", "weight", "template_id", "is_gift_packaging", "is_direct_send",
		}
		for _, str := range strs {
			fmt.Println(ToCamelName(str))
		}
	*/

}

func ToCamelName(name string) string {
	// name = strings.Replace(name, "_", " ", -1)
	nameArr := strings.Split(name, "_")
	for key, val := range nameArr {
		val = strings.Title(val)
		if val == "Id" {
			val = "ID"
		}
		nameArr[key] = val
	}
	return strings.Join(nameArr, "")
}

func sum(values []int, resultChan chan int) {
	sum := 0
	for _, value := range values {
		sum += value
	}

	resultChan <- sum
}
