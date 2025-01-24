package gojup

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"golang.org/x/time/rate"
)

const PRICE_API_URL = "https://api.jup.ag/price/v2"

type PriceConfidenceLevel string

const (
	PriceConfidenceLevelHigh   PriceConfidenceLevel = "high"
	PriceConfidenceLevelMedium PriceConfidenceLevel = "medium"
	PriceConfidenceLevelLow    PriceConfidenceLevel = "low"
)

type LastSwappedPrice struct {
	LastJupiterSellAt    int64   `json:"lastJupiterSellAt"`
	LastJupiterSellPrice float64 `json:"lastJupiterSellPrice,string"`
	LastJupiterBuyAt     int64   `json:"lastJupiterBuyAt"`
	LastJupiterBuyPrice  float64 `json:"lastJupiterBuyPrice,string"`
}

type QuotedPrice struct {
	BuyPrice  float64 `json:"buyPrice,string"`
	BuyAt     int64   `json:"buyAt"`
	SellPrice float64 `json:"sellPrice,string"`
	SellAt    int64   `json:"sellAt"`
}

type PriceDepth struct {
	Ten      float64 `json:"10"`
	Hundred  float64 `json:"100"`
	Thousand float64 `json:"1000"`
}

type BuyPriceImpactRatio struct {
	Depth     PriceDepth `json:"depth"`
	Timestamp int64      `json:"timestamp"`
}

type SellPriceImpactRatio struct {
	Depth     PriceDepth `json:"depth"`
	Timestamp int64      `json:"timestamp"`
}

type PriceDepthInfo struct {
	BuyPriceImpactRatio  BuyPriceImpactRatio  `json:"buyPriceImpactRatio"`
	SellPriceImpactRatio SellPriceImpactRatio `json:"sellPriceImpactRatio"`
}

type PriceExtraInfo struct {
	LastSwappedPrice LastSwappedPrice     `json:"lastSwappedPrice"`
	QuotedPrice      QuotedPrice          `json:"quotedPrice"`
	ConfidenceLevel  PriceConfidenceLevel `json:"confidenceLevel"`
	Depth            PriceDepthInfo       `json:"depth"`
}

type PriceInfo struct {
	ID        string          `json:"id"`
	Type      string          `json:"type"`
	Price     float64         `json:"price,string"`
	ExtraInfo *PriceExtraInfo `json:"extraInfo,omitempty"`
}

type PriceRespData struct {
	Data      map[string]PriceInfo `json:"data"`
	TimeTaken float64              `json:"timeTaken"`
}

var priceLimiter = rate.NewLimiter(rate.Every(time.Minute), 600)

func GetPrices(showExtraInfo bool, ids ...string) (PriceRespData, error) {
	if !priceLimiter.Allow() {
		return PriceRespData{}, fmt.Errorf("jupiter: rate limit exceeded")
	}
	err := priceLimiter.Wait(context.Background())
	if err != nil {
		return PriceRespData{}, fmt.Errorf("jupiter: rate limit exceeded")
	}
	url := fmt.Sprintf("%s?ids=%s&showExtraInfo=%t", PRICE_API_URL, strings.Join(ids, ","), showExtraInfo)
	body, err := simpleGet(url)
	if err != nil {
		return PriceRespData{}, err
	}
	var priceRespData PriceRespData
	if err := json.Unmarshal(body, &priceRespData); err != nil {
		return PriceRespData{}, err
	}
	return priceRespData, nil
}

func GetPricesVsToken(vsToken string, ids ...string) (PriceRespData, error) {
	if !priceLimiter.Allow() {
		return PriceRespData{}, fmt.Errorf("jupiter: rate limit exceeded")
	}
	err := priceLimiter.Wait(context.Background())
	if err != nil {
		return PriceRespData{}, fmt.Errorf("jupiter: rate limit exceeded")
	}
	url := fmt.Sprintf("%s?vsToken=%s&ids=%s", PRICE_API_URL, vsToken, strings.Join(ids, ","))
	body, err := simpleGet(url)
	if err != nil {
		return PriceRespData{}, err
	}
	var priceRespData PriceRespData
	if err := json.Unmarshal(body, &priceRespData); err != nil {
		return PriceRespData{}, err
	}
	return priceRespData, nil
}
