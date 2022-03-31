package main

import (
	"io"
	"fmt"
	"log"
	"strings"
	"net/http"
	"encoding/json"
	"github.com/k0kubun/pp"
	"astuart.co/goq"
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

type productAPIResponse map[string]struct {
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

type productHTMLTableKV struct {
	Key string `goquery:"th"`
	Values []string `goquery:"td a:not(.btn_follow)"`
}

type productHTMLTable struct {
	Outline []productHTMLTableKV `goquery:"table#work_outline > tbody > tr"`
	Maker []productHTMLTableKV `goquery:"table#work_maker > tbody > tr"`
}

type Chobit struct {
	struct {
	Count int `json:"count"`
	Works []struct {
		WorkID       string `json:"work_id"`
		DlsiteWorkID string `json:"dlsite_work_id"`
		WorkName     string `json:"work_name"`
		WorkNameKana string `json:"work_name_kana"`
		URL          string `json:"url"`
		EmbedURL     string `json:"embed_url"`
		Thumb        string `json:"thumb"`
		MiniThumb    string `json:"mini_thumb"`
		FileType     string `json:"file_type"`
		EmbedWidth   int    `json:"embed_width"`
		EmbedHeight  int    `json:"embed_height"`
	} `json:"works"`
}

type Product struct {
	productAPIResponse
	productTable map[string][]string
	chobit Chobit
}

func main() {
	productid := "BJ301460"
	url := fmt.Sprintf("https://www.dlsite.com/maniax/product/info/ajax?product_id=%s&cdn_cache_min=1",productid)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	product := Product{}
	json.NewDecoder(res.Body).Decode(&product.productAPIResponse)

	res, err = http.Get("https://www.dlsite.com/maniax/work/=/product_id/BJ301460.html")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	var table productHTMLTable
	err = goq.NewDecoder(res.Body).Decode(&table)
	if err != nil {
		log.Fatal(err)
	}

	product.productTable = map[string][]string{}
	for _, KV := range table.Maker {
		product.productTable[KV.Key] = KV.Values
	}
	for _, KV := range table.Outline {
		product.productTable[KV.Key] = KV.Values
	}
	pp.Println(product)

	res, err = http.Get("https://chobit.cc/api/v1/dlsite/embed?workno=RJ302659")
	defer res.Body.Close()
	bytes, _ := io.ReadAll(res.Body)
	str := strings.TrimPrefix(string(bytes),"response(")
	str = strings.TrimSuffix(str,")")
	strings.New
	pp.Println(str)
}
