    function thanks() {
        var name=document.getElementById("fir").value;
        var email=document.getElementById("sec").value;
        var message=document.getElementById("thir").value;
        var n=name.length;
        var e=email.length;
        var m=message.length;
        if(n==0 || e==0 || m==0){
              swal("Error!", "Please fill correctly", "error"); 
        }
        else{
      swal("Thanks!", "You message has been sent", "success");
            }
    }

// for hiding the sidenav on scroll
            function openNav1() {
              document.getElementById("myNav1").style.width = "100%";
            }
            function openNav2() {
              document.getElementById("myNav2").style.width = "100%";
            }
            function openNav3() {
              document.getElementById("myNav3").style.width = "100%";
            }
              function openNav4() {
              document.getElementById("myNav4").style.width = "100%";
            }

            function closeNav() {
              document.getElementById("myNav1").style.width = "0%";
              document.getElementById("myNav2").style.width = "0%";
              document.getElementById("myNav3").style.width = "0%";
              document.getElementById("myNav4").style.width = "0%";
            }    
