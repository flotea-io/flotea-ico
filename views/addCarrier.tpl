<div id="add-carrier">
	
</div>

<script type="text/x-template" id="add-carrier-template" class="hide">
	<div class="container add-carrier">
		<br />
		<div class="text-right">
			<MetamaskButton @metamaskUpdated="(s) => metamaskState = s"></MetamaskButton>
		</div>
		<div v-if="tx.length==0">
			<h3>Dołacz ${ metamaskState.userType == 1? "ponownie":"" } do blockchain</h3>
			<div class="row" id="form">
				<div class="col-xs-12 col-sm-10 col-md-8 col-lg-6">
					<div v-if="metamaskState.userType == 0">
						<div class="form-group">
							<label>Company</label>
							<div class="input-group">
								<input type="text" v-model="company" class="form-control" />
								<span class="input-group-addon">
									<span class="badge">${ max32(company) }</span>
								</span>
							</div>
							<p class="text-danger" v-if="company.trim().length == 0 && !firstTime">Please fill your company name</p>
						</div>
						<div class="form-group">
							<label>Web page</label>
							<div class="input-group">
								<input type="text" v-model="web" class="form-control" />
								<span class="input-group-addon">
									<span class="badge">${ max32(web) }</span>
								</span>
							</div>
							<p class="text-danger" v-if="web.trim().length == 0 && !firstTime">Please fill your web page, or Facebook profile</p>
						</div>
					</div>
				</div>
			</div>
			<p v-if="errorText.length > 0" class="text-danger">${ errorText }</p>
			<button class="btn btn-primary" :disabled="validate()" @click="submit">Dołacz</button>
		</div>

		<div v-if="tx.length>0">
			<p class="text-success tx">
				<a class="btn btn-primary" href="/vote/applicants" data-toggle="tooltip" data-placement="bottom" title="Aplikujący przewoźnicy">Aplikanci</a>
			</p><br>
			Success, you can verify transaction: 
			<a :href="'https://kovan.etherscan.io/tx/' + tx">${tx}</a></p>
		</div>
	</div>
</script>