{{ template "layout.tpl" . }}

{{ define "content" }}
<div id="login-page" class="row">

	<div class="col s12 z-depth-4 card-panel">

		<form class="login-form" action="/auth/forgot" method="POST">
            <input type="hidden" name="gorilla.csrf.Token" value="{{ .CSRFToken }}">

			<div class="row">
				<div class="input-field col s12 center">
					<h4>Forgot Password</h4>
					<p class="center">You can reset your password</p>
				</div>
			</div>

			{{ template "flash.tpl" . }}

			<div class="row margin">
				<div class="input-field col s12">
					<i class="mdi-communication-email prefix"></i>
					<input id="email" type="email" name="email" required>
					<label for="email" class="center-align">Email</label>
				</div>
			</div>


			<div class="row">
				<div class="input-field col s12">
						<button class="btn waves-effect waves-light col s12" type="submit">Forgot Password</button>
				</div>

				<div class="input-field col s6 m6 l6">
					<p class="margin medium-small"><a href="/auth/login">Login</a></p>
				</div>
				<div class="input-field col s6 m6 l6">
					<p class="margin right-align medium-small"><a href="/auth/register">Register</a></p>
				</div>
			</div>


		</form>

	</div>
	
</div>
{{ end }}