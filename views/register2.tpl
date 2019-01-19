{{ template "layout.tpl" . }}

{{ define "content" }}
<div id="login-page" class="row">

	<div class="col s12 z-depth-4 card-panel">

		<form method="POST" class="login-form" name="register" action="/register">
			<input id="step" name="step" type="hidden" ref="step" value="2">
            <input type="hidden" name="gorilla.csrf.Token" value="{{ .CSRFToken }}">
            <input type="hidden" name="token" value="{{ .Token }}">

			<div class="row">
				<div class="input-field col s12 center">
					<h4>Register (cont.)</h4>
				</div>
			</div>

			<div class="row margin">
				<div class="input-field col s12">
					<i class="mdi-social-cake prefix"></i>
					<input id="birthday" name="birthday" type="date" class="datepicker" ref="birthday">
					<label for="birthday" class="center-align">Date of Birth</label>
				</div>
			</div>

			<div class="row margin">
				<div class="input-field col s12">
					<input id="address" name="address" type="text" ref="address" required>
					<label for="address" class="center-align">Address</label>
				</div>
			</div>

			<div class="row margin">
				<div class="input-field col s12">
					<input id="city" name="city" type="text" ref="city" required>
					<label for="city" class="center-align">City</label>
				</div>
			</div>

			<div class="row margin">
				<div class="input-field col s6">
					<input id="state" name="state" type="text" ref="state" required>
					<label for="state" class="center-align">State</label>
				</div>
				<div class="input-field col s6">
					<input id="country" name="country" type="text" ref="country" required>
					<label for="country" class="center-align">Country</label>
				</div>
			</div>
			
			<div class="row">
				<div class="input-field col s12">
					<button class="btn waves-effect waves-light col s12" type="submit">Next</button>
				</div>
				<div class="input-field col s12">
					<p class="margin center medium-small sign-up">Already have an account? <a href="/login">Login</a></p>
				</div>
			</div>
			
		</form>

	</div>
</div>
{{ end }}