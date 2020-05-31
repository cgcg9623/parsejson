//http 불러오기
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//type 만들기
type shopinfo struct {
	ShopID             int            `json:"shopid"`
	ShopName           string         `json:"shopname"`
	ShopNamealias      string         `json:"shopname_alias"`
	Addressjibun       string         `json:"address_jibun"`
	Addressdoro        string         `json:"address_doro"`
	Addressid          int            `json:"address_id"`
	MobileNum          string         `json:"mobilenum"`
	Callmentnum        string         `json:"callmentnum"`
	PointX             float32        `json:"point_x"`
	PointY             float32        `json:"point_y"`
	ImgURL             string         `json:"imgurl"`
	PicName            string         `json:"pic_name"`
	PicMobileNum       string         `json:"pic_mobilenum"`
	CorpNumber         string         `json:"corp_number"`
	VenderID           int            `json:"vendor_id"`
	AccountID          int            `json:"accountid"`
	AdminID            int            `json:"adminid"`
	Type               string         `json:"type"`
	ShopCategory       int            `json:"shop_category"`
	ActiveYN           string         `json:"activeyn"`
	Deleted            int            `json:"deleted"`
	DeliveryMemo       string         `json:"delivery_memo"`
	SVCStartDttm       string         `json:"svc_start_dttm"`
	ReversableInterval int            `json:"reservable_interval"`
	SVCAreaID          int            `json:"svc_area_id"`
	HrDayResvNoInv     int            `json:"hr_day_resv_no_inv"`
	OptionAc22YN       string         `json:"option_ac22_yn"`
	AdSentence         string         `json:"ad_sentence"`
	BlogURL            string         `json:"blog_url"`
	DawnDelivery       int            `json:"dawn_delivery"`
	IsEvent            int            `json:"is_event"`
	Capacity           int            `json:"capacity"`
	BusinessHour       string         `json:"business_hour"`
	Alias              string         `json:"alias"`
	OwnerName          string         `json:"ownername"`
	CorpAddress        string         `json:"corp_address"`
	UpTae              string         `json:"uptae"`
	UpJong             string         `json:"upjong"`
	AccountNumber      string         `json:"account_number"`
	Bank               string         `json:"bank"`
	AccountHolder      string         `json:"account_holder"`
	Email              string         `json:"email"`
	Memo               string         `json:"memo"`
	Options            Option         `json:"options"`
	SupportBrandIDS    SupportBrandID `json:"support_brand_ids"`
	ReviewCount        int            `json:"review_cnt"`
	AccountInfo        AccInfo        `json:"account_info"`
}

//타입 안에 타입
type Option struct {
	ShopID   int    `json:"shopid"`
	OptionID int    `json:"option_id"`
	SVCType  string `json:"svc_type"`
	Desc     string `json:"desc"`
}

//타입 안에 타입
type SupportBrandID struct {
	SupportBrandID []int
}

//타입 안에 타입
type AccInfo struct {
	AccountID int    `json:"accountid"`
	LoginID   string `json:"loginid"`
}

//시작은 main
func main() {
	response, err := http.Get("http://cardoc.co.kr/v2/mt_shop/42")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	DataInBytes, err := ioutil.ReadAll(response.Body)
	//PageContent := string(DataInBytes)

	var shop1 map[string]interface{}

	jsonErr := json.Unmarshal(DataInBytes, &shop1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	fmt.Println(shop1)

	//여기에 뭔가가 있어야함

	output, err := json.Marshal(shop1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(output))

	ioutil.WriteFile("shopinfo.json", output, 0644)
}
