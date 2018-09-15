<div class="col-sm-3 col-sm-offset-1 blog-sidebar">
  <div class="sidebar-module sidebar-module-inset">
    <h4>About</h4>
    <p>这是学习golang用beego搭建的个人blog</p>
  </div>
  <div class="sidebar-module">
    <h4>标签</h4>
    <ol class="list-unstyled">
      {{range .Tags}}
      <li><a href="/tag/{{.Name}}">{{.Name}}</a></li>              
      {{end}}
    </ol>
  </div>   
  <div class="sidebar-module">
    <h4>分类</h4>
    <ol class="list-unstyled">
      {{range .Categories}}
      <li><a href="/category/{{.Id}}">{{.Name}}</a></li>              
      {{end}}
    </ol>
  </div>          
</div><!-- /.blog-sidebar -->