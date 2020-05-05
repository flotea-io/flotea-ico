<div id="applicants">
	
</div>

<script type="text/x-template" id="applicants-template" class="hide">
	<div>
		<div class="container carriers">
			<br />
			<div class="text-right">
				<MetamaskButton @metamaskUpdated="(s) => metamaskState = s"></MetamaskButton>
			</div>
			<h1>Przewoźnicy aplikujący</h1>
	<h2 class="subhead">Lista przewoźników, którzy chcą dołączyć do grona firm oferujących przewozy</h2>

	<div class="nav-buttons">
		<a class="btn btn-primary" href="/vote/add" data-toggle="tooltip" data-placement="bottom" title="Dołącz jako nowy przewoźnik">Dołącz</a>
		<a class="btn btn-primary" href="/vote/carriers" data-toggle="tooltip" data-placement="bottom" title="Lista przewoźników">Przewoznicy</a>
	</div>

			<div class="row">
				<div class="col-xs-12" id="listContainer">
					<div class="carrier-block" v-for="(applicant, index) in applicants" v-bind:class="{ended:applicant.result != 2, remove: applicant.actionType == 1}">
						<div class="flex-columns">
							<div class="limit-avatar">
								<span class="label label-info info" v-if="applicant.result==2 && applicant.actionType==0" v-tooltip.right="'przewoźnik w trakcie dodawania - głosujemy'">Dodawany</span>
								<span class="label label-warning info" v-if="applicant.result==2 && applicant.actionType==1" v-tooltip.right="'przewoźnik w trakcie usuwania - głosujemy'">Usuwany</span>

								<span class="label label-success info" v-if="applicant.result==1 && applicant.actionType==0" v-tooltip.right="'dodany do grupy przewoźników'">Dodany</span>

								<span class="label label-danger info" v-if="applicant.result==0 && applicant.actionType==0" v-tooltip.right="'przewoźnik nie został dodany'">Nie dodany</span>

								<span class="label label-success info" v-if="applicant.result==1 && applicant.actionType==1" v-tooltip.right="'przewoźnik usunięty'">Usunięty</span>

								<span class="label label-danger info" v-if="applicant.result==0 && applicant.actionType==1" v-tooltip.right="'przewoźnik zostaje'">Nie usunięty</span>
								
								<div class="avatar">
									<div>
										${avatarChars(applicant.name)}
									</div>
								</div>
							</div>
							<div class="row">
								<div class="col-md-8">
									<p class="end">
										<a @click="detail(index)">Szczegóły głosowania</a>
										${applicant.result==2? (ended(applicant.endTime)? "Głosowanie #"+ applicant.index +" kończy się" : "Głosowanie #"+applicant.index+" skączyło się") : "Głosowanie #"+applicant.index+" zakończone" } 
										<span>${applicant.result==2? toEndText(applicant.endTime) : ""}</span>
									</p>
									<h4 class="name">${applicant.name}</h4>
									<div class="wallet break-word">${applicant.actionAddress}</div>
									<a class="web" :href="applicant.web" target="_blank">${applicant.web}</a>
								</div>
								<div class="col-md-4">
									<label class="sum">Suma głosów: ${applicant.voted}</label>
									<div class="progress-detail">
										<div>Tak</div>
										<div class="progress">
											<div class="progress-bar progress-bar-info" role="progressbar" aria-valuenow="20" aria-valuemin="0" aria-valuemax="100" v-bind:style="{ width: applicant.voted==0? 0 : Math.round(applicant.yes/applicant.voted*100) + '%'}">
											</div>
										</div>
										<div>
											${ applicant.voted==0? 0 : Math.round(applicant.yes/applicant.voted*100) }%
										</div>
									</div>
									<div class="progress-detail">
										<div>Nie</div>
										<div class="progress">
											<div class="progress-bar progress-bar-info" role="progressbar" aria-valuenow="20" aria-valuemin="0" aria-valuemax="100" v-bind:style="{ width: applicant.voted==0? 0 : Math.round(applicant.no/applicant.voted*100) + '%'}">
											</div>
										</div>
										<div>
											${ applicant.voted==0? 0 : Math.round(applicant.no/applicant.voted*100) }%
										</div>
									</div>
									<div class="progress-detail">
										<div></div>
										<div>
											<button v-if="ended(applicant.endTime) && applicant.result == 2 && applicant.tx.length == 0 && isVoted(applicant) == -1" class="btn btn-primary btn-block vote-button" :disabled="!metamaskState.enabledMetamask" @click="applicant.showVoteButtons = !applicant.showVoteButtons">
												Głosuję
											</button>
											<button v-if="(!ended(applicant.endTime) || isVoted(applicant) != -1) && applicant.result == 2" class="btn btn-primary btn-block vote-button" @click="finish(applicant.index)">
												Zakończ
											</button>
										</div>
										<div>
										</div>
									</div>
								</div>
							</div>
						</div>
						<div class="vote-buttons" v-if="applicant.showVoteButtons && applicant.tx.length == 0 && isVoted(applicant) == -1">
							<hr />
							<div>
								<button class="btn btn-yes" @click="vote(1, applicant.index)">Tak</button>
								<button class="btn btn-no" @click="vote(0, applicant.index)">Nie</button>
								<button class="btn btn-resigned" @click="vote(2, applicant.index)">Wstrzymuje się</button>
							</div>
						</div>
						<p v-if="applicant.tx.length>0" class="text-success tx">
							Success, you can verify transaction: 
							<a :href="'https://kovan.etherscan.io/tx/' + applicant.tx" target="_blank" class="break-word">${applicant.tx}</a>
						</p>
					</div>

					<div v-if="loading" class="progress">
						<div class="progress-bar progress-bar-success progress-bar-striped active" role="progressbar" aria-valuenow="40" aria-valuemin="0" aria-valuemax="100" style="width: 100%">
						</div>
					</div>
					<p class="text-center">
						<button class="btn btn-primary" @click="loadNext" v-if="!loading && applicantsCount > applicants.length">Doładuj kolejne</button>
					</p>
				</div>
			</div>
		</div>
		<modal class="carriers-modal" v-if="showModal" @close="showModal = false">
			<h3 slot="header">Szczegóły głosowania dla ${applicants[selectedApplicant].name}</h3>
			<div slot="body" v-if="selectedApplicant!=-1">
				<div class="results" v-for="(voter, index) in applicants[selectedApplicant].voters" v-if="voter!='0x0000000000000000000000000000000000000000'">
					<span class="wallet break-word">${voter}</span>
					<span>
						<span v-if="applicants[selectedApplicant].votersVotes[index]==0" class="no">Nie</span>
						<span v-if="applicants[selectedApplicant].votersVotes[index]==1" class="yes">Tak</span>
						<span v-if="applicants[selectedApplicant].votersVotes[index]==2" class="resigned">Wstrzymany</span>
					</span>
				</div>
			</div>
		</modal>
	</div>
</script>


<script type="text/x-template" id="modal-template">
	<transition name="modal">
		<div class="modal-mask">
			<div class="modal-container">

				<div class="modal-header">
					<slot name="header">
					</slot>
				</div>

				<div class="modal-body">
					<slot name="body">
					</slot>
				</div>

				<div class="modal-footer">
					<slot name="footer">
						<button class="btn btn-primary" @click="$emit('close')">
							Zamknij
						</button>
					</slot>
				</div>
			</div>
		</div>
	</transition>
</script>