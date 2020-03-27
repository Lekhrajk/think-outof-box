<!-- =========================================================================================
    File Name: AgGridTable.vue
    Description: Ag Grid table
    ----------------------------------------------------------------------------------------
    Item Name: Vuexy - Vuejs, HTML & Laravel Admin Dashboard Template
    Author: Pixinvent
    Author URL: http://www.themeforest.net/user/pixinvent
========================================================================================== -->


<template>
  <div id="ag-grid-demo">
    <vx-card>

      <!-- TABLE ACTION ROW -->
      <div class="flex flex-wrap justify-between items-center">

        <!-- ITEMS PER PAGE -->
        <div class="mb-4 md:mb-0 mr-4 ag-grid-table-actions-left">
          <vs-dropdown vs-trigger-click class="cursor-pointer">
            <div class="p-4 border border-solid d-theme-border-grey-light rounded-full d-theme-dark-bg cursor-pointer flex items-center justify-between font-medium">
              <span class="mr-2">{{ currentPage * paginationPageSize - (paginationPageSize - 1) }} - {{ contacts.length - currentPage * paginationPageSize > 0 ? currentPage * paginationPageSize : contacts.length }} of {{ contacts.length }}</span>
              <feather-icon icon="ChevronDownIcon" svgClasses="h-4 w-4" />
            </div>
            <!-- <vs-button class="btn-drop" type="line" color="primary" icon-pack="feather" icon="icon-chevron-down"></vs-button> -->
            <vs-dropdown-menu>

              <vs-dropdown-item @click="gridApi.paginationSetPageSize(20)">
                <span>20</span>
              </vs-dropdown-item>
              <vs-dropdown-item @click="gridApi.paginationSetPageSize(50)">
                <span>50</span>
              </vs-dropdown-item>
              <vs-dropdown-item @click="gridApi.paginationSetPageSize(100)">
                <span>100</span>
              </vs-dropdown-item>
              <vs-dropdown-item @click="gridApi.paginationSetPageSize(150)">
                <span>150</span>
              </vs-dropdown-item>
            </vs-dropdown-menu>
          </vs-dropdown>
        </div>

        <!-- TABLE ACTION COL-2: SEARCH & EXPORT AS CSV -->
        <div class="flex flex-wrap items-center justify-between ag-grid-table-actions-right">
          <vs-input class="mb-4 md:mb-0 mr-4" v-model="searchQuery" @input="updateSearchQuery" placeholder="Search..." />
          <vs-button class="mb-4 md:mb-0" @click="gridApi.exportDataAsCsv()">Export as CSV</vs-button>
        </div>
      </div>
      <ag-grid-vue
        ref="agGridTable"
        :gridOptions="gridOptions"
        class="ag-theme-material w-100 my-4 ag-grid-table"
        :columnDefs="columnDefs"
        :defaultColDef="defaultColDef"
        :rowData="contacts"
        rowSelection="multiple"
        colResizeDefault="shift"
        :animateRows="true"
        :floatingFilter="true"
        :pagination="true"
        :paginationPageSize="paginationPageSize"
        :suppressPaginationPanel="true"
        :enableRtl="$vs.rtl">
      </ag-grid-vue>
      <vs-pagination
        :total="totalPages"
        :max="maxPageNumbers"
        v-model="currentPage" />

    </vx-card>
  </div>
</template>

<script>
import { AgGridVue } from 'ag-grid-vue'
import contacts from './data.json'

import '@/assets/scss/vuexy/extraComponents/agGridStyleOverride.scss'

export default {
  components: {
    AgGridVue
  },
  data () {
    return {
      searchQuery: '',
      gridOptions: {},
      maxPageNumbers: 7,
      gridApi: null,
      defaultColDef: {
        sortable: true,
        editable: true,
        resizable: true,
        suppressMenu: true
      },
      columnDefs: [
        {
          headerName: 'First Name',
          field: 'firstname',
          width: 175,
          filter: true,
          checkboxSelection: true,
          headerCheckboxSelectionFilteredOnly: true,
          headerCheckboxSelection: true
        },
        {
          headerName: 'Last Name',
          field: 'lastname',
          filter: true,
          width: 175
        },
        {
          headerName: 'Email',
          field: 'email',
          filter: true,
          width: 250,
          pinned: 'left'
        },
        {
          headerName: 'Company',
          field: 'company',
          filter: true,
          width: 250
        },
        {
          headerName: 'City',
          field: 'city',
          filter: true,
          width: 150
        },
        {
          headerName: 'Country',
          field: 'country',
          filter: true,
          width: 150
        },
        {
          headerName: 'State',
          field: 'state',
          filter: true,
          width: 125
        },
        {
          headerName: 'Zip',
          field: 'zip',
          filter: true,
          width: 125
        },
        {
          headerName: 'Followers',
          field: 'followers',
          filter: 'agNumberColumnFilter',
          width: 125
        }
      ],
      contacts
    }
  },
  watch: {
    '$store.state.windowWidth' (val) {
      if (val <= 576) {
        this.maxPageNumbers = 4
        this.gridOptions.columnApi.setColumnPinned('email', null)
      } else this.gridOptions.columnApi.setColumnPinned('email', 'left')
    }
  },
  computed: {
    paginationPageSize () {
      if (this.gridApi) return this.gridApi.paginationGetPageSize()
      else return 50
    },
    totalPages () {
      if (this.gridApi) return this.gridApi.paginationGetTotalPages()
      else return 0
    },
    currentPage: {
      get () {
        if (this.gridApi) return this.gridApi.paginationGetCurrentPage() + 1
        else return 1
      },
      set (val) {
        this.gridApi.paginationGoToPage(val - 1)
      }
    }
  },
  methods: {
    updateSearchQuery (val) {
      this.gridApi.setQuickFilter(val)
    }
  },
  mounted () {
    this.gridApi = this.gridOptions.api

    /* =================================================================
      NOTE:
      Header is not aligned properly in RTL version of agGrid table.
      However, we given fix to this issue. If you want more robust solution please contact them at gitHub
    ================================================================= */
    if (this.$vs.rtl) {
      const header = this.$refs.agGridTable.$el.querySelector('.ag-header-container')
      header.style.left = `-${  String(Number(header.style.transform.slice(11, -3)) + 9)  }px`
    }
  }
}

</script>
