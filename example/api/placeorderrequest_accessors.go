// Code generated by "requestgen -type PlaceOrderRequest ./example/api"; DO NOT EDIT.

package api

import (
	"fmt"
)

func (p *PlaceOrderRequest) ClientOrderID(clientOrderID string) *PlaceOrderRequest {
	p.clientOrderID = &clientOrderID
	return p
}

func (p *PlaceOrderRequest) Symbol(symbol string) *PlaceOrderRequest {
	p.symbol = symbol
	return p
}

func (p *PlaceOrderRequest) Tag(tag string) *PlaceOrderRequest {
	p.tag = &tag
	return p
}

func (p *PlaceOrderRequest) Side(side SideType) *PlaceOrderRequest {
	p.side = side
	return p
}

func (p *PlaceOrderRequest) OrdType(ordType OrderType) *PlaceOrderRequest {
	p.ordType = ordType
	return p
}

func (p *PlaceOrderRequest) Size(size string) *PlaceOrderRequest {
	p.size = size
	return p
}

func (p *PlaceOrderRequest) Price(price string) *PlaceOrderRequest {
	p.price = &price
	return p
}

func (p *PlaceOrderRequest) TimeInForce(timeInForce TimeInForceType) *PlaceOrderRequest {
	p.timeInForce = &timeInForce
	return p
}

func (p *PlaceOrderRequest) ComplexArg(complexArg ComplexArg) *PlaceOrderRequest {
	p.complexArg = complexArg
	return p
}

func (p *PlaceOrderRequest) getParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}
	if p.clientOrderID != nil {
		clientOrderID := *p.clientOrderID
		if len(clientOrderID) == 0 {
			return params, fmt.Errorf("clientOid is required, empty string given")
		}
		params["clientOid"] = clientOrderID
	}
	symbol := p.symbol
	if len(symbol) == 0 {
		return params, fmt.Errorf("symbol is required, empty string given")
	}
	params["symbol"] = symbol
	if p.tag != nil {
		tag := *p.tag
		params["tag"] = tag
	}
	side := p.side
	if len(side) == 0 {
		return params, fmt.Errorf("side is required, empty string given")
	}
	switch side {
	case "buy", "sell":
		params["side"] = side

	default:
		return params, fmt.Errorf("side value %v is not valid", side)

	}
	params["side"] = side
	ordType := p.ordType
	switch ordType {
	case "limit", "market":
		params["ordType"] = ordType

	default:
		return params, fmt.Errorf("ordType value %v is not valid", ordType)

	}
	params["ordType"] = ordType
	size := p.size
	params["size"] = size
	if p.price != nil {
		price := *p.price
		params["price"] = price
	}
	if p.timeInForce != nil {
		timeInForce := *p.timeInForce
		switch timeInForce {
		case "GTC", "GTT", "FOK":
			params["timeInForce"] = timeInForce

		default:
			return params, fmt.Errorf("timeInForce value %v is not valid", timeInForce)

		}
		params["timeInForce"] = timeInForce
	}
	complexArg := p.complexArg
	params["complexArg"] = complexArg
	return params, nil
}
