package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Fundo struct {
	Nome               string `json:"nome"`
	Codigo             string `json:"codigo"`
	Preco              string `json:"preco"`
	Dividendo          string `json:"dividendo"`
	LiquidezDiaria     string `json:"liquidezDiaria"`
	DividendYield      string `json:"dividendYield"`
	PatrimonioLiquido  string `json:"patrimonioLiquido"`
	ValorPatrimonial   string `json:"valorPatrimonial"`
	RentabilidadeNoMes string `json:"rentabilidadeNoMes"`
	PVP                string `json:"p/vp"`
}

// FundsApi recebe um código de fundo imobiliário ou fiagro e retorna informações úteis.
func FundsApi(codigo string) (Fundo, error) {
	url := "https://www.fundsexplorer.com.br/funds/" + codigo

	req, err := http.Get(url)
	if err != nil {
		return Fundo{}, fmt.Errorf("erro ao criar requisição http... %v", err)
	}

	if strings.Split(req.Status, " ")[0] != strconv.Itoa(http.StatusOK) {
		return Fundo{}, fmt.Errorf("status code não é de sucesso... %v", err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36")

	defer req.Body.Close()

	doc, err := goquery.NewDocumentFromReader(req.Body)
	if err != nil {
		return Fundo{}, fmt.Errorf("erro ao obter o corpo da requisição... %v", err)
	}

	//código do fundo
	codigoDoFundo := doc.Find(`#head > div > div > div > div.ticker-wrapper > h1`).Text()

	//nome do fundo
	nomeDoFundo := doc.Find(`#head > div > div > div > h3`).Text()

	//preço
	preco := strings.TrimSpace(doc.Find(`#stock-price > span.price`).Text())

	//liquidez diária
	liquidezDiaria := strings.TrimSpace(doc.Find(`#main-indicators-carousel .carousel-cell:nth-child(1) .indicator-value`).Text())

	//último dividendo mensal
	dividendo := strings.TrimSpace(doc.Find(`#main-indicators-carousel .carousel-cell:nth-child(2) .indicator-value`).Text())

	//dividend yield
	dividendYield := strings.TrimSpace(doc.Find(`#main-indicators-carousel .carousel-cell:nth-child(3) .indicator-value`).Text())

	//patrimônio líquido
	patrimonioLiquido := strings.TrimSpace(doc.Find(`#main-indicators-carousel .carousel-cell:nth-child(4) .indicator-value`).Text())

	//valor patrimonial
	valorPatrimonial := strings.TrimSpace(doc.Find(`#main-indicators-carousel .carousel-cell:nth-child(5) .indicator-value`).Text())

	//rendabilidade no mês
	rentabilidadeNoMes := strings.TrimSpace(doc.Find(`#main-indicators-carousel div:nth-of-type(6) span:nth-of-type(2)`).Text())

	//p/vp
	pvp := strings.TrimSpace(doc.Find(`#main-indicators-carousel .carousel-cell:nth-child(7) .indicator-value`).Text())

	fundo := Fundo{
		Nome:               nomeDoFundo,
		Codigo:             codigoDoFundo,
		Preco:              preco,
		Dividendo:          dividendo,
		LiquidezDiaria:     liquidezDiaria,
		DividendYield:      dividendYield,
		PatrimonioLiquido:  patrimonioLiquido,
		ValorPatrimonial:   valorPatrimonial,
		RentabilidadeNoMes: rentabilidadeNoMes,
		PVP:                pvp,
	}

	return fundo, nil
}
