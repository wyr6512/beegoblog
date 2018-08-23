{{if .Pager}}
<nav aria-label="Page navigation" class="text-center">
  {{if gt .Pager.PageNums 1}}
  <ul class="pagination">
    {{if .Pager.HasPrev}}    
    <li>
      <a href="{{.Pager.PageLinkPrev}}" aria-label="Previous">
        <span aria-hidden="true">&laquo;</span>
      </a>
    </li>
    {{else}}
    <li class="disabled">
      <a href="#" aria-label="Previous">
        <span aria-hidden="true">&laquo;</span>
      </a>
    </li>
    {{end}}
    {{range $index, $page := .Pager.Pages}}
    <li {{if $.Pager.IsActive $page}}class="active"{{end}}><a href="{{$.Pager.PageLink $page}}">{{$page}}</a></li>
    {{end}}
    {{if .Pager.HasNext}}
    <li>
      <a href="{{.Pager.PageLinkNext}}" aria-label="Next">
        <span aria-hidden="true">&raquo;</span>
      </a>
    </li>
    {{else}}
    <li class="disabled">
      <a href="#" aria-label="Next">
        <span aria-hidden="true">&raquo;</span>
      </a>
    </li>    
    {{end}} 
    <li class="disabled" style="float: right;">
      <a href="#" aria-label="Next">
        <span aria-hidden="true">共{{.Pager.Nums }}条 {{.Pager.Page}}/{{.Pager.PageNums}}页</span>
      </a>
    </li>  
  </ul> 
  {{end}}
</nav>
{{end}}