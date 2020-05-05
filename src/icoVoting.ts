import Vue from "vue";
import moment from 'moment';
import Blockchain from './components/blockchain';
declare var jGrowl: any;
declare var $: any;
declare var window: any;

const sharedData = {
    proposalsLength: 0,
    proposalShowing: false,
    selectedProposal: {
        description: "",
        endTime: 0,
        amount: 0,
        actionAddress: "",
        actionType: 0,
        result: 2,
        name: "",
        index: 0
    },
    selectedStatus: <any[]>[],
    userAddress: "",
    actionType: 0,
    description: "",
    durationHours: 2,
    address: "",
    name: "",
    amount: 0,
    showModal: false,
    proposals: <any[]>[],
    metamaskState: { enabledMetamask: false, userAddress: "" },
}

class Voting {
    blockchain: Blockchain;
    vue: Vue;

    constructor() {
        this.blockchain = new Blockchain(false);
        this.initVue();
        this.blockchain.loadContract("FloteaICO", (c:any)=> this.registerEvents(c));
    }

    initVue(){
        let voting = this;
        this.vue = new Vue({
            el: "#voting",
            template: "#app-template",
            delimiters: ['${', '}'],
            data: sharedData,
            watch: {
                actionType: function () {
                    switch(this.actionType){
                        case 0: 
                        this.amount = 0;
                        break;
                        case 1:
                        this.name = "";
                        this.amount = 0;
                        break;
                        case 2:
                        this.name = "";
                        break;
                    }
                }
            },
            methods: {
                loadProposals: this.loadProposals.bind(this),
                showProposal: function(proposal: any){
                    this.selectedProposal = proposal;
                    this.proposalShowing = true;
                    voting.loadProposalStatus(this.selectedProposal.index);
                },
                createProposal: function(){
                    voting.blockchain.sendContract("FloteaICO","createProposal", { "from": this.metamaskState.userAddress },
                        (err: any) => this.resolveError(err.message), 
                        voting.blockchain.toHex(this.description),
                        this.durationHours,
                        this.actionType,
                        this.address,
                        voting.blockchain.toHex(this.name),
                        this.amount
                        );
                },
                vote: function(type: number){
                    voting.blockchain.sendContract("FloteaICO", "voteInProposal", { "from": this.userAddress }, 
                        (err: any) => this.resolveError(err.message),
                        this.selectedProposal.index,
                        type, 
                        );
                },
                finish: function() {
                    voting.blockchain.sendContract("FloteaICO", "finishProposal", { "from": this.userAddress },
                        (err: any) => this.resolveError(err.message), 
                        this.selectedProposal.index
                        );
                },
                resolveError: function(str: string){
                    var re1= /revert (.*)"/;
                    var re2= /message":"(.*\.)\\n/;
                    var result = str.match(re1);
                    if(result && typeof result[1] != "undefined"){
                        $.jGrowl(result[1], { header: 'Error', life: 10000 });
                    }
                    result = str.match(re2);
                    if(result && typeof result[1] != "undefined"){
                        $.jGrowl(result[1], { header: 'Error', life: 10000 });
                    }
                },
                listStyle: function(i: any){
                    if(i.result == 2 && this.timeout(i.endTime)){
                        return "list-group-item-warning";
                    } else if (i.result != 2) {
                        return (i.result == 0? "list-group-item-danger" : "list-group-item-success");
                    } else return "";
                },
                enableActionButtons: function(){
                    return this.selectedProposal.result != 2 || this.timeout(this.selectedProposal.endTime);
                },
                timeLeft: function(val: number){
                    return moment().to(val * 1000);
                },
                timeout: function(item: number){
                    return moment(item * 1000) < moment();
                },
            }
        });
    }

    registerEvents(contract: any){
        this.loadProposals();
        contract.events.allEvents(null, (a:any,b:any)=>{
            console.log(a,b);
            if(a != false) return; 
            switch(b.event){
                case "ProposalCreated": 
                $.jGrowl("Proposal created", { header: 'Success', life: 20000 });
                sharedData.proposals = [{
                    description: this.blockchain.toHex(b.returnValues.description),
                    endTime: b.returnValues.endTime,
                    amount: b.returnValues.amount,
                    actionAddress: b.returnValues.actionAddress,
                    actionType: b.returnValues.actionType,
                    result: 2,
                    name: this.blockchain.toHex(b.returnValues.description),
                    index: b.returnValues.proposalIndex
                }].concat(sharedData.proposals);
                sharedData.description = "";
                sharedData.durationHours = 2;
                sharedData.amount = 0;
                sharedData.address = "";
                sharedData.name = "";
                sharedData.showModal = false;
                break;
                case "Vote":
                $.jGrowl("Somebody "+
                    (b.returnValues.vote == 0? "voted No" : (b.returnValues.vote==1? "voted Yes" : "Resignated")) +
                    " in proposal: '" + 
                    this.blockchain.toHex(b.returnValues.description) +
                    "' from address: " +
                    b.returnValues.addr, { header: 'Success', life: 20000 });
                if(sharedData.selectedProposal.index == b.returnValues.proposalIndex){
                    this.loadProposalStatus(sharedData.selectedProposal.index);
                }
                break;
                case "FinishVoting":
                let index = b.returnValues.proposalIndex;
                let result = b.returnValues.result? 1:0;
                $.jGrowl("Voting '"+ 
                    this.blockchain.toHex(b.returnValues.description) +
                    "' finished " 
                    , { header: 'Success', life: 20000 });
                if(sharedData.selectedProposal.index == index){
                    sharedData.selectedProposal.result = result;
                }
                for (var i = 0; i < sharedData.proposals.length; i++) {
                    if(sharedData.proposals[i].index == index){
                        sharedData.proposals[i].result = result;
                    }
                }
                break;
            }
        });
    }

    loadProposals(){
        let f = (index: number, proposalsLength: number) => {
            this.blockchain.callContract("FloteaICO", "proposals", (r: any) => {
                //console.log(r, index);
                sharedData.proposals.push({
                    description: this.blockchain.toHex(r.description),
                    endTime: r.endTime,
                    amount: r.amount,
                    actionAddress: r.actionAddress,
                    actionType: r.actionType,
                    result: r.result,
                    name: this.blockchain.toHex(r.name),
                    index: index
                });
                if(index > 0 && index > proposalsLength - 3)
                    f(index - 1, proposalsLength);
            }, null, [index]);
        }

        if(sharedData.proposals.length == 0){
            this.blockchain.callContract("FloteaICO", "proposalsLength", (r: any) => {
                sharedData.proposalsLength = r;
                if(sharedData.proposalsLength > 0){
                    f(sharedData.proposalsLength-1, sharedData.proposalsLength);
                }
            }, null);
        } else {
            f(sharedData.proposalsLength - sharedData.proposals.length -1 , sharedData.proposalsLength - sharedData.proposals.length);
        }
    }

    loadProposalStatus(proposalIndex: number){
        this.blockchain.callContract("FloteaICO", "statusOfProposal", (r: any) => {
            sharedData.selectedStatus = [];
            if(typeof r[0] == "object"){
                let empty = '0x'+'0'.repeat(40);
                for (var i = 0; i < r[0].length; i++) {
                    if( r[0][i] != empty ) {
                        sharedData.selectedStatus.push({
                            "address": r[0][i],
                            "name": this.blockchain.toHex(r[1][i]),
                            "vote": r[2][i]
                        });
                    }
                }
            }
        }, null, proposalIndex);
    }
}


$(document).ready(()=>{
    if($("#voting").length > 0)
        new Voting();
});