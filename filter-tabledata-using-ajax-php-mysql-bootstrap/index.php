<?php

  define("DB_HOST", "localhost"); //hostname
  define("DB_USER", "root");      //username
  define("DB_PASS", "");          //password
  define("DB_NAME", "outofbox");  //give your dbname
  $conn = mysqli_connect(DB_HOST,DB_USER,DB_PASS,DB_NAME);
  
  // Check connection
  if ($conn->connect_error) {
  die("Connection failed: " . $conn->connect_error);
  }

  $query = "SELECT * FROM itemlist ORDER BY id desc";
  $result = mysqli_query($conn, $query);

?>
<!DOCTYPE html>
<html lang="en">

<head>
    <title>Filter using ajax</title>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <!-- bootstrap cdn links for styling -->
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css">
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.4.1/jquery.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.7/umd/popper.min.js"></script>
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js"></script>
    <style type="text/css">
    .table-responsive {
        display: table;
    }
    </style>
</head>

<body>
    <div class="jumbotron text-center">
        <h2 class="text-center text-primary">Filter using Ajax</h2>
        <p class="text-center text-info">Select from and to date and see filter output</p>
    </div>
    <div class="container">
        <div class="row">
            <div class="col-lg-4 col-xl-4 col-sm-12 col-md-4">
                <input type="date" name="from_date" id="from_date" value="" placeholder="enter fromdate" class="form-control" required>
            </div>
            <div class="col-lg-4 col-xl-4 col-sm-12 col-md-4">
                <input type="date" name="to_date" id="to_date" value="" placeholder="enter todate" class="form-control" required>
            </div>
            <div class="col-4">
                <input type="button" name="filter" class="btn btn-warning btn-sm" id="filter" value="Filter">
            </div>
            <div style="clear: both;"></div>
            <div id="order_table" class="container">
                <table class="table table-bordered mx-auto w-100 d-block d-md-table mt-5">
                    <thead class="bg-dark text-white">
                        <tr>
                            <th>OrderId</th>
                            <th>Customer name</th>
                            <th>Item</th>
                            <th>Price</th>
                            <th>Order Date</th>
                        </tr>
                    </thead>
                    <?php 
           while($row = mysqli_fetch_array($result))
           {
          ?>
                    <tbody>
                        <tr>
                            <td>
                                <?php echo $row["id"]; ?>
                            </td>
                            <td>
                                <?php echo $row["Customer name"]; ?>
                            </td>
                            <td>
                                <?php echo $row["item"]; ?>
                            </td>
                            <td>
                                <?php echo $row["value"]; ?>
                            </td>
                            <td>
                                <?php echo $row["order_date"]; ?>
                            </td>
                        </tr>
                    </tbody>
                    <?php
         }
         ?>
                </table>
            </div>
        </div>
    </div>
    <!-- ajax post request -->
    <script type="text/javascript">
        $("#filter").click(function() {
          var from_date  = $('#from_date').val();
          var to_date  = $('#to_date').val();
          if(from_date!='' && to_date!='')
          {
            $.ajax({
              url:"filter.php",
              method:"POST",
              data:{from_date:from_date, to_date:to_date},
              success:function(data){
                $('#order_table').html(data);
              }
            });
          }
          else
          {
            alert("please select the id");
          }
        });
        
    </script>

</body>
</html>