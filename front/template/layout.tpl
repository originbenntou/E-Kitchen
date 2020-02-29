{{define "layout"}}
<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<meta http-equiv="x-ua-compatible" content="ie=edge">
	<title>{{.PageName}} | E-KITCHEN </title>
	<link rel="stylesheet" type="text/css" href="./static/style.css">

	<!-- UIkit CSS -->
	<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/uikit/3.2.0/css/uikit.min.css" />

	<!-- UIkit JS -->
	<script src="https://cdnjs.cloudflare.com/ajax/libs/uikit/3.2.0/js/uikit.min.js"></script>
	<script src="https://cdnjs.cloudflare.com/ajax/libs/uikit/3.2.0/js/uikit-icons.min.js"></script>

	<!-- Vue-cdn -->
	<script src="https://cdn.jsdelivr.net/npm/vue@2.6.11"></script>

	<!-- axios-cdn -->
	<script src="https://unpkg.com/axios/dist/axios.min.js"></script>

	<!-- VeeValidate-cdn -->
	<script src="https://cdn.jsdelivr.net/npm/vee-validate@3.2.3/dist/vee-validate.js"></script>
	<script src="https://cdn.jsdelivr.net/npm/vee-validate@3.2.3/dist/rules.umd.min.js"></script>
</head>
<body class="app">
<div class="app__main" id="appMain">
    {{template "header" . }}
    {{template "content" . }}
</div>
</body>
</html>
{{end}}
