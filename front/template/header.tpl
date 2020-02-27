{{define "header"}}
<h1 class="uk-text-muted uk-padding uk-text-uppercase uk-navbar-container tm-navbar-container uk-sticky">
	<span class="uk-text-background">E-KITCHEN</span>
</h1>
{{if .User}}
<div class="uk-padding-small uk-position-small uk-position-top-right">
	<p class="uk-text-emphasis uk-text-small uk-margin-remove-bottom">{{.User}}</p>
	<p class="uk-text-emphasis uk-text-small uk-align-right"><a href="/signout"><span uk-icon="sign-out"></span>Signout</a></p>
</div>
{{end}}
{{end}}
