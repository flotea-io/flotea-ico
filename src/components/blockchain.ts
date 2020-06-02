/*
* Project: FLOTEA - Decentralized passenger transport system
* Copyright (c) 2020 Flotea, All Rights Reserved
* For conditions of distribution and use, see copyright notice in LICENSE
*/

import Vue from "vue";
import Web3 from 'web3';
declare var web3: any;
declare var web3wss: any;
declare var jGrowl: any;
declare var ethereum: any;
declare var window: any;
declare var document: any;
declare var $: any;

// @ts-ignore
import Popper from 'vue-popperjs';


const addresses: any = {
	wssUrl: "wss://kovan.infura.io/ws/v3/eabd1b929937436ea9a0529c3eae21c8",
	FloteaToken: "0x1C94D0D2F4ED3E8Cd38B0cBA9C8d05e6648fdaa2",
	VotingCarrier: "0xa71eE8c6fb6079FCB8b13Fa81Bb554950809A74C",
	Carriers: "0xC9C5c6Ac5836AaCEA50Fe7fcf286071032cE5451",
	FloteaICO: "0xDa8370898A95446875DaC7AeAd92E5fd5851174D",
	Transport: "0x0D2C2582F68f792E4f7c6b8045D5a9Ae70806526",
};

const blockchainSharedData = {
	enabledMetamask: <boolean> false,
	userAddress: <string> null,
	carrierName: <string> "",
	carrierWeb: <string> "",
	userType: <number> 0,
	flt: <number> 0,
	eth: <number> 0,
	isMetamask: <boolean> true,
	contractsLoaded: <boolean> false,
	waitForResponse: <boolean> false,
};

export default class Blockchain {
	contracts: any = [];
	contractsCall: any = [];
	initializedEvents: boolean = false;
	onlyRead: boolean = false;
	constructor(onlyRead: boolean) {
		this.onlyRead = onlyRead;
		if(this.onlyRead){
			this.initWebsocketProvider();
		}
		this.register();
	}

	register(){
		let bl = this;
		Vue.component('MetamaskButton', {
			components: {
				'popper': Popper
			},
			mounted() {
				bl.initWebsocketProvider();
				bl.mainInit();
				this.$emit('getObject', bl);
			},
			data: function(){
				return blockchainSharedData;
			},
			computed: {
				metamaskState(){
					return {
						enabledMetamask: blockchainSharedData.enabledMetamask,
						userAddress: blockchainSharedData.userAddress,
						carrierName: blockchainSharedData.carrierName,
						carrierWeb: blockchainSharedData.carrierWeb,
						userType: blockchainSharedData.userType,
						flt: blockchainSharedData.flt,
						eth: blockchainSharedData.eth,
						isMetamask: blockchainSharedData.isMetamask,
					};
				}
			},
			watch:{
				metamaskState: function(newState){
					this.$emit('metamaskUpdated', newState);
				}
			},

			template: '<popper> <div class="popper"> <h3 v-if="carrierName!=\'\'">{{carrierName}}</h3> <div>{{eth.toFixed(4)}} ETH </div> <div> {{flt.toFixed(2)}} FLT </div> {{userAddress}} </div> <div slot="reference" class="metamask" @click="enableMetamask"> <img v-if="!enabledMetamask" src="/static/img/metamask-logo-92-a-73-d-44-kopia@3x.png" /> <img src="/static/img/user-solid.svg" class="user-icon" v-if="enabledMetamask" /> <span>{{enabledMetamask? "Połączony":"Odblokuj Metamask"}}</span> <span v-if="enabledMetamask">{{userType==0? "pasażer": "przewoźnik"}}</span> </div> </popper>',
			methods:{
				enableMetamask: function(){
					if(!this.isMetamask){
						let link = "https://metamask.io/";
						if(navigator.userAgent.indexOf("Opera") != -1 || navigator.userAgent.indexOf('OPR') != -1)
							link = "https://addons.opera.com/en/extensions/details/metamask/";
						else if(navigator.userAgent.indexOf("Chrome") != -1 ) // Chrome and Brave have same addon
							link = "https://chrome.google.com/webstore/detail/nkbihfbeogaeaoehlefnkodbefgpgknn";
						else if(navigator.userAgent.indexOf("Firefox") != -1 )
							link = "https://addons.mozilla.org/en-US/firefox/addon/ether-metamask/";
						window.open(link, '_blank');
					}
					else if(!this.enabledMetamask){
						bl.init();
					}
				}
			},
		});
	}

	mainInit() {
		this.getAccounts();
		this.init();
		this.loadContract("FloteaToken", ()=> this.loadContract("Carriers", ()=>{
			blockchainSharedData.contractsLoaded = true;
			if(blockchainSharedData.enabledMetamask){
				this.getAccountInfo();
			}
		}));
	}

