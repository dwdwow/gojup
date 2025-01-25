package gojup

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/dwdwow/golimiter"
)

const TOKENS_BASE_URL = "https://tokens.jup.ag"

type Token struct {
	Address           string     `json:"address"`
	Name              string     `json:"name"`
	Symbol            string     `json:"symbol"`
	Decimals          int        `json:"decimals"`
	LogoURI           string     `json:"logoURI"`
	Tags              []string   `json:"tags"`
	DailyVolume       float64    `json:"daily_volume"`
	CreatedAt         string     `json:"created_at"`
	FreezeAuthority   string     `json:"freeze_authority"`
	MintAuthority     string     `json:"mint_authority"`
	PermanentDelegate string     `json:"permanent_delegate"`
	MintedAt          string     `json:"minted_at"`
	Extensions        Extensions `json:"extensions"`
}

type Extensions struct {
	CoingeckoID string `json:"coingeckoId"`
}

type TokenTag string

const (
	// Verified tokens displayed on jup.ag (superset of community and lst tags)
	TagVerified TokenTag = "verified"
	// Untagged tokens that display warnings
	TagUnknown TokenTag = "unknown"
	// Community verified tokens
	TagCommunity TokenTag = "community"
	// Previously validated tokens from strict-list (deprecated)
	TagStrict TokenTag = "strict"
	// Sanctum's LST list
	TagLST TokenTag = "lst"
	// Top 100 trending tokens from Birdeye
	TagBirdeyeTrending TokenTag = "birdeye-trending"
	// Tokens from Clone protocol
	TagClone TokenTag = "clone"
	// Pump tokens
	TagPump TokenTag = "pump"
)

var tokensLimiter = golimiter.NewReqLimiter(time.Minute, 30)

func GetTokensByTags(tags ...TokenTag) ([]Token, error) {
	if !tokensLimiter.Allow() {
		return nil, fmt.Errorf("jupiter: rate limit exceeded")
	}
	tokensLimiter.Wait(context.Background())
	tagStrs := make([]string, len(tags))
	for i, tag := range tags {
		tagStrs[i] = string(tag)
	}
	tagsParam := strings.Join(tagStrs, ",")
	url := fmt.Sprintf("%s/tokens?tags=%s", TOKENS_BASE_URL, tagsParam)
	body, err := simpleGet(url)
	if err != nil {
		return nil, err
	}
	var tokens []Token
	if err := json.Unmarshal(body, &tokens); err != nil {
		return nil, fmt.Errorf("jupiter: failed to decode response: %w", err)
	}
	return tokens, nil
}

func GetTokenByMint(mint string) (Token, error) {
	if !tokensLimiter.Allow() {
		return Token{}, fmt.Errorf("jupiter: rate limit exceeded")
	}
	tokensLimiter.Wait(context.Background())
	url := fmt.Sprintf("%s/token/%s", TOKENS_BASE_URL, mint)
	body, err := simpleGet(url)
	if err != nil {
		return Token{}, err
	}
	var token Token
	if err := json.Unmarshal(body, &token); err != nil {
		return Token{}, err
	}
	return token, nil
}

func GetTradableTokens() ([]Token, error) {
	if !tokensLimiter.Allow() {
		return nil, fmt.Errorf("jupiter: rate limit exceeded")
	}
	tokensLimiter.Wait(context.Background())
	url := fmt.Sprintf("%s/tokens_with_markets", TOKENS_BASE_URL)
	body, err := simpleGet(url)
	if err != nil {
		return nil, err
	}
	var tokens []Token
	if err := json.Unmarshal(body, &tokens); err != nil {
		return nil, fmt.Errorf("jupiter: failed to decode response: %w", err)
	}
	return tokens, nil
}
