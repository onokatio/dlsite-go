package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/k0kubun/pp"
)


type Rank []struct {
	Term     string `json:"term"`
	Category string `json:"category"`
	Rank     int    `json:"rank"`
	RankDate string `json:"rank_date"`
}

type RateCountDetail []struct {
	ReviewPoint int `json:"review_point"`
	Count       int `json:"count"`
	Ratio       int `json:"ratio"`
}

type LocalePrice struct {
	Usd float64 `json:"USD"`
	Eur float64 `json:"EUR"`
	Gbp float64 `json:"GBP"`
	Twd float64 `json:"TWD"`
	Cny float64 `json:"CNY"`
	Krw float64 `json:"KRW"`
}

type LocalePriceStr struct {
	Usd string `json:"USD"`
	Eur string `json:"EUR"`
	Gbp string `json:"GBP"`
	Twd string `json:"TWD"`
	Cny string `json:"CNY"`
	Krw string `json:"KRW"`
}

type ProductResponse map[string]struct {
	SiteID        string `json:"site_id"`
	SiteIDTouch   string `json:"site_id_touch"`
	MakerID       string `json:"maker_id"`
	AffiliateDeny int    `json:"affiliate_deny"`
	DlCount       string `json:"dl_count"` // string or null
	WishlistCount string `json:"wishlist_count"` // string or int 0
	DlFormat      int    `json:"dl_format"`
	Rank          Rank `json:"rank"`
	RateAverage     int     `json:"rate_average"`
	RateAverage2Dp  float64 `json:"rate_average_2dp"`
	RateAverageStar int     `json:"rate_average_star"`
	RateCount       int     `json:"rate_count"`
	RateCountDetail RateCountDetail`json:"rate_count_detail"`
	ReviewCount      string        `json:"review_count"` // string or int 0
	Price            int           `json:"price"`
	PriceWithoutTax  int           `json:"price_without_tax"`
	PriceStr         string        `json:"price_str"`
	DefaultPointRate int           `json:"default_point_rate"`
	DefaultPoint     int           `json:"default_point"`
	ProductPointRate interface{}   `json:"product_point_rate"`
	DlsiteplayWork   bool          `json:"dlsiteplay_work"`
	IsSale           bool          `json:"is_sale"`
	OnSale           int           `json:"on_sale"`
	IsDiscount       bool          `json:"is_discount"`
	IsPointup        bool          `json:"is_pointup"`
	Gift             []interface{} `json:"gift"`
	IsRental         bool          `json:"is_rental"`
	WorkRentals      []interface{} `json:"work_rentals"`
	UpgradeMinPrice  int           `json:"upgrade_min_price"`
	DownURL          string        `json:"down_url"`
	IsTartget        string        `json:"is_tartget"`
	TitleID          string        `json:"title_id"`
	TitleName        string        `json:"title_name"`
	IsTitleCompleted bool          `json:"is_title_completed"`
	BulkbuyKey       interface{}   `json:"bulkbuy_key"`
	Bonuses          []interface{} `json:"bonuses"`
	IsLimitWork      bool          `json:"is_limit_work"`
	IsSoldOut        bool          `json:"is_sold_out"`
	LimitStock       int           `json:"limit_stock"`
	IsReserveWork    bool          `json:"is_reserve_work"`
	IsReservable     bool          `json:"is_reservable"`
	IsTimesale       bool          `json:"is_timesale"`
	TimesaleStock    int           `json:"timesale_stock"`
	IsFree           bool          `json:"is_free"`
	IsOly            bool          `json:"is_oly"`
	IsLed            bool          `json:"is_led"`
	WorkName         string        `json:"work_name"`
	WorkImage        string        `json:"work_image"`
	DefaultPointStr string `json:"default_point_str"`
	LocalePrice     LocalePrice `json:"locale_price"`
	LocalePriceStr  LocalePriceStr`json:"locale_price_str"`
}

func main() {
	productid := "RJ302659"
	url := fmt.Sprintf("https://www.dlsite.com/maniax/product/info/ajax?product_id=%s&cdn_cache_min=1",productid)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	var product ProductResponse

	json.NewDecoder(res.Body).Decode(&product)

	pp.Print(url)
	pp.Print(product)
}
