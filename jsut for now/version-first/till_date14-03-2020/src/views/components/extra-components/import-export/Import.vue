<!-- =========================================================================================
    File Name: i18n.vue
    Description: I18n demo
    ----------------------------------------------------------------------------------------
    Item Name: Vuexy - Vuejs, HTML & Laravel Admin Dashboard Template
      Author: Pixinvent
    Author URL: http://www.themeforest.net/user/pixinvent
========================================================================================== -->

<template>
  <div>
    <div class="mb-8">
      <import-excel :onSuccess="loadDataInTable" />
    </div>
    <vx-card v-if="tableData.length && header.length">
      <vs-table stripe pagination :max-items="10" search :data="tableData">
        <template slot="header">
          <h4>{{ sheetName }}</h4>
        </template>

        <template slot="thead">
          <vs-th :sort-key="heading" v-for="heading in header" :key="heading">{{ heading }}</vs-th>
        </template>

        <template slot-scope="{data}">
          <vs-tr :data="tr" :key="indextr" v-for="(tr, indextr) in data">
            <vs-td :data="col" v-for="(col, indexcol) in data[indextr]" :key="indexcol + col">
              {{ col }}
            </vs-td>
          </vs-tr>
        </template>
      </vs-table>
    </vx-card>
  </div>
</template>

<script>
import ImportExcel from "@/components/excel/ImportExcel.vue"

export default {
  components: {
    ImportExcel
  },
  data() {
    return {
      tableData: [],
      header: [],
      sheetName: ""
    }
  },
  methods: {
    loadDataInTable({ results, header, meta }) {
      this.header = header
      this.tableData = results
      this.sheetName = meta.sheetName
    }
  }
}
</script>
