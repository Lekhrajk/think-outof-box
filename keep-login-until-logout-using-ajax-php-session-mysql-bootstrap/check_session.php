<?php

session_start();

if(isset($_SESSION['name']))
{
	echo '0';
}
else {
	echo '1';
}

?>