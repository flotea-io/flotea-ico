<div class="container-fluid" id="first-section">
	<div class="header">
		<h2>Zdecentralizowany system zarządzania <b>transportem przewozu osób</b></h2>
		<a href="/ico/get" class="btn btn-danger">KUP TOKEN FLT</a>
	</div>
	<div class="container">
		<div class="row white-blocks">
			<div class="col-md-3 col-sm-6">
				<div class="panel panel-default">
					<div class="panel-heading">
						<div></div>
					</div>
					<div class="panel-body">
						<h3 class="panel-title">Dla przewoźników</h3>
						<p>Przewoźnicy mogą umieszczać swoje trasy bezpośrednio w publicznej zdecentralizowanej bazie danych opartej o standard GTFS, rozszerzając znacznie zakres dostępności tras.</p>
					</div>
				</div>
			</div>
			<div class="col-md-3 col-sm-6">
				<div class="panel panel-default">
					<div class="panel-heading">
						<div></div>
					</div>
					<div class="panel-body">
						<h3 class="panel-title">Dla agentów</h3>
						<p>Dowolni agenci bez koniecznosci podpisywania umowy mogą sprzedawać bilety i miejsca na trasy pobierając określoną zaliczkę w momencie kiedy bilet jest rozliczony.</p>
					</div>
				</div>
			</div>
			<div class="col-md-3 col-sm-6">
				<div class="panel panel-default">
					<div class="panel-heading">
						<div></div>
					</div>
					<div class="panel-body">
						<h3 class="panel-title">Dla programistów</h3>
						<p>Otwarte API, narzędzia i kod źródłowy umożliwia dowolnie rozbudowywać oprogramowanie, uwalniać je od potencjalnych błędów oraz integrować.</p>
						<a href="https://github.com/flotea-io">Link do github</a>
					</div>
				</div>
			</div>
			<div class="col-md-3 col-sm-6">
				<div class="panel panel-default">
					<div class="panel-heading">
						<div></div>
					</div>
					<div class="panel-body">
						<h3 class="panel-title">Token FLT</h3>
						<p>Token FLT to wartość, która służy do błyskawicznego rozliczania zakupionych biletów. FLT jest automatycznie rozliczany podczas zakupu biletu pomiędzy przewoźnika i agenta sprzedaży..</p>
					</div>
				</div>
			</div>
		</div>
	</div>
</div>
<div id="token-info">

</div>

<div id="what-is-ftn" class="container">
	<div class="row">
		<div class="col-md-6">
			<h2>Co to jest<br>FLT token</h2>
			<p>Token Flotea FLT zapewnia niezakłócone transakcje peer-to-peer na całym świecie dzięki zdecentralizownemu ekosystemowi w technologii blockchain zaprojektowanemu specjalnie dla transportu przewozu osób.</p>
			<p>Dzięki tokenowi FLT rozliczysz się w kilka sekund w trakcie każdej transakcji bezpośrednio z Twoim klientem i nie zapłacisz zbędnych prowizji naliczanych przez banki lub pośredników.</p>
			<a href="/ico/get" class="blue-outline-button">Zobacz więcej</a>
		</div>
		<div class="col-md-6">
			<img src="/static/img/flotea-token (2).jpg" class="center-block img-responsive">
		</div>
	</div>
</div>


<div id="how-it-works">
	<div class="container">
		<span>ROADMAP</span>
		<h2>How it works</h2>
		<img src="/static/img/roadmap-flotea.png" class="center-block img-responsive">
	</div>
</div>



<script type="text/x-template" id="app-template" class="hide">
	<div class="container-fluid" id="token-info">
		<div class="dark-blue row">
			<h3>Wyemitowane tokeny w fazie</h3>
				<div class="col-sm-4">${ icoBalanceWei/10**18 } ETH<p>ETH RECEIVED</p></div>
				<div class="col-sm-4">${ icoBalancePln } PLN<p>PLN RECEIVED</p></div>
				<div class="col-sm-4">${ (total/1000).toLocaleString('en').replace(/,/g," ",) } FLT<p>TOTAL</p></div>
		</div>
		<div class="white container">
			<h3>Etap: ${ phase + 1 } / ${ phases.length }</h3>
			<div class="progress">
				<div class="progress-bar progress-bar-info progress-bar-striped active" role="progressbar" :aria-valuenow="percentTokensInPhase" aria-valuemin="0" aria-valuemax="100" v-bind:style="{ width: percentTokensInPhase + '%' }">
					${ percentTokensInPhase.toFixed(2) }% Complete
				</div>
			</div>
			<div class="head">Fazy</div>
			<div class="tokens">
				${ ((phases[phase]-leftInActualPhase)/1000).toLocaleString('en').replace(/,/g," ",) } FLT 
				<span>/ ${ (phases[phase]/1000).toLocaleString('en').replace(/,/g," ",)  } FLT</span>
			</div>
			<br>
			<button class="btn btn-link" @click="showPhases = !showPhases">${showPhases? "Zamkni":"Pokaz"} wszystkie etapy</button>
			<div v-if="showPhases" class="table-responsive phases">
				<br>
				<table class="table table-condensed table-hover">
					<thead>
						<tr>
							<th colspan="2"></th>
							<th colspan="2">Cena za token w</th>
						</tr>
						<tr>
							<th>Etapy</th>
							<th>Tokenów</th>
							<th>PLN</th>
							<th>ETH</th>
						</tr>
					</thead>
					<tbody>
						<tr v-for="(tokens, index) in phases" v-bind:class="{complete : index<phase, actual: index==phase}">
							<td>${index + 1 }</td>
							<td><span>${(tokens/1000).toLocaleString('en').replace(/,/g," ",)}</span></td>
							<td><span>${Math.round(prices[index] / 10**16 * ethPrice*100)/100}</span></td>
							<td><span>${prices[index] / 10**16}</span></td>
						</tr>
					</tbody>
				</table>
			</div>
		</div>
	</div>
</script>