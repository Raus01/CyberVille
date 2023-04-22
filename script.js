$(document).ready(function() {
    $("#login-form").submit(function(event) {
      event.preventDefault();
      var username = $("#username").val();
      var password = $("#password").val();
      if (username === "ss" && password === "ss") {
        window.location.href = "dashboard.html";
      } else {
        alert("Invalid username or password. Please try again.");
      }
    });
  });
  