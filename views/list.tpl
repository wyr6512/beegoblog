{{range .Articles}}
<div class="blog-post">
	<h2 class="blog-post-title"><a href="/article/{{.Id}}">{{.Title}}</a></h2>
	<p class="blog-post-meta">{{.CreatedAt | formatTime}} by <a href="https://github.com/wyr6512/beegoblog">wyr6512</a></p>
	<p>{{.Abstract}}</p>
</div><!-- /.blog-post -->
{{end}}