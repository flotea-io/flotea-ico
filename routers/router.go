/*
* Project: FLOTEA - Decentralized passenger transport system
* Copyright (c) 2020 Flotea, All Rights Reserved
* For conditions of distribution and use, see copyright notice in LICENSE
*/

package routers

import (
    "token/controllers"

    "github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("white-paper", &controllers.MainController{}, "*:WhitePaper")
    beego.Router("send-contact", &controllers.MainController{}, "post:SendContact")

    beego.AutoRouter(&controllers.VoteController{})

    beego.Router("/ico/price/:tokens", &controllers.IcoController{}, "*:Price")
    beego.Router("/ico/ethprice", &controllers.IcoController{}, "*:EthPrice")
    beego.Router("/ico/tokens/:amount", &controllers.IcoController{}, "*:PriceForAmount")
    beego.Router("/ico/checkaddress/:address", &controllers.IcoController{}, "*:CheckAddress")
    beego.AutoRouter(&controllers.IcoController{})
    beego.Router("/stripe/buy", &controllers.StripeController{}, "*:Buy")
    beego.AutoRouter(&controllers.StripeController{})
    beego.Router("/transport", &controllers.TransportController{})
}
