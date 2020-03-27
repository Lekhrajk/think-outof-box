    <!-- =========================================================================================
    File Name: CardAnalytics.vue
    Description: Analytic Cards
    ----------------------------------------------------------------------------------------
    Item Name: Vuexy - Vuejs, HTML & Laravel Admin Dashboard Template
      Author: Pixinvent
    Author URL: http://www.themeforest.net/user/pixinvent
========================================================================================== -->

<template>
    <div id="demo-card-analytics">
        <!-- ROW 1 -->
        <div class="vx-row">
            <!-- SESSIONS BY DEVICES -->
            <div class="vx-col w-full lg:w-1/3 mb-base">
                <vx-card title="Sessions By Device">
                    <!-- SLOT = ACTION -->
                    <template slot="actions">
                        <change-time-duration-dropdown />
                    </template>

                    <!-- CHART -->
                    <div slot="no-body">
                        <vue-apex-charts type="donut" height="340" class="mb-12 mt-4" :options="analyticsData.sessionsByDeviceDonut.chartOptions" :series="sessionsData.series" />
                    </div>

                    <!-- CHART DATA -->
                    <ul>
                        <li v-for="deviceData in sessionsData.analyticsData" :key="deviceData.device" class="flex mb-3">
                            <feather-icon :icon="deviceData.icon" :svgClasses="[`h-5 w-5 stroke-current text-${deviceData.color}`]"></feather-icon>
                            <span class="ml-2 inline-block font-semibold">{{ deviceData.device }}</span>
                            <span class="mx-2">-</span>
                            <span class="mr-4">{{ deviceData.sessionsPercentage }}%</span>
                            <div class="ml-auto flex -mr-1">
                            <span class="mr-1">{{ deviceData.comparedResultPercentage }}%</span>
                            <feather-icon :icon=" deviceData.comparedResultPercentage < 0 ? 'ArrowDownIcon' : 'ArrowUpIcon'" :svgClasses="[deviceData.comparedResultPercentage < 0 ? 'text-danger' : 'text-success'  ,'stroke-current h-4 w-4 mb-1 mr-1']"></feather-icon>
                            </div>
                        </li>
                    </ul>
                </vx-card>
            </div>

            <!-- PRODUCT ORDERS -->
            <div class="vx-col w-full lg:w-1/3 mb-base">
                <vx-card title="Product Orders">
                    <!-- SLOT = ACTIONS -->
                    <template slot="actions">
                        <change-time-duration-dropdown />
                    </template>

                    <!-- CHART -->
                    <div slot="no-body">
                        <vue-apex-charts type="radialBar" height="420" :options="analyticsData.productOrdersRadialBar.chartOptions" :series="productsOrder.series" />
                    </div>

                    <!-- CHART DATA -->
                    <ul>
                        <li v-for="orderData in productsOrder.analyticsData" :key="orderData.orderType" class="flex mb-3 justify-between">
                            <span class="flex items-center">
                                    <span class="inline-block h-4 w-4 rounded-full mr-2 bg-white border-3 border-solid" :class="`border-${orderData.color}`"></span>
                                    <span class="font-semibold">{{ orderData.orderType }}</span>
                            </span>
                            <span>{{ orderData.counts }}</span>
                        </li>
                    </ul>
                </vx-card>
            </div>

            <!-- CUSTOMERS CHART -->
            <div class="vx-col w-full lg:w-1/3 mb-base">
                <vx-card title="Customers">
                    <!-- SLOT = ACTIONS -->
                    <template slot="actions">
                        <change-time-duration-dropdown />
                    </template>

                    <div slot="no-body">
                        <!-- CHART -->
                        <vue-apex-charts type="pie" height="334" class="my-12" :options="analyticsData.customersPie.chartOptions" :series="customersData.series" />

                        <!-- CHART DATA -->
                        <ul class="mb-1">
                            <li v-for="customerData in customersData.analyticsData" :key="customerData.customerType" class="flex justify-between py-3 px-6 border d-theme-border-grey-light border-solid border-r-0 border-l-0 border-b-0">
                                <span class="flex items-center">
                                    <span class="inline-block h-3 w-3 rounded-full mr-2" :class="`bg-${customerData.color}`"></span>
                                    <span class="font-semibold">{{ customerData.customerType }}</span>
                                </span>
                                <span>{{ customerData.counts }}</span>
                            </li>
                        </ul>
                    </div>

                </vx-card>
            </div>

        </div>

        <!-- ROW 2 -->
        <div class="vx-row">
            <!-- SALES CHART - RADIAL -->
            <div class="vx-col w-full md:w-1/2 lg:w-1/2 xl:w-1/2 mb-base">
                <vx-card title="Sales" subtitle="Last 6 Months">
                    <!-- SLOT = ACTION -->
                    <template slot="actions">
                        <feather-icon icon="MoreVerticalIcon" svgClasses="w-6 h-6 text-grey"></feather-icon>
                    </template>
                    <!-- LABELS -->
                    <div class="flex">
                        <span class="flex items-center"><div class="h-3 w-3 rounded-full mr-2 bg-primary"></div><span>Sales</span></span>
                        <span class="flex items-center ml-5"><div class="h-3 w-3 rounded-full mr-2 bg-success"></div><span>Visits</span></span>
                    </div>
                    <!-- CHART -->
                    <div slot="no-body-bottom">
                        <vue-apex-charts type="radar" :options="analyticsData.statisticsRadar.chartOptions" :series="salesRadar.series" />
                    </div>
                </vx-card>
            </div>

            <!-- SUPPORT TRACKER CHART+- -->
            <div class="vx-col w-full md:w-1/2 lg:w-1/2 xl:w-1/2 mb-base">
                <vx-card title="Support Tracker">
                    <!-- CARD ACTION -->
                    <template slot="actions">
                        <change-time-duration-dropdown />
                    </template>

                    <div slot="no-body" v-if="supportTracker.analyticsData">
                        <div class="vx-row text-center">

                            <!-- Open Tickets Heading -->
                            <div class="vx-col w-full lg:w-1/5 md:w-full sm:w-1/5 flex flex-col justify-between mb-4 lg:order-first md:order-last sm:order-first order-last">
                                <div class="lg:ml-6 lg:mt-6 md:mt-0 md:ml-0 sm:ml-6 sm:mt-6">
                                    <h1 class="font-bold text-5xl">{{ supportTracker.analyticsData.openTickets }}</h1>
                                    <small>Tickets</small>
                                </div>
                            </div>

                            <!-- Chart -->
                            <div class="vx-col w-full lg:w-4/5 md:w-full sm:w-4/5 justify-center mx-auto lg:mt-0 md:mt-6 sm:mt-0 mt-6">
                                <vue-apex-charts type="radialBar" height="385" :options="analyticsData.supportTrackerRadialBar.chartOptions" :series="supportTracker.series" />
                            </div>
                        </div>

                        <!-- Support Tracker Meta Data -->
                        <div class="flex flex-row justify-between px-8 pb-4 mt-2">
                            <p class="text-center" v-for="(val, key) in supportTracker.analyticsData.meta" :key="key">
                              <span class="block">{{ key }}</span>
                              <span class="text-2xl font-semibold">{{ val }}</span>
                            </p>
                        </div>
                    </div>
                </vx-card>
            </div>
        </div>

        <!-- ROW 1 -->
        <div class="vx-row">

            <!-- LINE CHART -->
            <div class="vx-col w-full md:w-2/3 mb-base">
                <vx-card title="Revenue">

                    <template slot="actions">
                        <feather-icon icon="SettingsIcon" svgClasses="w-6 h-6 text-grey"></feather-icon>
                    </template>

                    <div slot="no-body" class="p-6 pb-0">

                        <div class="flex" v-if="revenueComparisonLine.analyticsData">
                            <div class="mr-6">
                                <p class="mb-1 font-semibold">This Month</p>
                                <p class="text-3xl text-success"><sup class="text-base mr-1">$</sup>{{ revenueComparisonLine.analyticsData.thisMonth.toLocaleString() }}</p>
                            </div>
                            <div>
                                <p class="mb-1 font-semibold">Last Month</p>
                                <p class="text-3xl"><sup class="text-base mr-1">$</sup>{{ revenueComparisonLine.analyticsData.lastMonth.toLocaleString() }}</p>
                            </div>
                        </div>

                        <vue-apex-charts
                          type="line"
                          height="266"
                          :options="analyticsData.revenueComparisonLine.chartOptions"
                          :series="revenueComparisonLine.series" />
                    </div>
                </vx-card>
            </div>

            <!-- RADIAL CHART -->
            <div class="vx-col w-full md:w-1/3 mb-base">
                <vx-card title="Goal Overview">
                    <template slot="actions">
                        <feather-icon icon="HelpCircleIcon" svgClasses="w-6 h-6 text-grey"></feather-icon>
                    </template>

                    <!-- CHART -->
                    <template slot="no-body">
                        <div class="mt-10">
                            <vue-apex-charts
                              type="radialBar"
                              height="240"
                              :options="analyticsData.goalOverviewRadialBar.chartOptions"
                              :series="goalOverview.series" />
                        </div>
                    </template>

                    <!-- DATA -->
                    <div
                      v-if="goalOverview.analyticsData"
                      class="flex justify-between text-center mt-4"
                      slot="no-body-bottom">

                        <div class="w-1/2 border border-solid d-theme-border-grey-light border-r-0 border-b-0 border-l-0">
                            <p class="mt-4">Completed</p>
                            <p class="mb-4 text-3xl font-semibold">{{ goalOverview.analyticsData.completed.toLocaleString() }}</p>
                        </div>
                        <div class="w-1/2 border border-solid d-theme-border-grey-light border-r-0 border-b-0">
                            <p class="mt-4">In Progress</p>
                            <p class="mb-4 text-3xl font-semibold">{{ goalOverview.analyticsData.inProgress.toLocaleString() }}</p>
                        </div>
                    </div>
                </vx-card>
            </div>
        </div>

        <!-- ROW 4 -->
        <div class="vx-row">
            <div class="vx-col w-full md:w-1/2 mb-base">
                <vx-card>
                    <div class="vx-row flex-col-reverse lg:flex-row">

                        <!-- LEFT COL -->
                        <div class="vx-col w-full lg:w-1/2 xl:w-1/2 flex flex-col justify-between" v-if="salesBarSession.analyticsData">
                            <div>

                                <!-- Avg Session -->
                                <h2 class="mb-1 font-bold">{{ salesBarSession.analyticsData.session | k_formatter }}</h2>
                                <span class="font-medium">Avg Sessions</span>

                                <!-- Previous Data comparison -->
                                <p class="mt-2 text-xl font-medium">
                                  <span :class="salesBarSession.analyticsData.comparison.result >= 0 ? 'text-success' : 'text-danger'">
                                    <span v-if="salesBarSession.analyticsData.comparison.result > 0">+</span>
                                    <span>{{ salesBarSession.analyticsData.comparison.result }}</span>
                                  </span>
                                  <span> vs </span>
                                  <span>{{ salesBarSession.analyticsData.comparison.str }}</span>
                                </p>

                            </div>
                            <vs-button icon-pack="feather" icon="icon-chevrons-right" icon-after class="shadow-md w-full lg:mt-0 mt-4">View Details</vs-button>
                        </div>

                        <!-- RIGHT COL -->
                        <div class="vx-col w-full lg:w-1/2 xl:w-1/2 flex flex-col lg:mb-0 mb-base">
                            <change-time-duration-dropdown class="self-end" />
                            <vue-apex-charts
                              type="bar"
                              height="200"
                              :options="analyticsData.salesBar.chartOptions"
                              :series="salesBarSession.series"
                              v-if="salesBarSession.series" />
                        </div>

                    </div>

                    <vs-divider class="my-6"></vs-divider>

                    <div class="vx-row">
                        <div class="vx-col w-1/2 mb-3">
                            <small>Goal: $100000</small>
                            <vs-progress class="block mt-1" :percent="50" color="primary"></vs-progress>
                        </div>
                        <div class="vx-col w-1/2 mb-3">
                            <small>Users: 100K</small>
                            <vs-progress class="block mt-1" :percent="60" color="warning"></vs-progress>
                        </div>
                        <div class="vx-col w-1/2 mb-3">
                            <small>Retention: 90%</small>
                            <vs-progress class="block mt-1" :percent="70" color="danger"></vs-progress>
                        </div>
                        <div class="vx-col w-1/2 mb-3">
                            <small>Duration: 1yr</small>
                            <vs-progress class="block mt-1" :percent="90" color="success"></vs-progress>
                        </div>
                    </div>
                </vx-card>
            </div>
            <div class="vx-col w-full md:w-1/2 mb-base">
                <vx-card class="overflow-hidden">
                    <template slot="no-body">
                        <div class="flex justify-between items-center p-6 border border-solid d-theme-border-grey-light border-r-0 border-l-0 border-t-0">
                            <div>
                                <p><span class="font-semibold">{{ todoToday.numComletedTasks }} task completed</span> out of {{ todoToday.totalTasks }}</p>
                                <vs-progress :percent="20" color="primary"></vs-progress>
                            </div>
                            <span>{{ todoToday.date }}</span>
                        </div>
                        <ul class="tasks-today-container">
                            <li class="px-6 py-4 tasks-today__task" v-for="todo in todoToday.tasksToday" :key="todo.id">
                                <div class="vx-row">
                                    <div class="vx-col w-full sm:w-auto">
                                        <p class="font-semibold text-lg">{{ todo.task }}</p>
                                        <small>{{ todo.date }}</small>
                                    </div>
                                    <div class="tasks-today__actions vx-col w-full sm:w-auto ml-auto mt-2 sm:mt-0">
                                        <vs-button radius color="primary" type="border" icon-pack="feather" icon="icon-check" size="small"></vs-button>
                                        <vs-button radius color="primary" type="border" icon-pack="feather" icon="icon-edit-2" size="small" class="ml-3"></vs-button>
                                        <vs-button radius color="primary" type="border" icon-pack="feather" icon="icon-trash" size="small" class="ml-3"></vs-button>
                                    </div>
                                </div>
                            </li>
                        </ul>
                    </template>
                </vx-card>
            </div>
        </div>

        <!-- ROW 3 -->
        <div class="vx-row">
            <div class="vx-col w-full md:w-2/3 lg:w-3/4">
                <vx-card title="Sales" class="mb-base">
                    <template slot="actions">
                        <feather-icon icon="SettingsIcon" svgClasses="w-6 h-6 text-grey"></feather-icon>
                    </template>
                    <div slot="no-body" class="p-6 pb-0">
                        <vue-apex-charts type="line" height="266" :options="analyticsData.salesLine.chartOptions" :series="salesLine.series" />
                    </div>
                </vx-card>
            </div>
            <div class="vx-col w-full md:w-1/3 lg:w-1/4 xl:w-1/4">
                <vx-card>
                    <template slot="no-body" v-if="Object.entries(funding).length">
                        <div class="p-8 clearfix">
                            <div>
                                <h1>
                                  <sup class="text-lg">$</sup>
                                  <span>{{ funding.currBalance.toLocaleString() }}</span>
                                </h1>
                                <small>
                                  <span class="text-grey">Deposits: </span>
                                  <span>$</span>
                                  <span>{{ funding.depostis.toLocaleString() }}</span>
                                </small>
                            </div>
                            <p
                              class="mt-2 mb-8 text-xl font-medium"
                              :class="funding.comparison.resultPerc >= 0 ? 'text-success' : 'text-danger'"
                              >
                              <span v-if="funding.comparison.resultPerc > 0">+</span>
                              <span>{{ funding.comparison.resultPerc }}%</span>
                              <span class="ml-1">(${{ funding.comparison.pastData }})</span>
                            </p>
                            <vs-button icon-pack="feather" icon="icon-chevrons-right" icon-after class="shadow-md w-full">Add Funds</vs-button>
                        </div>
                        <div class="p-8 border d-theme-border-grey-light border-solid border-r-0 border-l-0 border-b-0">
                            <div class="mb-4">
                                <small>Earned: ${{ funding.meta.earned.val }}</small>
                                <vs-progress :percent="funding.meta.earned.progress" color="success"></vs-progress>
                            </div>
                            <div>
                                <small>Duration: {{ funding.meta.duration.val }}</small>
                                <vs-progress :percent="funding.meta.duration.progress" color="warning"></vs-progress>
                            </div>
                        </div>
                    </template>
                </vx-card>
            </div>
        </div>

        <div class="vx-row">

            <div class="vx-col w-full md:w-1/3 lg:w-1/3 xl:w-1/3">
                <vx-card title="Browser Statistics" class="mt-base">
                    <div v-for="(browser, index) in browserStatistics" :key="browser.id" :class="{'mt-4': index}">
                        <div class="flex justify-between">
                            <div class="flex flex-col">
                                <span class="mb-1">{{ browser.name }}</span>
                                <h4>{{ browser.ratio }}%</h4>
                            </div>
                            <div class="flex flex-col text-right">
                                <span class="flex -mr-1">
                                    <span class="mr-1">{{ browser.comparedResult }}</span>
                                    <feather-icon :icon=" browser.comparedResult < 0 ? 'ArrowDownIcon' : 'ArrowUpIcon'" :svgClasses="[browser.comparedResult < 0 ? 'text-danger' : 'text-success'  ,'stroke-current h-4 w-4 mb-1 mr-1']"></feather-icon>
                                </span>
                                <span class="text-grey">{{ browser.time | time(true) }}</span>
                            </div>
                        </div>
                        <vs-progress :percent="browser.ratio"></vs-progress>
                    </div>
                </vx-card>
            </div>

            <div class="vx-col w-full md:w-2/3">
                <vx-card title="Client Retention" class="mt-base">
                    <div class="flex items-center mb-3">
                        <span class="flex items-center"><div class="h-3 w-3 rounded-full mr-1 bg-warning"></div><span>New Clients</span></span>
                        <span class="flex items-center ml-4"><div class="h-3 w-3 rounded-full mr-1 bg-danger"></div><span>Retained Clients</span></span>
                    </div>
                    <vue-apex-charts type="bar" height="266" :options="analyticsData.clientRetentionBar.chartOptions" :series="clientRetentionBar.series" />
                </vx-card>
            </div>
        </div>
    </div>
