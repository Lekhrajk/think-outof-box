<!-- =========================================================================================
    File Name: CardStatistics.vue
    Description: Statistics Card
    ----------------------------------------------------------------------------------------
    Item Name: Vuexy - Vuejs, HTML & Laravel Admin Dashboard Template
      Author: Pixinvent
    Author URL: http://www.themeforest.net/user/pixinvent
========================================================================================== -->


<template>
    <div>

        <!-- ROW 1-->
        <div class="vx-row">
            <div class="vx-col w-1/2 md:w-1/3 xl:w-1/6">
                <statistics-card-line
                  hideChart
                  class="mb-base"
                  icon="EyeIcon"
                  statistic="36.9k"
                  statisticTitle="Views" />
            </div>

            <div class="vx-col w-1/2 md:w-1/3 xl:w-1/6">
                <statistics-card-line
                  hideChart
                  class="mb-base"
                  icon="MessageSquareIcon"
                  statisticTitle="Comments"
                  statistic="12k"
                  color="success" />
            </div>

            <div class="vx-col w-1/2 md:w-1/3 xl:w-1/6">
                <statistics-card-line
                  hideChart
                  class="mb-base"
                  icon="ShoppingBagIcon"
                  statistic="978"
                  statisticTitle="Orders"
                  color="warning" />
            </div>

            <div class="vx-col w-1/2 md:w-1/3 xl:w-1/6">
                <statistics-card-line
                  hideChart
                  class="mb-base"
                  icon="HeartIcon"
                  statistic="26.8k"
                  statisticTitle="Favorited"
                  color="danger" />
            </div>

            <div class="vx-col w-1/2 md:w-1/3 xl:w-1/6">
                <statistics-card-line
                  hideChart
                  class="mb-base"
                  icon="SmileIcon"
                  statistic="689"
                  statisticTitle="Reviews"
                  color="success" />
            </div>

            <div class="vx-col w-1/2 md:w-1/3 xl:w-1/6">
                <statistics-card-line
                  hideChart
                  class="mb-base"
                  icon="TruckIcon"
                  statistic="2"
                  statisticTitle="Returns"
                  color="warning" />
            </div>
        </div>

        <!-- ROW 2 -->
        <div class="vx-row">

            <div class="vx-col w-full sm:w-1/2 md:w-1/2 lg:w-1/4 xl:w-1/4">
                <statistics-card-line
                  hideChart
                  class="mb-base"
                  icon="CpuIcon"
                  icon-right
                  statistic="86%"
                  statisticTitle="CPU Usage" />
            </div>

            <div class="vx-col w-full sm:w-1/2 md:w-1/2 lg:w-1/4 xl:w-1/4">
                <statistics-card-line
                  hideChart
                  class="mb-base"
                  icon="ServerIcon"
                  icon-right
                  statistic="1.2gb"
                  statisticTitle="Memory Usage"
                  color="success" />
            </div>

            <div class="vx-col w-full sm:w-1/2 md:w-1/2 lg:w-1/4 xl:w-1/4">
                <statistics-card-line
                  hideChart
                  class="mb-base"
                  icon="ActivityIcon"
                  icon-right
                  statistic="0.1%"
                  statisticTitle="Downtime Ratio"
                  color="danger" />
            </div>

            <div class="vx-col w-full sm:w-1/2 md:w-1/2 lg:w-1/4 xl:w-1/4">
                <statistics-card-line
                  hideChart
                  class="mb-base"
                  icon="AlertOctagonIcon"
                  icon-right
                  statistic="13"
                  statisticTitle="Issues Found"
                  color="warning" />
            </div>
        </div>

        <!-- ROW 3 -->
        <div class="vx-row">
            <div class="vx-col w-full sm:w-1/2 md:w-1/2 lg:w-1/4 xl:w-1/4">
                <statistics-card-line
                  class="mb-base"
                  v-if="subscribersGained.analyticsData"
                  icon="UsersIcon"
                  :statistic="subscribersGained.analyticsData.subscribers | k_formatter"
                  statisticTitle="Subscribers Gained"
                  :chartData="subscribersGained.series"
                  type="area" />
            </div>

            <div class="vx-col w-full sm:w-1/2 md:w-1/2 lg:w-1/4 xl:w-1/4">
                <statistics-card-line
                  class="mb-base"
                  v-if="revenueGenerated.analyticsData"
                  icon="DollarSignIcon"
                  :statistic="revenueGenerated.analyticsData.revenue | k_formatter"
                  statisticTitle="Revenue Generated"
                  :chartData="revenueGenerated.series"
                  color="success"
                  type="area" />
            </div>

            <div class="vx-col w-full sm:w-1/2 md:w-1/2 lg:w-1/4 xl:w-1/4">
                <statistics-card-line
                  class="mb-base"
                  v-if="quarterlySales.analyticsData"
                  icon="ShoppingCartIcon"
                  :statistic="quarterlySales.analyticsData.sales"
                  statisticTitle="Quarterly Sales"
                  :chartData="quarterlySales.series"
                  color="danger"
                  type="area" />
            </div>
            <div class="vx-col w-full sm:w-1/2 md:w-1/2 lg:w-1/4 xl:w-1/4">
                <statistics-card-line
                  class="mb-base"
                  v-if="ordersRecevied.analyticsData"
                  icon="ShoppingBagIcon"
                  :statistic="ordersRecevied.analyticsData.orders | k_formatter"
                  statisticTitle="Orders Received"
                  :chartData="ordersRecevied.series"
                  color="warning"
                  type="area" />
            </div>
        </div>

        <!-- ROW 4 -->
        <div class="vx-row">

            <div class="vx-col w-full md:w-1/3 lg:w-1/3 xl:w-1/3">
                <statistics-card-line
                  class="md:mb-0 mb-base"
                  icon="MonitorIcon"
                  icon-right
                  statistic="78.9k"
                  statisticTitle="Site Traffic"
                  :chartData="siteTraffic.series" />
            </div>

            <div class="vx-col w-full md:w-1/3 lg:w-1/3 xl:w-1/3">
                <statistics-card-line
                  class="md:mb-0 mb-base"
                  icon="UserCheckIcon"
                  icon-right
                  statistic="659.8k"
                  statisticTitle="Active Users"
                  :chartData="activeUsers.series"
                  color="success" />
            </div>

            <div class="vx-col w-full md:w-1/3 lg:w-1/3 xl:w-1/3">
                <statistics-card-line
                  icon="MailIcon"
                  icon-right
                  statistic="28.7k"
                  statisticTitle="Newsletter"
                  :chartData="newsletter.series"
                  color="warning" />
            </div>
        </div>

    </div>
