<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <title>Login page</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <!-- bootstrap cdn links for styles -->
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css">
    <link rel="stylesheet" type="text/css" href="style.css"
    <script type="text/javascript" src="main.js"></script>
    <script src="http://code.jquery.com/jquery-1.9.1.js"></script>
</head>

<body>
    <div class="container">
        <div class="mt-lg-5 col-lg-4 col-xl-4 col-md-5 col-sm-12 mx-auto heading">
            <h1 class="big-title text-center backcolor">OUT OF BOX</h1>
        </div>
        <div class="col-lg-5 col-xl-5 col-md-6 col-sm-12 mx-auto cardbox">
            <!-- <span id="error_message" class="text-danger h4 font-weight-lighter"></span> -->
            <form id="loginform" method="post" action="loginverify.php" class="p-3 mx-3">
                <div class="form-group">
                    <label class="form-control-label font-weight-bolder font">Email id</label>
                    <input type="email" name="email" id="email" class="form-control" placeholder="enter emailid" required autocomplete="off">
                </div>
                <div class="form-group">
                    <label class="form-control-label font-weight-bolder font">Password</label>
                    <input type="password" name="password" id="password" class="form-control" placeholder="enter your password" required>
                </div>
                <div class="text-center">
                    <input type="submit" name="submit" value="Login" id="submit" class="lbtn lbtn-accent-color">
                </div>
                <p class="justify-content-start">forgot password?<a href="resetpasswd.html">reset here</a></p>
            </form>
        </div>
        <mark>
          **email id : thinkoutofbox@gmail.com **<br>
          **password: demouser123 **
        </mark>
    </div>

    <!-- modal for showing error -->

    <div class="modal fade" id="error_modal">
        <div class="modal-dialog modal-dialog-centered">
            <div class="modal-content">
                <!-- Modal body -->
                <div class="modal-body">
                    <button type="button" class="close" data-dismiss="modal">&times;</button>
                    <span id="error_message" class="text-danger h4 font-weight-lighter text-center"></span>
                </div>
            </div>
        </div>
    </div>

  <!--ajax call for login  -->
<script type='text/javascript'>
    /* attach a submit handler to the form */
    $("#loginform").submit(function(event) {

      /* stop form from submitting normally */
      event.preventDefault();

      /* get the action attribute from the <form action=""> element */
      var $form = $( this ),
          url = $form.attr( 'action' );

      /* Send the data using post with element id name and name2*/
      var posting = $.post( url, { email: $('#email').val(), password: $('#password').val() } );

      /* Alerts the results */
      posting.done(function( data ) {
        if(data!=''){
        	// $('#error_message').html(data);
        	$('#error_modal').modal({
             			backdrop:'static',
             			keyboard:false,
             		});
        	$('#error_message').html(data);

        }
        else
        {
        	window.location='home.php';
        }
      });
    });
</script>


	<!--  js code link-->
	<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.4.1/jquery.min.js"></script>
	<script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.7/umd/popper.min.js"></script>
	<script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js"></script>
	  
</body>
</html>