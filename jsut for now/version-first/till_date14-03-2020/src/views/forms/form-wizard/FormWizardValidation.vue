<!-- =========================================================================================
  File Name: FormWizardValidation.vue
  Description: Form wizard with validation
  ----------------------------------------------------------------------------------------
  Item Name: Xongl user control panel
  Author: Xongl cloud private limited
========================================================================================== -->


<template>
  <vx-card title="CREATE VIRTUAL MACHINE">
    <div class="mt-5"  id="div-with-loading">
          <!-- tab 1 content -->
        <form id="createvm" method="post" @submit.prevent="deploy">
          <div class="vx-row py-4">
            <div class="vx-col md:w-1/3 w-full mt-5 p-3">
              <p>Enter Virtual Machine Name</p>
            </div>
            <div class="vx-col md:w-2/3 w-full mt-5">
              <vs-input  v-model="vmName" class="w-full" name="vm_name" v-validate="'required'" autocomplete="off"/>
              <span class="text-danger">{{ error }}</span>
            </div>
          </div>
          <div class="vx-row pb-4">
            <div class="vx-col md:w-1/2 w-full">
                <div class="mt-5" title="Memory"> 
                    <label>Select Memory</label>
                    <vs-slider color="dark" text-fixed=GB ticks max="8" v-model="value1"/>
                    <p>{{ value1 }} GB-RAM</p><span class="hidden">{{ value1*1024 }}</span>
                </div>
            </div>
            <div class="vx-col md:w-1/2 w-full">
                  <div class="mt-5" title="CPU">
                      <label>Select VCPU</label>
                      <vs-slider color="dark" text-fixed=vCPU ticks max=4 v-model="value3"/>
                      <p>{{ value3 }} VCPU</p>
                  </div>
            </div>
          </div>
          <div class="vx-row py-4">
            <div class="vx-col w-full">
                <div class="mt-5" title="Storage memory">
                    <label>Select Storage(Disks)</label>
                    <vs-slider color="dark" text-fixed=GB ticks max="120" v-model="value2"/>
                    <p>{{ value2 }} GB</p><span class="hidden">{{ value2*1024 }}</span>
                </div>
            </div>
          </div>
          <div class="vx-row pb-5">
            <div class="vx-col w-full">
              <vs-tabs alignment="center">
                <vs-tab label="Normal Password">
                    <div class="con-tab-ejemplo">
                        <div class="vx-row">
                          <div class="vx-col md:w-1/4 w-full mt-5 p-3">
                            <p>Enter Password for VM</p>
                          </div>
                          <div class="vx-col md:w-3/4 w-full mt-5">
                            <vs-input type="password" v-model="password" class="w-full" name="password" v-validate="'required'" autocomplete="off"/>
                            <span class="text-danger">{{ error}}</span>
                          </div>
                        </div>
                    </div>
                </vs-tab>
                <vs-tab label="SSH Key">
                    <div class="con-tab-ejemplo">
                        <vs-textarea label="Paste your SSH here" v-model="sshKey" v-validate="'required'"/>
                    </div>
                </vs-tab>
            </vs-tabs>
            </div>
          </div>
          <div class="content-center">
            <vs-button name="submit"  id="submit"  @click="deploy('primary')"  color="primary" type="filled">Deploy VM</vs-button>
          </div>
        </form>
    </div>
  </vx-card>
</template>

<script>
import { FormWizard, TabContent } from 'vue-form-wizard'
import TabsAlignments from '../../components/vuesax/tabs/TabsAlignments.vue'
import 'vue-form-wizard/dist/vue-form-wizard.min.css'

// For custom error message
import axios from 'axios';
import router from '@/router';
import { Validator } from 'vee-validate';
export default {
    data() {
        return {
            vmName: "",
            password: "",
            sshKey: "",
            value1: 2,
            value2: 30,
            value3: 1,
            message: null,
            error: null,
        }
    },
    computed: {
       validateForm() 
      {
      if(this.input.vmName == "" && this.input.password == "")
        {
         return this.error = "vmname and password is required";
        }
      },
    },
    methods: {
            deploy(color) {
              this.colorAlert = color
              if(this.firstName != "" && this.value1 != "" && this.value2 !="" && this.value3 !="" && this.password !="") {
                const tokenurl = '/vms';
                var deploydata = new FormData();
                deploydata.set('tid', this.$route.params.id);
                deploydata.set('name', this.vmName);
                deploydata.set('cpu', this.value3);
                deploydata.set('memory', this.value1*1024);
                deploydata.set('disk', this.value2*1024);
                deploydata.set('password', this.password);
                
                axios({
                    method: 'post',
                    url: tokenurl,
                    data: deploydata,
                })
                .then((response) => {
                  this.message =  response.data.message;
                   if (response.status == 201) {
                      this.$vs.notify({
                      color: this.colorAlert,
                      title: 'Success',
                      text: "VM is created successfully",
                      })
                    this.$vs.loading({
                      container: '#div-with-loading',
                      scale: 0.6
                    })
                    setTimeout( ()=> {
                      this.$vs.loading.close('#div-with-loading > .con-vs-loading');
                      router.push({ path:"/home",name: "home" }); 
                    }, 3000);  
                  }
                })
            }
            else
            {
              this.error = "vmName and password is required";
              this.$vs.notify({
                      title:'Error',
                      text: this.error,
                      color:'danger',
                      iconPack: 'feather',
                      icon:'icon-alert-circle',
                      position:'top-right'
              })

            }
        },
    },
    components: {
        FormWizard,
        TabsAlignments,
        TabContent
    }
}
</script>