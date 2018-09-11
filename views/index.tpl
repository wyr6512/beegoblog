<!DOCTYPE html>
<html lang="zh-CN">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <!-- 上述3个meta标签*必须*放在最前面，任何其他内容都*必须*跟随其后！ -->
    <meta name="description" content="">
    <meta name="author" content="">

    <title>beego blog</title>

    <!-- Bootstrap core CSS -->
    <link href="https://cdn.bootcss.com/bootstrap/3.3.7/css/bootstrap.min.css" rel="stylesheet">

    <!-- Custom styles for this template -->
    <link href="/static/css/blog.css" rel="stylesheet">
    
  </head>

  <body>

    <div class="blog-masthead">
      <div class="container">
        <nav class="blog-nav">
          <a class="blog-nav-item active" href="/">Home</a>          
        </nav>
      </div>
    </div>

    <div class="container">

      <div class="blog-header">
      </div>

      <div class="row">

        <div class="col-sm-8 blog-main">
          {{range .Articles}}
          <div class="blog-post">
            <h2 class="blog-post-title"><a href="/article/{{.Id}}">{{.Title}}</a></h2>
            <p class="blog-post-meta">{{.CreatedAt | formatTime}} by <a href="https://github.com/wyr6512/beegoblog">wyr6512</a></p>
            <p>{{.Abstract}}</p>
          </div><!-- /.blog-post -->
          {{end}}
          {{if .Pager}}
          <nav>
            {{if gt .Pager.PageNums 1}}
            <ul class="pager">
              {{if .Pager.HasPrev}}  
              <li><a href="{{.Pager.PageLinkPrev}}">Previous</a></li>
              {{else}}
              <li class="disabled"><a href="#">Previous</a></li>
              {{end}}
              {{if .Pager.HasNext}}
              <li><a href="{{.Pager.PageLinkNext}}">Next</a></li>
              {{else}}
              <li class="disabled"><a href="#">Next</a></li>
              {{end}}
            </ul>
            {{end}}
          </nav>
          {{end}}

        </div><!-- /.blog-main -->

        <div class="col-sm-3 col-sm-offset-1 blog-sidebar">
          <div class="sidebar-module sidebar-module-inset">
            <h4>About</h4>
            <p>这是学习golang用beego搭建的个人blog</p>
          </div>
          <div class="sidebar-module">
            <h4>标签</h4>
            <ol class="list-unstyled">
              {{range .Tags}}
              <li><a href="/tag/{{.Id}}">{{.Name}}</a></li>              
              {{end}}
            </ol>
          </div>          
        </div><!-- /.blog-sidebar -->

      </div><!-- /.row -->

    </div><!-- /.container -->

    <footer class="blog-footer">
      <p>Blog built by <a href="https://github.com/wyr6512">@wyr6512</a>.</p>
      <p>
        <a href="#">Back to top</a>
      </p>
    </footer>


  </body>
</html>
