/*
* Project: FLOTEA - Decentralized passenger transport system
* Copyright (c) 2020 Flotea, All Rights Reserved
* For conditions of distribution and use, see copyright notice in LICENSE
*/

import Vue from "vue";
declare var jGrowl: any;
declare var $: any;
declare var window: any;
import Blockchain from './components/blockchain';

const sharedData = {
	removeAddress: "",
	reason: "Przewoznik nie komunikuje",
	web: "",
	errorText: "",
	tx: "",
	metamaskState: { enabledMetamask: false, userAddress: "" },
}
class RemoveCarrier {

	blockchain: Blockchain;
	vue: Vue;

	constructor() {
		let p = location.pathname.lastIndexOf("/");
		sharedData.removeAddress = location.pathname.substr(p+1, location.pathname.length - p);
		this.blockchain = new Blockchain(false);
		this.initVue();
		this.blockchain.loadContract("VotingCarrier");
	}

	initVue(){
		this.vue = new Vue({
			el: "#remove-carrier",
			template: "#remove-carrier-template",
			delimiters: ['${', '}'],
			data: sharedData,
			watch:{
				reason: (after: string, before: string) => this.limit32(after, before, "reason"),
				web: (after: string, before: string) => this.limit32(after, before, "web"),
			},
			methods: {
				validate: function(){
					return (!sharedData.metamaskState.enabledMetamask || sharedData.reason.trim().length == 0 );
				},
				max32: function(str: string){
					return 32-str.length;
				},
				submit: this.submit.bind(this)
			}
		});
	}

	limit32(after: string, before: string, name: string){
		if(after.length > 32){
			if(name == "web"){
				$.jGrowl("Ten adres internetowy jest za długi, użyj narzędzia takiego jak <a href='https://bitly.com' target='_blank'>https://bitly.com</a>", { header: "Błąd", life: 10000 });
			}
			(<any>sharedData)[name] = before;
		}
	}

	submit(){
		this.blockchain.callContract("VotingCarrier", "beforeCreateCarrierProposal",
			(error: any) => {
				if(error[0]){
					sharedData.errorText = error[1];
				} else {
					sharedData.errorText = "";
					this.blockchain.sendContract("VotingCarrier", "createCarrierProposal",
						{ "from": sharedData.metamaskState.userAddress }, (success: any) => {
							sharedData.tx = success;
						}, null,
						this.blockchain.toHex(sharedData.reason),
						this.blockchain.toHex(sharedData.web),
						1,
						sharedData.removeAddress
						);
				}
			}, null,
			1,
			sharedData.removeAddress
			);
	}
}

$(document).ready(()=>{
	if($("#remove-carrier").length > 0)
		new RemoveCarrier();
});
