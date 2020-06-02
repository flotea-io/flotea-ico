/*
* Project: FLOTEA - Decentralized passenger transport system
* Copyright (c) 2020 Flotea, All Rights Reserved
* For conditions of distribution and use, see copyright notice in LICENSE
*/

package controllers

import (
    "github.com/astaxie/beego"
    //  "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "fmt"
    //"os"
    //"log"
    "context"
    "crypto/ecdsa"
    "encoding/json"
    "io/ioutil"
    "math/big"
    "net/http"
    "regexp"
    "strconv"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/crypto"

    //"github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"

    floico "token/contract/ico"
    flotoken "token/contract/token"
)

const FloteaTokenAddress = "0x1C94D0D2F4ED3E8Cd38B0cBA9C8d05e6648fdaa2"
const FloteaICOAddress = "0xDa8370898A95446875DaC7AeAd92E5fd5851174D"

const FloteaICOOwnerPrivate = ""

type IcoController struct {
    beego.Controller
}

type jsonStruct struct {
    Data  string `json:"data"`
    Error bool   `json:"error"`
}

type infoStruct struct {
    InitialTokens     *big.Int   `json:"initialTokens"`
    TokensAvailable   *big.Int   `json:"tokensAvailable"`
    Phase             *big.Int   `json:"phase"`
    LeftInActualPhase *big.Int   `json:"leftInActualPhase"`
    Initialized       bool       `json:"initialized"`
    Enabled           bool       `json:"enabled"`
    Prices            []*big.Int `json:"prices"`
    Tokens            []*big.Int `json:"tokens"`
    BalanceWei        *big.Int   `json:"balanceWei"`
    Error             bool       `json:"error"`
}

func (c *IcoController) setLayout() {
    c.Layout = "layouts/main.tpl"
    c.LayoutSections = make(map[string]string)
    c.LayoutSections["Header"] = "layouts/header.tpl"
    c.LayoutSections["Footer"] = "layouts/footer.tpl"
    c.Data["TplName"] = c.TplName[:len(c.TplName)-4]
}

func (c *IcoController) getClient() *ethclient.Client {
    //client, err := ethclient.Dial("http://localhost:7545")
    client, err := ethclient.Dial("https://kovan.infura.io/v3/90952d0af11c4bccbcee97c79100061d")
    if err != nil {
        c.Response(err.Error(), true)
    }
    fmt.Println("we have a connection")

    return client
}

func (c *IcoController) getIco() *floico.Floico {
    client := c.getClient()

    address := common.HexToAddress(FloteaICOAddress)
    instance, err := floico.NewFloico(address, client)
    if err != nil {
        c.Response(err.Error(), true)
    }
    fmt.Println("ico is loaded")

    return instance
}

func (c *IcoController) getToken() *flotoken.Flotoken {
    //client, err := ethclient.Dial("http://localhost:7545")
    client, err := ethclient.Dial("https://kovan.infura.io/v3/90952d0af11c4bccbcee97c79100061d")
    if err != nil {
        c.Response(err.Error(), true)
    }
    fmt.Println("we have a connection")

    address := common.HexToAddress(FloteaTokenAddress)
    instance, err := flotoken.NewFlotoken(address, client)
    if err != nil {
        c.Response(err.Error(), true)
    }
    fmt.Println("token is loaded")

    return instance
}

func (c *IcoController) Response(s string, error bool) {
    jsonStruct := jsonStruct{Data: s, Error: error}
    c.Data["json"] = &jsonStruct
    c.ServeJSON()
}

func (c *IcoController) Get() {
    c.TplName = "ico.tpl"
    var Javascript = []string{
        // "https://cdn.jsdelivr.net/npm/vue",
        "https://js.stripe.com/v3/",
        "https://cdnjs.cloudflare.com/ajax/libs/jquery-jgrowl/1.4.6/jquery.jgrowl.min.js",
        "/static/js/vendors_applicants_carriers_icoBuy_icoVoting_main_voting.js",
        "/static/js/vendors_applicants_carriers_icoBuy_icoVoting_voting.js",
        "/static/js/vendors_icoBuy_main.js",
        "/static/js/vendors_icoBuy.js",
        "/static/js/applicants_carriers_icoBuy_icoVoting_voting.js",
        "/static/js/icoBuy.js",
    }
    c.Data["Javascript"] = Javascript
    c.setLayout()
}

func (c *IcoController) Vote() {
    c.TplName = "vote.tpl"
    var Javascript = []string{
        "/static/js/vendors_applicants_carriers_icoBuy_icoVoting_main_voting.js",
        "/static/js/vendors_applicants_carriers_icoBuy_icoVoting_voting.js",
        "/static/js/vendors_applicants_icoVoting.js",
        "/static/js/applicants_carriers_icoBuy_icoVoting_voting.js",
        "/static/js/icoVoting.js",
    }
    c.Data["Javascript"] = Javascript
    c.setLayout()
}

func (c *IcoController) Info() {
    ret1, ret2, ret3, ret4, ret5, ret6, ret7, ret8, ret9, err := c.getIco().Info(nil)
    if err != nil {
        c.Response(err.Error(), true)
    } else {

        infoStruct := infoStruct{
            InitialTokens:     ret1,
            TokensAvailable:   ret2,
            Phase:             ret3,
            LeftInActualPhase: ret4,
            Initialized:       ret5,
            Enabled:           ret6,
            Prices:            ret7,
            Tokens:            ret8,
            BalanceWei:        ret9,
            Error:             false}
        c.Data["json"] = &infoStruct
        c.ServeJSON()
    }
}

