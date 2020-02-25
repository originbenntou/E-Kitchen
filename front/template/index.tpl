{{define "content"}}
<!-- index start -->
<table class="uk-table uk-table-hover uk-table-divider">
	<thead>
	<tr>
		<th>Id</th>
		<th>Name</th>
		<th>Status</th>
		<th>CategoryId</th>
		<th>UserId</th>
		<th>Created_At</th>
		<th>Updated_At</th>
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
			<a class="uk-button uk-button-default" href="#modal-edit-{{$v.Id}}" uk-toggle>Edit</a>
			<a class="uk-button uk-button-default" href="#modal-delete-{{$v.Id}}" uk-toggle>Delete</a>
		</td>
	</tr>
	{{end}}
	</tbody>
</table>
<!-- index end -->

<!-- Modal Edit start -->
{{range $v := .Shops}}
<div id="modal-edit-{{$v.Id}}" uk-modal>
	<div class="uk-modal-dialog">
		<button class="uk-modal-close-default" type="button" uk-close></button>
		<form class="uk-form-stacked" action="/edit-shop" method="post">
			<div class="uk-modal-header">
				<h2 class="uk-modal-title">Edit {{$v.Name}}</h2>
			</div>
			<input class="uk-input" id="form-stacked-text-{{$v.Id}}" type="hidden" name="Id" value="{{$v.Id}}">
			<div class="uk-modal-body" uk-overflow-auto>
				<div class="uk-margin">
					<label class="uk-form-label" for="form-stacked-text-{{$v.Id}}">Name</label>
					<div class="uk-form-controls">
						<input class="uk-input" id="form-stacked-text-{{$v.Id}}" type="text" name="Name" value="{{$v.Name}}">
					</div>
				</div>
				<div class="uk-margin">
					<div class="uk-form-label">Status</div>
					<div class="uk-form-controls">
						<label><input class="uk-radio-0" type="radio" name="Status" value="0" {{checked $v.Status "PUBLIC"}}>公開</label>
						<br>
						<label><input class="uk-radio-1" type="radio" name="Status" value="1" {{checked $v.Status "PRIVATE"}}>非公開</label>
					</div>
				</div>
				<div class="uk-margin">
					<label class="uk-form-label" for="form-stacked-select-{{$v.Id}}">CategoryId</label>
					<div class="uk-form-controls">
						<select class="uk-select" id="form-stacked-select-{{$v.Id}}">
							<option>1</option>
							<option>2</option>
							<option>3</option>
							<option>4</option>
							<option>5</option>
						</select>
					</div>
				</div>
			</div>
			<div class="uk-modal-footer uk-text-right">
				<input class="uk-button uk-button-default uk-modal-close" type="button" value="Cancel">
				<input class="uk-button uk-button-primary" type="submit" value="Update">
			</div>
		</form>
	</div>
</div>
{{end}}
<!-- Modal Edit end -->

<!-- Modal Delete start -->
{{range $v := .Shops}}
<div id="modal-delete-{{$v.Id}}" uk-modal>
	<div class="uk-modal-dialog">
		<button class="uk-modal-close-default" type="button" uk-close></button>
		<form class="uk-form-stacked" action="/delete-shop" method="post">
			<div class="uk-modal-header">
				<h2 class="uk-modal-title">Delete {{$v.Name}}</h2>
			</div>
			<input class="uk-input" id="form-stacked-text-{{$v.Id}}" type="hidden" name="Id" value="{{$v.Id}}">
			<div class="uk-modal-body" uk-overflow-auto>
				<div class="uk-margin">
					削除しますか？
				</div>
			<div class="uk-modal-footer uk-text-right">
				<input class="uk-button uk-button-default uk-modal-close" type="button" value="Cancel">
				<input class="uk-button uk-button-primary" type="submit" value="Delete">
			</div>
		</form>
	</div>
</div>
{{end}}
<!-- Modal Delete end -->
{{end}}