</template>

<script>
import StatisticsCardLine from '@/components/statistics-cards/StatisticsCardLine.vue'

export default{
  components: {
    StatisticsCardLine
  },
  data () {
    return {
      // Area charts
      subscribersGained: {},
      revenueGenerated: {},
      quarterlySales: {},
      ordersRecevied: {},

      // Line Charts
      siteTraffic: {},
      activeUsers: {},
      newsletter: {}
    }
  },
  created () {
    // Subscribers gained
    this.$http.get('/api/card/card-statistics/subscribers')
      .then((response) => { this.subscribersGained = response.data })
      .catch((error) => { console.log(error) })

      // Revenue Generated
    this.$http.get('/api/card/card-statistics/revenue')
      .then((response) => { this.revenueGenerated = response.data })
      .catch((error) => { console.log(error) })

      // Sales
    this.$http.get('/api/card/card-statistics/sales')
      .then((response) => { this.quarterlySales = response.data })
      .catch((error) => { console.log(error) })

      // Orders
    this.$http.get('/api/card/card-statistics/orders')
      .then((response) => { this.ordersRecevied = response.data })
      .catch((error) => { console.log(error) })

      // Site Traffic gained
    this.$http.get('/api/card/card-statistics/site-traffic')
      .then((response) => { this.siteTraffic = response.data })
      .catch((error) => { console.log(error) })

      // Active Users
    this.$http.get('/api/card/card-statistics/active-users')
      .then((response) => { this.activeUsers = response.data })
      .catch((error) => { console.log(error) })

      // Site Traffic gained
    this.$http.get('/api/card/card-statistics/newsletter')
      .then((response) => { this.newsletter = response.data })
      .catch((error) => { console.log(error) })
  }
}
</script>
