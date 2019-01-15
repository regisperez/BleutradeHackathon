package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func Decodificar() {
	// Request the HTML page.
	res, err := http.Get("https://etherscan.io/tx/0xb1ed364e4333aae1da4a901d5231244ba6a35f9421d4607f7cb90d60bf45578a")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Carrega a pagina
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Encontra a textarea da página

	texto := doc.Find("textarea").Text()

	decoded, err := hex.DecodeString(strings.TrimLeft(texto, "0x"))

	fmt.Printf("%s\n", decoded) // fmt.Println(strings.TrimLeft(teste,"0x"))

}

func main() {
	Decodificar()
}
