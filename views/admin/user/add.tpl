  <div class="content-wrapper">
    <!-- Content Header (Page header) -->
    <section class="content-header">
      <h1>
        用户管理
        <small>Preview</small>
      </h1>
      <ol class="breadcrumb">
        <li><a href="/admin"><i class="fa fa-dashboard"></i>首页</a></li>
        <li><a href="/admin/user">用户管理</a></li>
        <li class="active">用户列表</li>
      </ol>
    </section>

    <!-- Main content -->
    <section class="content">
      <div class="row">
        <div class="col-xs-12">
          <div class="box box-info">
            <div class="box-header with-border">
               <ul class="nav nav-tabs">
                 <li><a href="/admin/user"><i class="fa fa-list"></i> 用户列表</a></li>
                 <li class="active"><a href="/admin/user/add"><i class="fa fa-plus"></i> 新增用户</a></li>
               </ul>
            </div>

            <!-- /.box-header -->
            <div class="box-body pad">
              <form action="/admin/user/add" method="post" class="form-horizontal">
                    <div class="form-group">
                       <label for="username" class="col-sm-2 control-label">用户名</label>
                       <div class="col-sm-10">
                         <input type="text" class="form-control" id="username"  name="username" placeholder="Username">
                       </div>
                    </div>
                    <div class="form-group">
                       <label for="password" class="col-sm-2 control-label">用户密码</label>
                       <div class="col-sm-10">
                       <input class="form-control" name="password" type="password">
                       </div>
                    </div>

                    <div class="form-group">
                       <label for="password" class="col-sm-2 control-label">用户密码</label>
                       <div class="col-sm-10">
                       <input class="form-control" name="password" type="password">
                       </div>
                    </div>

                    <div class="form-group">
                       {{ .xsrfdata }}
                    </div>
                    <div class="box-footer">
                      <button type="reset" class="btn btn-default">Cancel</button>
                      <button type="submit" class="btn btn-info pull-right">保存</button>
                    </div>
              </form>
            </div>
          </div>
          <!-- /.box -->
        </div>
      </div>


          </div>
       </div>
     </div>
   </section>

  </div>

