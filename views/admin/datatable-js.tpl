<!-- DataTables -->
<script src="/static/js/admin/jquery.dataTables.min.js"></script>
<script src="/static/js/admin//dataTables.bootstrap.min.js"></script>
<!-- SlimScroll -->
<script src="/static/js/admin/jquery.slimscroll.min.js"></script>


<script>

  $(document).ready(function() {
    $('#article1').DataTable({
      'paging'      : true,
      'lengthChange': true,
      'searching'   : true,
      'ordering'    : true,
      'info'        : true,
      'autoWidth'   : true
    });
  });

</script>


