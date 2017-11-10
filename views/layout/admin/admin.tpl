<!DOCTYPE html>
<html>
<head>
   <title>{{ .Pagetitle }}</title>
   <meta charset="utf-8">
   {{ template "layout/admin/css.tpl" .}}
   {{.HtmlHead}}

</head>


<body class="hold-transition skin-blue sidebar-mini">
   <div class="wrapper">
       {{ template "layout/admin/header.tpl" .}}
       {{ template "layout/admin/leftbar.tpl" .}}
       {{.LayoutContent}}
    </div>
       {{ template "layout/admin/footer_js.tpl" .}}
    {{ .Scripts }}
</body>
</html>
