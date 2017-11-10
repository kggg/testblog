  <div class="content-wrapper">
    <!-- Content Header (Page header) -->
    <section class="content-header">
      <h1>
        文章管理
        <small>Preview</small>
      </h1>
      <ol class="breadcrumb">
        <li><a href="/admin"><i class="fa fa-dashboard"></i>首页</a></li>
        <li><a href="#">文章管理</a></li>
        <li class="active">文章列表</li>
      </ol>
    </section>

    <!-- Main content -->
    <section class="content">
      <div class="row">
        <div class="col-xs-12">
          <div class="box box-info">
            <div class="box-header with-border">
              <a href="/admin/article" class="btn btn-default"><i class="fa fa-list"></i>  文章列表</a>
              <a href="/admin/article/add" class="btn btn-info"><i class="fa fa-plus"></i>  </a>
            </div>            
            <!-- /.box-header -->
            <div class="box-body">
              <table id="article1" class="table table-bordered table-striped table-hover">
                <thead>
                <tr>
                  <th>ID</th>
                  <th>主题</th>
                  <th>板块</th>
                  <th>标签</th>
                  <th>操作</th>
                </tr>
                </thead>
                <tbody>
                {{ range .Blog }}
                <tr>
                  <td>{{ .Id }}</td>
                  <td>{{ .Title }}</td>
                  <td>{{ .Section.Section }}</td>
                  <td><span class="label label-info">{{ .Label.Label }}</span></td>
                  <td><a herf="" class="btn btn-info btn-xs"><i class="fa fa-edit">  </i></a> 
                      <form action="/admin/article/delete/{{ .Id }}" method="post">
                      {{ $.xsrfdata }}
                      <button type="submit"  class="btn btn-danger btn-run btn-xs"><i class="fa fa-trash">  </i></button></form>
                  </td>
                </tr>
                 {{ end }}
                </tbody>
              </table>
            </div>
            <!-- /.box-body -->
          </div>
          <!-- /.box -->
        </div>
      </div>


          </div>
       </div>
     </div>
   </section>

  </div>

