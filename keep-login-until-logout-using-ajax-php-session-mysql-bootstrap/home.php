<?php
session_start();
?>

<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>Login page</title>
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<!-- bootstrap cdn links for styling -->
	<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css">
	<script src="https://unpkg.com/sweetalert/dist/sweetalert.min.js"></script>
	<link rel="stylesheet" type="text/css" href="style.css">
	<link rel="stylesheet" href="https://use.fontawesome.com/releases/v5.7.0/css/all.css" integrity="sha384-lZN37f5QGtY3VHgisS14W3ExzMWZxybE1SJSEsQp9S+oqd12jhcu+A56Ebc1zFSJ" crossorigin="anonymous">
	<script type="text/javascript" src="main.js"></script>
	<script src="http://code.jquery.com/jquery-1.9.1.js"></script>
</head>
<body>
<div class="container">
	<?php 

   if(isset($_SESSION['name']))
   {  
   	$name = $_SESSION['name'];
   	echo '<h1 class="fontbig text-center">Welcome to the home page</h1>';
   	echo '<h2 class="float-left font-weight-lighter text-primary">'.$name.'</h2>';
   	echo '<a href="logout.php" class="lbtn lbtn-light-color float-right">Logout</a>';
   }

   else
   {
     header("location:index.php");
   }
	?>
	<div class="mt-lg-5 col-lg-4 col-xl-4 col-md-5 col-sm-12 mx-auto heading">
		<h1 class="big-title text-center backcolor">OUT OF BOX</h1>
	</div>
	<div>
		<div class="mt-md-5 col-md-8 mx-auto">
			<img src="https://jooinn.com/images/welcome-1.jpg" alt="welcome" class="img-fluid">
		</div>
	</div>
</div>
<div class="modal bg-transparent border-0" id="login_modal">
    <div class="modal-dialog modal-dialog-centered">
        <div class="modal-content bg-transparent border-0">
            <h4 class="modal-title text-center backcolor font p-3">Session Expired login again</h4>
            <!-- Modal body -->

                <div class="col-lg-12 col-xl-11 col-md-12 col-sm-12 mx-auto cardbox">
                    <span id="error_message" class="text-danger h4 font-weight-lighter"></span>
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
                        <p class="justify-content-start">forgot password?<a href="reset_password.html">reset here</a></p>
                    </form>
                </div>
        </div>
    </div>
</div>

<script type="text/javascript">
	$(document).ready(function() {
		var is_session_expired = 'no';
		function check_session() 
		{
			$.ajax({
             url:"check_session.php",
             method:"POST",
             success:function(data)
             {
             	if(data == '1')
             	{
             		$('#login_modal').modal({
             			backdrop:'static',
             			keyboard:false,
             		});
             		is_session_expired = 'yes';
             	}
             }
			})
		}
		var count_interval = setInterval(function()
		{
			check_session();
			if(is_session_expired == 'yes')
			{
				clearInterval(count_interval);
			}
		}, 5000);
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
		  posting.done(function( data ) 
		  {
		    if(data!='')
		    {
		    	$('#error_message').html(data);
		    }
		    else
		    {
		    	location.reload();
		    }
		  });
		});
	});
</script>

	<!--  js code link-->
	<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.4.1/jquery.min.js"></script>
	<script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.7/umd/popper.min.js"></script>
	<script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js"></script>
	<link href="https://fonts.googleapis.com/css?family=PT+Sans+Narrow&display=swap" rel="stylesheet">
	<link href="https://fonts.googleapis.com/css?family=Kodchasan&display=swap" rel="stylesheet">
</body>
</html>