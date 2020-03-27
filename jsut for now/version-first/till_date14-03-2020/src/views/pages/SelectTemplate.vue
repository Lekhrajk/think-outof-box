<!-- =========================================================================================
    File Name: SelectTemplate.vue
    Description: List of Available templates
    ----------------------------------------------------------------------------------------
    Item Name: Xongl user control panel
    Author: Xongl cloud private Limited
========================================================================================== -->


<template>
    <div id="select-template-page">
        <!-- Templates CARDS  -->
        <div class="py-8">
            <h1>List of Available Templates</h1>
        </div>
        <div class="vx-row" v-if="vms">
            <div class="vx-col w-full md:w-1/5 sm:w-1/2 mb-base"  v-for="vm in vms.templates" :key="vm.id" @click="$router.push('/select-template/category/'+vm.id).catch(() => {})">
                <vx-card class="text-center cursor-pointer">
                    <img :src="vm.logo" alt="graphic" width="80" height="70" class="mx-auto mb-4">
                    <h4 class="mb-2">{{ vm.name }}</h4>
                </vx-card>
            </div>
        </div>
    </div>
</template>

<script>
import axios from 'axios';
export default{
    data() {
        return {
           vms: null,
        }
    },
    mounted() {

                const tokenurl = '/templates';
                axios({
                    method: 'get',
                    url: tokenurl,
                })
                .then((response) => {
                  this.vms =  response.data;
                  console.log(response.data);
                })

      },
    components: {}
}
</script>

<style lang="scss">

</style>
