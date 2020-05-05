!function(e){function n(n){for(var r,i,_=n[0],a=n[1],c=n[2],p=0,d=[];p<_.length;p++)i=_[p],Object.prototype.hasOwnProperty.call(s,i)&&s[i]&&d.push(s[i][0]),s[i]=0;for(r in a)Object.prototype.hasOwnProperty.call(a,r)&&(e[r]=a[r]);for(u&&u(n);d.length;)d.shift()();return o.push.apply(o,c||[]),t()}function t(){for(var e,n=0;n<o.length;n++){for(var t=o[n],r=!0,_=1;_<t.length;_++){var a=t[_];0!==s[a]&&(r=!1)}r&&(o.splice(n--,1),e=i(i.s=t[0]))}return e}var r={},s={icoBuy:0},o=[];function i(n){if(r[n])return r[n].exports;var t=r[n]={i:n,l:!1,exports:{}};return e[n].call(t.exports,t,t.exports,i),t.l=!0,t.exports}i.m=e,i.c=r,i.d=function(e,n,t){i.o(e,n)||Object.defineProperty(e,n,{enumerable:!0,get:t})},i.r=function(e){"undefined"!=typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},i.t=function(e,n){if(1&n&&(e=i(e)),8&n)return e;if(4&n&&"object"==typeof e&&e&&e.__esModule)return e;var t=Object.create(null);if(i.r(t),Object.defineProperty(t,"default",{enumerable:!0,value:e}),2&n&&"string"!=typeof e)for(var r in e)i.d(t,r,function(n){return e[n]}.bind(null,r));return t},i.n=function(e){var n=e&&e.__esModule?function(){return e.default}:function(){return e};return i.d(n,"a",n),n},i.o=function(e,n){return Object.prototype.hasOwnProperty.call(e,n)},i.p="/static/js/";var _=window.webpackJsonp=window.webpackJsonp||[],a=_.push.bind(_);_.push=n,_=_.slice();for(var c=0;c<_.length;c++)n(_[c]);var u=a;o.push(["./src/icoBuy.ts","vendors_applicants_carriers_icoBuy_icoVoting_main_voting","vendors_applicants_carriers_icoBuy_icoVoting_voting","vendors_icoBuy_main","vendors_icoBuy","applicants_carriers_icoBuy_icoVoting_voting"]),t()}({"./src/ico.service.ts":
/*!****************************!*\
  !*** ./src/ico.service.ts ***!
  \****************************/
/*! exports provided: default, icoService */function(module,__webpack_exports__,__webpack_require__){"use strict";eval('__webpack_require__.r(__webpack_exports__);\n/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "icoService", function() { return icoService; });\n/* harmony import */ var axios__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! axios */ "./node_modules/axios/index.js");\n/* harmony import */ var axios__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(axios__WEBPACK_IMPORTED_MODULE_0__);\n\nvar api = \'/ico\';\nvar IcoService = /** @class */ (function () {\n    function IcoService() {\n    }\n    IcoService.prototype.getInfo = function () {\n        return axios__WEBPACK_IMPORTED_MODULE_0___default.a.get(api + "/info");\n    };\n    IcoService.prototype.getTotalPrice = function (tokens) {\n        return axios__WEBPACK_IMPORTED_MODULE_0___default.a.get(api + "/price/" + tokens);\n    };\n    IcoService.prototype.getTokensForAmount = function (tokens) {\n        return axios__WEBPACK_IMPORTED_MODULE_0___default.a.get(api + "/tokens/" + tokens);\n    };\n    IcoService.prototype.getEthPrice = function () {\n        return axios__WEBPACK_IMPORTED_MODULE_0___default.a.get(api + "/ethprice");\n    };\n    IcoService.prototype.checkAddress = function (address) {\n        return axios__WEBPACK_IMPORTED_MODULE_0___default.a.get(api + "/checkaddress/" + address);\n    };\n    IcoService.prototype.getAddress = function () {\n        return axios__WEBPACK_IMPORTED_MODULE_0___default.a.get(api + "/address");\n    };\n    IcoService.prototype.getBalancePln = function () {\n        return axios__WEBPACK_IMPORTED_MODULE_0___default.a.get(api + "/balancepln");\n    };\n    return IcoService;\n}());\n/* harmony default export */ __webpack_exports__["default"] = (IcoService);\nvar icoService = new IcoService();\n\n\n//# sourceURL=webpack:///./src/ico.service.ts?')},"./src/icoBuy.ts":
/*!***********************!*\
  !*** ./src/icoBuy.ts ***!
  \***********************/
/*! no exports provided */function(module,__webpack_exports__,__webpack_require__){"use strict";eval('__webpack_require__.r(__webpack_exports__);\n/* harmony import */ var vue__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! vue */ "./node_modules/vue/dist/vue.esm.js");\n/* harmony import */ var lodash__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! lodash */ "./node_modules/lodash/lodash.js");\n/* harmony import */ var lodash__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(lodash__WEBPACK_IMPORTED_MODULE_1__);\n/* harmony import */ var _components_blockchain__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./components/blockchain */ "./src/components/blockchain.ts");\n/* harmony import */ var _ico_service__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! ./ico.service */ "./src/ico.service.ts");\n/* harmony import */ var _stripe_service__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! ./stripe.service */ "./src/stripe.service.ts");\n\n\n\nvar Clipboard = __webpack_require__(/*! clipboard */ "./node_modules/clipboard/dist/clipboard.js");\n\n\nvar blockchain = new _components_blockchain__WEBPACK_IMPORTED_MODULE_2__["default"](false);\nvar v = new vue__WEBPACK_IMPORTED_MODULE_0__["default"]({\n    el: "#app",\n    template: "#app-template",\n    delimiters: [\'${\', \'}\'],\n    data: {\n        domLoaded: false,\n        availableTokens: 0,\n        address: "",\n        actualPrice: 0,\n        plnInput: "",\n        tx: "",\n        ethInput: "0.000142857143",\n        phase: 0,\n        leftInActualPhase: 0,\n        inputType: 0,\n        payType: 0,\n        customerAddress: "",\n        customerTokens: 0,\n        customerValid: true,\n        statusTokens: 0,\n        statusAmount: NaN,\n        statusEth: 0,\n        ethPrice: 0,\n        stop: false,\n        prices: [],\n        tokens: [],\n        enabledMetamask: false,\n        debouncedGetPrice: function () { },\n        debouncedGetAmount: function () { },\n        debouncedCheckAddress: function () { }\n    },\n    created: function () {\n        this.debouncedGetAmount = lodash__WEBPACK_IMPORTED_MODULE_1__["debounce"](this.getAmount, 700);\n        this.debouncedCheckAddress = lodash__WEBPACK_IMPORTED_MODULE_1__["debounce"](this.checkAddress, 700);\n        this.getInfo();\n    },\n    updated: function () {\n        var _this = this;\n        this.$nextTick(function () {\n            if (!_this.domLoaded && $(\'#card-element\').length > 0) {\n                _this.domLoaded = true;\n                _stripe_service__WEBPACK_IMPORTED_MODULE_4__["stripeService"].init(_this);\n                var copyAddress = new Clipboard("#copyAddress, #copyAmount");\n                copyAddress.on(\'success\', function (e) {\n                    $(e.trigger).siblings(".tooltip").find(".tooltip-inner").text("Copied");\n                });\n            }\n            $(\'[data-toggle="tooltip"]\').tooltip({ trigger: "hover" });\n        });\n    },\n    watch: {\n        ethInput: function (newQuestion, oldQuestion) {\n            this.statusAmount = NaN;\n            this.debouncedGetAmount();\n        },\n        plnInput: function (newPln, oldPln) {\n            if (newPln.indexOf(",") != -1)\n                newPln = newPln.replace(\',\', \'.\');\n            var val = parseFloat(newPln);\n            if (isNaN(val) || val <= 0)\n                return;\n            this.statusAmount = NaN;\n            val *= 1 / this.ethPrice;\n            this.ethInput = val.toString();\n        },\n        payType: function (newType, oldType) {\n            this.statusAmount = NaN;\n            this.debouncedGetAmount();\n        },\n        customerAddress: function (newAddress, oldAddress) {\n            this.customerValid = true;\n            this.customerTokens = NaN;\n            this.debouncedCheckAddress();\n        }\n    },\n    methods: {\n        ethDecimals: function (value) {\n            return value.toFixed(12);\n        },\n        getInfo: function () {\n            var _this = this;\n            _ico_service__WEBPACK_IMPORTED_MODULE_3__["icoService"].getInfo()\n                .then(function (response) {\n                _this.availableTokens = response.data.tokensAvailable;\n                _this.phase = response.data.phase;\n                _this.leftInActualPhase = response.data.leftInActualPhase;\n                _this.prices = response.data.prices;\n                _this.tokens = response.data.tokens;\n            });\n            _ico_service__WEBPACK_IMPORTED_MODULE_3__["icoService"].getTotalPrice(100)\n                .then(function (response) {\n                if (response.data.error)\n                    $.jGrowl(response.data.data, { header: \'Error\', life: 20000 });\n                else\n                    _this.actualPrice = response.data.data / Math.pow(10, 18);\n            });\n            _ico_service__WEBPACK_IMPORTED_MODULE_3__["icoService"].getAddress()\n                .then(function (response) {\n                if (response.data.error)\n                    $.jGrowl(response.data.data, { header: \'Error\', life: 20000 });\n                else\n                    _this.address = response.data.data;\n            });\n            _ico_service__WEBPACK_IMPORTED_MODULE_3__["icoService"].getEthPrice()\n                .then(function (response) {\n                if (response.data.error)\n                    $.jGrowl(response.data.data, { header: \'Error\', life: 20000 });\n                else {\n                    _this.ethPrice = response.data.data;\n                    _this.plnInput = "1";\n                }\n            });\n        },\n        getAmount: function () {\n            var _this = this;\n            var val = parseFloat(this.ethInput);\n            if (isNaN(val) || val <= 0)\n                return;\n            _ico_service__WEBPACK_IMPORTED_MODULE_3__["icoService"].getTokensForAmount(Math.round(val * 1000000000000000000))\n                .then(function (response) {\n                if (response.data.error)\n                    $.jGrowl(response.data.data, { header: \'Error\', life: 20000 });\n                else {\n                    _this.statusAmount = response.data.data / 1000;\n                }\n            });\n        },\n        checkAddress: function () {\n            var _this = this;\n            _ico_service__WEBPACK_IMPORTED_MODULE_3__["icoService"].checkAddress(this.customerAddress)\n                .then(function (response) {\n                if (response.data.error)\n                    $.jGrowl(response.data.data, { header: \'Error\', life: 20000 });\n                else {\n                    _this.customerValid = response.data.valid;\n                    if (_this.customerValid) {\n                        _this.customerTokens = response.data.amount;\n                    }\n                }\n            });\n        },\n        ethToPLN: function (eth) {\n            if (isNaN(eth))\n                return "~";\n            else\n                return Math.ceil(this.ethPrice * eth * 100) / 100;\n        },\n        fillPrice: function () {\n            this.payType = this.inputType == 2 ? 0 : this.inputType;\n            if (this.inputType == 0) {\n                this.statusTokens = this.statusTokens;\n            }\n            else {\n                this.statusTokens = parseFloat(this.ethInput);\n            }\n        },\n        payWithWallet: function () {\n            var _this = this;\n            this.tx = "";\n            blockchain.sendTransaction(this.address, parseFloat(this.ethInput) * Math.pow(10, 18), 600000, function (tx) {\n                _this.tx = tx;\n            });\n        }\n    }\n});\n\n\n//# sourceURL=webpack:///./src/icoBuy.ts?')},"./src/stripe.service.ts":
/*!*******************************!*\
  !*** ./src/stripe.service.ts ***!
  \*******************************/
/*! exports provided: default, stripeService */function(module,__webpack_exports__,__webpack_require__){"use strict";eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, \"stripeService\", function() { return stripeService; });\n/* harmony import */ var axios__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! axios */ \"./node_modules/axios/index.js\");\n/* harmony import */ var axios__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(axios__WEBPACK_IMPORTED_MODULE_0__);\n\nvar api = '/stripe';\nvar StripeService = /** @class */ (function () {\n    function StripeService() {\n    }\n    StripeService.prototype.init = function (main) {\n        var _this = this;\n        this.main = main;\n        this.stripe = Stripe('pk_test_FG1xyNruCNcwSqyk8cFx2FzW005M3TV6uR');\n        var elements = this.stripe.elements();\n        var style = {\n            invalid: {\n                color: '#fa755a',\n                iconColor: '#fa755a'\n            }\n        };\n        this.card = elements.create('card', { style: style });\n        this.card.mount($('#main #card-element')[0]);\n        this.card.addEventListener('change', function (event) { return _this.change(event); });\n        $('#main #payment-form')[0].addEventListener('submit', function (event) { return _this.submit(event); });\n    };\n    StripeService.prototype.change = function (event) {\n        $('#main #card-errors').text(event.error ? event.error.message : '');\n    };\n    StripeService.prototype.submit = function (event) {\n        var _this = this;\n        event.preventDefault();\n        if (this.main.customerAddress == \"\" || !this.main.customerValid) {\n            if (this.main.customerAddress == \"\")\n                this.main.customerAddress = \"0x\";\n            $.jGrowl(\"Please fill your ETH address\", { header: 'Error', life: 20000 });\n        }\n        else {\n            this.stripe.createToken(this.card).then(function (result) {\n                if (result.error) {\n                    $('#main #card-errors').text(result.error.message);\n                }\n                else {\n                    _this.checkout(result.token.id);\n                    console.log(result);\n                }\n            });\n        }\n    };\n    StripeService.prototype.checkout = function (token) {\n        var qs = __webpack_require__(/*! qs */ \"./node_modules/qs/lib/index.js\");\n        axios__WEBPACK_IMPORTED_MODULE_0___default.a.post(api + \"/buy\", qs.stringify({\n            tokens: (parseFloat(this.main.statusAmount) * 100).toString(),\n            source: token,\n            address: this.main.customerAddress\n        }), {\n            headers: {\n                'Content-Type': 'application/x-www-form-urlencoded'\n            }\n        })\n            .then(function (response) {\n            console.log(response);\n            if (response.data.error) {\n                if (response.data.data.indexOf(\"{\") == 0) {\n                    $.jGrowl(JSON.parse(response.data.data).message, { header: 'Error', life: 20000 });\n                    console.log(JSON.parse(response.data.data), response.data.data);\n                }\n                else {\n                    $.jGrowl(response.data.data, { header: 'Error', life: 20000 });\n                }\n            }\n            else {\n                $.jGrowl(\"You bought tokens, You can refresh token status on the right of address input.\", { header: 'Congratulation', life: 20000 });\n            }\n        })\n            .catch(function (error) {\n            console.log(error);\n        });\n    };\n    return StripeService;\n}());\n/* harmony default export */ __webpack_exports__[\"default\"] = (StripeService);\nvar stripeService = new StripeService();\n\n\n//# sourceURL=webpack:///./src/stripe.service.ts?")}});