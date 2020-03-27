<!-- =========================================================================================
  File Name: Dashboard
  Description: Dashboard
  ----------------------------------------------------------------------------------------
  Item Name: Vuexy - Vuejs, HTML & Laravel Admin Dashboard Template
  Author: Xongl Cloud private limited
========================================================================================== -->

<template>
  <div id="dashboard-analytics">    
    <div class="vx-row">
      <!-- CARD List of  alll vms -->
      <div class="vx-col w-full">
        <vx-card title="List of Deployed VMs" refresh-content-action @refresh="acceptAlert">
          <div slot="no-body" class="mt-4" id="div-with-loading">
            <vs-table class="table-dark-inverted" ref="vmstable" v-if="vms">
              <template slot="thead">
                <vs-th>ID</vs-th>
                <vs-th>STATUS</vs-th>
                <vs-th>HOSTNAME</vs-th>
                <vs-th>OS</vs-th>
                <vs-th>Disk</vs-th>
                <vs-th>IP's</vs-th>
                <vs-th>Action</vs-th>
              </template>

              <template>
                <vs-tr v-for="vm in vms.vms" :key="vm.id">
                  <vs-td>
                    <span></span>
                  </vs-td>
                  <vs-td>
                    <span>{{vm.state}}</span>
                  </vs-td>
                  <vs-td>
                    <span>{{ vm.name }}</span>
                  </vs-td>
                  <vs-td>
                        <!-- <vs-avatar :src="vm.logo" size="30px" class="border-2 border-white border-solid -m-1">
                        </vs-avatar> -->
                         <img :src="vm.logo" alt="vm" width="60" height="40">
                  </vs-td>
                  <vs-td>
                    <span>{{ vm.disk/1024 }}GB</span>
                  </vs-td>
                  <vs-td>
                    <span>{{vm.ipv4}}</span>
                  </vs-td>
                  <vs-td>
                    <div class="dropdown-button-container">
                        <vs-dropdown class="lg:w-1/3 xl:w-1/3 my-0 p-0">
                            <vs-button class="btn-drop" color="danger" type="flat" icon="more_horiz"></vs-button>
                            <vs-dropdown-menu class="lg:w-1/4">
                              <div class="vx-row">
                                <div class="vx-col sm:w-full md:w-3/6 lg:w-1/2 xl:w-1/2 mx-auto">
                                  <vs-dropdown-item @click="actionvm(vm.id, name='reboot-hard')"  name="reboot-hard" value="reboot-hard"><p class="text-left py-2 px-2 cursor-pointer hover:bg-primary hover:text-white"><i class="feather icon-loader mr-4 "></i>Reboot</p></vs-dropdown-item>
                                  <vs-dropdown-item  @click="actionvm(vm.id, name='poweroff-hard')"  name="poweroff-hard" value="poweroff-hard"><p class="text-left py-2 px-2 cursor-pointer hover:bg-primary hover:text-white" ><i class="feather icon-power mr-4 "></i>PowerOff</p></vs-dropdown-item>
                                  <vs-dropdown-item @click="actionvm(vm.id, name='resume')"  name="resume" value="resume" ><p class="py-2 px-2 cursor-pointer hover:bg-primary hover:text-white" ><i class="feather icon-play-circle mr-4"></i>Resume</p></vs-dropdown-item>
                                  <vs-dropdown-item  @click="changeName=true" ><p  @click="getVmName(vm.name,vm.id)" class="py-2 px-2 cursor-pointer hover:bg-primary hover:text-white"><i class="feather icon-file-text mr-4"></i>Rename</p></vs-dropdown-item>
                                </div>
                                <div class="vx-col sm:w-full md:w-3/6 lg:w-1/2 xl:w-1/2 mx-auto">
                                  <vs-dropdown-item> 
                                    <p @click="confirmDelete=true"  color="danger" type="flat" class="text-left py-2 px-2 cursor-pointer hover:bg-danger hover:text-white" ><span  @click="getVmId(vm.id)"><i class="feather icon-trash mr-4"></i> Delete</span></p>
                                  </vs-dropdown-item >
                                  <vs-dropdown-item @click="ResizeDisk=true" type="filled" >
                                    <p  @click="getdiskSize(vm.disk,vm.id)" class=" text-left py-2 px-2 cursor-pointer hover:bg-danger hover:text-white"><i class="feather icon-maximize-2 mr-4"></i> Resize disk</p>
                                  </vs-dropdown-item>
                                  <vs-dropdown-item @click="ResizeMemory=true" type="filled" >
                                    <p  @click="getdiskSize(vm.disk,vm.id)" class="py-2 px-2 cursor-pointer hover:bg-danger hover:text-white"><i class="feather icon-repeat mr-4"></i>Resize RAM/vCPU</p>
                                  </vs-dropdown-item>
                                  <vs-dropdown-item>
                                    <vs-button color="dark" type="filled" @click="$router.replace({ path:'/select-template/category/vmstatus/'+vm.id })"><i class="feather icon-square mr-4"></i>More</vs-button>
                                  </vs-dropdown-item>
                                
                                </div>
                              </div>
                          </vs-dropdown-menu>
                        </vs-dropdown>
                    </div>
                  </vs-td>
                </vs-tr>
              </template>
            </vs-table>
          </div>
        </vx-card>
      </div>
    </div>
    <!-- popup for resize disk -->
    <vs-popup background-color="rgba(152,152,152,.7)" class=""  title="Resize disk" :active.sync="ResizeDisk">
      <mark>** Please make sure vm is powered off before applying this action **</mark>
      <form>
        <div class="vx-row">
          <div class="vx-col md:w-1/2 w-full mt-5 p-3">
              <p>Current disk size</p>
          </div>
          <div class="vx-col md:w-1/2 w-full mt-5">
            <vs-input  v-model="diskSize" class="w-full" name="previousSize"  autocomplete="off" :value = "diskSize" readonly/>
          </div>
        </div>
        <div class="vx-row">
          <div class="vx-col md:w-1/2 w-full mt-5 p-3">
              <p>Enter resize value</p>
          </div>
          <div class="vx-col md:w-1/2 w-full mt-5">
            <!-- <vs-input-number v-model="newSize" min="30" max="120"  label="Size in GB:"/> -->
            <vs-input  v-model="newDiskSize" class="w-full" name="newSize"  autocomplete="off" label-placeholder="must be greater then current size" />
          </div>
        </div>
        <vs-button name="submit"  @click="resizeVM('primary',name='disk-resize')"  color="danger" type="filled">Resize Disk</vs-button>
      </form>  
    </vs-popup>
    <!-- popup for resize memory and vcpu -->
    <vs-popup background-color="rgba(152,152,152,.7)" title="Resize RAM and VCPU" :active.sync="ResizeMemory">
      <mark>** Please make sure vm is powered off before performing this action **</mark>
      <form>
        <div class="vx-row">
          <div class="vx-col md:w-1/2 w-full mt-5 p-3">
              <p>Enter resize value for RAM</p>
          </div>
          <div class="vx-col md:w-1/2 w-full mt-5">
            <!-- <vs-input-number v-model="newSize" min="30" max="120"  label="Size in GB:"/> -->
            <vs-input  v-model="newRAMSize" class="w-full" name="newSize"  autocomplete="off" label-placeholder="must be greater then current size" />
          </div>
        </div>
        <div class="vx-row">
          <div class="vx-col md:w-1/2 w-full mt-5 p-3">
              <p>Enter resize value for VCPU</p>
          </div>
          <div class="vx-col md:w-1/2 w-full mt-5">
            <!-- <vs-input-number v-model="newSize" min="30" max="120"  label="Size in GB:"/> -->
            <vs-input  v-model="newVCPUSize" class="w-full" name="newSize"  autocomplete="off" label-placeholder="must be greater then current size" />
          </div>
        </div>
        <vs-button name="submit"  @click="resizeMemory('primary',name='resize')"  color="danger" type="filled">Resize</vs-button>
      </form>  
    </vs-popup>
    <vs-popup class="holamundo" title="Rename VM" :active.sync="changeName">
      <p>Current VM name :<strong>{{ vmName }}</strong></p>
      <form>
         <vs-input  v-model="Rename" class="w-full" name="newName"  autocomplete="off" label-placeholder="enter new name for vm" />
         <div class="text-center mt-5 p-3">
           <vs-button name="submit"  @click="renameVM()"   color="primary" type="border">Rename VM</vs-button>
         </div>
      </form>
    </vs-popup>
    <vs-popup class="holamundo" title="Confirmation" :active.sync="confirmDelete">
      <strong>Are you sure you want to delete selected VM</strong>
      <form>
         <div class="text-center d-flex mt-5 p-3">
           <vs-button name="submit"  @click="DeleteVm()"   color="primary" type="border">Confirm</vs-button>
           <vs-button @click="confirmDelete=false"  class="mx-5" color="dark" type="border">Cancel</vs-button>
         </div>
      </form>
    </vs-popup>
  </div>
