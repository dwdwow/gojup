package gojup

import (
	"fmt"
	"testing"
)

func TestGetPrices(t *testing.T) {
	prices, err := GetPrices(true, "JUPyiwrYJFskUPiHa7hkeR8VUtAeFoSYbKedZNsDvCN", "So11111111111111111111111111111111111111112")
	if err != nil {
		t.Fatalf("failed to get prices: %v", err)
	}
	fmt.Printf("%+v\n", prices)
}

func TestGetPricesVsToken(t *testing.T) {
	prices, err := GetPricesVsToken("JUPyiwrYJFskUPiHa7hkeR8VUtAeFoSYbKedZNsDvCN", "So11111111111111111111111111111111111111112", "HeLp6NuQkmYB4pYWo2zYs22mESHXPQYzXbB8n4V98jwC")
	if err != nil {
		t.Fatalf("failed to get prices: %v", err)
	}
	fmt.Printf("%+v\n", prices)
}
