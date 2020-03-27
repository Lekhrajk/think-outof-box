<!-- =========================================================================================
	File Name: EchartsRadarChart.vue
	Description: Create Radar chart
	----------------------------------------------------------------------------------------
	Item Name: Vuexy - Vuejs, HTML & Laravel Admin Dashboard Template
	Author: Pixinvent
	Author URL: http://www.themeforest.net/user/pixinvent
========================================================================================== -->


<template>
	<vx-card title="Radar Chart" class="" code-toggler>

		<div class="mt-5">
			<e-charts autoresize :options="radar" ref="radar" auto-resize />
		</div>

		<template slot="codeContainer">
&lt;template&gt;
  &lt;e-charts autoresize :options=&quot;radar&quot; ref=&quot;radar&quot; auto-resize /&gt;
&lt;/template&gt;

&lt;script&gt;
import ECharts from 'vue-echarts/components/ECharts'
import 'echarts/lib/component/tooltip'
import 'echarts/lib/component/legend'
import 'echarts/lib/chart/radar'

export default {
  data() {
    return {
      radar:{
            tooltip: {},
            radar: {
              indicator: [
              {name: 'Attack', max: 20},
              {name: 'Defensive', max: 20},
              {name: 'Speed', max: 20},
              {name: 'Power', max: 20},
              {name: 'Endurance', max: 20},
              {name: 'Agile', max: 20}
            ]
            },
            series: [{
              name: 'Ability value',
              type: 'radar',
              data: [{value: [19,9,18,16,16,20]}]
            }]
      },
    }
  },
  components: {
    ECharts
  }
}
&lt;/script&gt;
		</template>
	</vx-card>
</template>

<script>
import ECharts from 'vue-echarts/components/ECharts'
import 'echarts/lib/component/tooltip'
import 'echarts/lib/component/legend'
import 'echarts/lib/chart/radar'

export default {
    data() {
        return {
            radar: {
                tooltip: {},
                radar: {
                    indicator: [
                        { name: 'Attack', max: 20 },
                        { name: 'Defensive', max: 20 },
                        { name: 'Speed', max: 20 },
                        { name: 'Power', max: 20 },
                        { name: 'Endurance', max: 20 },
                        { name: 'Agile', max: 20 }
                    ]
                },
                series: [{
                    name: 'Ability value',
                    type: 'radar',
                    data: [{ value: [19, 9, 18, 16, 16, 20] }]
                }]
            },
        }
    },
    components: {
        ECharts
    }
}
</script>
