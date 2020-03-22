<template>
    <div class="sidepanel">
      <!-- toggle icon for small devices -->
        <div class="d-md-none d-sm-block float-right p-2 cardbox" @click="toggleMenu">
            <i class="fas fa-bars fa-2x text-primary"></i>
        </div>
        <!-- sidebar starts here -->
        <transition name="slide-fade">
            <div class="sidenav" v-if="isActive">
                <div class="brand-name">
                    <router-link to="/" id="brand"> <img src="@/assets/img/home.svg" class="img-fluid" /></router-link>
                </div>
                <router-link to="/Class I" class="sidelink"><i class="fas fa-school mr-4 float-left mt-1"></i> Class I</router-link>
                <router-link to="/Class II" class="sidelink"><i class="fas fa-school mr-4 float-left mt-1"></i> Class II</router-link>
                <router-link to="/Class III" class="sidelink"><i class="fas fa-school mr-4 float-left mt-1"></i> Class III</router-link>
                <router-link to="/Class IV" class="sidelink"><i class="fas fa-school mr-4 float-left mt-1"></i> Class IV</router-link>
                <router-link to="/Class V" class="sidelink"><i class="fas fa-school mr-4 float-left mt-1"></i> Class V</router-link>
                <router-link to="/Class VI" class="sidelink"><i class="fas fa-school mr-4 float-left mt-1"></i> Class VI</router-link>
                <router-link to="/Class VII" class="sidelink"><i class="fas fa-school mr-4 float-left mt-1"></i> Class VII</router-link>
                <a href="#" class="text-center text-white">Total <span class="badge badge-light">{{ getClassCount }}</span></a>
            </div>
        </transition>
        <!-- end of sidebar -->
    </div>
</template>
<script>
import { mapGetters } from 'vuex';
export default {
    name: 'Sidepanel',
    data() {
        return {
            isActive: true,
        }
    },
    computed: {
        ...mapGetters(['getClassCount']),
    },
    methods: {
        toggleMenu() {
            this.isActive = !this.isActive;
        }
    }
}
</script>
<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
/* Sidebar Styles */
* {
    padding: 0;
    margin: 0;
}

.sidenav {
    height: 100%;
    width: 220px;
    position: fixed;
    z-index: 1;
    top: 0;
    left: 0;
    background-color: #111;
    overflow-x: hidden;
    padding-top: 20px;
}

.sidenav .brand-name,
#brand {
    text-align: center;
    margin-top: 0px !important;
}

.sidenav a {
    margin-top: 10px;
    padding: 8px 16px;
    text-decoration: none;
    font-size: 20px;
    color: #818181;
    display: block;
}

.sidelink:hover,
.sidelink.router-link-active,
.sidelink.router-link-exact-active {
    background-color: #0000ff;
    color: #ffffff;
    cursor: pointer;
}


.main {
    margin-left: 160px;
    /* Same as the width of the sidenav */
    font-size: 28px;
    /* Increased text to enable scrolling */
    padding: 0px 10px;
}

.slide-fade-enter-active {
    transition: all .5s ease;
}

.slide-fade-leave-active {
    transition: all .8s cubic-bezier(1.0, 0.5, 0.8, 1.0);
}

.slide-fade-enter,
.slide-fade-leave-to {
    transform: translateX(-150px);
    opacity: 0;
}

@media screen and(max-height: 468px) {
    .sidenav {
        padding-top: 15px;
    }

    .sidenav a {
        font-size: 18px;
    }
}
</style>