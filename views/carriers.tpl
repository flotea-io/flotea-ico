
<div id="carriers"></div>

<script type="text/x-template" id="carriers-template" class="hide">
	<div class="container carriers">
	<h1>Lista przewoźników</h1>
	<h2 class="subhead">Przewoźnicy którzy zostali zaakceptowani i dołączeni przes innych przewoźników</h2>

	<div class="nav-buttons">
		<a class="btn btn-primary" href="/vote/add" data-toggle="tooltip" data-placement="bottom" title="Dołącz jako nowy przewoźnik">Dołącz</a>
		<a class="btn btn-primary" href="/vote/applicants" data-toggle="tooltip" data-placement="bottom" title="Aplikujący przewoźnicy">Aplikanci</a>
	</div>
		<div class="row">
			<div class="col-xs-12" id="listContainer">
				<div v-if="loading" class="progress">
					<div class="progress-bar progress-bar-success progress-bar-striped active" role="progressbar" aria-valuenow="40" aria-valuemin="0" aria-valuemax="100" style="width: 100%">
					</div>
				</div>
				<div class="carrier-block" v-for="(carrier, index) in carriers" v-if="carrier.enabled">
					<div class="flex-columns">
						<div class="limit-avatar">
							<div class="avatar">
								<div>
									${avatarChars(carrier.name)}
								</div>
							</div>
						</div>
						<div>
							<h4 class="name">${carrier.name}</h4>
							<div class="wallet break-word">${carrier.companyWallet}</div>
							<a class="web" :href="carrier.web" target="_blank">${carrier.web}</a>
						</div>
						<div>
							<a :href="'/vote/remove/'+carrier.companyWallet" class="btn btn-primary">Zgłoś do usunięcia</a>
						</div>
					</div>
				</div>
			</div>
		</div>
		<p class="text-center">
			<button class="btn btn-primary" @click="loadNext" v-if="!loading && carriersCount > carriers.length">Doładuj kolejne</button>
		</p>
	</div>
</script>