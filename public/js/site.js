$(function() {
  $('a[href*="#"]:not([href="#"])').click(function() {
    if (location.pathname.replace(/^\//,'') == this.pathname.replace(/^\//,'') && location.hostname == this.hostname) {
      var target = $(this.hash);
      target = target.length ? target : $('[name=' + this.hash.slice(1) +']');
      if (target.length) {
        $('html, body').animate({
          scrollTop: target.offset().top
        }, 1000);
        return false;
      }
    }
  });
});

$(function() {
  $('#bookService').submit(function() {

    var name = $("#name").val();
    var email = $("#email").val();
    var phone = $("#phone").val();
    var fromDate = $("#fromDate").val();
    var toDate = $("#toDate").val();
    var address = $("#address").val();
    var details = $("#details").val();
    var recaptchaResponse = grecaptcha.getResponse();
    
    var data = {name: name, email: email, phone: phone, address: address, details: details, fromDate: fromDate, toDate: toDate, recaptchaResponse: recaptchaResponse};

    $.ajax({
      type: "POST",
      url: '/App/BookService',
      data: data,
      success: function(data, textStatus, jqXHR){
        $("#bookSuccess").show();
        $("#bookError").hide();

        $('#bookService')[0].reset();
      },
      error: function(jqXHR, textStatus, errorThrown){
        if (jqXHR.status === 400){
          $("#errorMessage").html("You must verify you are human by ticking the I'm not a robot checkbox");
        }
        else{
          $("#errorMessage").html("We were unable to receive your details at this time. Please try again or give us a call.");
        }

        $("#bookSuccess").hide();
        $("#bookError").show();
      }
    });

    return false;
  });
});

$(function() {
  $( ".date" ).datepicker({
    dateFormat: "yy-mm-dd",
    minDate: new Date(),
  });
});