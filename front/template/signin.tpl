{{define "content"}}
<script src="https://cdn.jsdelivr.net/npm/vue"></script>
<script src="vuelidate/dist/vuelidate.min.js"></script>
<!-- The builtin validators is added by adding the following line. -->
<script src="vuelidate/dist/validators.min.js"></script>

<div id="app">
</div>

<script>
    import Vue from "vue";
    import App from "./App";
    import Vuelidate from "vuelidate";
    import "./assets/styles/app.scss";

    Vue.config.productionTip = false;

    Vue.use(Vuelidate);

    /* eslint-disable no-new */
    new Vue({
        el: "#app",
        components: { App },
        template: "<App/>"
    });
</script>

<section class="uk-margin-large-top uk-width-2-5@l uk-width-4-5 uk-margin-auto uk-margin uk-card uk-card-default uk-card-body uk-text-center">
	<form action="/user-verify" method="post">
		<div class="uk-margin">
			<label for="mail">Email</label><br>
			<div class="uk-width-4-5 uk-inline">
				<span class="uk-form-icon" uk-icon="icon: user"></span>
				<input class="uk-input" type="text" name="email" required>
			</div>
		</div>
		<div class="uk-margin">
			<label for="password">Password</label><br>
			<div class="uk-width-4-5 uk-inline">
				<span class="uk-form-icon" uk-icon="icon: lock"></span>
				<input class="uk-input" type="password" name="password">
			</div>
		</div>
		<div class="button-aria mg-t-20 uk-margin-medium-top">
			<input type="submit" class="uk-button uk-button-primary" value="Login"><br>
			<a href="/signup">Sign up now</a>
		</div>
	</form>
</section>
{{end}}

