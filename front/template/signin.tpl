{{define "content"}}

<!-- Vue-cdn -->
<script src="https://cdn.jsdelivr.net/npm/vue@2.6.11"></script>

<!-- VeeValidate-cdn -->
<script src="https://cdn.jsdelivr.net/npm/vee-validate@3.2.3/dist/vee-validate.js"></script>
<script src="https://cdn.jsdelivr.net/npm/vee-validate@3.2.3/dist/rules.umd.min.js"></script>

<div id="app">
	<section class="uk-margin-large-top uk-width-2-5@l uk-width-4-5 uk-margin-auto uk-margin uk-card uk-card-default uk-card-body uk-text-center">
		<form action="/user-verify" method="post">
			<fieldset class="uk-fieldset">
				<legend class="uk-legend uk-margin-bottom">サインイン</legend>
				<validation-observer ref="obs" v-slot="ObserverProps">
					<div class="uk-margin">
						<label for="mail">Email</label><br>
						<validation-provider name="email" rules="required|email">
							<div slot-scope="ProviderProps">
								<div class="uk-width-4-5 uk-inline">
									<div>
										<span class="uk-form-icon" uk-icon="icon: user"></span>
										<input v-model="email" class="uk-input" type="text" name="email" required>
									</div>
									<p v-show="ProviderProps.errors[0]" class="uk-position-absolute uk-alert-warning uk-align-center uk-width-1-1 uk-margin-remove-top">${ ProviderProps.errors[0] }</p>
								</div>
							</div>
						</validation-provider>
					</div>
					<div class="uk-margin uk-margin-medium-top">
						<label for="password">Password</label><br>
						<validation-provider name="password" rules="required|password">
							<div slot-scope="ProviderProps">
								<div class="uk-width-4-5 uk-inline">
									<div>
										<span class="uk-form-icon" uk-icon="icon: lock"></span>
										<input v-model="password" class="uk-input" type="password" name="password" required>
									</div>
									<p v-show="ProviderProps.errors[0]" class="uk-position-absolute uk-alert-warning uk-align-center uk-width-1-1 uk-margin-remove-top">${ ProviderProps.errors[0] }</p>
								</div>
							</div>
						</validation-provider>
					</div>
					<div class="button-aria mg-t-20 uk-margin-medium-top">
						<input type="submit" @click="submit" :disabled="ObserverProps.invalid || !ObserverProps.validated" class="uk-button uk-button-primary" value="Sign in"><br>
						<a href="/signup">Sign up now</a>
					</div>
				</validation-observer>
			</fieldset>
		</form>
	</section>
</div>

<script src="./static/index.js" type="text/javascript"></script>
{{end}}

