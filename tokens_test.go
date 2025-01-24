package gojup

import (
	"fmt"
	"testing"
)

func TestGetTokensByTags(t *testing.T) {
	tokens, err := GetTokensByTags(TagStrict)
	if err != nil {
		t.Fatalf("failed to get tokens: %v", err)
	}
	println(len(tokens))
	fmt.Printf("%+v\n", tokens[0])
}

func TestGetTokenByMint(t *testing.T) {
	token, err := GetTokenByMint("So11111111111111111111111111111111111111112")
	if err != nil {
		t.Fatalf("failed to get token: %v", err)
	}
	fmt.Printf("%+v\n", token)
}

func TestGetTradableTokens(t *testing.T) {
	tokens, err := GetTradableTokens()
	if err != nil {
		t.Fatalf("failed to get tokens: %v", err)
	}
	println(len(tokens))
	fmt.Printf("%+v\n", tokens[0])
}
