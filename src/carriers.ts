/*
* Project: FLOTEA - Decentralized passenger transport system
* Copyright (c) 2020 Flotea, All Rights Reserved
* For conditions of distribution and use, see copyright notice in LICENSE
*/

import Web3 from 'web3';
import Vue from "vue";
declare var web3: any;
declare var $: any;
declare var window: any;
import Blockchain from './components/blockchain';

const firstLoadCount = 3;

const sharedData = {
	carriers: <any[]>[],
	loading: <boolean>false,
	carriersCount: <number>0,
};

class Carriers {
	blockchain: Blockchain;
	vue: Vue;

	constructor(){
		this.blockchain = new Blockchain(true);
		this.blockchain.loadContract("Carriers", (contract: any) => {
			contract.events.allEvents({},(error: boolean, event: any)=>{
				if(typeof event.returnValues != "undefined")
					this.resolveEvent(event);
			});
			this.loadCarriersCount();
		})
		this.initVue();
		$('[data-toggle="tooltip"]').tooltip();
	}

	initVue(){
		this.vue = new Vue({
			el: "#carriers",
			template: "#carriers-template",
			delimiters: ['${', '}'],
			data: sharedData,
			methods: {
				avatarChars: function(name: string){
					return name[0]+ name[name.length-1];
				},
				loadNext: () => {
					this.loadCarriers(sharedData.carriersCount - sharedData.carriers.length -1, firstLoadCount, false);
				}
			}
		});
	}

	resolveEvent(e: any){
		console.log(e);
		let index;
		switch(e.event){
			case "FinishVotingCarrier":
			index = this.getIndexById(e.returnValues.proposalIndex);
			if(index != -1){
			//	sharedData.carriers[index].result = e.returnValues.result? 1 : 0;
			}
			break;
		}
	}

	getIndexById(id: number){
		return sharedData.carriers.findIndex((el)=> el.index === id);
	}

	loadCarriersCount(){
		this.blockchain.callContract("Carriers", "carrierLength", (response: any) => {
			sharedData.carriersCount = response;
			if( response == 0 ) sharedData.loading = false;
			this.loadCarriers(response-1, firstLoadCount, false);
		}, 0);
	}

	loadCarriers(index: number, max: number, toUp: boolean){
		if(index < 0 || max <= 0) {
			sharedData.loading = false;
			return;
		}
		sharedData.loading = true;
		this.blockchain.callContract("Carriers", "carriers", (response: any) => {
				let data = {
					companyWallet: response.companyWallet,
					name: this.blockchain.fromHex(response.company),
					web: this.blockchain.fromHex(response.web),
					index: index,
					enabled: response.enabled
				};

				if(toUp)
					sharedData.carriers.unshift(data);
				else
					sharedData.carriers.push(data);

			this.loadCarriers(index-1, max-1, false);
		}, null, index);
	}

}

$(document).ready(()=>{
	new Carriers();
});
