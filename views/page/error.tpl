{{ template "layout/layout.tpl" . }}

{{ define "content" }}
<div id="login-page" class="row">

	<div class="col s12 z-depth-4 card-panel">

		<form class="login-form" action="/auth/login">

			<div class="row">
				<div class="input-field col s12 center">
					<img src="{{ logo }}" alt="" class="circle responsive-img valign profile-image-login">
					<h4 class="header">Error !!</h4>
				</div>
			</div>

			{{ template "partial/flash.tpl" . }}
			{{ template "partial/errmsg.tpl" . }}

			<div class="row">
				<div class="input-field col s12">
					<a href="/auth/forgot" class="btn waves-effect waves-light col s12">Forgot Password?</a>
				</div>

				<div class="input-field col s6">
					<a href="/auth/login" class="btn waves-effect waves-light col s12">Login</a>
				</div>

				<div class="input-field col s6">
					<a href="/" class="btn waves-effect waves-light col s12">Home</a>
				</div>
			</div>



		</form>

	</div>
	
</div>
{{ end }}