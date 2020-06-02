/*
* Project: FLOTEA - Decentralized passenger transport system
* Copyright (c) 2020 Flotea, All Rights Reserved
* For conditions of distribution and use, see copyright notice in LICENSE
*/

package controllers

import (
    "crypto/tls"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net"
    "net/http"
    "net/mail"
    "net/smtp"
    "net/url"
    "time"
    global "token/global"

    "github.com/astaxie/beego"
    "github.com/astaxie/beego/validation"
)

type JSONAPIResponse struct {
    Success     bool      `json:"success"`
    ChallengeTS time.Time `json:"challenge_ts"` // timestamp of the challenge load (ISO format yyyy-MM-dd'T'HH:mm:ssZZ)
    Hostname    string    `json:"hostname"`     // the hostname of the site where the reCAPTCHA was solved
    ErrorCodes  []int     `json:"error-codes"`  //optional
}

type MainController struct {
    beego.Controller
}

func (c *MainController) setLayout() {
    c.Layout = "layouts/main.tpl"
    c.LayoutSections = make(map[string]string)
    c.LayoutSections["Header"] = "layouts/header.tpl"
    c.LayoutSections["Footer"] = "layouts/footer.tpl"
    c.Data["TplName"] = c.TplName[:len(c.TplName)-4]

    if c.Ctx.Request.Method == "POST" {
        u := global.Form{}
        if err := c.ParseForm(&u); err != nil {
            fmt.Println(err)
        }
        v := global.Validate(u)
        c.Data["errors"] = v
        c.Data["form"] = u
    }
}

func (c *MainController) Get() {
    c.TplName = "index.tpl"
    var Javascript = []string{
        "https://cdn.jsdelivr.net/npm/vue",
        "/static/js/vendors_applicants_carriers_icoBuy_icoVoting_main_voting.js",
        "/static/js/vendors_icoBuy_main.js",
        "/static/js/main.js",
    }
    c.Data["Javascript"] = Javascript
    c.setLayout()
}

func (c *MainController) WhitePaper() {
    c.TplName = "whitePaper.tpl"
    var Javascript = []string{}
    c.Data["Javascript"] = Javascript
    c.setLayout()
}

func (c *MainController) SendContact() {
    c.TplName = "sendedContact.tpl"
    c.setLayout()

    if len(c.Data["errors"].(map[string][]*validation.Error)) > 0 { // user press submit button without passing reCAPTCHA test
        c.Redirect(c.Ctx.Request.Referer()+"#footer", 307)
        return // return control to stop execution, otherwise it will continue
    }

    data := c.Data["form"].(global.Form)

    postStr := url.Values{"secret": {beego.AppConfig.String("reCAPTCHA")}, "response": {data.Recaptcha}, "remoteip": {beego.AppConfig.String("remoteip")}}
    responsePost, err := http.PostForm("https://www.google.com/recaptcha/api/siteverify", postStr)
    if err != nil {
        fmt.Println(err)
        return
    }

    defer responsePost.Body.Close()
    body, err := ioutil.ReadAll(responsePost.Body)
    if err != nil {
        fmt.Println(err)
        return
    }

    var APIResp JSONAPIResponse
    json.Unmarshal(body, &APIResp)
    if !APIResp.Success {
        c.Redirect(c.Ctx.Request.Referer()+"#footer", 307)
        return
    }

    from := mail.Address{"", "contact-flt@flotea.pl"}
    to := mail.Address{"", beego.AppConfig.String("contactEmailRecipient")}
    pass := "cq68RWmQoVf9C3k"

    subj := "FLT Contact form"
    messageBody := "Imię i nazwisko: " + data.NameSurname +
        "\nE-mail: " + data.Email +
        "\nNr telefonu: " + data.Phone +
        "\nWiadomość: " + data.Message

    // Setup headers
    headers := make(map[string]string)
    headers["From"] = from.String()
    headers["To"] = to.String()
    headers["Subject"] = subj

    // Setup message
    message := ""
    for k, v := range headers {
        message += fmt.Sprintf("%s: %s\r\n", k, v)
    }
    message += "\r\n" + messageBody

    servername := "flotea.pl:465"
    host, _, _ := net.SplitHostPort(servername)

    tlsconfig := &tls.Config{
        InsecureSkipVerify: true,
        ServerName:         host,
    }

    conn, err := tls.Dial("tcp", servername, tlsconfig)
    if err != nil {
        fmt.Println(err)
    }

    client, err := smtp.NewClient(conn, host)
    if err != nil {
        fmt.Println(err)
    }

    if err = client.Auth(smtp.PlainAuth("", from.Address, pass, host)); err != nil {
        fmt.Println(err)
    }

    // To && From
    if err = client.Mail(from.Address); err != nil {
        fmt.Println(err)
    }

    if err = client.Rcpt(to.Address); err != nil {
        fmt.Println(err)
    }

    w, err := client.Data()
    if err != nil {
        fmt.Println(err)
    }

    _, err = w.Write([]byte(message))
    if err != nil {
        fmt.Println(err)
    }

    err = w.Close()
    if err != nil {
        fmt.Println(err)
    }

    client.Quit()

    c.Data["form"] = global.EmptyForm()
}
