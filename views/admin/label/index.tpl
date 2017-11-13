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
        <li class="active">标签列表</li>
      </ol>
    </section>

    <!-- Main content -->
    <section class="content">
      <div class="row">
        <div class="col-xs-12">
          <div class="box box-info">
            <div class="box-header with-border">
               <ul class="nav nav-tabs">
                 <li class="active"><a href="/admin/label"><i class="fa fa-list"></i> 标签列表</a></li>
                 <li><a href="/admin/label/add"><i class="fa fa-plus"></i> 新增标签</a></li>
               </ul>
            </div>            
            <!-- /.box-header -->
            <div class="box-body">
              <table id="article1" class="table table-bordered table-striped table-hover">
                <thead>
                <tr>
                  <th>ID</th>
                  <th>板块</th>
                  <th>标签</th>
                  <th>操作</th>
                </tr>
                </thead>
                <tbody>
                {{ range .Label }}
                <tr>
                  <td>{{ .Id }}</td>
                  <td>{{ .Section.Section }}</td>
                  <td><span class="label label-info">{{ .Label }}</span></td>
                  <td><a href="/admin/label/edit/{{ .Id }}" class="btn btn-info btn-xs"><i class="fa fa-edit"> 编辑</i></a> 
                       <a href="#myModal{{ .Id }}" class="btn btn-xs btn-danger" data-toggle="modal"><i class="fa fa-trash"></i> 删除</a>
                  <div class="modal fade" id="myModal{{ .Id }}" tabindex="-1" role="dialog" aria-labelledby="myModalLabel">
                     <div class="modal-dialog modal-sm" role="document">
                       <div class="modal-content">
                        <div class="modal-body">
                            <div class="form-group">
                              确定删除吗?
                            </div>
                        </div>
                       <div class="modal-footer">
                         <form method="post" action="/admin/label/delete/{{ .Id }}">
                           <button type="button" class="btn btn-default" data-dismiss="modal"> 取消</button>
                           {{ $.xsrfdata }}
                           <button type="submit" class="btn btn-primary">删除</button>
                         </form>
                       </div>
                    </div>
                 </div>
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

