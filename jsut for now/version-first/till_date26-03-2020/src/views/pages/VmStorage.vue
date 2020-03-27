<!-- =========================================================================================
    File Name: KnowledgeBaseCategoryQuestion.vue
    Description: Knowledge Base Question - Answer Article
    ----------------------------------------------------------------------------------------
    Item Name: Vuexy - Vuejs, HTML & Laravel Admin Dashboard Template
      Author: Pixinvent
    Author URL: http://www.themeforest.net/user/pixinvent
========================================================================================== -->
<template>
  <div id="VMs-Storage">
    <div class="vx-row">
      <!-- CARD 9: DISPATCHED ORDERS -->
      <div class="vx-col w-full">
        <vx-card title="Storage">
          <div slot="no-body" class="mt-4">
            <div class="float-right">
               <vs-button @click="RefreshSnapshotlist()"  color="primary" type="flat" cursor="pointer">Refresh</vs-button>
               <vs-button @click="popupActive=true"  color="success" type="border" cursor="pointer">Attach</vs-button>
            </div>
            <vs-table class="table-dark-inverted" ref="storagetable">
              <template slot="thead">
                <vs-th>ID</vs-th>
                <vs-th>Target</vs-th>
                <vs-th>Image/Size-format</vs-th>
                <vs-th>Persistent</vs-th>
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
    <vs-popup class="holamundo text-primary" title="Snapshot" :active.sync="popupActive">
      <form  method="post" @submit.prevent="CreateSnapshot">
        <div class="vx-row">
            <div class="vx-col md:w-1/4 w-full mt-5 p-3">
              <p>Enter Snapshot Name</p>
            </div>
            <div class="vx-col md:w-3/4 w-full mt-5">
              <vs-input  v-model="SnapshotName" class="w-full" name="snapshot_name" v-validate="'required'" autocomplete="off"/>
            </div>
            <vs-button color="primary" type="filled" @click="CreateSnapshot()">Take Snapshot</vs-button>
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
          SnapshotName: '',

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

    created() {
      this.interval = setInterval(() => this.silentRefresh(), 5000);
    },
}
</script>
