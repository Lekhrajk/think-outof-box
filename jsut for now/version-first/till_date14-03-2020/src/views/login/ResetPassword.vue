<!-- =========================================================================================
    File Name: ResetPassword.vue
    Description: Reset Password Page
    ----------------------------------------------------------------------------------------
    Item Name: Xongl user control panel
    Author: Xongl cloud private Limited
========================================================================================== -->


<template>
    <div class="h-screen flex w-full bg-img" id="reset-password">
        <div class="vx-col sm:w-3/5 md:w-3/5 lg:w-3/4 xl:w-3/5 mx-auto self-center">
            <vx-card>
                <div slot="no-body" class="full-page-bg-color">
                    <div class="vx-row">
                        <div class="vx-col hidden sm:hidden md:hidden lg:block lg:w-1/2 mx-auto self-center">
                            <img src="@/assets/images/pages/change-password.svg" alt="login" class="mx-auto"  width="400">
                        </div>
                        <div class="vx-col sm:w-full md:w-full lg:w-1/2 mx-auto self-left  d-theme-dark-bg">
                            <div class="p-8">
                                 <h3 class="text-danger">{{ msg }}</h3>
                                <form id="resetform" method="post" @submit.prevent="reset">
                                    <div class="vx-card__title mb-8">
                                        <h4 class="mb-4">Reset Password</h4>
                                        <p>Please enter your new password.</p>
                                    </div>
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
                                    <vs-input type="password" label-placeholder="Confirm Password" v-model="input.confirmpassword" class="w-full mb-8" />

                                    <div class="flex flex-wrap justify-between flex-col-reverse sm:flex-row">
                                        <vs-button type="border" to="/" class="w-full sm:w-auto mb-8 sm:mb-auto mt-3 sm:mt-auto">Go Back To Login</vs-button>
                                        <vs-button class="w-full sm:w-auto" name="submit" id="submit" v-on:click="reset()">Reset</vs-button>
                                    </div>
                                </form>
                            </div>
                        </div>
                    </div>
                </div>
            </vx-card>
        </div>
    </div>
</template>

<script>
/* eslint-disable */
import axios from 'axios';
import router from '@/router';
export default {
  data () {
    return {
        input: {
            password: "",
            confirmpassword: ""
        },
        msg: '',
    }
  },
  methods: 
    {
      reset() 
      {
        if(this.input.password != "" && this.input.password === this.input.confirmpassword)
        {
          const tokenurl = '/14/resetpassword';
          var resetdata = new FormData();
          resetdata.set('password', this.input.password);
          axios({
          method: 'post',
          url: tokenurl,
          data: resetdata,
          })
          .then((response) => {
              this.msg  = response.data.message;
              this.$vs.notify({  
              color: 'danger',
              title: "Status",
              text: response.data.message
              }) 
          })
          .catch((error) => {
              this.$vs.notify({  
              color: 'danger',
              title: "Status",
              text: error.response.data
            }) 
          })  
        }
        else
        {
          this.server_err = "**password does't match**";
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
