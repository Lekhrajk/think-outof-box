<?php

if(isset($_POST['from_date'],$_POST['to_date']))
{
	define("DB_HOST", "localhost"); //hostname
	define("DB_USER", "root");      //username
	define("DB_PASS", "");          //password
	define("DB_NAME", "outofbox");  //give your dbname
	$conn = mysqli_connect(DB_HOST,DB_USER,DB_PASS,DB_NAME);
	if($conn->connect_error){
		die("Connection failed".$conn->connect_error);
	}
	$output = '';

	// query for getting data from database

	$query = "SELECT * FROM itemlist WHERE order_date BETWEEN '".$_POST["from_date"]."' AND '".$_POST["to_date"]."' ";
	$result = mysqli_query($conn ,$query);
	$output .='
	    <table class="table table-bordered mx-auto w-100 d-block d-md-table mt-5">
	        <thead class="bg-dark text-white">
	          <tr>
	            <th>OrderId</th>
	            <th>Customer name</th>
	            <th>Item</th>
	            <th>Price</th>
	            <th>Order Date</th>
	          </tr>
	        </thead>';
			if(mysqli_num_rows($result)>0)
			{
				while ($row = mysqli_fetch_array($result)) {
					$output .= '
		            <tbody>
		              <tr>
		                <td> '.$row["id"].'</td>
		                <td> '.$row["Customer name"].'</td>
		                <td> '.$row["item"].'</td>
		                <td> '.$row["value"].'</td>
		                <td> '.$row["order_date"].'</td>
		              </tr>
		            </tbody>

					';
				    
				}
			}
			else
			{
		        $output .='
		        <tbody>
		        	<tr>
		        	    <td colspan="5">No data found</td>
		        	</tr>
		        </tbody>
		          ';
			} 
	$output .='</table>';
	echo $output;       

}

?>