<?php

	session_start();

if(isset($_POST['email']))
{
	define("DB_HOST", "localhost"); //hostname
	define("DB_USER", "root");      //username
	define("DB_PASS", "");          //password
	define("DB_NAME", "outofbox");  //give your dbname

	$connect = mysqli_connect(DB_HOST,DB_USER,DB_PASS,DB_NAME);

	if($connect->connect_error){
		die("Connection failed".$conn->connect_error);
	}
    // sql query for checking user email
	$query = "SELECT * FROM demouser WHERE email = '".$_POST['email']."'";
    $results = mysqli_query($connect, $query);
	$output = '';
	$row = mysqli_fetch_array($results,MYSQLI_ASSOC);
    if (mysqli_num_rows($results) ==1)
    {
        if($_POST['password']==$row['password'])	
		   {
		   	$_SESSION['name'] = $row['firstname'];
		   }
		else
		{
			$output .= '<label class="h3 text-danger font-weight-lighter">Wrong Password</label>';
		}   
    }
	else
	{
       $output .= '<label class="h3 text-danger font-weight-lighter">Email doesnot exists</label>';
	}
    echo $output;	
}	

?>