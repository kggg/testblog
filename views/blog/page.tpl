          <div class="list-group navigation">
              <ul class="pagination">
                   <li> <a href="/page/1"> 首页 </a></li>
                     <li class="disabled">
                     <span>
                         <span aria-hidden="true">&laquo;</span>
                     </span>
                    </li>

                   {{if gt .Page 3}}
                   <li> <a href="/page/1">1</a></li>
                   <li> <a href="/page/2">2</a></li>
                   <li> <a href="/page/3">3</a></li>
                      {{ if and (ge .Page 4)  (le .Page .Pagetotal) }}
                          <li> <a href="javascript:void(0);">...</a></li>
                          <li> <a href="/page/{{ .Page }} ">{{ .Page }}</a></li>
                          <li> <a href="/page/{{adds .Page 1 }} ">{{adds  .Page 1}}</a></li>
                      {{ end }}

                   {{else  }}
                   <li> <a href="/page/1">1</a></li>
                   <li> <a href="/page/2">2</a></li>
                   <li> <a href="/page/3">3</a></li>
                   {{ end }}
                          <li> <a href="javascript:void(0);">...</a></li>
                    <li> <a href="/page/{{adds .Page 1 }}">&raquo;</a>
                    </li>
                   <li> <a href="/page/{{ .Pagetotal }}">尾页</a></li>

              </ul>
          </div>


