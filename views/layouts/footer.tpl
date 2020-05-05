<div id="footer">
		<div class="white-bg">
				<img src="/static/img/FE_logo.png">
				<img src="/static/img/RP_logo.png">
				<img src="/static/img/Slaskie_logo.png">
				<img src="/static/img/UE_logo2.png">
		</div>
	<div class="container blue">
		<div class="row">
			<div class="col-md-6 col-lg-7">
				<label>Chcesz się dowiedzieć więcej?<br>Skontaktuj się z nami</label>
				<p>ul. Zamkowa 3A<br>43-400 Cieszyn</p>
				<p>
					Email: info@flotea.pl<br />
					<span>Kom:</span> +48 881 327 624<br />
					<span></span> +48 791 139 005
				</p>
				<p class="social">
					<img src="/static/img/inteligentny-obiekt-wektorowy-2@2x.png" />
					<img src="/static/img/inteligentny-obiekt-wektorowy-3@2x.png" />
					<img src="/static/img/inteligentny-obiekt-wektorowy-4@2x.png" />
					<img src="/static/img/inteligentny-obiekt-wektorowy-5@2x.png" />
				</p>
			</div>
			<div class="col-md-6 col-lg-5 max-460">
				<form method="POST" action="/send-contact">
					{{if .errors.NameSurname}}
						<div class="form-group has-error">
					{{else}}
						<div
						 class="form-group">
					{{end}}
						<input type="text" name="nameSurname" value="{{.form.NameSurname}}" class="form-control" placeholder="Imię i nazwisko">
						{{if .errors.NameSurname}}
							{{ index .errors.NameSurname 0}}
						{{end}}
					</div>

					{{if .errors.Email}}
						<div class="form-group has-error">
					{{else}}
						<div class="form-group">
					{{end}}
							<input type="text" name="email" value="{{.form.Email}}" class="form-control" placeholder="E-mail">
							{{if .errors.Email}}
								{{ index .errors.Email 0}}
							{{end}}
					</div>
					<div class="form-group">
						<input type="text" name="phone" value="{{.form.Phone}}" class="form-control" placeholder="Nr telefonu">
					</div>
					{{if .errors.Message}}
						<div class="form-group has-error">
					{{else}}
						<div class="form-group">
					{{end}}
						<textarea name="message" class="form-control" rows="3" placeholder="Twoja wiadomość">{{.form.Message}}</textarea>
						{{if .errors.Message}}
							{{ index .errors.Message 0}}
						{{end}}
					</div>
					{{if .errors.Recaptcha}}
						<div class="form-group has-error">
					{{else}}
						<div class="form-group">
					{{end}}
						<div class="g-recaptcha" data-sitekey="6LfgaDMUAAAAAEHrQRBs2Ohf4z2_1yuFOh4buiVa"></div>
					</div>
						<button type="submit" class="btn btn-white btn-outline">Wyślij</button>
					</form>
				</div>
			</div>

			<div class="text-center copyright">2019 © Flotea</div>
		</div>
	</div>