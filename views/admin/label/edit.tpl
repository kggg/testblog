  <div class="content-wrapper">
    <!-- Content Header (Page header) -->
    <section class="content-header">
      <h1>
        文章管理
        <small>Preview</small>
      </h1>
      <ol class="breadcrumb">
        <li><a href="/admin"><i class="fa fa-dashboard"></i>首页</a></li>
        <li><a href="/admin/label">标签管理</a></li>
        <li class="active">编辑标签</li>
      </ol>
    </section>

    <!-- Main content -->
    <section class="content">
      <div class="row">
        <div class="col-xs-12">
          <div class="box box-info">
            <div class="box-header with-border">
               <ul class="nav nav-tabs">
                 <li><a href="/admin/label"><i class="fa fa-list"></i> 标签列表</a></li>
                 <li><a href="/admin/label/add"><i class="fa fa-plus"></i> 新增标签</a></li>
                 <li class="active"><a href="/admin/label/edit/{{ .Label.Id }}"><i class="fa fa-edit"></i> 编辑标签</a></li>
               </ul>
            </div>
            <!-- /.box-header -->
            <div class="box-body pad">
              <form action="/admin/label/edit/{{ .Label.Id }}" method="post">
                    <div class="form-group">
                       <label for="label">标签</label>
                       <input type="text" class="form-control" id="label" name="label" value="{{ .Label.Label }}">
                    </div>
                    <div class="form-group">
                       <label>板块</label>
                       <select class="form-control" name="section">
                       <option value="{{ .Label.Section.Id }}">{{ .Label.Section.Section }}</option>
                       {{ range .Sections }}
                       <option value="{{ .Id }}">{{ .Section }}</option>
                       {{ end }}
                       </select>
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

