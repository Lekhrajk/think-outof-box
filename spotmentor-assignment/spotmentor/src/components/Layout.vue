<template>
    <div class="layout">
        <div class="row container-fluid mx-0 px-0">
            <div class="col-md-2 col-12">
                <Sidepanel></Sidepanel>
            </div>
            <div class="col-md-10 col-12">
                <router-view />
            </div>
            
        </div>
    </div>
</template>
<script>
import Sidepanel from './Sidepanel.vue'
export default {
    name: 'Layout',
    components: {
        Sidepanel,
    },
    
}
</script>
<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
* {
    margin-left: 0px !important;
    min-height: 650px;
    background-color: #F2F3F4;
}

@media only screen and (max-width: 600px) {
    *{
      min-height: 5px;
    }
}
</style>