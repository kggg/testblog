<div class="row">
  {{template  "blog/left.tpl" .}}

  <div class="col-sm-9 padding-auto-5">
      <div class="bg-color-white border-grey-1 border-radius-3 padding-auto-15">
           {{ range .Blog }}
            <div class='border-grey-1 border-radius-3 margin-tb-10'>
                <div class='padding-auto-15 padding-tb-15'>
                    <a href="/show/{{ .Id }}"> <span class="lead">{{ .Title }}<span></a>
                </div>
                <div class='padding-auto-15 padding-tb-15 text-success'>
                       
                </div>
                <div class="padding-auto-15 text-info">
                     <p>{{$var := html2str .Content  }}{{ $var2 := htmlunquote $var }}{{ substr $var2 0 200 }}...</p>
                </div>

            </div>
           {{ end }}
      </div>
      {{template "blog/page.tpl" .}}
  </div>
</div>