func (c *IcoController) Initialized() bool {
    _, _, _, _, ret5, _, _, _, _, err := c.getIco().Info(nil)
    if err != nil {
        return false
    }
    return ret5
}

func (c *IcoController) Address() {
    c.Response(FloteaICOAddress, false)
}

func (c *IcoController) Price() {
    param := c.Ctx.Input.Param(":tokens")
    tokens, ok := new(big.Int).SetString(param, 10)
    if !ok || tokens.Cmp(big.NewInt(0)) < 1 {
        c.Response("Error in parse BigInt", true)
        return
    }
    if tokens.Cmp(big.NewInt(72*1000000*1000)) > 0 {
        c.Response("Max tokens is 72 M", true)
        return
    }

    if !c.Initialized() {
        c.Response("Token is not initialized", true)
        return
    }

    price, err := c.getIco().GetPrice(nil, tokens)
    if err != nil {
        c.Response(err.Error(), true)
    } else {
        c.Response(price.String(), false)
    }
}

func (c *IcoController) PriceForAmount() {
    param := c.Ctx.Input.Param(":amount")
    amount, ok := new(big.Int).SetString(param, 10)
    if !ok {
        c.Response("Error in parse BigInt", true)
        return
    }
    if amount.Cmp(big.NewInt(1428571430000)) < 0 {
        c.Response("Minimum price for token is 1428571430000 WEI", true)
        return
    }

    if !c.Initialized() {
        c.Response("Token is not initialized", true)
        return
    }

    tokens, err := c.getIco().GetTokensForWei(nil, amount)
    if err != nil {
        c.Response(err.Error(), true)
    } else {
        c.Response(tokens.String(), false)
    }
}

type addressStruct struct {
    Amount *big.Int `json:"amount"`
    Valid  bool     `json:"valid"`
    Error  bool     `json:"error"`
}

func (c *IcoController) CheckAddress() {
    s := c.Ctx.Input.Param(":address")
    re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
    valid := re.MatchString(s)
    addressStruct := addressStruct{
        Amount: big.NewInt(0),
        Valid:  valid,
        Error:  false}
    if valid {
        amount, err := c.getToken().BalanceOf(nil, common.HexToAddress(s))
        if err != nil {
            c.Response(err.Error(), true)
            return
        } else {
            addressStruct.Amount = amount
        }
    }
    c.Data["json"] = &addressStruct
    c.ServeJSON()
}

func (c *IcoController) EthPrice() {
    c.Response(fmt.Sprintf("%f", c.getEthPrice()), false)
}

/*type ticker struct {
    Max     float64 `json:"max"`
    Min     float64 `json:"min"`
    Last    float64 `json:"last"`
    Bid     float64 `json:"bid"`
    Ask     float64 `json:"ask"`
    Vwap    float64 `json:"vwap"`
    Average float64 `json:"average"`
    Volume  float64 `json:"volume"`
}*/

type TickerItem struct {
    TradeId  float64 `json:"trade_id"`
    Type     string  `json:"type"`
    Quantity string  `json:"quantity"`
    Price    string  `json:"price"`
    Amount   string  `json:"amount"`
    Date     float64 `json:"date"`
}

func (c *IcoController) getEthPrice() float64 {

    //url := "https://bitbay.net/API/Public/ETHPLN/ticker.json"
    url := "https://api.exmo.com/v1/trades/?pair=ETH_PLN&limit=1"

    res, err := http.Get(url)
    if err != nil {
        fmt.Println(err.Error())
        c.Response(err.Error(), true)
    }
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        fmt.Println(err.Error())
        c.Response(err.Error(), true)
    }

    var tick map[string][]TickerItem
    jsonErr := json.Unmarshal([]byte(body), &tick)
    if jsonErr != nil {
        fmt.Println(err.Error())
        c.Response(jsonErr.Error(), true)
    } else {
        f, _ := strconv.ParseFloat(tick["ETH_PLN"][0].Price, 64)
        return f
    }
    return -1
}

func (c *IcoController) buyTokens(address string, tokens *big.Int, s *StripeController) {
    privateKey, err := crypto.HexToECDSA(FloteaICOOwnerPrivate)
    if err != nil {
        s.Response(err.Error(), true)
    }

    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        s.Response("error casting public key to ECDSA", true)
    }

    client := c.getClient()
    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        s.Response(err.Error(), true)
    }

    auth := bind.NewKeyedTransactor(privateKey)
    auth.Nonce = big.NewInt(int64(nonce))
    auth.Value = big.NewInt(0)
    auth.From = fromAddress

    ret, err := c.getIco().BuyOverBackend(auth, common.HexToAddress(address), tokens)
    if err != nil {
        s.Response(err.Error(), true)
        return
    } else {
        s.Response(ret.Hash().String(), false)
        return
    }
}

func (c *IcoController) Balancepln() {
    stripe := StripeController{}
    balance := stripe.Balance()
    c.Response(strconv.FormatInt(balance, 10), false)
}
