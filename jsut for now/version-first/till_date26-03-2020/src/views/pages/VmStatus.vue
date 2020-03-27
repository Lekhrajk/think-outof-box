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
        <!-- ///////////////////////
         left side card for vm info
         ////////////////////////// -->
        <vx-card :title="vms.state" refresh-content-action @refresh="acceptAlert">
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
      <!-- ///////////////////////////////////
            card containing action buttons 
           //////////////////////////////////-->
      <div class="vx-col lg:w-1/2 md:w-1/2 sm:w-full">
        <vx-card>
          <div class="vx-row">
            <div class="vx-col lg:w-1/2 md:w-1/2 sm:w-full ">
              <ul class="list-unstyled">
                <li @click="actionvm(vms.id, name='reboot-hard')" name="reboot-hard" value="reboot-hard">
                  <p class="text-left py-2 px-2 cursor-pointer hover:bg-primary hover:text-white"><i class="feather icon-loader mr-4 "></i>Reboot</p>
                </li>
                <li @click="actionvm(vms.id, name='poweroff-hard')" name="poweroff-hard" value="poweroff-hard">
                  <p class="text-left py-2 px-2 cursor-pointer hover:bg-primary hover:text-white"><i class="feather icon-power mr-4 "></i>PowerOff</p>
                </li>
                <li @click="actionvm(vms.id, name='resume')" name="resume" value="resume">
                  <p class="py-2 px-2 cursor-pointer hover:bg-primary hover:text-white"><i class="feather icon-play-circle mr-4"></i>Resume</p>
                </li>
                <li @click="changeName=true">
                  <p @click="getVmName(vms.name,vms.id)" class="py-2 px-2 cursor-pointer hover:bg-primary hover:text-white"><i class="feather icon-file-text mr-4"></i>Rename</p>
                </li>
                <li>
                  <p @click="confirmDelete=true" color="danger" type="flat" class="text-left py-2 px-2 cursor-pointer hover:bg-danger hover:text-white"><span @click="getVmId(vms.id)"><i class="feather icon-trash mr-4"></i> Delete</span></p>
                </li>
                <li @click="ResizeDisk=true" type="filled">
                  <p @click="getdiskSize(vms.disk,vms.id)" class=" text-left py-2 px-2 cursor-pointer hover:bg-danger hover:text-white"><i class="feather icon-maximize-2 mr-4"></i> Resize disk</p>
                </li>
                <li @click="ResizeMemory=true" type="filled">
                  <p @click="getdiskSize(vms.disk,vms.id)" class="py-2 px-2 cursor-pointer hover:bg-danger hover:text-white"><i class="feather icon-repeat mr-4"></i>Resize RAM/vCPU</p>
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
    <!-- ///////////////////////////
         tabs 
        ////////////////////////////  -->
    <vx-card class="mt-5">
      <div class="vx-row">
        <div class="vx-col w-full">
          <vs-tabs alignment="fixed">
            <vs-tab label="Snapshots">
              <Snapshots />
            </vs-tab>
            <!-- <vs-tab label="Storage">
                            Storage
                        </vs-tab> -->
          </vs-tabs>
        </div>
      </div>
    </vx-card>
    <!-- ////////////////////////////////
         popup for resize disk 
         ///////////////////////////////-->
    <vs-popup background-color="rgba(152,152,152,.7)" class="" title="Resize disk" :active.sync="ResizeDisk">
      <mark>** Please make sure vm is powered off before applying this action **</mark>
      <form>
        <div class="vx-row">
          <div class="vx-col md:w-1/2 w-full mt-5 p-3">
            <p>Current disk size</p>
          </div>
          <div class="vx-col md:w-1/2 w-full mt-5">
            <vs-input v-model="diskSize" class="w-full" name="previousSize" autocomplete="off" :value="diskSize" readonly />
          </div>
        </div>
        <div class="vx-row">
          <div class="vx-col md:w-1/2 w-full mt-5 p-3">
            <p>Enter resize value</p>
          </div>
          <div class="vx-col md:w-1/2 w-full mt-5">
            <!-- <vs-input-number v-model="newSize" min="30" max="120"  label="Size in GB:"/> -->
            <vs-input v-model="newDiskSize" class="w-full" name="newSize" autocomplete="off" label-placeholder="must be greater then current size" />
          </div>
        </div>
        <vs-button name="submit" @click="resizeVM('primary',name='disk-resize')" color="danger" type="filled">Resize Disk</vs-button>
      </form>
    </vs-popup>

    <!-- //////////////////////////////////////
          popup for resize memory and vcpu
        ///////////////////////////////////////   -->
    <vs-popup background-color="rgba(152,152,152,.7)" title="Resize RAM and VCPU" :active.sync="ResizeMemory">
      <mark>** Please make sure vm is powered off before performing this action **</mark>
      <form>
        <div class="vx-row">
          <div class="vx-col md:w-1/2 w-full mt-5 p-3">
            <p>Enter resize value for RAM</p>
          </div>
          <div class="vx-col md:w-1/2 w-full mt-5">
            <!-- <vs-input-number v-model="newSize" min="30" max="120"  label="Size in GB:"/> -->
            <vs-input v-model="newRAMSize" class="w-full" name="newSize" autocomplete="off" label-placeholder="must be greater then current size" />
          </div>
        </div>
        <div class="vx-row">
          <div class="vx-col md:w-1/2 w-full mt-5 p-3">
            <p>Enter resize value for VCPU</p>
          </div>
          <div class="vx-col md:w-1/2 w-full mt-5">
            <!-- <vs-input-number v-model="newSize" min="30" max="120"  label="Size in GB:"/> -->
            <vs-input v-model="newVCPUSize" class="w-full" name="newSize" autocomplete="off" label-placeholder="must be greater then current size" />
          </div>
        </div>
        <vs-button name="submit" @click="resizeMemory('primary',name='resize')" color="danger" type="filled">Resize</vs-button>
      </form>
    </vs-popup>
    <!-- /////////////////////////////////////
         Pop up for Rename vm 
        /////////////////////////////////////  -->
    <vs-popup class="holamundo" title="Rename VM" :active.sync="changeName">
      <p>Current VM name :<strong>{{ vmName }}</strong></p>
      <form>
        <vs-input v-model="Rename" class="w-full" name="newName" autocomplete="off" label-placeholder="enter new name for vm" />
        <div class="text-center mt-5 p-3">
          <vs-button name="submit" @click="renameVM()" color="primary" type="border">Rename VM</vs-button>
        </div>
      </form>
    </vs-popup>
    <!-- /////////////////////////////////
         Popup for Delete VM
        /////////////////////////////////  -->
    <vs-popup class="holamundo" title="Confirmation" :active.sync="confirmDelete">
      <strong>Are you sure you want to delete selected VM</strong>
      <form>
        <div class="text-center d-flex mt-5 p-3">
          <vs-button name="submit" @click="DeleteVm()" color="primary" type="border">Confirm</vs-button>
          <vs-button @click="confirmDelete=false" class="mx-5" color="dark" type="border">Cancel</vs-button>
        </div>
      </form>
    </vs-popup>
  </div>
