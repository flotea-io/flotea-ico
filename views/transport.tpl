<div id="transport"></div>

<script type="text/x-template" id="app-template" class="hide">
	<div class="container">
		<br>
		<div class="row">
			<div class="col-md-4">
				<label>Your address: &nbsp; ${ typeof userAddress != "undefined"? userAddress.substr(0,10)+"..":"" }</label>&nbsp; &nbsp;
				<button class="btn btn-default btn-xs" v-on:click="loadAddress">
					<span class="glyphicon glyphicon-refresh"></span>
				</button>
				<label>Token address: ${ tokenAddress }</label>
				<div class="panel panel-primary">
					<div class="panel-heading">Flotea transport</div>
					<div class="panel-body">
						<button class="btn btn-sm btn-default" v-on:click="transportAction=0" :class="{ active: transportAction==0 }">create carrier</button>
						<button class="btn btn-sm btn-default" v-on:click="getCarriers" :class="{ active: transportAction==1 }">get carriers</button>
						<button class="btn btn-sm btn-default" v-on:click="getTrips" :class="{ active: transportAction==2 }">get trips</button>
						<hr>
						<div :class="{ hide: transportAction != 0}">
							<div class="form-group">
								<label>Company</label>
								<input v-model="company" class="form-control" />
							</div>
							<div class="form-group">
								<label>Company address</label>
								<input v-model="companyAddress" class="form-control" />
							</div>
							<div class="form-group">
								<label>Phone</label>
								<input v-model="phone" class="form-control" />
							</div>
							<button class="btn btn-sm btn-success" v-on:click="createCarrier">Create</button>
						</div>
						<div :class="{ hide: transportAction != 1}">
							<div v-for="c in carriers">${ c.company } - ${ c.addr.substr(0, 8) }..</div>
						</div>
						<div :class="{ hide: transportAction != 2}">
							<div v-for="t in trips">${ t.addr }</div>
						</div>
					</div>
				</div>
			</div>
			<div class="col-md-8">
				<div class="form-inline">
					<div class="form-group">
						<label>List of trips</label>
						<select v-model="selectedTrip" class="form-control">
							<option v-for="(t, index) in trips" v-bind:value="index" :class="{ 'bg-danger' : !t.enabled }">${ t.addr }</option>
						</select>
					</div>
					<hr>
				</div>
				<br>
				<div class="row">
					<div class="col-md-6">
						<div class="panel panel-info">
							<div class="panel-heading">Carrier</div>
							<div class="panel-body">
								<button class="btn btn-sm btn-default" v-on:click="carrierAction=0" :class="{ active: carrierAction==0 }">create trip</button>
								<button class="btn btn-sm btn-default" v-on:click="carrierAction=9" :disabled="selectedTrip==null" :class="{ active: carrierAction==9 }">info</button>
								<button class="btn btn-sm btn-default" v-on:click="carrierAction=1" :disabled="selectedTrip==null" :class="{ active: carrierAction==1 }">refund</button>
								<button class="btn btn-sm btn-default" v-on:click="getTickets" :disabled="selectedTrip==null" :class="{ active: carrierAction==2 }">get tickets</button>
								<button class="btn btn-sm btn-default" v-on:click="availableChargeAmount" :disabled="selectedTrip==null" :class="{ active: carrierAction==3 }">available charge amount</button>
								<button class="btn btn-sm btn-default" v-on:click="carrierAction=4" :disabled="selectedTrip==null" :class="{ active: carrierAction==4 }">charge</button>
								<button class="btn btn-sm btn-default" v-on:click="carrierAction=5" :disabled="selectedTrip==null" :class="{ active: carrierAction==5 }">transfer ownership carrier</button>
								<button class="btn btn-sm btn-default" v-on:click="carrierAction=6" :disabled="selectedTrip==null" :class="{ active: carrierAction==6 }">set atribute</button>
								<button class="btn btn-sm btn-default" v-on:click="carrierAction=8" :disabled="selectedTrip==null" :class="{ active: carrierAction==8 }">set route type</button>
								<hr>
								<div :class="{ hide: carrierAction != 0}">
									<div class="form-group">
										<label>From Latitude</label>
										<input v-model="fromLat" class="form-control" />
									</div>
									<div class="form-group">
										<label>From Longitude</label>
										<input v-model="fromLon" class="form-control" />
									</div>
									<div class="form-group">
										<label>To Latitude</label>
										<input v-model="toLat" class="form-control" />
									</div>
									<div class="form-group">
										<label>To Longitude</label>
										<input v-model="toLon" class="form-control" />
									</div>
									<div class="form-group">
										<label>Price</label>
										<input type="number" min="0" v-model="price" class="form-control" />
									</div>
									<div class="form-group">
										<label>Schedule</label>
										<textarea v-model="schedule" class="form-control" ></textarea>
									</div>
									<div class="form-group">
										<label>Places</label>
										<input v-model="places" class="form-control" />
									</div>
									<div class="form-group">
										<label>Description</label>
										<input v-model="description" class="form-control" />
									</div>
									<div class="form-group">
										<label>Vehicle type</label>
										<select v-model="vehicleType" class="form-control">
											<option v-for="(item, index) in vehicles" v-bind:value="index">${ item.replace(/_/g," ") }</option>
										</select>
									</div>
									<div class="form-group">
										<label>Enabled</label>
										<input type="checkbox" v-model="enabled" />
									</div>
									<button class="btn btn-sm btn-success" v-on:click="createTrip">Create</button>
								</div>
								<div :class="{ hide: carrierAction != 1}">
									<div class="form-group">
										<label>Buyer</label>
										<input v-model="buyer" class="form-control"/>
									</div>
									<div class="form-group">
										<label>Time</label>
										<input v-model="time" class="form-control"/>
									</div>
									<button class="btn btn-sm btn-success" v-on:click="createTrip">Refund</button>
								</div>
								<div :class="{ hide: carrierAction != 2}">
									<p v-for="(t, index) in ticketsKeys"> ${ formatTime(t) } : ${ tickets[index] }</p>
								</div>
								<div :class="{ hide: carrierAction != 3}">${ chargeAmount } FLT</div>
								<div :class="{ hide: carrierAction != 4}">
									<div class="form-group">
										<label>Address</label>
										<input v-model="chargeAddress" class="form-control"/>
									</div>
									<button class="btn btn-sm btn-success" v-on:click="charge">Charge</button>
								</div>
								<div :class="{ hide: carrierAction != 5}">
									<div class="form-group">
										<label>Address</label>
										<input v-model="ownerAddress" class="form-control"/>
									</div>
									<button class="btn btn-sm btn-success" v-on:click="transferOwnership">TransferOwnership</button>
								</div>
								<div :class="{ hide: carrierAction != 6}">
									<div class="form-group">
										<label>From Latitude</label>
										<input v-model="fromLat" class="form-control" />
									</div>
									<div class="form-group">
										<label>From Longitude</label>
										<input v-model="fromLon" class="form-control" />
									</div>
									<div class="form-group">
										<label>To Latitude</label>
										<input v-model="toLat" class="form-control" />
									</div>
									<div class="form-group">
										<label>To Longitude</label>
										<input v-model="toLon" class="form-control" />
									</div>
									<div class="form-group">
										<label>Price</label>
										<input type="number" min="0" v-model="price" class="form-control" />
									</div>
									<div class="form-group">
										<label>Schedule</label>
										<textarea v-model="schedule" class="form-control" ></textarea>
									</div>
									<div class="form-group">
										<label>Places</label>
										<input v-model="places" class="form-control" />
									</div>
									<div class="form-group">
										<label>Description</label>
										<input v-model="description" class="form-control" />
									</div>
									<button class="btn btn-sm btn-success" v-on:click="setAttributes">Set</button>
								</div>
								<div :class="{ hide: carrierAction != 7}">
									<div class="form-group">
										<label>Address</label>
										<input v-model="ownerAddress" class="form-control"/>
									</div>
									<button class="btn btn-sm btn-success" v-on:click="transferOwnership">TransferOwnership</button>
								</div>
								<div :class="{ hide: carrierAction != 8}">
									<div class="form-group">
										<label>Route type</label>
										<select v-model="vehicleType" class="form-control">
											<option v-for="(item, index) in vehicles" v-bind:value="index">${ item.replace(/_/g," ") }</option>
										</select>
									</div>
									<button class="btn btn-sm btn-success" v-on:click="setVehicle">Set</button>
								</div>
								<div :class="{ hide: carrierAction != 9}" v-if="tripContract && tripContract">
									Description: ${ fromHex(tripContract._description) }
									<p>Trip: <span class="label label-success">${ fromHex(tripContract._fromLat) }, ${ fromHex(tripContract._fromLng) }</span> &nbsp;
									- &nbsp;
									<span class="label label-danger">${ fromHex(tripContract._toLat) }, ${ fromHex(tripContract._toLng) }</span><p/>
									Places: ${ tripContract._places } Route type: ${ vehicles[tripContract._routeType] }
									<p class="wrap">Schedule: ${ fromHex(tripContract._schedule) }</p>
								</div>
							</div>
						</div>
					</div>
					<div class="col-md-6">
						<div class="panel panel-success">
							<div class="panel-heading">Passager</div>
							<div class="panel-body">
								<div>
									<div class="form-group">
										<label>Price</label>
										<input v-model="price" class="form-control"/>
									</div>
									<div class="form-group">
										<label>Time</label>
										<input v-model="time" class="form-control"/>
									</div>
									<button class="btn btn-sm btn-success" v-on:click="buyTicket" :disabled="selectedTrip==null">Buy</button>
								</div>
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
</script>