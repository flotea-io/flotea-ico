<!DOCTYPE html>
<html>
<head>
    <title>Flotea Token</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <link href="https://fonts.googleapis.com/css?family=Open+Sans:300,400,600,700,800" rel="stylesheet">
    <link rel="stylesheet" type="text/css" href="https://cdnjs.cloudflare.com/ajax/libs/jquery-jgrowl/1.4.6/jquery.jgrowl.min.css">
    <link rel="stylesheet" href="/static/css/index.css">
    <link rel="stylesheet" id="contrast" type="text/css" href="/static/css/contrast.css" disabled>
    <link rel="icon" href="/static/img/favicon.ico" type="image/x-icon"/>
    {{.HtmlHead}}
</head>
<body>
	{{.Header}}
    
    <div class="content">        
    {{.LayoutContent}}
    </div>

    {{.Footer}}

   
    <script type="text/javascript" src="https://code.jquery.com/jquery-2.0.3.min.js"></script>
    <script type="text/javascript" src="https://www.google.com/recaptcha/api.js"></script>
    <script type="text/javascript" src="https://stackpath.bootstrapcdn.com/bootstrap/3.4.1/js/bootstrap.min.js"></script>
    {{ range $val := .Javascript }}
        <script src="{{ $val }}"></script>
    {{ end }} 
</body>
</html>