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
        <li class="active">新增文章</li>
      </ol>
    </section>

    <!-- Main content -->
    <section class="content">
      <div class="row">
        <div class="col-xs-12">
          <div class="box box-info">
            <div class="box-header with-border">
              <label class="btn btn-default">新增文章</label>
              <a href="/admin/article" class="btn btn-info"><i class="fa fa-list"></i> </a>
              <!-- tools box -->
              <div class="pull-right box-tools">
                <button type="button" class="btn btn-info btn-sm" data-widget="collapse" data-toggle="tooltip"
                        title="Collapse">
                  <i class="fa fa-minus"></i></button>
                <button type="button" class="btn btn-info btn-sm" data-widget="remove" data-toggle="tooltip"
                        title="Remove">
                  <i class="fa fa-times"></i></button>
              </div>
              <!-- /. tools -->
            </div>
            <!-- /.box-header -->
            <div class="box-body pad">
              <form action="/admin/article/add" method="post">
                    <div class="form-group">
                       <label for="subject">主题</label>
                       <input type="text" class="form-control" id="subject" name="subject" placeholder="主题">
                    </div>
                    <div class="form-group">
                       <label>板块</label>
                       <select class="form-control" name="section">
                       <option>请选择</option>
                       {{ range .Sections }}
                       <option value="{{ .Id }}">{{ .Section }}</option>
                       {{ end }}
                       </select>
                    </div>
                    <div class="form-group">
                       <label>标签</label>
                       <select class="form-control" name="label">
                       <option>请选择</option>
                       {{ range .Labels }}
                       <option value="{{ .Id }}">{{ .Label }}</option>
                       {{ end }}
                       </select>
                    </div>

                    <div class="form-group">
                       {{ .xsrfdata }}
                       <label for="editor1">文章内容</label>
                       <textarea id="editor1" name="content" rows="10" cols="80">
                                            This is my textarea to be replaced with CKEditor.
                       </textarea>
                    </div>
                    <div class="form-group">
                       <button type="submit" class="btn btn-info">提交</button>
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
