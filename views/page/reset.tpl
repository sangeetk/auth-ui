{{ template "layout/layout.tpl" . }}

{{ define "content" }}
<div id="login-page" class="row">

	<div class="col s12 z-depth-4 card-panel">
	
		<form class="login-form" action="/auth/login" method="POST">
            <input type="hidden" name="gorilla.csrf.Token" value="{{ .CSRFToken }}">

			<div class="row">
				<div class="input-field col s12 center">
					<a href="/"><img src="{{ logo }}" alt="" class="" height="100px" width="auto"></a>
					<p class="center login-form-text">{{ siteName }}</p>
				</div>
			</div>

			{{ template "partial/flash.tpl" . }}
			{{ template "partial/errmsg.tpl" . }}
			
			<div class="row margin">
				<div class="input-field col s12">
					<i class="mdi-action-lock-outline prefix"></i>
					<input id="password" name="password" type="password" ref="password" required>
					<label for="password">New Password</label>
				</div>
			</div>

			<div class="row margin">
				<div class="input-field col s12">
					<i class="mdi-action-lock-outline prefix"></i>
					<input id="password2" name="password2" type="password" ref="password2" required>
					<label for="password2">New Password again</label>
				</div>
			</div>

			<div class="row">
				<div class="input-field col s12">
					<!-- a href="/user/login" class="btn waves-effect waves-light col s12">Login</a -->
						<button class="btn waves-effect waves-light col s12" type="submit">Reset Password</button>
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