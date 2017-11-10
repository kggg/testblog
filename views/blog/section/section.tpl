<div class="row">
  {{template  "blog/left.tpl" .}}

  <div class="col-md-9 col-xs-12 col-sm-12">
           <div class="alert alert-danger" role="alert">
                非淡泊无以明志，非宁静无以致远。夫学须静也，才须学也，非学无以广才，非志无以成学。淫慢不能励精，险躁则不能冶性。年与时驰，意与日去，遂成枯落，多不接世，悲守穷庐，将复何及！ 
           </div>
           {{ range .Blog }}
            <div class='panel panel-info'>
                <div class='panel-heading'>
                    <a href="/show/{{ .Id }}"> <span class="lead">{{ .Title }}<span></a>
                </div>
                <div class="panel-body">
                  <div class="row">
                    <div class="col-md-9">
                     {{$var := html2str .Content  }}{{ $var2 := htmlunquote $var }}{{ substr $var2 0 200 }}...
                     <p>...</p>
                     <p><a href="/show/{{ .Id }}" class="btn btn-primary">查看全文</a></p> 
                     <p><span><small>发表于: {{ .Created_at }} 文章标签: {{ .Label.Label }}</small></span></p>
                     
                    </div>
                    <div class="col-md-3">
  		       <a href="/show/{{ .Id }}" class="thumbnail">
                         <img src="/static/images/article/main/{{randpicnum }}.jpg"  alt="{{.Title }}" width="160px" height="120px">
		       </a>
                    </div>
                  </div>
                </div>
            </div>
           {{ end }}
      {{if gt .Pagetotal 1 }}
         {{template "blog/section/page.tpl" .}}
      {{ end }}

  </div>
</div>