</template>


<script>
/* eslint-disable */
import axios from 'axios'
import Snapshots from './Snapshots.vue'
export default{
    data() {
        return {
            vms: {},
            error: null,
            statusobj: null,
            diskSize:null,
            vmid: null,
            vmName:null,
            Rename: '',
            newDiskSize: 40,
            newRAMSize: null,
            newVCPUSize: null,
            backgroundLoading:'primary',
            colorLoading:'#fff',
            ResizeDisk: false,
            ResizeMemory: false,
            changeName: false, 
            confirmDelete: false,
        }
    },
    created() {

        ////////////////////////
        // authenticated or not
        // ////////////////////
        if(localStorage.getItem('token')==null)
        {
        this.$router.replace('/').catch(() => {})  
        } 

       /////////////////////////////////////
       ///api call for selected vm
       ////////////////////////////////////
        this.$emit('changeRouteTitle', 'VM Status');
        let temp_id = this.$route.params.id;
        console.log('The id is: ' + this.$route.params.id);
        const tokenurl = '/api/vms/'+temp_id;
                axios({
                    method: 'get',
                    url: tokenurl,
                })
                .then((response) => {
                  this.vms =  response.data;
                  console.log(response.data);
                })
        // const monitorurl = '/api/vms/'+temp_id+'/monitor';
        //         axios({
        //             method: 'get',
        //             url: monitorurl,
        //         })
        //         .then((response) => {
        //           this.monitordata =  response.data;
        //           console.log(response.data);
        //         })        

    },

    methods:{
         ///////////////////////////////////////////////////////////////////////////
         // API call for performing action like poweroff, reboot,suspend,resume etc
         // ///////////////////////////////////////////////////////////////////////
        actionvm(id, name) 
        {
          let temp_id = id;
          var actiondata = new FormData();
          actiondata.set('type', name);
          actiondata.set('name', name);
          const tokenurl = '/api/vms/'+temp_id+'/actions';
            axios({
              method: 'post',
              url: tokenurl,
              data: actiondata,
            })
            .then((response) => {
              this.$vs.notify({
              color: "primary",
              title: 'Server reponse',
              text: response.data.message,
            })
            }) 
            .catch((error) => {
              this.$vs.notify({  
              color: 'danger',
              title: "Status",
              text: error.response.data.error
            })  
            })
        },
        /////////////////////////////
        // API call for vms refresh
        // //////////////////////////
        acceptAlert(card) 
        { 
          let temp_id = this.$route.params.id;  
          const tokenurl = '/api/vms/'+temp_id;
            axios({
              method: 'get',
              url: tokenurl,
            })
            .then((response) => {
		    	  this.vms =  response.data;
			      card.removeRefreshAnimation(3000);
            })
            this.$vs.notify({
                color: 'primary',
                title: 'Success',
                text: 'VMs data is refreshed successfully'
            })
        },
        ///////////////////////////////////////
        // for getting name of selected vm
        // //////////////////////////////////
        getVmName(name,id)
        {
         this.vmName = name;
         this.vmid = id;
        },
        /////////////////////////////////////////////////
         // For getting size of disk of each vm on click
         // //////////////////////////////////////////////
        getdiskSize(size,id)
        {
         this.diskSize = size/1024+"GB";
         this.vmid = id;
        },
        ////////////////////////////
        // fucntion for resize disk
        // //////////////////////////
        resizeVM(color,name){
          this.colorAlert = color;
          let temp_id = this.vmid;
          var actiondata = new FormData();
          let resizeValue = this.newDiskSize;
          let sendValue = resizeValue*1024;
          actiondata.set('type', name);
          actiondata.set('disk', sendValue);
          const tokenurl = '/api/vms/'+temp_id+'/actions';
            axios({
              method: 'post',
              url: tokenurl,
              data: actiondata,
            })
            .then((response) => {
              this.$vs.notify({
              color: "primary",
              title: 'Server reponse',
              text: response.data.message
            })
            })  
            .catch((error) => {
              this.$vs.notify({  
              color: 'danger',
              title: "Status",
              text: error.response.data.error
            }) 
            }),
             this.ResizeDisk = false;
        },
        ///////////////////////////////////////
        // fucntion for resize Memory/vcpu
        // ///////////////////////////////////
          resizeMemory(color,name)
          {
            this.colorAlert = color;
            let temp_id = this.vmid;
            var actiondata = new FormData();
            let resizeRamValue = this.newRAMSize;
            let resizeVcpuValue = this.newVCPUSize;
            let sendRamValue = resizeRamValue*1024;
            let sendVcpuValue = resizeVcpuValue;
            actiondata.set('type', name);
            actiondata.set('memory', sendRamValue);
            actiondata.set('cpu', sendVcpuValue);
            const tokenurl = '/api/vms/'+temp_id+'/actions';
              axios({
                method: 'post',
                url: tokenurl,
                data: actiondata,
              })
              .then((response) => {
                this.$vs.notify({
                color: "primary",
                title: 'Server reponse',
                text: response.data.message
              }) 
              })  
              .catch((error) => {
                this.$vs.notify({  
                color: 'danger',
                title: "Status",
                text: error.response.data.error
              })
              }),
               this.ResizeMemory = false;
            },
        ////////////////////////////////    
        // fucntion for renaming the vm
        // /////////////////////////////
        renameVM() 
        {
          let temp_id = this.vmid;
          var actiondata = new FormData();
		      actiondata.set('type', 'rename');
		      actiondata.set('name', this.Rename);
          const tokenurl = '/api/vms/'+temp_id+'/actions';
            axios({
              method: 'post',
              url: tokenurl,
              data: actiondata,
            })
            .then((response) => {
              this.$vs.notify({
              color: "primary",
              title: 'Server reponse',
              text: response.data.message
            })

          })
          .catch((error) => {
              this.$vs.notify({  
              color: 'danger',
              title: "Status",
              text: error.response.data.error
            })   
            }),
            this.changeName = false;
        },
        ///////////////////////////////////////
        // for getting vm id for confirmDelete
        // //////////////////////////////////////
        getVmId(id){
         this.vmid = id;
        },
        ////////////////////////////////////////////
        // API cal for deleting vm with notification
        ////////////////////////////////////////////
        DeleteVm() 
        {
            let temp_id = this.vmid;
            const tokenurl = '/api/vms/'+temp_id;
            axios({
              method: 'delete',
              url: tokenurl,
             })
            .then((response) => {
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
                text: error.response.data.error
                }) 
            })
            this.confirmDelete = false;
        }
    },
    components: {
        Snapshots,
    },
}
</script>
