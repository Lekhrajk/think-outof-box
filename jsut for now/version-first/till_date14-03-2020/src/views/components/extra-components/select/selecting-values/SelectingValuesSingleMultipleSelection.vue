<!-- =========================================================================================
    File Name: SelectingValuesSingleMultipleSelection.vue
    Description: Selecting single and multiple values
    ----------------------------------------------------------------------------------------
    Item Name: Vuexy - Vuejs, HTML & Laravel Admin Dashboard Template
      Author: Pixinvent
    Author URL: http://www.themeforest.net/user/pixinvent
========================================================================================== -->


<template>
    <vx-card title="Single/Multiple Selection" code-toggler no-shadow card-border>

        <div>

            <p class="mb-3">By default, vue-select supports choosing a single value. If you need multiple values, use the multiple prop</p>

            <v-select multiple :closeOnSelect="false" v-model="selected" :options="options" :dir="$vs.rtl ? 'rtl' : 'ltr'" /><br>

        </div>

        <template slot="codeContainer">
&lt;template&gt;
  &lt;v-select multiple :closeOnSelect=&quot;false&quot; v-model=&quot;selected&quot; :options=&quot;options&quot; :dir=&quot;$vs.rtl ? 'rtl' : 'ltr'&quot; /&gt;&lt;br&gt;
&lt;/template&gt;

&lt;script&gt;
import vSelect from 'vue-select'

export default{
  data() {
    return {
      selected: ['foo','bar'],
      options: ['foo','bar','baz']
    }
  },
  components: {
    'v-select': vSelect,
  }
}
&lt;/script&gt;
        </template>
    </vx-card>
</template>

<script>
import vSelect from 'vue-select'

export default{
    data() {
        return {
            selected: ['foo','bar'],
            options: ['foo','bar','baz']
        }
    },
    components: {
        'v-select': vSelect,
    }
}
</script>
