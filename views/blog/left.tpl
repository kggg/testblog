  <div class="col-sm-3">
      <div class='panel panel-primary'>
         <div class="panel-heading">
            <i class="glyphicon glyphicon-menu-hamburger"></i> <strong>分类</strong>
         </div>
         <div class="panel-body">
             <ul class="list-group">
             {{range .section}}
               <li>{{ .Section }}</li>
             {{ end }}
            
         </div>
      </div>


      <div class='panel panel-primary'>
         <div class="panel-heading ">
             <i class="glyphicon glyphicon-menu-hamburger"></i> <strong>标签</strong>
         </div>
         <div class="panel-body">
            <ul class='list-group list-inline'>
                {{ range .label }}
                 <li><a href='/categoryshow/{{ .Id }}'>{{ .Label }}</a></li>
                {{ end }}
            </ul>
         </div>
      </div>

      <div class='panel panel-primary'>
         <div class="panel-heading ">
             <i class="glyphicon glyphicon-menu-hamburger"></i> <strong>热门文章</strong>
         </div>
         <ul class='list-group'>
         </ul>
      </div>      

  </div>

