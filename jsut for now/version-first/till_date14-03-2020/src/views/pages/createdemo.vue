<!-- =========================================================================================
  File Name: FormWizardValidation.vue
  Description: Form wizard with validation
  ----------------------------------------------------------------------------------------
  Item Name: Xongl user control panel
  Author: Xongl cloud private limited
========================================================================================== -->


<template>

    <div class="mt-5" id="div-with-loading">
      <form @click.prevent="preventform($event)">
        <div class="vx-row py-4">
          <div class="vx-col md:w-1/3 w-full mt-5 p-3">
            <p>Enter Virtual Machine Name</p>
          </div>
          <div class="vx-col md:w-2/3 w-full mt-5">
            <vs-input v-model="vmName" class="w-full" name="vm_name" v-validate="'required'" autocomplete="off" />
            <span class="text-danger">{{ error }}</span>
          </div>
        </div>
        <div class="vx-row pb-4">
          <div class="vx-col md:w-1/2 w-full">
            <div class="mt-5" title="Memory">
              <label>Select Memory</label>
              <vs-slider color="dark" text-fixed=GB ticks max="8" v-model="value1" />
              <p>{{ value1 }} GB-RAM</p><span class="hidden">{{ value1*1024 }}</span>
            </div>
          </div>
          <div class="vx-col md:w-1/2 w-full">
            <div class="mt-5" title="CPU">
              <label>Select VCPU</label>
              <vs-slider color="dark" text-fixed=vCPU ticks max=4 v-model="value3" />
              <p>{{ value3 }} VCPU</p>
            </div>
          </div>
        </div>
        <div class="vx-row py-4">
          <div class="vx-col w-full">
            <div class="mt-5" title="Storage memory">
              <label>Select Storage(Disks)</label>
              <vs-slider color="dark" text-fixed=GB ticks max="120" v-model="value2" v-on:keydown.enter.prevent='stopautosubmit' />
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
                      <div class="password">
                        <input :class='{valid:passwordValidation.valid}' :type="passwordVisible ? 'text' : 'password'" v-model="password">
                        <button class="visibility" tabindex='-1' @click='togglePasswordVisibility' :arial-label='passwordVisible ? "Hide password" : "Show password"'>
                          <i class="material-icons">{{ passwordVisible ? "visibility" : "visibility_off" }}</i>
                        </button>
                      </div>
                      <input type="password" v-model.lazy='checkPassword'>
                      <transition name="hint" appear>
                        <div v-if='passwordValidation.errors.length > 0 && !submitted' class='hints'>
                          <h2>Hints</h2>
                          <p v-for='error in passwordValidation.errors' :key="error">{{error}}</p>
                        </div>
                      </transition>
                      <div class="matches" v-if='notSamePasswords'>
                        <p>Passwords don't match.</p>
                      </div>
                      <!-- <vs-input type="password" v-model="password" class="w-full" name="password" v-validate="'required'" autocomplete="off"/> -->
                      <!-- <span class="text-danger">{{ error}}</span> -->
                    </div>
                  </div>
                </div>
              </vs-tab>
              <vs-tab label="SSH Key">
                <div class="con-tab-ejemplo">
                  <vs-textarea label="Paste your SSH here" v-model="sshKey" v-validate="'required'" />
                </div>
              </vs-tab>
            </vs-tabs>
          </div>
        </div>
        <div class="content-center">
          <vs-button name="submit" id="submit" @click='resetPasswords' v-if='passwordsFilled && !notSamePasswords && passwordValidation.valid' @click.prevent="deploy('primary')" color="primary" type="filled">Deploy VM</vs-button>
        </div>
      </form>
    </div>
  </vx-card>
</template>


