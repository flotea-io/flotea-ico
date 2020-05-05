import Vue from "vue";
declare var jGrowl: any;
declare var $: any;
declare var window: any;
import moment from 'moment';

import Blockchain from './components/blockchain';
const VTooltip = require('v-tooltip');

const firstLoadCount = 3;

const sharedData = {
	applicants: <any[]>[],
	enabledMetamask: <boolean>false,
	loading: <boolean>false,
	votersCount: <number>0,
	showModal: <boolean>false,
	selectedApplicant: <number>-1,
	applicantsCount: <number>0,
	metamaskState: <any>{ enabledMetamask: false, userAddress: "" },
};

class Applicants {
	blockchain: Blockchain;
	vue: Vue;

	constructor(){
		this.blockchain = new Blockchain(false);
		this.blockchain.loadContract("VotingCarrier", (contract: any) => {
			contract.events.allEvents({},(error: boolean, event: any)=>{ 
				if(typeof event.returnValues != "undefined")
					this.resolveEvent(event);
			});
			this.loadApplicantsCount();
		});
		moment.locale('pl')
		this.initVue();
	}

	initVue(){
		Vue.use(VTooltip);
		Vue.component('modal', {
			template: '#modal-template'
		});

		this.vue = new Vue({
			el: "#applicants",
			template: "#applicants-template",
			delimiters: ['${', '}'],
			data: sharedData,
			created: function() {
				setInterval(()=>{
					(<any>this).$forceUpdate();
				}, 30 * 1000);
			},
			watch: {
			},
			methods: {
				isVoted: function(applicant: any){
					return applicant.voters.findIndex((el: any)=> el.toUpperCase() == this.metamaskState.userAddress.toUpperCase() );
				},
				toEndText: function(time: number){
					return moment(time*1000).fromNow();
				},
				ended: function(time: number){
					return moment().isBefore( moment(time*1000) )
				},
				avatarChars: function(name: string){
					return name[0]+ name[name.length-1];
				},
				finish: (index: number) => { this.finish(index)},
				vote: (index: number, vote: number) => { this.vote(index, vote)},
				detail(index: number){
					this.selectedApplicant = index;
					this.showModal = true;
				},
				loadNext: () => {
					this.loadAplicants(sharedData.applicantsCount - sharedData.applicants.length -1, firstLoadCount, false);
				}
			}
		});
	}

	getIndexById(id: number){
		return sharedData.applicants.findIndex((el)=> el.index === id);
	}

	resolveEvent(e: any){
		console.log(e);
		let index;
		switch(e.event){
			case "VoteCarrier":
			index = this.getIndexById(e.returnValues.proposalIndex);
			if(index != -1){
				sharedData.applicants[index].voted +=1;
				sharedData.applicants[index].voters.push(e.returnValues.addr);
				sharedData.applicants[index].votersVotes.push(e.returnValues.vote);
				if(e.returnValues.vote < 2){
					sharedData.applicants[index][e.returnValues.vote == 1? "yes" : "no"] += 1;
				}
			}
			break;
			case "FinishVotingCarrier":
			index = this.getIndexById(e.returnValues.proposalIndex);
			if(index != -1){
				sharedData.applicants[index].result = e.returnValues.result? 1 : 0;
			}
			break;
			case "CarrierProposalCreated":
			this.loadAplicants(e.returnValues.proposalIndex, 1, true);
			break;
		}
	}

	finish(index: number){
		let id = this.getIndexById(index);
		sharedData.applicants[id].tx = "";
		this.blockchain.callContract("VotingCarrier", "beforeFinishCarrierProposal", (error: any) => {
			if(error[0]){
				$.jGrowl(error[1], { header: "error", life: 10000 });
			} else {
				this.blockchain.sendContract("VotingCarrier", "finishCarrierProposal",
					{ "from": sharedData.metamaskState.userAddress }, (success: any) => {
						if(id != -1){
							sharedData.applicants[id].tx = success;
						}
					}, null,
					this.blockchain.numberToHex(index),
					);
			}
		}, null,
		this.blockchain.numberToHex(index),
		);
	}

	vote(vote: number, index: number){
		let id = this.getIndexById(index);
		sharedData.applicants[id].tx = "";
		this.blockchain.callContract("VotingCarrier", "beforeVoteInCarrierProposal", (error: any) => {
			if(error[0]){
				$.jGrowl(error[1], { header: "error", life: 10000 });
			} else {
				this.blockchain.sendContract("VotingCarrier", "voteInCarrierProposal",
					{ "from": sharedData.metamaskState.userAddress }, (success: any) => {
						if(id != -1){
							sharedData.applicants[id].tx = success;
						}
					}, null,
					this.blockchain.numberToHex(index),
					this.blockchain.numberToHex(vote),
					);
			}
		}, null,
		index,
		sharedData.metamaskState.userAddress
		);
	}

	loadApplicantsCount(){
		this.blockchain.callContract("VotingCarrier", "proposalsLength", (response: any) => {
			sharedData.applicantsCount = response;
			if( response == 0 ) sharedData.loading = false;
			this.loadAplicants(response-1, firstLoadCount, false);
		});
		this.blockchain.callContract("VotingCarrier", "VotersLength", (response: any) => {
			sharedData.votersCount = response;
		});
	}

	loadAplicants(index: number, max: number, toUp: boolean){
		if(index < 0 || max <= 0) {
			sharedData.loading = false;
			return;
		}
		sharedData.loading = true;
		this.blockchain.callContract("VotingCarrier", "carrierProposals", (response: any) => {
			this.blockchain.callContract("VotingCarrier", "statusOfCarrierProposal", (responseStatus: any) => {
				let yes = 0, no = 0, voted = 0;
				for(let i = 0; i < responseStatus[0].length; i++){
					if(responseStatus[0][i] != "0x0000000000000000000000000000000000000000"){
						voted++;
						if(responseStatus[1][i] == 0)
							no++;
						else if(responseStatus[1][i] == 1)
							yes++;
					}
				}

				let data = {
					actionAddress: response.actionAddress,
					actionType: response.actionType,
					endTime: response.endTime,
					name: this.blockchain.fromHex(response.name),
					web: this.blockchain.fromHex(response.web),
					index: index,
					voters: responseStatus[0],
					votersVotes: responseStatus[1],
					voted: voted,
					yes: yes,
					no: no,
					result: response.result,
					tx: "",
					showVoteButtons: false,
				};	

				if(toUp)
					sharedData.applicants.unshift(data);
				else
					sharedData.applicants.push(data);
				
			}, null, index);
			this.loadAplicants(index-1, max-1, toUp);
		}, null, index);
	}
}

$(document).ready(()=>{
	new Applicants();
});