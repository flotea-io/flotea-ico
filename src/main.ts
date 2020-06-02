/*
* Project: FLOTEA - Decentralized passenger transport system
* Copyright (c) 2020 Flotea, All Rights Reserved
* For conditions of distribution and use, see copyright notice in LICENSE
*/

import Vue from "vue";
import { icoService } from "./ico.service";

let v = new Vue({
    el: "#token-info",
    template: "#app-template",
    delimiters: ['${', '}'],
    data: {
        tokensMax: 0,
        availableTokens: 0,
        phase: 0,
        icoBalanceWei: 0,
        icoBalancePln: 0,
        leftInActualPhase: 0,
        prices: [],
        phases: [],
        total: 0,
        showPhases: false,
        ethPrice: 1
    },
    mounted: function () {
        (this as any).getInfo();
        (this as any).loadPlnBalance();
    },
    methods: {
        loadPlnBalance: function () {
            this.icoBalancePln = 0;
            return;
            icoService.getBalancePln()
            .then((response: any) => {
                this.icoBalancePln = response.data.data/100;
            });
        },
        getInfo: function () {
            icoService.getInfo()
            .then((response: any) => {
                console.log(response);
                this.availableTokens = response.data.tokensAvailable;
                this.icoBalanceWei = response.data.balanceWei;
                this.phase = response.data.phase;
                this.leftInActualPhase = response.data.leftInActualPhase;
                this.prices = response.data.prices;
                this.tokensMax = response.data.initialTokens;
                this.phases = response.data.tokens;
                let buyedTokens = this.tokensMax - this.availableTokens;
                let p = 0;
                let find = false;
                while (this.phases.length -1 > p && !find) {
                    if(this.phases[p] >= buyedTokens){
                        find = true;
                    }
                    else {
                        buyedTokens -= this.phases[p];
                        p++;
                    }
                }
                this.phase = p;
                this.leftInActualPhase = this.phases[p] - buyedTokens;

                for(let i = 0; i < this.phases.length; i++)
                    this.total += this.phases[i];
            });
            icoService.getEthPrice()
            .then((response: any) => {
                this.ethPrice = response.data.data;
            });
        },
    },
    computed: {
        percentTokensInPhase: function() {
            return 100 - this.leftInActualPhase / this.phases[this.phase] * 100;
        }
    }
});
