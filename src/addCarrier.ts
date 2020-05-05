import Vue from "vue";
declare var jGrowl: any;
declare var $: any;
import Blockchain from './components/blockchain';

const sharedData = {
	firstTime: true,
	company: "Acme co.",
	web: "https://acme.co",
	errorText: "",
	tx: "",
	metamaskState: { enabledMetamask: false, userAddress: "" },
}
class AddCarrier {

	blockchain: Blockchain;
	vue: Vue;

	constructor() {
		$("#form input").on("keyup", this.inputChanged);
		this.blockchain = new Blockchain(false);
		this.initVue();
		this.blockchain.loadContract("VotingCarrier");
	}

	initVue(){
		this.vue = new Vue({
			el: "#add-carrier",
			template: "#add-carrier-template",
			delimiters: ['${', '}'],
			data: sharedData,
			watch:{
				company: (after: string, before: string) => this.limit32(after, before, "company"),
				web: (after: string, before: string) => this.limit32(after, before, "web"),
				metamaskState: function(state: any){
					if(state.userType == 0){
						this.company = "";
						this.web = "";
					} else{
						this.company = state.carrierName;
						this.web = state.carrierWeb;
					}
				},
			},
			methods: {
				validate: function(){
					return (!this.metamaskState.enabledMetamask || this.company.trim().length == 0 || this.web.trim().length == 0);
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

	inputChanged(event: any){
		let el = $(event.currentTarget);
		if(el.val().trim().length > 0)
			el.parent().removeClass("has-error");
		else
			el.parent().addClass("has-error");
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
						}, null ,
						this.blockchain.toHex(sharedData.company),
						this.blockchain.toHex(sharedData.web),
						0,
						sharedData.metamaskState.userAddress
						);
				}
			}, null,
			this.blockchain.numberToHex(0),
			sharedData.metamaskState.userAddress,
			);
	}

	
}

$(document).ready(()=>{
	if($("#add-carrier").length > 0)
		new AddCarrier();
});