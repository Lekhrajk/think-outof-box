<!-- =========================================================================================
    File Name: Login.vue
    Description: Login Page
    ----------------------------------------------------------------------------------------
    Item Name: Vuexy - Vuejs, HTML & Laravel Admin Dashboard Template
      Author: Pixinvent
    Author URL: http://www.themeforest.net/user/pixinvent
========================================================================================== -->


<template>
  <div class="h-screen flex w-full bg-img vx-row no-gutter items-center justify-center" id="page-login">
    <div class="vx-col sm:w-1/2 md:w-1/2 lg:w-full xl:w-3/4 sm:m-0 m-4">
      <vx-card>
        <div slot="no-body" class="full-page-bg-color">

          <div class="vx-row no-gutter justify-center items-center">

            <div class="vx-col hidden lg:block lg:w-1/2">
              <img src="@/assets/images/pages/login/login.svg" alt="login" class="mx-auto img-fluid" width="400">
            </div>

            <div class="vx-col sm:w-full md:w-full lg:w-1/2 d-theme-dark-bg">
              <div class="p-8 login-tabs-container">

                <div class="vx-card__title mb-4">
                  <h4 class="mb-4">Login</h4>
                  <p>Welcome back, please login to your account.</p>
                </div>

                <div>
                  <form id="loginform" method="post" @submit.prevent="login">
                  <vs-input
                    type="text"
                    icon-no-border
                    icon="icon icon-user"
                    icon-pack="feather"
                    label-placeholder="Username"
                    v-model="input.username"
                    autocomplete = '(suggested: "current-username")'
                    class="w-full"/>
                  <vs-input
                    type="password"
                    name="password"
                    icon-no-border
                    icon="icon icon-lock"
                    icon-pack="feather"
                    label-placeholder="Password"
                    v-model="input.password"
                    autocomplete = '(suggested: "current-password")'
                    class="w-full mt-6" />
                  <div class="flex flex-wrap justify-between my-5">
                      <vs-checkbox v-model="checkbox_remember_me" class="mb-3">Remember Me</vs-checkbox>
                      <!-- <router-link to="">Forgot Password?</router-link> -->
                  </div>
                  <div class="justify-between text-center mb-3">
                    <vs-button name="submit" id="submit" v-on:click="login()" >Login</vs-button>
                  </div>
                  </form>
                  <div class="text-danger p-2 text-center">{{ server_err }}</div>
                </div>

              </div>
            </div>
          </div>
        </div>
      </vx-card>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
import router from '@/router';
export default{
    data() {
      return {
      input: {
              username: "",
              password: ""
              },
      checkbox_remember_me: false,
      error: null,
      server_err: '',
    }
    },
    methods: 
    {
      login() 
      {
        if(this.input.username != "" && this.input.password != "")
        {
          const tokenurl = '/login';
          var logindata = new FormData();
          localStorage.setItem('name', this.input.username);
          localStorage.setItem('pw', this.input.password);
          logindata.set('username', this.input.username);
          logindata.set('password', this.input.password);

          axios({
          method: 'post',
          url: tokenurl,
          data: logindata,
          })
          .then(function(response) {
              if (response.status == 200) 
              {
              localStorage.setItem('token',response.data.token)
              router.replace({ name: "home", path: '/dashboard/home' });
              }
          })
          .catch((error) => {
              this.$vs.notify({  
              color: 'danger',
              title: "Status",
              text: error.response.data.error
            }) 
          })  
        }
        else
        {
          this.server_err = "**username and password is required**";
          this.$vs.notify({
                  title:'Error',
                  text: this.server_err,
                  color:'danger',
                  iconPack: 'feather',
                  icon:'icon-alert-circle',
                  position:'top-right'
          })

        }
      },
    }
}




</script>

<style lang="scss">
#page-login {
  .social-login-buttons {
    .bg-facebook { background-color: #1551b1 }
    .bg-twitter { background-color: #00aaff }
    .bg-google { background-color: #4285F4 }
    .bg-github { background-color: #333 }
  }
}
</style>
