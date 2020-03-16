<?php

define("DB_HOST", "localhost"); //hostname
define("DB_USER", "root");      //username
define("DB_PASS", "");          //password
define("DB_NAME", "outofbox");  //give your dbname
$conn = mysqli_connect(DB_HOST,DB_USER,DB_PASS,DB_NAME);
if($conn->connect_error){
	die("Connection failed".$conn->connect_error);
}
if(isset($_POST['name'])){
	$output = '';
	
		// query for getting data from database
		
	$query = "SELECT * FROM itemlist WHERE item LIKE '%".$_POST['name']."%' ";
	$result = mysqli_query($conn, $query);
	$output ='<ul class="list-unstyled">';
	if(mysqli_num_rows($result)>0)
	{
		while($row = mysqli_fetch_array($result))
		{
			$output .='<li>'.$row["item"].'</li>';
		}
	} 
	else
	{
		$output .='<li>Item not found</li>';
	}      
	$output .='</ul>';
	echo $output; 
}
?>