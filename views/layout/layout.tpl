<!DOCTYPE html>
<html lang="en">

<head>
<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<meta name="msapplication-tap-highlight" content="no">
<meta name="description" content="">
<meta name="keywords" content="">
<title>{{ siteName }}</title>

<!-- Favicons-->
<link rel="icon" href="/auth/images/favicon-32x32.png" sizes="32x32">
<!-- Favicons-->
<link rel="apple-touch-icon-precomposed" href="/auth/images/apple-touch-icon-152x152.png">
<!-- For iPhone -->
<meta name="msapplication-TileColor" content="#00bcd4">
<meta name="msapplication-TileImage" content="/user/images/mstile-144x144.png">
<!-- For Windows Phone -->

<style>
.top-aside,
.banner,
.small_banner,
.banner-slider,
footer {
{{ with backgroundImage }}
	background-image: url("{{ . }}");
{{ end  }}
background-position: center center;
background-repeat: repeat;
{{ with backgroundColor }}
	background-color: {{ . }} ;
	/* background: -webkit-gradient(radial, center center, 0, center center, 460, from(#FFFFFF), to({{.}}));
	background: -webkit-radial-gradient(circle, #FFFFFF, {{.}});
	background: -moz-radial-gradient(circle, #FFFFFF, {{.}});
	background: -ms-radial-gradient(circle, #FFFFFF, {{.}}); */
{{ end  }}
}
</style>

<!-- CORE CSS-->
<link href="/auth/css/materialize.css" type="text/css" rel="stylesheet" media="screen,projection">
<link href="/auth/css/style.css" type="text/css" rel="stylesheet" media="screen,projection">

<!-- Custome CSS--> 
<link href="/auth/css/custom.css" type="text/css" rel="stylesheet" media="screen,projection">
<link href="/auth/css/page-center.css" type="text/css" rel="stylesheet" media="screen,projection">

<!-- INCLUDED PLUGIN CSS ON THIS PAGE -->
<link href="/auth/css/prism.css" type="text/css" rel="stylesheet" media="screen,projection">
<link href="/auth/css/perfect-scrollbar.css" type="text/css" rel="stylesheet" media="screen,projection">

</head>

<body class="banner">

{{ block "content" . }}{{ end }}

<!-- ================================================
Scripts
================================================ -->

<!-- jQuery Library -->
<script type="text/javascript" src="/auth/js/jquery-1.11.2.min.js"></script>
<!--materialize js-->
<script type="text/javascript" src="/auth/js/materialize.js"></script>
<!--prism-->
<script type="text/javascript" src="/auth/js/prism.js"></script>
<!--scrollbar-->
<script type="text/javascript" src="/auth/js/perfect-scrollbar.min.js"></script>

<!--plugins.js - Some Specific JS codes for Plugin Settings-->
<script type="text/javascript" src="/auth/js/plugins.js"></script>
<!--custom-script.js - Add your own theme custom JS-->
<script type="text/javascript" src="/auth/js/custom-script.js"></script>

<!-- vue.js and axiom for ajax requests -->
<script type="text/javascript" src="/auth/js/vue.min.js"></script>
<script type="text/javascript" src="/auth/js/vee-validate.js"></script>
<!-- script type="text/javascript" src="/auth/js/axiom.min.js"></script -->

{{ block "javascript" . }} {{ end }}

</body>
</html>