	init() {
		if(blockchainSharedData.waitForResponse) return;
		if (window.ethereum) {
			blockchainSharedData.waitForResponse = true;
			window.ethereum.autoRefreshOnNetworkChange = false;
			window.web3 = new Web3(window.ethereum);
			const main = async () => {
				await window.ethereum.enable();
			};
			main().catch(err => {
				blockchainSharedData.waitForResponse = false;
				$.jGrowl("Please enable MetaMask, or Opera Wallet.", { position:"center", header: 'Error', life: 10000 });
			});
		}
		else {
			blockchainSharedData.isMetamask = false;
			$.jGrowl("Non-Ethereum browser detected. You should consider trying MetaMask or Opera browser!", { position:"center", header: 'Error', life: 10000 });
		}
	}

	getAccountInfo() {
		this.callContract("Carriers", "getCarrierData", (response: any)=>{
			if(response.exist){
				blockchainSharedData.carrierName = this.fromHex(response.company);
			 	blockchainSharedData.carrierWeb = this.fromHex(response.web);
				blockchainSharedData.userType = 1;
			} else {
				blockchainSharedData.carrierName = "";
				blockchainSharedData.carrierWeb = "";
				blockchainSharedData.userType = 0;
			}
		}, null, blockchainSharedData.userAddress);
		this.callContract("FloteaToken", "balanceOf", (balance: number)=>{
			blockchainSharedData.flt = balance/100;
		}, null, blockchainSharedData.userAddress);
		web3.eth.getBalance(blockchainSharedData.userAddress).then((amount: number)=>{
			blockchainSharedData.eth = amount/10**18;
		});
	}

	getAccounts() {
		var accountInterval = setInterval(()=>{
			if (web3.currentProvider.selectedAddress !== blockchainSharedData.userAddress) {
				blockchainSharedData.userAddress = web3.currentProvider.selectedAddress;
				if(blockchainSharedData.userAddress != null){
					blockchainSharedData.enabledMetamask = true;
					if(blockchainSharedData.contractsLoaded){
						this.getAccountInfo();
					}
				} else{
					blockchainSharedData.enabledMetamask = false;
				}
			}
		}, 500);
	}

	initWebsocketProvider() {
		window.web3wss = new Web3(new Web3.providers.WebsocketProvider(addresses.wssUrl));
	}

	callContract(contract: string, method: string, success: any, error: any = null, ...params: any[]) {
		this.contractsCall[contract].methods[method](...params).call().then((result: any) => {
			if(typeof success == "function")
				success(result);
		}).catch((err: any) => {
			console.log(err);
			if(typeof error == "function")
				error(err);
			this.handleErrorMessages(err);
		});
	}

	sendContract(contract: string, method: string, sendParams: any, success: any, error: any = null, ...params: any[]) {
		this.contracts[contract].methods[method](...params).send(sendParams)
		.on('transactionHash', function(hash: string){
			if(typeof success == "function")
				success(hash);
			console.log(hash);
		})
		.on('receipt', function(receipt: string){
			console.log(receipt);
		})
		.on('error', (err: any) => {
			if(typeof error == "function")
				error(err);
			this.handleErrorMessages(err);
		});
	}

	sendTransaction(to: string, value: number, gas: number, success: any, error: any = null) {
		web3.eth.sendTransaction({
			from: blockchainSharedData.userAddress,
			to: to,
			gas: gas,
			value: value
		}).on('transactionHash', function(hash: string){
			if(typeof success == "function")
				success(hash);
			console.log(hash);
		})
		.on('receipt', function(receipt: string){
			console.log(receipt);
		})
		.on('error', (err: any) => {
			if(typeof error == "function")
				error(err);
			this.handleErrorMessages(err);
		});
	}

	loadContract(name: string, fn: any = null) {
		if(typeof this.contracts[name] != "undefined"){
			if(typeof fn == "function")
				fn(this.contracts[name]);
			return;
		}
		$.getJSON("/static/"+ name +".json", (e: any) => {
			if (typeof window.web3 == "object" && !this.onlyRead)
				this.contracts[name] = new window.web3.eth.Contract(e.abi, addresses[name]);

			this.contractsCall[name] = new window.web3wss.eth.Contract(e.abi, addresses[name]);
			if(typeof fn == "function"){
				fn(this.contractsCall[name]);
			}
		});
	}

	loadContractAtAddress(name: string, address: string, fn: any) {
		if(typeof this.contracts[name+address] != "undefined"){
			if(typeof fn == "function")
				fn(this.contracts[name+address]);
			return;
		}
		$.getJSON("/static/"+ name +".json", (e: any) => {
			if (typeof window.web3 == "object" && !this.onlyRead )
				this.contracts[name+address] = new window.web3.eth.Contract(e.abi, address);

			this.contractsCall[name+address] = new window.web3wss.eth.Contract(e.abi, address);
			if(typeof fn == "function"){
				fn(this.contractsCall[name+address]);
			}
		});
	}

	fromHex(hex: string){
		return window.web3wss.utils.hexToUtf8(hex);
	}

	toHex(str: string){
		let hex = window.web3wss.utils.utf8ToHex(str.toString());
		return hex + '0'.repeat(66-hex.length);
	}

	numberToHex(number: number){
		return window.web3wss.utils.numberToHex(number);
	}

	hexToNumber(number: number){
		return window.web3wss.utils.hexToNumber(number);
	}

	handleErrorMessages(error: any){
		$.jGrowl(error.stack, { position:"center", header: error.message, life: 10000 });
	}
}
