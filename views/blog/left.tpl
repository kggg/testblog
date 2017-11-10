  
  <div class="col-sm-3">
     <div class="margin-bottom-20">
     <form id="searchform" method="get" action="/search">
        <div class="input-group">
            <span class="input-group-addon"><span class="glyphicon glyphicon-search"></span></span>
                <input type="search" class="form-control" placeholder="输入关键词" name="search" id="s">
            <span class="input-group-btn">
                <button class="btn btn-info" type="submit" value="Search">搜索</button>
            </span>
        </div>
      </form>
      </div>


      <div class='panel panel-info'>
         <div class="panel-heading">
            <i class="glyphicon glyphicon-menu-hamburger"></i> <strong>分类</strong>
         </div>
         <div class="panel-body">
             <ul class="list-group list-unstyled">
             {{range .section}}
               <li class="list-group-item"><a href="/section/{{ .Id }}">{{ .Section }}</a><span class="badge">{{ .Count }}</span></li>
             {{ end }}

         </div>
      </div>

      <div class='panel panel-info'>
         <div class="panel-heading ">
             <i class="glyphicon glyphicon-menu-hamburger"></i> <strong>标签</strong>
         </div>
         <div class="panel-body">
            <ul class='list-group list-inline'>
                {{ range .label }}
                 <li><a href='/label/{{ .Id }}'>{{ .Label }}</a></li>
                {{ end }}
            </ul>
         </div>
      </div>

      <div class='panel panel-info'>
         <div class="panel-heading ">
             <i class="glyphicon glyphicon-menu-hamburger"></i> <strong>热门文章</strong>
         </div>
         <ul class='list-group'>
         </ul>
      </div>      

  </div>

