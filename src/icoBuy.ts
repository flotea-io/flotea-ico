import Vue from "vue";
import Web3 from 'web3';
import * as _ from "lodash";
import Blockchain from './components/blockchain';

const Clipboard = require('clipboard');
import { icoService } from "./ico.service";
import { stripeService } from "./stripe.service";
declare var web3: any;
declare var jGrowl: any;
declare var $: any;
declare var ethereum: any;
let blockchain = new Blockchain(false);
let v = new Vue({
    el: "#app",
    template: "#app-template",
    delimiters: ['${', '}'],
    data: {
        domLoaded: false,
        availableTokens: 0,
        address: "",
        actualPrice: 0,
        plnInput: "",
        tx: "",
        ethInput: "0.000142857143",
        phase: 0,
        leftInActualPhase: 0,
        inputType: 0,
        payType: 0,
        customerAddress: "",
        customerTokens: 0,
        customerValid: true,
        statusTokens: 0,
        statusAmount: NaN,
        statusEth: 0,
        ethPrice: 0,
        stop: false,
        prices: [],
        tokens: [],
        enabledMetamask: false,
        debouncedGetPrice: ()=>{},
        debouncedGetAmount: ()=>{},
        debouncedCheckAddress: ()=>{}
    },
    created: function () {
    	this.debouncedGetAmount = _.debounce(this.getAmount, 700);
        this.debouncedCheckAddress = _.debounce(this.checkAddress, 700);
        this.getInfo();
    },
    updated: function(){
        this.$nextTick(() => {
            if(!this.domLoaded && $('#card-element').length > 0){
                this.domLoaded = true;
                stripeService.init(this);
                let copyAddress = new Clipboard("#copyAddress, #copyAmount");
                copyAddress.on('success', function(e:any) {
                    $(e.trigger).siblings(".tooltip").find(".tooltip-inner").text("Copied");
                });
            }
            (<any>$('[data-toggle="tooltip"]')).tooltip({ trigger: "hover" });

        });
    },
    watch: {
    	ethInput: function (newQuestion, oldQuestion) {
            this.statusAmount = NaN;
            this.debouncedGetAmount();
        },
        plnInput: function(newPln, oldPln) {
            if(newPln.indexOf(",") != -1)
                newPln = newPln.replace(',','.');
            let val = parseFloat(newPln);
            if(isNaN(val) || val <= 0) return;
            
            this.statusAmount = NaN;
            
            val *= 1 / this.ethPrice;
            this.ethInput = val.toString();
        },
        payType: function(newType, oldType){
            this.statusAmount = NaN;
            this.debouncedGetAmount();
        },
        customerAddress: function (newAddress, oldAddress) {
            this.customerValid = true;
            this.customerTokens = NaN
            this.debouncedCheckAddress();
        }
    },
    methods: {
        ethDecimals: function(value : number){
            return value.toFixed(12)
        },
        getInfo : function () {
          icoService.getInfo()
          .then((response: any) => {
           this.availableTokens = response.data.tokensAvailable;
           this.phase = response.data.phase;
           this.leftInActualPhase = response.data.leftInActualPhase;
           this.prices = response.data.prices;
           this.tokens = response.data.tokens;
       });
          icoService.getTotalPrice(100)
          .then((response: any) => {
            if(response.data.error)
                $.jGrowl(response.data.data, { header: 'Error', life: 20000 });
            else 
                this.actualPrice = response.data.data/ 10**18;
        });
          icoService.getAddress()
          .then((response: any) => {
            if(response.data.error)
                $.jGrowl(response.data.data, { header: 'Error', life: 20000 });
            else 
                this.address = response.data.data;
        });
          icoService.getEthPrice()
          .then((response: any) => {
            if(response.data.error)
                $.jGrowl(response.data.data, { header: 'Error', life: 20000 });
            else {
                this.ethPrice = response.data.data;
                this.plnInput = "1";
            }
        });
          
      },
      getAmount: function () {
        let val = parseFloat(this.ethInput);
        if(isNaN(val) || val <= 0) return;
        icoService.getTokensForAmount(Math.round(val * 1000000000000000000))
        .then((response: any) => {
           if(response.data.error)
            $.jGrowl(response.data.data, { header: 'Error', life: 20000 });
        else {
            this.statusAmount = response.data.data/1000;
        }
        });
    },
    checkAddress: function() {
        icoService.checkAddress(this.customerAddress)
        .then((response: any) => {
            if(response.data.error)
                $.jGrowl(response.data.data, { header: 'Error', life: 20000 });
            else{
                this.customerValid = response.data.valid;
                if(this.customerValid){
                    this.customerTokens = response.data.amount; 
                }
            }
        });
    },
    ethToPLN: function(eth: number){
        if(isNaN(eth))
            return "~";
        else
            return Math.ceil(this.ethPrice * eth * 100) / 100;
    },
    fillPrice: function(){
        this.payType = this.inputType == 2 ? 0 : this.inputType;
        if(this.inputType == 0){
            this.statusTokens = this.statusTokens;
        }
        else{
            this.statusTokens = parseFloat(this.ethInput);
        }
    },
    payWithWallet: function(){
        this.tx = "";
        blockchain.sendTransaction(this.address, parseFloat(this.ethInput) * 10**18, 600000, (tx: string)=>{
            this.tx = tx
        });
    }
}
});
