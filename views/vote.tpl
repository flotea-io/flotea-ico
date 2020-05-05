

<script type="text/x-template" id="app-template" class="hide">
	<div class="container">
		<div class="text-right">
			<MetamaskButton @metamaskUpdated="(s) => metamaskState = s"></MetamaskButton>
		</div>
		<h1>Voting system</h1>
		<hr>
		<div class="row">
			<div class="col-md-5">
				<div class="panel panel-default">
					<div class="panel-heading clearfix">
						<h4 class="panel-title pull-left">List of proposals</h4>
						<button class="btn btn-success pull-right btn-sm" @click="showModal = true" style="margin: -7px 0px;">
							<span class="glyphicon glyphicon-plus-sign"></span> &nbsp;
							Add proposal
						</button>
					</div>
					<div class="panel-body">
						<div class="list-group">
							<div v-for="item in proposals" class="list-group-item" v-bind:class="listStyle(item)" v-on:click="showProposal(item)">
								${ item.description }
								<div class="pull-right">${ timeLeft(item.endTime) }</div>
							</div>
						</div>
						<div class="text-center" style="margin-top: -14px;">
							<button class="btn btn-link btn-xs" :disabled="this.proposalsLength==this.proposals.length" v-on:click="loadProposals">Load more</button>
						</div>
					</div>
				</div>
			</div>

			<div class="col-md-7">
				<div class="panel panel-default" v-if="proposalShowing">
					<div class="panel-heading clearfix">
						<div class="pull-left">
							${ selectedProposal.description }
						</div>
						<div class="pull-right">
						<span class="label label-info">
							${ selectedProposal.actionType==0? 'Add participant' : (selectedProposal.actionType==1? 'Remove participant' : 'Transfer ether') }
						</span> &nbsp;
						<span class="label label-danger" v-if="selectedProposal.result != 2">finished</span>
						<span class="label label-warning" v-if="selectedProposal.result == 2 && timeout(selectedProposal.endTime)">time out</span>
						</div>
					</div>
					<div class="panel-body">
						<p>End of voting: ${ timeLeft(selectedProposal.endTime) }</p>
						<h4 v-if="selectedStatus.length>0">Data:</h4>
						<blockquote>
							<p v-if="selectedProposal.actionType==2">amount: ${ selectedProposal.amount }</p>
							<p>Address to ${selectedProposal.actionType==0? "add participant" : (selectedProposal.actionType==1? "remove participant": "transfer Wei")}: ${ selectedProposal.actionAddress }</p>
		
							<p v-if="selectedProposal.actionType==0">Participant name: ${ selectedProposal.name }</p>
						</blockquote>
						
						<h4 v-if="selectedStatus.length>0">Voting status:</h4>
						<table class="table table-striped table-bordered table-hover">
							<tbody>
						<tr v-for="i in selectedStatus">
							<td class="text-center" style="width: 86px;">
								<span class="label" v-bind:class="{'label-success': i.vote == 1, 'label-danger': i.vote == 0, 'label-info': i.vote == 2 }">${ i.vote == 1 ? 'Yes' : (i.vote == 2? 'Resignated' : 'No') }</span>
							</td>
							<td>
								${i.name}
							</td>
							<td style="width: 340px;"> 
								${i.address}
							</td>
						</tr>
					</tbody>
						</table>
					</div>
					<div class="panel-footer">
						<button class="btn btn-success" :disabled="enableActionButtons()" v-on:click="vote(1)">Yes</button>
						<button class="btn btn-danger" :disabled="enableActionButtons()" v-on:click="vote(0)">No</button>
						<button class="btn btn-info" :disabled="enableActionButtons()" v-on:click="vote(2)">Resignate</button> &nbsp; &nbsp; 
						<button class="btn btn-warning" :disabled="enableActionButtons()" v-on:click="finish">Finish</button>
						<button class="btn btn-default" v-on:click="proposalShowing=false">Close</button>
					</div>
				</div>
			</div>
		</div>
		
		<div v-if="showModal">
			<transition name="modal">
				<div class="modal-mask">
					<div class="modal-wrapper">
						<div class="modal-dialog">
							<div class="modal-content">
								<div class="modal-header">
									<button type="button" class="close" @click="showModal=false">
										<span aria-hidden="true">&times;</span>
									</button>
									<h4 class="modal-title">Create new proposal</h4>
								</div>
								<div class="modal-body">
									<div class="form-horizontal">
										<div class="form-group">
											<label class="col-sm-2 control-label">Description</label>
											<div class="col-sm-10">
												<input v-model="description" type="text" class="form-control"></input>
											</div>
										</div>
										<div class="form-group">
											<label class="col-sm-2 control-label">Duration (h)</label>
											<div class="col-sm-10">
												<input v-model="durationHours" type="number" class="form-control"></input>
											</div>
										</div>
										<div class="form-group">
											<label class="col-sm-2 control-label">Type</label>
											<div class="col-sm-10">
												<div class="btn-group">
													<label class="btn btn-default" v-bind:class="{ active: actionType==0 }" v-on:click="actionType=0">
														Add participant
													</label>
													<label class="btn btn-default" v-bind:class="{ active: actionType==1 }" v-on:click="actionType=1">
														Remove participant
													</label>
													<label class="btn btn-default" v-bind:class="{ active: actionType==2 }" v-on:click="actionType=2">
														Transfer Ether
													</label>
													<label class="btn btn-default" v-bind:class="{ active: actionType==3 }" v-on:click="actionType=3">
														Transfer FLT
													</label>
												</div>
											</div>
										</div>
										<div class="form-group">
											<label class="col-sm-2 control-label">Address</label>
											<div class="col-sm-10">
												<input v-model="address" type="text" class="form-control"></input>
											</div>
										</div>
										<div class="form-group" v-if="actionType==0">
											<label class="col-sm-2 control-label">Name</label>
											<div class="col-sm-10">
												<input v-model="name" type="text" class="form-control"></input>
											</div>
										</div>
										<div class="form-group" v-if="actionType==2">
											<label class="col-sm-2 control-label">Amount</label>
											<div class="col-sm-10">
												<input v-model="amount" type="number" class="form-control"></input>
											</div>
										</div>
										<div class="form-group">
											<div class="col-sm-2"></div>
											<div class="col-sm-10">
												<button class="btn btn-default" v-on:click="createProposal">Create</button>
											</div>
										</div>

									</div>
								</div>
							</div>
						</div>
					</div>
				</div>
			</transition>
		</div>





	</div>
</script>




<div id="voting"></div>