</template>

<script>
/* eslint-disable */
import axios from 'axios';

export default {
    data() {
        return {
            vms: {},
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
    mounted() {
      if(localStorage.getItem('name')=='' || localStorage.getItem('pw') ==''){
        this.$router.replace('/').catch(() => {})  
      } 
    },
    created(){
        //API call for list of vms  
      const tokenurl = '/vms';
        axios({
            method: 'get',
             url: tokenurl,
        })
        .then((response) => {
          this.vms =  response.data;
          console.log(response.data);
        })
        .catch((error) => {
              this.$vs.notify({  
              color: 'danger',
              title: "Status",
              text: error.response.data.error
            }) 
            if(error.response.data.error =="Missing cookie or authorization header!"){
              this.$router.replace('/').catch(() => {}) 
            }
          })  
    },
    methods:{
      // API call for vms refresh
        acceptAlert(card) 
        {
          const tokenurl = '/vms';
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
        // for getting name of selected vm
        getVmName(name,id)
        {
         this.vmName = name;
         this.vmid = id;
        },
         // For getting size of disk of each vm on click
        getdiskSize(size,id)
        {
         this.diskSize = size/1024+"GB";
         this.vmid = id;
        },
        // fucntion for resize disk
        resizeVM(color,name){
          this.colorAlert = color;
          let temp_id = this.vmid;
          var actiondata = new FormData();
          let resizeValue = this.newDiskSize;
          let sendValue = resizeValue*1024;
          actiondata.set('type', name);
          actiondata.set('disk', sendValue);
          const tokenurl = '/vms/'+temp_id+'/actions';
            axios({
              method: 'post',
              url: tokenurl,
              data: actiondata,
            })
            .then((response) => {
              this.$vs.notify({
              color: "warning",
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
            this.$vs.loading({
              container: '#div-with-loading',
              scale: 0.6
            })
            setTimeout( ()=> {
              this.$vs.loading.close('#div-with-loading > .con-vs-loading')
            }, 3000);
        },
                // fucntion for resize Memory/vcpu
          resizeMemory(color,name){
          this.colorAlert = color;
          let temp_id = this.vmid;
          var actiondata = new FormData();
          let resizeRamValue = this.newRAMSize;
          let resizeVcpuValue = this.newVCPUSize;
          let sendRamValue = resizeRamValue*1024;
          let sendVcpuValue = resizeVcpuValue*1024;
          actiondata.set('type', name);
          actiondata.set('memory', sendRamValue);
          actiondata.set('cpu', sendVcpuValue);
          const tokenurl = '/vms/'+temp_id+'/actions';
            axios({
              method: 'post',
              url: tokenurl,
              data: actiondata,
            })
            .then((response) => {
              this.$vs.notify({
              color: "warning",
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
            this.$vs.loading({
              container: '#div-with-loading',
              scale: 0.6
            })
            setTimeout( ()=> {
              this.$vs.loading.close('#div-with-loading > .con-vs-loading')
            }, 3000);
        },
        // API call for performing action like poweroff, reboot,suspend,resume etc
        actionvm(id, name) 
        {
          let temp_id = id;
          var actiondata = new FormData();
          actiondata.set('type', name);
          actiondata.set('name', name);
          const tokenurl = '/vms/'+temp_id+'/actions';
            axios({
              method: 'post',
              url: tokenurl,
              data: actiondata,
            })
            .then((response) => {
              this.$vs.notify({
              color: "dark",
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
            }),
            this.$vs.loading({
              container: '#div-with-loading',
              scale: 0.6
            })
            setTimeout( ()=> {
              this.$vs.loading.close('#div-with-loading > .con-vs-loading')
            }, 3000);
        },
        // fucntion for renaming the vm
        renameVM() 
        {
          let temp_id = this.vmid;
          var actiondata = new FormData();
		      actiondata.set('type', 'rename');
		      actiondata.set('name', this.Rename);
          const tokenurl = '/vms/'+temp_id+'/actions';
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
            this.$vs.loading({
              container: '#div-with-loading',
              scale: 0.6
            })
            setTimeout( ()=> {
              this.$vs.loading.close('#div-with-loading > .con-vs-loading')
            }, 3000);
        },
        // for getting vm id for confirmDelete
        getVmId(id){
         this.vmid = id;
        },
        // API cal for deleting vm with notification
        DeleteVm() 
        {
            let temp_id = this.vmid;
            const tokenurl = '/vms/'+temp_id;
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
}
</script>

<style lang="scss">
/*! rtl:begin:ignore */
#dashboard-analytics {
  .greet-user{
    position: relative;

    .decore-left{
      position: absolute;
      left:0;
      top: 0;
    }
    .decore-right{
      position: absolute;
      right:0;
      top: 0;
    }
  }

  @media(max-width: 576px) {
    .decore-left, .decore-right{
      width: 140px;
    }
  }
}
body{
  counter-reset: Serial;    
}

tr td:first-child:before
{
  counter-increment: Serial;      /* Increment the Serial counter */
  content: counter(Serial); /* Display the counter */
}
/*! rtl:end:ignore */
</style>