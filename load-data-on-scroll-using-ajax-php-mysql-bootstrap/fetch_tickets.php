<?php

	session_start();
	define("DB_HOST", "localhost");
	define("DB_USER", "root");
	define("DB_PASS", "");
	define("DB_NAME", "outofbox");
	$connect = mysqli_connect(DB_HOST,DB_USER,DB_PASS,DB_NAME);
	if($connect->connect_error){
		die("Connection failed".$conn->connect_error);
	}

// fetching list of users on page load
if(isset($_POST['limit'], $_POST['start']))
{
	$query = "SELECT * FROM users ORDER BY id DESC LIMIT ".$_POST['start'].",".$_POST['limit']."";
    $results = mysqli_query($connect, $query);
	$output = '';
	if(mysqli_num_rows($results)>0){
	 		$number = 1;
	 		while($row = mysqli_fetch_array($results)){
	 			$ticket_id = $row['id'];
	 			$output .='<tr>
	 			    <td>'.$row['id'].'</td>
	 			    <td>'.$row['user_name'].'</td>
	 			    <td>'.$row['user_email'].'</td>
	 			    <td>'.$row['user_phone'].'</td>
	 			    <td>'.$row['user_address'].'</td>
	 			    <td>'.$row['user_country'].'</td>
	 			    <td>'.$row['user_postcode'].'</td>
	 			    <td>'.$row['user_ip'].'</td>
	 			    <td class="text-center float-middle">
	 			        <a href="#" class="hover lbtn lbtn-light-color lbtn-small" onclick='."get_full_details($ticket_id)".'>Take-action</a></td>
	 			   </tr> ';
	 			   $number++;
	 		}
	 	} 
	 			   echo $output;  
	 		
	
}

	//popup dynamic modal on click for each row of table 
if(isset($_POST['detailid']))
{
	$query = "SELECT * FROM users WHERE id='".$_POST["detailid"]."'";
    $results = mysqli_query($connect, $query);
	$output = '';
	if(mysqli_num_rows($results)>0){
	 		while($row = mysqli_fetch_array($results)){
	 			$output .='<tr>
	 			   <th>User-ID</th><td>'.$row['id'].'</td></tr><tr>
	 			   <th>User Name</th><td>'.$row['user_name'].'</td></tr><tr>
	 			   <th>Email id</th><td>'.$row['user_email'].'</td></tr><tr>
	 			   <th>Phone no.</th><td>'.$row['user_phone'].'</td></tr><tr>
	 			   <th>Password</th><td>'.$row['user_password'].'</td></tr><tr>
	 			   <th>Address</th><td>'.$row['user_address'].'</td></tr><tr>
	 			   <th>Country name</th><td>'.$row['user_country'].'</td></tr><tr>
	 			   <th>Postel Code</th><td>'.$row['user_postcode'].'</td></tr><tr>
	 			   <th>IP Address</th><td>'.$row['user_ip'].'</td></tr><tr>
	 			   <th>Created On</th><td>'.$row['created'].'</td></tr><tr>
	 			   </tr> ';
	 	} 
	 			   echo $output;  
	
}
}
?>