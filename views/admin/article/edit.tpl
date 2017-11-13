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
        <li class="active">编辑文章</li>
      </ol>
    </section>

    <!-- Main content -->
    <section class="content">
      <div class="row">
        <div class="col-xs-12">
          <div class="box box-info">
            <div class="box-header with-border">
               <ul class="nav nav-tabs">
               <li><a href="/admin/article"><i class="fa fa-list"></i> 文章列表</a></li>
               <li><a href="/admin/article/add"><i class="fa fa-plus"></i> 新增文章</a></li>
               <li class="active"><a href="/admin/article/edit/{{ .Blog.Id }}"><i class="fa fa-edit"></i> 编辑文章</a></li>
               </ul>
            </div>
            <!-- /.box-header -->
            <div class="box-body pad">
              <form action="/admin/article/edit/{{ .Blog.Id }}" method="post">
                    <div class="form-group">
                       <label for="subject">主题</label>
                       <input type="text" class="form-control" id="subject" name="subject" value="{{ .Blog.Title }}">
                    </div>
                    <div class="form-group">
                       <label>板块</label>
                       <select class="form-control" name="section">
                       <option value="{{ .Blog.Section.Id }}">{{ .Blog.Section.Section }}</option>
                       {{ range .Sections }}
                       <option value="{{ .Id }}">{{ .Section }}</option>
                       {{ end }}
                       </select>
                    </div>
                    <div class="form-group">
                       <label>标签</label>
                       <select class="form-control" name="label">
                       <option value="{{ .Blog.Label.Id }}">{{ .Blog.Label.Label }}</option>
                       {{ range .Labels }}
                       <option value="{{ .Id }}">{{ .Label }}</option>
                       {{ end }}
                       </select>
                    </div>

                    <div class="form-group">
                       {{ .xsrfdata }}
                       <label for="editor1">文章内容</label>
                       <textarea id="editor1" name="content" rows="10" cols="80">{{ .Blog.Content }}</textarea>
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

