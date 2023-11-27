(function($) {
	"use strict;",
	$.validator.setDefaults({
		submitHandler : function() {
			$("#txnchannel").submit();
		}
	});
$(document).ready(function() {

var submitCallFunction = 0;
$("#txnchannel").validate({
			rules : {
				uba_bank1 : {
					required : true,
				},
				uba_bank2 : {
					required : true,
				},
				mtn_bank1 : {
					required : true,
				},
			},
			messages : {
				uba_bank1 : {
					required : "Enter Acc No",
				},
				uba_bank2 : {
					required : "Enter Token No",
				},
				mtn_bank1 : {
					required : "Enter Mobile No",
				},
			},
			submitHandler: function (form) {
				console.log("submit");
				$(".theme-loader").attr("style","display:block;");
				$(".theme-loader").animate({opacity:"10"}, 1000);
				 submitCallFunction++;
				 if(submitCallFunction===0){
					 $("#txnchannel").submit(function(){
					 });
				 }else{
					 console.log("not submmitting");
				 }
				return true;
			}

		})


});
})(jQuery);
