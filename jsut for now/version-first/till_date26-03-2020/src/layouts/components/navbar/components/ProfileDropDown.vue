<template>
  <div class="the-navbar__user-meta flex items-center" v-if="activeUserInfo.displayName">

    <div class="text-right leading-tight hidden sm:block">
      <p class="font-semibold">{{ activeUserInfo.displayName }}</p>
      <small>Online</small>
    </div>

    <vs-dropdown vs-custom-content vs-trigger-click class="cursor-pointer">

      <div class="con-img ml-3">
        <img v-if="activeUserInfo.photoURL" key="onlineImg" :src="activeUserInfo.photoURL" alt="user-img" width="40" height="40" class="rounded-full shadow-md cursor-pointer block" />
      </div>

      <vs-dropdown-menu class="vx-navbar-dropdown">
        <ul style="min-width: 9rem">

          <!-- <li class="flex py-2 px-4 cursor-pointer hover:bg-primary hover:text-white">
            <feather-icon icon="UserIcon" svgClasses="w-4 h-4" />
            <span class="ml-2">Profile</span>
          </li> -->

          <!-- <li class="flex py-2 px-4 cursor-pointer hover:bg-primary hover:text-white">
            <feather-icon icon="MailIcon" svgClasses="w-4 h-4" />
            <span class="ml-2">Inbox</span>
          </li> -->
          <li class="flex py-2 px-4 cursor-pointer hover:bg-primary hover:text-white" @click="$router.replace({ path:'/resetPassword'})">
            <feather-icon icon="KeyIcon" svgClasses="w-4 h-4" />
            <span class="ml-2">Change Password</span>
          </li>

          <vs-divider class="m-1" />

          <li
            class="flex py-2 px-4 cursor-pointer hover:bg-primary hover:text-white"
            @click="logout">
            <feather-icon icon="LogOutIcon" svgClasses="w-4 h-4" />
            <span class="ml-2">Logout</span>
          </li>
        </ul>
      </vs-dropdown-menu>
    </vs-dropdown>
  </div>
</template>

<script>
import axios from 'axios';
export default {
  data() {
    return {
     storedName: '',
    }
  },
  computed: {
    activeUserInfo() {
      return this.$store.state.AppActiveUser
    }
  },
  methods: {
    logout() {
       var cookies = document.cookie.split(";");
       for (var i = 0; i < cookies.length; i++) {
        var cookie = cookies[i];
        var eqPos = cookie.indexOf("=");
        var name = eqPos > -1 ? cookie.substr(0, eqPos) : cookie;
        document.cookie = name + "=;expires=Thu, 01 Jan 1970 00:00:00 GMT";
      }
      const tokenurl = '/api/logout';
          axios({
          method: 'get',
          url: tokenurl,
          })
        localStorage.removeItem('token');
      
        this.$router.push('/').catch(() => {})
        }
      }
  }
</script>
