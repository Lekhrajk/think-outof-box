<template>
  <div class="layout">
     <nav class="navbar navbar-expand-lg navbar-light bg-light shadow">
        <div class="container-fluid">
            <span class="navbar-brand font-weight-bolder">TODO ASSIGNMENT</span>
            <div class="navbar-toggler border-0" data-toggle="collapse" data-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
                <span ><i class="fas fa-bars fa-lg text-dark"></i></span>
            </div>
            <div class="collapse navbar-collapse" id="navbarSupportedContent">
                <ul class="navbar-nav ml-auto">
                    <li class="nav-item">
                        <router-link class="nav-link active_class" :to="{ name: 'home' }">TODOS</router-link>
                    </li>
                    <li class="nav-item">
                        <button class="nav-link btn btn-primary btn-sm text-white"  data-toggle="modal" data-target="#create">Create ToDo</button>
                    </li>
                </ul>
            </div>
        </div>
    </nav>
    <div id="layoutSidenav_content">
      <main>
        <div class="container-fluid">
          <router-view></router-view>
        </div>
      </main>
      <footer class="footer mt-auto footer-light">
        <div class="container-fluid">
          <div class="row">
            <div class="col-md-6 small">Copyright &#xA9; Haryanvi Developer 2020</div>
            <div class="col-md-6 text-md-right small">
              <a href="#" class="ext-links">Privacy Policy</a>
              &#xB7;
              <a href="#" class="ext-links">Terms &amp; Conditions</a>
            </div>
          </div>
        </div>
      </footer>
    </div>
    <CreateModal/>
  </div>
</template>

<script>
import CreateModal from './CreateModal'
  export default {
    name: 'Layout',
    components:{
      CreateModal
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  body {
    min-height: 100vh;
  }
.nav-link{
    padding: 8px 15px;
    margin-left: 10px;
    margin-right: 20px;
    border-bottom: 2px solid transparent;
}

.nav-link:hover{
    color: black;
    border-bottom: 2px solid #000;
}
.nav-link.router-link-active{
  color: black;
  border-bottom: 2px solid #000;
}

.footer {
  height: 5rem;
  display: -webkit-box;
  display: flex;
  -webkit-box-align: center;
          align-items: center;
}
.footer.footer-dark {
  color: rgba(255, 255, 255, 0.6);
}
.footer.footer-light {
  color: #a2acba;
}
.ext-links:hover{
  color: #000!important;
  cursor: pointer!important;
  text-decoration: none;
}

</style>