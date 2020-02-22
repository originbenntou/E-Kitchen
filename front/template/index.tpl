{{define "content"}}
<!-- index start -->
<table class="uk-table uk-table-hover uk-table-divider">
	<thead>
	<tr>
		<th>Id</th>
		<th>Name</th>
		<th>Satus</th>
		<th>CategoryId</th>
		<th>UserId</th>
		<th>CreatedAt</th>
		<th>UpdatedAt</th>
		<th></th>
	</tr>
	</thead>
	<tbody>
	{{range $v := .Shops}}
	<tr>
		<td>{{$v.Id}}</td>
		<td>{{$v.Name}}</td>
		<td>{{$v.Status}}</td>
		<td>{{$v.CategoryId}}</td>
		<td>{{$v.UserId}}</td>
		<td>{{$v.CreatedAt}}</td>
		<td>{{$v.UpdatedAt}}</td>
		<td>
			<button type="button" uk-toggle="target: #modal-overflow; cls: target-{{$v.Id}}">Edit</button>
			<a class="uk-button uk-button-default" href="#modal-overflow" uk-toggle>Edit</a>
			<a class="uk-button uk-button-default" href="#modal-overflow" uk-toggle>Delete</a>
		</td>
	</tr>
	{{end}}
	</tbody>
</table>
<!-- index end -->

<!-- Modal Edit start -->
<div id="modal-overflow" uk-modal>
	<div class="uk-modal-dialog">
		<button class="uk-modal-close-default" type="button" uk-close></button>
		<div class="uk-modal-header">
			<h2 class="uk-modal-title">Headline</h2>
		</div>
		<div class="uk-modal-body" uk-overflow-auto>
			<p>あああああああああ</p>
		</div>
		<div class="uk-modal-footer uk-text-right">
			<button class="uk-button uk-button-default uk-modal-close" type="button">Cancel</button>
			<button class="uk-button uk-button-primary" type="button">Save</button>
		</div>
	</div>
</div>
<!-- Modal Edit end -->

<!-- Modal Delete start -->

<!-- Modal Delete end -->
{{end}}
