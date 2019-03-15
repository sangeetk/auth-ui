{{ template "layout/layout.tpl" . }}

{{ define "content" }}
<div id="login-page" class="row">

	<div class="col s12 z-depth-4 card-panel">

		<form method="POST" class="login-form" name="register" action="/auth/register">
			<input id="step" name="step" type="hidden" ref="step" value="3">
            <input type="hidden" name="gorilla.csrf.Token" value="{{ .CSRFToken }}">
            <input type="hidden" name="token" value="{{ .Token }}">

			<div class="row">
				<div class="input-field col s12 center">
					<h4>Register (cont.)</h4>
				</div>
			</div>

			{{ template "partial/flash.tpl" . }}
			{{ template "partial/errmsg.tpl" . }}

			<div class="row margin">
				<div class="input-field col s12">
					<input id="profession" name="profession" type="text" v-model="profession" ref="profession" required>
					<label for="profession" class="center-align">Profession</label>
				</div>
			</div>

			<div class="row margin">
                <div class="input-field col s12">
                  <textarea id="introduction" name="introduction" class="materialize-textarea" v-model="introduction" ref="introduction" required></textarea>
                  <label for="introduction" class="">Introduce Yourself</label>

                </div>
			</div>

			{{ if getenv "TERMS" }}
			<div class="row margin">
                <div class="input-field col s12">
                  <input type="checkbox" name="acceptTerms" class="filled-in" v-model="acceptTerms" id="acceptTerms" ref="acceptTerms" required/>
                  <label for="acceptTerms">Accept <a href="{{ getenv "TERMS" }}">Terms &amp; Conditions</a></label>
                </div>
			</div>
			{{ end }}

			<br>

			<div class="row">
				<div class="input-field col s12">
					<button class="btn waves-effect waves-light col s12" type="submit">Register</button>
				</div>
				<div class="input-field col s12">
					<p class="margin center medium-small sign-up">Already have an account? <a href="/login">Login</a></p>
				</div>
			</div>
			
		</form>

	</div>
</div>
{{ end }}