  <div class="content-wrapper">
    <!-- Content Header (Page header) -->
    <section class="content-header">
      <h1>
        文章管理
        <small>Preview</small>
      </h1>
      <ol class="breadcrumb">
        <li><a href="/admin"><i class="fa fa-dashboard"></i>首页</a></li>
        <li><a href="/admin/section">板块管理</a></li>
        <li class="active">编辑板块</li>
      </ol>
    </section>

    <!-- Main content -->
    <section class="content">
      <div class="row">
        <div class="col-xs-12">
          <div class="box box-info">
            <div class="box-header with-border">
               <ul class="nav nav-tabs">
                 <li><a href="/admin/section"><i class="fa fa-list"></i> 板块列表</a></li>
                 <li><a href="/admin/section/add"><i class="fa fa-plus"></i> 新增板块</a></li>
                 <li class="active"><a href="/admin/section/edit/{{ .Sections.Id }}"><i class="fa fa-edit"></i> 编辑板块</a></li>
               </ul>
            </div>
            <!-- /.box-header -->
            <div class="box-body pad">
              <form action="/admin/section/edit/{{ .Sections.Id }}" method="post">
                    <div class="form-group">
                       <label>板块名称</label>
                       <input class="form-control" name="section" type="text" value="{{ .Sections.Section }}">
                    </div>

                    <div class="form-group">
                       {{ .xsrfdata }}
                       <button type="submit" class="btn btn-info">保存</button>
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

