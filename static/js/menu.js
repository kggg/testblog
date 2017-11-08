var timeout  = 500;
var closetimer = 0;
var ddmenuitem  = 0;

function cnnoc_open()
{
     cnnoc_canceltimer();
         cnnoc_close();
         ddmenuitem = $(this).find('ul').eq(0).css('visibility', 'visible');
}

function cnnoc_close()
{
     if(ddmenuitem) ddmenuitem.css('visibility', 'hidden');
}

function cnnoc_timer()
{
     closetimer = window.setTimeout(cnnoc_close, timeout);
}

function cnnoc_canceltimer()
{
      if(closetimer)
      {
          window.clearTimeout(closetimer);
          closetimer = null;
      }
}