<script>
import { FormWizard} from 'vue-form-wizard'
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
            sshKey: "",
            value1: 2,
            value2: 30,
            value3: 1,
            message: null,
            error: null,
                        rules: [
				{ message:'One lowercase letter required.', regex:/[a-z]+/ },
				{ message:"One uppercase letter required.",  regex:/[A-Z]+/ },
				{ message:"8 characters minimum.", regex:/.{8,}/ },
				{ message:"One number required.", regex:/[0-9]+/ }
			],
			password:'',
			checkPassword:'',
			passwordVisible:false,
			submitted:false
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
      		notSamePasswords () {
			if (this.passwordsFilled) {
				return (this.password !== this.checkPassword)
			} else {
				return false
			}
		},
		passwordsFilled () {
			return (this.password !== '' && this.checkPassword !== '')
		},
		passwordValidation () {
			let errors = []
			for (let condition of this.rules) {
				if (!condition.regex.test(this.password)) {
					errors.push(condition.message)
				}
			}
			if (errors.length === 0) {
				return { valid:true, errors }
			} else {
				return { valid:false, errors }
			}
		}
    },
    methods: {
            preventform(event){
             
              if (event) event.preventDefault();
              if (event) event.stopPropagation();
          },
            stopautosubmit(){
              this.error = "Please check all fields are required";
            },
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
        resetPasswords () 
        {
          this.password = ''
          this.checkPassword = ''
          this.submitted = true
          setTimeout(() => {
            this.submitted = false
          }, 2000)
        },
        togglePasswordVisibility () {
          this.passwordVisible = !this.passwordVisible
        }	
    },
    components: {
        FormWizard,
        TabsAlignments,
        TabContent
    }
}
</script>
<style lang="scss" scoped>
.hints {
	max-width:400px;
	padding:1em;
	background:whitesmoke;
	margin: 1em 0;
	font-size: .9em;
	color:darken(#D4DEDF, 50%);
	h2 {
		margin: .5em 0 .2em 0;
	}
	p {
		margin: 0;
		padding-left:1em;
		&::before {
			content:">";
			font-size:.9em;
			margin-right:6px;
			display:inline-block;
		}
	}
}

.visibility {
	all:unset;
	background:whitesmoke;
	display:inline-block;
	color:#3A81C8;
	padding:2px .4em 0;
	vertical-align:center;
}

.material-icons {
	font-size: 1.5em;
}
.password {
	display:flex;
	align-items:center;
}

input, button {
	padding:6px;
}
input + button {
	font-size: .85em;
}
[type=password], [type=text] {
	color:#D4DEDF;
	border:1px solid transparent;
	background:darken(#3A81C8, 5%);
	margin:6px 10px 6px 0;
	&.valid {
		border:1px solid whitesmoke;
	}
}
[type=password] {
	letter-spacing:.2em;
}

.hint {
	&-enter {
		opacity:0;
		transform:translate3d(-20px,0,0);
	}
	&-leave-to {
		opacity:0;
		transform:translate3d(0, 20px, 0);
	}
	&-enter-active {
		transition:300ms;
	}
	&-leave-active {
		transition:400ms;
	}
}
</style>





<!-- =========================================================================================
    File Name: CreateVM.vue
    Description: Deploy Virtual machine 
    ----------------------------------------------------------------------------------------
    Item Name: Xongl user control panel
    Author: Xongl cloud private limited
========================================================================================== -->


<template>
    <div id="deploy-vm-page">
        <vx-card title="CREATE VIRTUAL MACHINE">
            <div class="text-center cursor-pointer float-right">
                <img :src="vms.logo" alt="graphic" width="100" height="70" class="p-4">
            </div>
            <div class="mt-5" id="div-with-loading">
                <form v-on:submit.prevent ="validateForm" >
                    <div class="vx-row py-4">
                        <div class="vx-col md:w-1/3 w-full mt-5 p-3">
                            <p>Enter Virtual Machine Name</p>
                        </div>
                        <div class="vx-col md:w-2/3 w-full mt-5">
                            <vs-input placeholder="enter vm name" name="vm_name" v-model="vmName" autocomplete="off" class="w-full" v-on:keyup="validateForm"/>
                            <span class="text-danger text-sm">{{ vmNameError }}</span>
                        </div>
                    </div>
                    <div class="vx-row pb-4">
                        <div class="vx-col md:w-1/2 w-full">
                            <div class="mt-5" title="Memory">
                              <label>Select Memory</label>
                              <vs-slider color="dark" text-fixed=GB ticks max="8" v-model="selectedMemory" v-on:change="validateForm" />
                              <p>{{ selectedMemory }} GB-RAM</p><span class="hidden">{{ selectedMemory*1024 }}</span>
                            </div>
                            <span class="text-danger text-sm">{{ selectedMemoryError }}</span>
                        </div>
                        <div class="vx-col md:w-1/2 w-full">
                            <div class="mt-5" title="CPU">
                              <label>Select VCPU</label>
                              <vs-slider color="dark" text-fixed=vCPU ticks max=4 v-model="selectedVcpu" v-on:change="validateForm" />
                              <p>{{ selectedVcpu }} VCPU</p>
                            </div>
                            <span class="text-danger text-sm">{{ selectedVcpuError }}</span>
                        </div>
                    </div>
                    <vs-button type="filled" @click.prevent="submitForm" class="mt-5 text-center">Deploy VM</vs-button>
                    <div class="mb-5">
                        <vs-alert :active.sync="active1" v-show="error" color="danger" closable icon-pack="feather" close-icon="icon-x">
                          {{ error }}
                        </vs-alert>
<!--                         <vs-alert active="true" class="text-center" v-show="error">
                              {{ error }}
                        </vs-alert>  -->
                    </div>
                </form>
            </div> 
        </vx-card>
    </div>
</template>

<script>
import axios from 'axios';

export default{
    data(){
        return{ 
            vms: null,
            logo: "",
            vmName: "",
            error: null,
            vmNameError: null,
            selectedMemoryError: null,
            selectedDiskError: null,
            selectedVcpuError: null,
            active1: true,
            selectedMemory: 2,
            selectedDisk: 30,
            selectedVcpu: 1,

        }
    },
    computed: {
    },
    methods: {
        validateForm(){
            if(this.vmName == '')
            {
             this.vmNameError = "VM name is required"
            }
            else if(this.selectedMemory == 0 || this.selectedDisk == 0 || this.selectedVcpu == 0 && this.vmName == '')
            {
             this.selectedMemoryError = "Memory can't be zero ";
             this.selectedDiskError   = "Disk size must be greater than zero";
             this.selectedVcpuError   = "Vcpu can't be zero";
             this.vmNameError = "VM name is required";
            }
        },
        submitForm() 
        {
                if (this.vmName != '')
                {
                  alert("form submitted!");
                }
                else 
                {
                    this.error = "Please check all field is required";
                    this.$vs.notify({
                          title:'Error',
                          text: this.error,
                          color:'danger',
                          iconPack: 'feather',
                          icon:'icon-alert-circle',
                          position:'top-right'
                    })
                }
        }
    },
    components: {
    },
    mounted() {
        this.$emit('changeRouteTitle', 'Deploy VM');
        let temp_id = this.$route.params.id
        console.log('The id is: ' + this.$route.params.id);
        const tokenurl = '/templates/'+temp_id;
            axios({
                method: 'get',
                url: tokenurl,
            })
            .then((response) => {
                this.vms =  response.data;
                console.log(response.data);
            })
    }
}
</script>
