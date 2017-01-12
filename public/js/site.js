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

    var data = {name: name, email: email, phone: phone, address: address, details: details, fromDate: fromDate, toDate: toDate};

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