<!-- =========================================================================================
    File Name: VmStatus.vue
    Description: Virtual machine status
    ----------------------------------------------------------------------------------------
    Item Name: Xongl user control panel
    Author: Xongl cloud private limited
========================================================================================== -->

<template>
    <div id="vm-status-page">
        <div class="vx-row">
            <div class="vx-col lg:w-1/2 md:w-1/2 sm:w-full">
                <vx-card :title= "vms.state">
                    <ul class="bordered-items">
                        <li class="py-2" v-if="vms">
                                    <img :src="vms.logo" alt="graphic" width="100" height="100" class="mx-auto mb-4">
                                    <h4 class="mb-2">{{ vms.name}}</h4>
                            <div class="article-title mb-6">
                                <p class="mt-1">Last updated on <span>{{ vms.time }}</span></p>
                            </div>
                        </li>
                    </ul>
                </vx-card>
            </div>
            <div class="vx-col lg:w-1/2 md:w-1/2 sm:w-full">
                <vx-card>
                    <div class="vx-row">
                        <div class="vx-col lg:w-1/2 md:w-1/2 sm:w-full ">
                            <ul class="list-unstyled">                     
                                <li @click="actionvm(vm.id, name='reboot-hard')"  name="reboot-hard" value="reboot-hard"><p class="text-left py-2 px-2 cursor-pointer hover:bg-primary hover:text-white"><i class="feather icon-loader mr-4 "></i>Reboot</p></li>
                                <li  @click="actionvm(vm.id, name='poweroff-hard')"  name="poweroff-hard" value="poweroff-hard"><p class="text-left py-2 px-2 cursor-pointer hover:bg-primary hover:text-white" ><i class="feather icon-power mr-4 "></i>PowerOff</p></li>
                                <li @click="actionvm(vm.id, name='resume')"  name="resume" value="resume" ><p class="py-2 px-2 cursor-pointer hover:bg-primary hover:text-white" ><i class="feather icon-play-circle mr-4"></i>Resume</p></li>
                                <li  @click="changeName=true" ><p  @click="getVmName(vm.name,vm.id)" class="py-2 px-2 cursor-pointer hover:bg-primary hover:text-white"><i class="feather icon-file-text mr-4"></i>Rename</p></li>
                                <li> 
                                    <p @click="confirmDelete=true"  color="danger" type="flat" class="text-left py-2 px-2 cursor-pointer hover:bg-danger hover:text-white" ><span  @click="getVmId(vm.id)"><i class="feather icon-trash mr-4"></i> Delete</span></p>
                                </li >
                                <li @click="ResizeDisk=true" type="filled" >
                                    <p  @click="getdiskSize(vm.disk,vm.id)" class=" text-left py-2 px-2 cursor-pointer hover:bg-danger hover:text-white"><i class="feather icon-maximize-2 mr-4"></i> Resize disk</p>
                                </li>
                                <li @click="ResizeMemory=true" type="filled" >
                                    <p  @click="getdiskSize(vm.disk,vm.id)" class="py-2 px-2 cursor-pointer hover:bg-danger hover:text-white"><i class="feather icon-repeat mr-4"></i>Resize RAM/vCPU</p>
                                </li>
                            </ul>
                        </div>
                        <div class="vx-col lg:w-1/2 md:w-1/2 sm:w-full my-auto">
                            <img src="@/assets/images/pages/action.svg" alt="login" class="my-auto img-fluid" width="200">
                        </div>
                    </div>
                </vx-card>
            </div>  
        </div>         
        <vx-card class="mt-5">
            <div class="vx-row">
                <div class="vx-col w-full">
                    <vs-tabs alignment="fixed">
                        <vs-tab label="Snapshots">
                            <Snapshots />
                        </vs-tab>
                        <vs-tab label="Storage">
                            Storage
                        </vs-tab>
                    </vs-tabs>
                </div>
            </div>
       </vx-card>
    </div>
</template>

<script>
import axios from 'axios'
import Snapshots from './Snapshots.vue'
export default{
    data() {
        return {
            vms: {},
        }
    },
    mounted() {
        this.$emit('changeRouteTitle', 'VM Status');
        let temp_id = this.$route.params.id
        console.log('The id is: ' + this.$route.params.id);
        const tokenurl = '/vms/'+temp_id;
                axios({
                    method: 'get',
                    url: tokenurl,
                })
                .then((response) => {
                  this.vms =  response.data;
                  console.log(response.data);
                })
        const monitorurl = '/vms/'+temp_id+'/monitor';
                axios({
                    method: 'get',
                    url: monitorurl,
                })
                .then((response) => {
                  this.monitordata =  response.data;
                  console.log(response.data);
                })        

    },
    components: {
        Snapshots,
    },
        created() {
        }
}
</script>
