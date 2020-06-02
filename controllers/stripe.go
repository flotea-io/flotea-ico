/*
* Project: FLOTEA - Decentralized passenger transport system
* Copyright (c) 2020 Flotea, All Rights Reserved
* For conditions of distribution and use, see copyright notice in LICENSE
*/

package controllers

import (
    "fmt"
    "math"
    "math/big"
    "strconv"

    "github.com/astaxie/beego"
    "github.com/stripe/stripe-go"
    "github.com/stripe/stripe-go/charge"
    //  "os"
)

const stripeKey = "sk_test_347Dngpx4AVWUxFXAYrE557s"

type StripeController struct {
    beego.Controller
}

func (c *StripeController) Response(s string, error bool) {
    jsonStruct := jsonStruct{Data: s, Error: error}
    c.Data["json"] = &jsonStruct
    c.ServeJSON()
}

func (c *StripeController) Balance() int64 {
    stripe.Key = stripeKey
    amount := int64(0)
    params := &stripe.ChargeListParams{}
    params.Filters.AddFilter("limit", "", "100")
    i := charge.List(params)
    for i.Next() {
        c := i.Charge()
        if c.Description == "Buy FLotea token" {
            amount += c.Amount
        }
    }
    return amount
}

func (c *StripeController) Buy() {
    ico := IcoController{}

    param := c.GetString("tokens")
    source := c.GetString("source")
    address := c.GetString("address")

    tokens, ok := new(big.Int).SetString(param, 10)
    fmt.Println(param)
    if !ok {
        c.Response("Error in parse BigInt", true)
        return
    }
    if tokens.Cmp(big.NewInt(0)) < 1 {
        c.Response("Too small assset", true)
        return
    }
    if tokens.Cmp(big.NewInt(72*1000000*1000)) > 0 {
        c.Response("Max tokens is 72 M", true)
        return
    }

    if !ico.Initialized() {
        c.Response("Token is not initialized", true)
        return
    }

    ethPrice := ico.getEthPrice()

    price, err := ico.getIco().GetPrice(nil, tokens)
    if err != nil {
        c.Response(err.Error(), true)
    } else {
        floatPrice := new(big.Float).Mul( //*
            new(big.Float).Quo( // /
                new(big.Float).SetInt(price),
                big.NewFloat(math.Pow10(16))),
            big.NewFloat(ethPrice))
        intPrice := new(big.Int)
        floatPrice.Int(intPrice)
        if intPrice.Cmp(big.NewInt(1)) == -1 {
            c.Response("Too small assset", true)
            return
        }

        // Payment
        stripe.Key = stripeKey
        chargeParams := &stripe.ChargeParams{
            Amount:      stripe.Int64(intPrice.Int64()),
            Currency:    stripe.String(string(stripe.CurrencyPLN)),
            Description: stripe.String("Buy FLotea token"),
        }
        chargeParams.SetSource(source) // obtained with Stripe.js
        ch, err := charge.New(chargeParams)
        if err != nil {
            c.Response(err.Error(), true)
            return
        } else {
            fmt.Println(ch)
            if ch.Paid {
                // Transfer tokens
                ico.buyTokens(address, tokens, c)
                return
            }
            c.Response(strconv.FormatBool(ch.Paid), false)
        }
    }
}
