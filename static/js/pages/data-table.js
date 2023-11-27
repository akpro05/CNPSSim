//[Data Table Javascript]



$(function () {
    "use strict";

$('#example1').DataTable( {
				"processing": true,
				"serverSide": true,
				'paging'      : true,
				'lengthChange': false,
				'searching'   : false,
				'ordering'    : true,
				'info'        : true,
				'autoWidth'   : false,				
							
				
				
		        "ajax": {
		            "url": "../../server-json-data.php",
		            "type": "POST",
		            "dataSrc": "records"
		        },
		        "columns": [
		            { "data": "invoice_no" },
		            { "data": "product_name" },
		            { "data": "delivery_status" },
					{ "data": "pin_code" },
		        ],
				"columnDefs": [
					{
			        	"targets": 2,
			            "render": function ( data, type, row ) {
			                return data == 1 ? 'Delivered': 'Not delivered';
			            }
		        	},
		        	{
		      		 	"targets": 3,
		                "render": function ( data, type, row ) {
		                    return row["city"] +', ' + row["country"] +', '+data;
		                },
		              
		            },
		        ]
		    });
	
	
	
    $('#example2').DataTable({
      'paging'      : true,
      'lengthChange': false,
      'searching'   : false,
      'ordering'    : true,
      'info'        : true,
      'autoWidth'   : false
    });
	
	
	$('#example').DataTable( {
		dom: 'Bfrtip',
		buttons: [
			'copy', 'csv', 'excel', 'pdf', 'print'
		]
	} );
	
	$('#tickets').DataTable({
	  'paging'      : true,
	  'lengthChange': false,
	  'searching'   : false,
	  'ordering'    : true,
	  'info'        : true,
	  'autoWidth'   : false,
	});
	
	$('#employeelist').DataTable({
	  'paging'      : true,
	  'lengthChange': false,
	  'searching'   : true,
	  'ordering'    : true,
	  'info'        : true,
	  'autoWidth'   : false,
	});
	
  }); // End of use strict