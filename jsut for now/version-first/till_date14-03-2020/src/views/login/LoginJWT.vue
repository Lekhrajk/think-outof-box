<template>
  <div>
    <form id="loginform" method="post" @submit.prevent="login">
      <vs-input
        v-validate="'required'"
        data-vv-validate-on="blur"
        name="username"
        icon-no-border
        icon="icon icon-user"
        icon-pack="feather"
        label-placeholder="Username"
        v-model="input.username"
        class="w-full"/>
    <span class="text-danger text-sm">{{ errors.first('username') }}</span>

    <vs-input
        data-vv-validate-on="blur"
        v-validate="'required|min:3|max:30'"
        type="password"
        name="password"
        icon-no-border
        icon="icon icon-lock"
        icon-pack="feather"
        label-placeholder="Password"
         v-model="input.password"
        class="w-full mt-6" />
        <span class="text-danger text-sm">{{ errors.first('password') }}</span>

        <div class="flex flex-wrap justify-between my-5">
            <vs-checkbox v-model="checkbox_remember_me" class="mb-3">Remember Me</vs-checkbox>
            <router-link to="/pages/forgot-password">Forgot Password?</router-link>
        </div>
        <div class="flex flex-wrap justify-between mb-3">
          <vs-button :disabled="!validateForm" 
                name="submit"  id="submit" v-on:click="login()" >Login</vs-button>
        </div>
    </form>
  </div> 
</template>

<script>
import axios from 'axios';
import router from '@/router';
export default 
{
  data()
  {
    return 
    {
      input: {
              username: ""
              password: ""
            }
      checkbox_remember_me: false
    }
  },
  computed: {
      validateForm() {
        return !this.errors.any() && this.input.username != '' && this.input.password != '';
      },
  },
  methods: 
  {
    login() 
    {
      if(this.input.username != "" && this.input.password != "") 
      {
        const tokenurl = '/login';
        var logindata = new FormData();
        logindata.set('username', this.input.username);
        logindata.set('password', this.input.password);
        axios({
            method: 'post',
            url: tokenurl,
            data: logindata,
        })
        .then(function(response) 
        {
            if (response.status == 200) 
            {
              localStorage.setItem('token',response.data.token)
              router.replace({ name: "dashboard-analytics" }); 
            }
        })
      }
    },
  }
}

</script>

