<!-- =========================================================================================
    File Name: KnowledgeBaseCategoryQuestion.vue
    Description: Knowledge Base Question - Answer Article
    ----------------------------------------------------------------------------------------
    Item Name: Vuexy - Vuejs, HTML & Laravel Admin Dashboard Template
      Author: Pixinvent
    Author URL: http://www.themeforest.net/user/pixinvent
========================================================================================== -->
<template>
  <div id="VMs-Snapshots">
    <div class="vx-row">
      <!-- CARD 9: DISPATCHED ORDERS -->
      <div class="vx-col w-full">
        <vx-card title="List of Snapshots">
          <div slot="no-body" class="mt-4 px-3">
            <div class="float-right">
               <vs-button @click="RefreshSnapshotlist()"  color="primary" type="flat" cursor="pointer" class="mx-3" size="small"><i class="feather icon-rotate-cw mx-3"></i>Refresh</vs-button>
               <vs-button @click="popupActive=true"  color="success" type="border" cursor="pointer" size="small"><i class="feather icon-camera mx-3"></i>Take Snapshot</vs-button>
            </div>
            <vs-table class="table-dark-inverted" ref="snapshotstable">
              <template slot="thead">
                <vs-th>ID</vs-th>
                <vs-th>NAME</vs-th>
                <vs-th>TIME</vs-th>
                <vs-th>Action</vs-th>
              </template>
              <template>
                <vs-tr v-for="snapshot in snapshotslist.snapshots" :key="snapshot.id">
                  <vs-td>
                    <span></span>
                  </vs-td>
                  <vs-td>
                    <span>{{ snapshot.name }}</span>
                  </vs-td>
                  <vs-td>
                    <span>{{snapshot.time}}</span>
                  </vs-td>
                  <vs-td>
                    <div class="dropdown-button-container">
                      <vs-dropdown class="flex items-center">
                        <vs-button class="btn-drop" color="danger" type="flat" icon="more_horiz"></vs-button>
                        <vs-dropdown-menu>
                          <vs-dropdown-item>
                          <vs-button color="primary" type="filled" @click="Revertsnapshot(snapshot.id, name='revert')">Revert</vs-button>
                          </vs-dropdown-item>
                          <vs-dropdown-item>
                            <vs-button color="danger" type="filled"  @click="DeleteSnapshot(snapshot.id)" >Delete</vs-button>
                          </vs-dropdown-item>
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
    <vs-popup class="holamundo text-primary" title="Create Snapshot" :active.sync="popupActive">
      <form  method="post" @submit.prevent="CreateSnapshot">
        <div class="vx-row">
            <div class="vx-col md:w-1/2 w-full mt-5">
              <p class="text-dark">Enter Snapshot Name</p>
            </div>
            <div class="vx-col md:w-1/2 w-full mt-5">
              <vs-input  v-model="SnapshotName" class="w-full" name="snapshot_name"  autocomplete="off"/>
            </div>
        </div>
        <div class="text-center my-4 p-3">
            <vs-button color="danger" type="filled" @click="CreateSnapshot()">Take Snapshot</vs-button>
        </div>
      </form>
    </vs-popup>
  </div>
</template>


<script>
import axios from 'axios';
export default{
    data() {
        return {
          snapshotslist: {},
          popupActive: false,
          passedid: null,
          SnapshotName: null,

        }
    },
    mounted() {
        this.Snapshotlist();
    },

    methods: {
        CreateSnapshot() 
        {
          let temp_id = this.$route.params.id;
          var snapshotdata = new FormData();
          snapshotdata.set('name', this.SnapshotName);
          const tokenurl = '/api/vms/'+temp_id+'/snapshots';
            axios({
              method: 'post',
              url: tokenurl,
              data: snapshotdata
            })
            .then((response) => {
              this.$vs.notify({  
              color: 'primary',
              title: "Status",
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
              this.popupActive = false;  
        },
        DeleteSnapshot(snapid) 
        {
          let snap_id = snapid
          let temp_id = this.$route.params.id
          const tokenurl = '/api/vms/'+temp_id+'/snapshots/'+snap_id;
            axios({
              method: 'delete',
              url: tokenurl,
            })
            .then((response) => {
              this.$vs.notify({  
              color: 'danger',
              title: "Status",
              text: response.data.message,
              }) 
            })
                this.popupActive = false;
        },
        Snapshotlist()
        {
          let temp_id = this.$route.params.id;
          const tokenurl = '/api/vms/'+temp_id+'/snapshots';
          axios({
              method: 'get',
              url: tokenurl,
          })
          .then((response) => {
            this.snapshotslist =  response.data;
          })
        },
        RefreshSnapshotlist()
        {
          let temp_id = this.$route.params.id;
          const tokenurl = '/api/vms/'+temp_id+'/snapshots';
          axios({
              method: 'get',
              url: tokenurl,
          })
          .then((response) => {
            this.snapshotslist =  response.data;
          })
          this.$vs.notify({  
          color: 'primary',
          title: "Status",
          text: "Snapshotlist is refreshed",
          }) 
        },
        Revertsnapshot(snapid,action){
          let snap_id = snapid;
          var actiondata = new FormData();
          actiondata.set('type', action);
          let temp_id = this.$route.params.id
          const tokenurl = '/api/vms/'+temp_id+'/snapshots/'+snap_id;
            axios({
              method: 'post',
              url: tokenurl,
              data: actiondata,
            })
            .then((response) => {
              this.$vs.notify({  
              color: 'danger',
              title: "Status",
              text: response.data.message,
              }) 
            })
        },
        silentRefresh(){
          this.Snapshotlist();
        }
    },

    // created() {
    //   this.interval = setInterval(() => this.silentRefresh(), 5000);
    // },
}
</script>
