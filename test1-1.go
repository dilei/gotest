/*
this a doc for go
*/
package main

import (
	"fmt"
	"strings"
	"time"
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
		for _, str := range strs {
			fmt.Println(ToCamelName(str))
		}
	*/

	for i := 0; i < 10; i++ {
		/*
			go func(i int) {
				fmt.Println(i)
			}(i)
		*/
		go func() {
			fmt.Println(i)
		}()
	}

	time.Sleep(2 * time.Second)

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
