<nav class="navbar navbar-default navbar-fixed-top">
	<div class="container">
		<div class="navbar-header">
			<a class="navbar-brand" href="/">
				<img alt="Flotea" src="/static/img/logo_flotea.png">
			</a>

			<div class="help" tabindex="0"><span class="size" id="zoom" onclick="size()">A+</span> <span class="contrast" onclick="toggleContrast()"><span>A</span><span>A</span></span></div>

			<button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#colapse-navbar" aria-expanded="false">
				<span class="sr-only">Toggle navigation</span>
				<span class="icon-bar"></span>
				<span class="icon-bar"></span>
				<span class="icon-bar"></span>
			</button>
			<img class="visible-xs eu-head" src="/static/img/UE_logo2.png">
		</div>
		<div class="collapse navbar-collapse" id="colapse-navbar">
			<ul class="nav navbar-nav navbar-right">
				<li><a href="https://transport.flotea.com">Transport</a></li>
				<li class="{{if eq .TplName "carriers"}}active{{end}}" ><a href="/vote/carriers">Przewoźnicy</a></li>
				<li class="{{if eq .TplName "whitePaper"}}active{{end}}" ><a href="/white-paper">White paper</a></li>
				<li><a href="/#footer">Kontakt</a></li>
				<li><img class="hidden-xs eu-head" src="/static/img/UE_logo2.png"></li>
			</ul>
		</div>
	</div>
</nav>

<script type="text/javascript">
	var contrast = localStorage.getItem('contrast') == "true";
	var zoomSelect = parseInt(localStorage.getItem('zoom'));
	if(isNaN(zoomSelect))
		zoomSelect = 0;

	toggleBodyClass(true, "zoom-"+zoomSelect);
	applyContrast();

	function size(){
		toggleBodyClass(false, "zoom-"+zoomSelect);
		zoomSelect = zoomSelect > 1? 0 :zoomSelect+1;
		toggleBodyClass(true, "zoom-"+zoomSelect);
		localStorage.setItem('zoom', zoomSelect);
	}
	function toggleContrast() {
		contrast = !contrast;
		applyContrast();
	}
	function applyContrast(){
		document.getElementById("contrast").disabled = !contrast;
		localStorage.setItem('contrast', contrast);
	}
	function toggleBodyClass(addClass, className) {
		const el = document.body;
		if (addClass) {
			el.classList.add(className);
		} else {
			el.classList.remove(className);
		}
	}
</script>