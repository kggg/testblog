{{ template "admin/datatable-js.tpl" .}}

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

