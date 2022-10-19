package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	fmt.Println(" Barkod numarası:")
	var number string // neyi arayacağımızı soruyoruz
	fmt.Scanln(&number)
	url := "#" /* burraya arama yapacağınız site url ni yazıyoruz */ + number

	res, _ := http.Get(url)

	if res.StatusCode == 200 {
		fmt.Println("hata", res.StatusCode)
		return

	}

	doc, _ := goquery.NewDocumentFromReader(res.Body)

	title := strings.TrimSpace(doc.Find("#").Text()) //buraya sitede aradığuımızın id veya classını yazıyoruz
	price := strings.TrimSpace(doc.Find("#").Text()) //buraya sitede aradığuımızın id veya classını yazıyoruz
	price = strings.Replace(price, "\n", " ", 1)
	fmt.Println("Kitap Adı:", title)
	fmt.Println("Fiyatı:", price)

}
