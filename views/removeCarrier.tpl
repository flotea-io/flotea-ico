<div id="remove-carrier">
	
</div>

<script type="text/x-template" id="remove-carrier-template" class="hide">
	<div class="container remove-carrier">
		<br />
		<div class="text-right">
			<MetamaskButton @metamaskUpdated="(s) => metamaskState = s"></MetamaskButton>
		</div>
		<div v-if="tx.length==0">
			<h3>Zgłosz przewoznika</h3>
			<div class="row" id="form">
				<div class="col-xs-12 col-sm-10 col-md-8 col-lg-6">
					<div class="form-group">
						<label>Adres przewoznika</label>
						<input type="text" v-model="removeAddress" class="form-control">
					</div>
					<div class="form-group">
						<label>Powód</label>
						<div class="input-group">
							<input type="text" v-model="reason" class="form-control" />
							<span class="input-group-addon">
								<span class="badge">${ max32(reason) }</span>
							</span>
						</div>
						<p class="text-danger" v-if="reason.trim().length == 0">Prosze wypełnić powód</p>
					</div>
					<div class="form-group">
						<label>Adres s dokumentami</label>
						<div class="input-group">
							<input type="text" v-model="web" class="form-control" />
							<span class="input-group-addon">
								<span class="badge">${ max32(web) }</span>
							</span>
						</div>
					</div>
				</div>
			</div>
			<p v-if="errorText.length > 0" class="text-danger">${ errorText }</p>
			<button class="btn btn-primary" :disabled="validate()" @click="submit">Zgłosz</button>
		</div>
		<div v-if="tx.length>0">
			<p class="text-success tx">
				<a class="btn btn-primary" href="/vote/applicants" data-toggle="tooltip" data-placement="bottom" title="Aplikujący przewoźnicy">Aplikanci</a>
			</p><br>
			Success, you can verify transaction:
			<a :href="'https://kovan.etherscan.io/tx/' + tx">${tx}</a>
		</div>
	</div>
</script>