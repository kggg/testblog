<!DOCTYPE html>
<html>
  <head><title>TestBlog</title>
                <meta charset="utf-8">
                <meta name="viewport" content="width=device-width, initial-scale=1">      
  <link rel=stylesheet type=text/css href="/static/css/bootstrap.min.css">
  <link rel=stylesheet type=text/css href="/static/css/custom.css">
  <script type="text/javascript" src="/static/js/jquery-3.0.0.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>


 

  </head>

<body>
   <div class="row padding-bottom-200 padding-top-50">
              <div class="col-sm-4 col-sm-offset-3 form-box padding-tb-100 bg-color-white">
                      	<div class="form-top">
                        		<div class="form-top-left">
                        			<h3 class="text-success">请登录</h3>
                            		<p class="text-success">Enter your username and password to log on:</p>
                        		</div>
                        		<div class="form-top-right">
                        			<i class="fa fa-key"></i>
                        		</div>
                        </div>
                        <div class="form-bottom">
		                    <form role="form" action="/login" method="post" class="login-form">
		                    	<div class="form-group">
	                    		<label class="sr-only" for="form-username">Username</label>
 	                        	<input type="text" name="username" placeholder="Username" class="form-username form-control" id="form-username">
		                        </div>
		                        <div class="form-group">
		                        	<label class="sr-only" for="form-password">Password</label>
		                        	<input type="password" name="password" placeholder="Password" class="form-password form-control" id="form-password">
		                        </div>
		                        <button type="submit" class="btn btn-success">登录</button>
		                    </form>
                       </div>
			       <div style="padding-top:15px;">
		                          <p><a href="/register" class="text-success">注册新账户</a></p>
			       </div>
               </div>
   </div>

</body>
</html>

