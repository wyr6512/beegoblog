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