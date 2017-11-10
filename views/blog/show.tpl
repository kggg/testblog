<div class="row">
  {{template  "blog/left.tpl" .}}

  <div class="col-md-9 col-sm-12 col-xs-12">
           <div class="alert alert-danger" role="alert">
                非淡泊无以明志，非宁静无以致远。夫学须静也，才须学也，非学无以广才，非志无以成学。淫慢不能励精，险躁则不能冶性。年与时驰，意与日去，遂
成枯落，多不接世，悲守穷庐，将复何及！
           </div>
            <div class='panel panel-info'>
                <div class='panel-heading'>
                    <span class="lead panel-title">{{ .Blog.Title }}<span>
                </div>
                <div class="panel-body">
                     {{ str2html .Blog.Content }}               
                </div>

            </div>
  </div>
</div>

