{{define "content"}}
<div id="index">
	<!-- list start -->
	<table class="uk-table uk-table-hover uk-table-divider">
		<thead>
		<tr>
			<th>Id</th>
			<th>Name</th>
			<th>Star</th>
			<th>Tag</th>
			<th>Status</th>
			<th>UserId</th>
			<th>Created_At</th>
			<th>Updated_At</th>
			<th class="uk-padding-remove-bottom uk-padding-small uk-inline uk-light">
				<a href="#modal-create" uk-toggle uk-marker>
					<span class="uk-margin-small-bottom uk-icon uk-margin-small-right"></span>
					Create
				</a>
			</th>
		</tr>
		</thead>
		<tbody>
		{{range $v := .Shops}}
		<tr class="{{trClass $v.Status}}">
			<td>{{$v.Id}}</td>
			<td><a href="{{$v.Url}}">{{$v.Name}}</a></td>
			<td>★★★★☆</td>
			<td>
				{{(index $.Tags $v.Id)}}
			</td>
			<td>{{$v.Status}}</td>
			<td>{{$v.UserId}}</td>
			<td>{{convertTime $v.CreatedAt}}</td>
			<td>{{convertTime $v.UpdatedAt}}</td>
			<td>
				<a class="uk-button uk-button-default" href="#modal-edit-{{$v.Id}}" uk-toggle>Edit</a>
				<a class="uk-button uk-button-default" href="#modal-delete-{{$v.Id}}" uk-toggle>Delete</a>
			</td>
		</tr>
		{{end}}
		</tbody>
	</table>
	<!-- list end -->

	<!-- Modal Create start -->
	<div id="tag-create">
		<div id="modal-create" uk-modal>
			<div class="uk-modal-dialog">
				<button class="uk-modal-close-default" type="button" uk-close></button>
				<form class="uk-form-stacked" action="/create-shop" method="post" @submit="submitTag">
					<div class="uk-modal-header">
						<h2 class="uk-modal-title">Create</h2>
					</div>
					<div class="uk-modal-body" uk-overflow-auto>
						<div class="uk-margin">
							<label class="uk-form-label" for="form-stacked-text">Name</label>
							<div class="uk-form-controls">
								<input class="uk-input" id="form-stacked-text" type="text" name="Name" value="">
							</div>
						</div>
						<div class="uk-margin">
							<label class="uk-form-label" for="form-stacked-text">Url</label>
							<div class="uk-form-controls">
								<input class="uk-input" id="form-stacked-text" type="text" name="Url" value="">
							</div>
						</div>
						<div class="uk-margin">
							<label class="uk-form-label" for="form-stacked-text">Tag</label>
							<div class="uk-form-controls">
								<div class="tag-container cf">
									<div v-for="(tag, index) in tags" :key="index" class="tag-label">
										<span class="tag-label-text">${tag}</span>
										<a href="#" class="tag-remove" @click.stop.prevent="remove(index)">
											<svg xmlns="http://www.w3.org/2000/svg"
												 xmlns:xlink="http://www.w3.org/1999/xlink"
												 viewBox="0 0 50 50" version="1.1" width="15px" height="15px">
												<g id="surface1">
													<path style=" " d="M 7.71875 6.28125 L 6.28125 7.71875 L 23.5625 25 L 6.28125 42.28125 L 7.71875 43.71875 L 25 26.4375 L 42.28125 43.71875 L 43.71875 42.28125 L 26.4375 25 L 43.71875 7.71875 L 42.28125 6.28125 L 25 23.5625 Z "/>
												</g>
											</svg>
										</a>
									</div>
								</div>
								<div class="editor">
									<textarea id="form-stacked-text" class="uk-textarea" row="1" placeholder="add tags..." @keyup.enter="enter($event.target)" @keypress="canEnter = true"></textarea>
								</div>
							</div>
						</div>
						<div class="uk-margin">
							<div class="uk-form-label">Status</div>
							<div class="uk-form-controls">
								<label><input class="uk-radio-0" type="radio" name="Status" value="0">公開</label>
								<br>
								<label><input class="uk-radio-1" type="radio" name="Status" value="1">非公開</label>
							</div>
						</div>
					</div>
					<div class="uk-modal-footer uk-text-right">
						<button class="uk-button uk-button-default uk-modal-close">Cancel</button>
						<button class="uk-button uk-button-primary">Create</button>
					</div>
				</form>
			</div>
		</div>
	</div>
	<!-- Modal Create end -->

	<!-- Modal Edit start -->
	{{range $v := .Shops}}
	<div id="modal-edit-{{$v.Id}}" uk-modal>
		<div class="uk-modal-dialog">
			<button class="uk-modal-close-default" type="button" uk-close></button>
			<form class="uk-form-stacked" action="/edit-shop" method="post">
				<div class="uk-modal-header">
					<h2 class="uk-modal-title">Edit {{$v.Name}}</h2>
				</div>
				<input class="uk-input" type="hidden" name="Id" value="{{$v.Id}}">
				<div class="uk-modal-body" uk-overflow-auto>
					<div class="uk-margin">
						<label class="uk-form-label" for="form-stacked-text-{{$v.Id}}">Name</label>
						<div class="uk-form-controls">
							<input class="uk-input" id="form-stacked-text-{{$v.Id}}" type="text" name="Name" value="{{$v.Name}}">
						</div>
					</div>
					<div class="uk-margin">
						<label class="uk-form-label" for="form-stacked-text-{{$v.Id}}">Url</label>
						<div class="uk-form-controls">
							<input class="uk-input" id="form-stacked-text-{{$v.Id}}" type="text" name="Url" value="{{$v.Url}}">
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
				<input class="uk-input" type="hidden" name="Id" value="{{$v.Id}}">
				<div class="uk-modal-body" uk-overflow-auto>
					<div class="uk-margin">
						<label class="uk-form-label" for="form-stacked-text-{{$v.Id}}">Name</label>
						<div class="uk-form-controls">
							<p>{{$v.Name}}</p>
						</div>
					</div>
					<div class="uk-margin">
						<label class="uk-form-label" for="form-stacked-text-{{$v.Id}}">Url</label>
						<div class="uk-form-controls">
							<p>{{$v.Url}}</p>
						</div>
					</div>
					<div class="uk-margin">
						<div class="uk-form-label">Status</div>
						<div class="uk-form-controls">
							<p>{{$v.Status}}</p>
						</div>
					</div>
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
</div>

<script src="./static/tag-input.js" type="text/javascript"></script>
{{end}}

