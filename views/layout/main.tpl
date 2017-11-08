<!DOCTYPE html>
<html>
  <head><title>My Blog</title>
                <meta charset="utf-8">
                <meta name="viewport" content="width=device-width, initial-scale=1">	  
  <link rel=stylesheet type=text/css href="/static/css/bootstrap.min.css">
  <link rel=stylesheet type=text/css href="/static/css/custom.css">
  <script type="text/javascript" src="/static/js/jquery-3.0.0.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <script type="text/javascript" src="/static/js/menu.js"></script>


 

  </head>
  <nav class="navbar navbar-default margin-auto navbar-fixed-top">
      <div class="container-fluid">
            <div class="navbar-header">
                 <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#bs-example-navbar-collapse-1" aria-expanded="false">
                    <span class="sr-only">Toggle navigation</span>
                     <span class="icon-bar"></span>
                    <span class="icon-bar"></span>
                    <span class="icon-bar"></span>
                    </button>		    
		    <a href='/' class="navbar-brand active"><span class="text-primary">我的博客</span></a>
	    </div>	    

    <!-- Collect the nav links, forms, and other content for toggling -->
    <div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
      <ul class="nav navbar-nav">
        <li class="active"><a href="/">首页<span class="sr-only">(current)</span></a></li>
	<li><a href="/about">关于我</a></li>
      </ul>

      <ul class="nav navbar-nav navbar-right">
        <li><a href="/admin"><i class="glyphicon glyphicon-cog"></i>&nbsp; 后台管理</a></li>
      </ul>
    </div><!-- /.navbar-collapse -->
  </div><!-- /.container-fluid -->
  </nav>
  <body>
      <div class="container margin-top-5">	  
          {{.LayoutContent}}
      </div>
      <div class='col-sm-12 text-info bg-color-white padding-tb-15 text-center'>Steven Lu</div> 
  </body>    
</html>

