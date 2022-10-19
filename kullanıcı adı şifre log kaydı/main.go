package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	kullanıcıadı := "admin"
	şifre := "1234"
	girişhakkı := 5
	giriş := false
	fmt.Println("Hoş geldiniz kullanıcı adı ve şifre girerek sisteme giriş yapabilirsiniz:")
	fmt.Println("Toplam giriş deneme hakkınız:", girişhakkı)

	for girişhakkı > 0 {
		var kullanıcıadıgir string
		var şifregir string
		var durum string
		fmt.Print("Kullanıcı adı:")
		fmt.Scan(&kullanıcıadıgir)
		fmt.Print("Şifre:")
		fmt.Scan(&şifregir)
		time.Now()
		if kullanıcıadıgir == kullanıcıadı && şifregir == şifre {
			fmt.Println("Giriş başarılı")
			durum = "Başarılı"
			giriş = true

		} else {
			girişhakkı--
			fmt.Println("Kullanıcı adı veya şifre yanlış")
			fmt.Println("Kalan giriş hakkınız:", girişhakkı)
			durum = "Başarısız"
			giriş = false
		}
		file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			log.Default().Println(err)
		}
		defer file.Close()

		len, err := file.WriteString("Kullanıcı adı: " + kullanıcıadıgir + "\n" + "Şifre: " + şifregir + "\n" + "Zaman: " + time.Now().Local().String() + "\n" + "Durum: " + fmt.Sprint(durum) + "\n" + "----------------------------------------" + "\n")
		if err != nil {
			log.Default().Println(err)
		} else {
			fmt.Println(len, "karakter yazıldı.")
		}
		if giriş == true {
			break
		}

	}
	if girişhakkı == 0 {
		fmt.Println("Giriş hakkınız bitti")
	} else {
		fmt.Println("Giriş başarılı")
	}

}