</template>

<script>
import VueApexCharts from 'vue-apexcharts'

import StatisticsCardLine from '@/components/statistics-cards/StatisticsCardLine.vue'
import analyticsData from './analyticsData.js'
import ChangeTimeDurationDropdown from '@/components/ChangeTimeDurationDropdown.vue'

export default {
  data () {
    return {
      sessionsData: {},
      productsOrder: {},
      customersData: {},

      salesRadar: {},
      supportTracker: {},

      revenueComparisonLine: {},
      goalOverview: {},

      salesBarSession: {},
      sessionDataTime: 'lastWeek',
      todoToday: {},

      salesLine: {},
      funding: {},

      browserStatistics: [],
      clientRetentionBar: {},

      analyticsData
    }
  },
  components: {
    VueApexCharts,
    StatisticsCardLine,
    ChangeTimeDurationDropdown
  },
  created () {
    // Sessions By Device
    this.$http.get('/api/card/card-analytics/session-by-device')
      .then((response) => { this.sessionsData = response.data })
      .catch((error) => { console.log(error) })

      // Products Order
    this.$http.get('/api/card/card-analytics/products-orders')
      .then((response) => { this.productsOrder = response.data })
      .catch((error) => { console.log(error) })

      // Customers
    this.$http.get('/api/card/card-analytics/customers')
      .then((response) => { this.customersData = response.data })
      .catch((error) => { console.log(error) })

      // Sales Radar
    this.$http.get('/api/card/card-analytics/sales/radar')
      .then((response) => { this.salesRadar = response.data })
      .catch((error) => { console.log(error) })

      // Support Tracker
    this.$http.get('/api/card/card-analytics/support-tracker')
      .then((response) => { this.supportTracker = response.data })
      .catch((error) => { console.log(error) })

      // Revenue Comparison
    this.$http.get('/api/card/card-analytics/revenue-comparison')
      .then((response) => { this.revenueComparisonLine = response.data })
      .catch((error) => { console.log(error) })

      // Goal Overview
    this.$http.get('/api/card/card-analytics/goal-overview')
      .then((response) => { this.goalOverview = response.data })
      .catch((error) => { console.log(error) })

      // Sales bar
    this.$http.get('/api/card/card-analytics/sales/bar')
      .then((response) => { this.salesBarSession = response.data })
      .catch((error) => { console.log(error) })

      // Todo
    this.$http.get('/api/card/card-analytics/todo/today')
      .then((response) => { this.todoToday = response.data })
      .catch((error) => { console.log(error) })

      // Funding
    this.$http.get('/api/card/card-analytics/funding')
      .then((response) => { this.funding = response.data })
      .catch((error) => { console.log(error) })

      // Sales line
    this.$http.get('/api/card/card-analytics/sales/line')
      .then((response) => { this.salesLine = response.data })
      .catch((error) => { console.log(error) })

      // Browser Analytics
    this.$http.get('/api/card/card-analytics/browser-analytics')
      .then((response) => { this.browserStatistics = response.data })
      .catch((error) => { console.log(error) })

      // Client Retention
    this.$http.get('/api/card/card-analytics/client-retention')
      .then((response) => { this.clientRetentionBar = response.data })
      .catch((error) => { console.log(error) })
  }
}
</script>


<style lang="scss">
#demo-card-analytics {
    .tasks-today-container {
        .tasks-today__task {
            transition: background .15s ease-out;
            .tasks-today__actions {
                display: none;
            }

            &:hover {
                background: rgba(var(--vs-primary),.04) ;
                .tasks-today__actions {
                    display: flex;
                }
            }
        }
    }
}
</style>
