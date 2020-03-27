<!-- =========================================================================================
    File Name: DropdownOptionsBasic.vue
    Description: Basic dropdown options in select component
    ----------------------------------------------------------------------------------------
    Item Name: Vuexy - Vuejs, HTML & Laravel Admin Dashboard Template
      Author: Pixinvent
    Author URL: http://www.themeforest.net/user/pixinvent
========================================================================================== -->


<template>
    <vx-card title="Basic" code-toggler no-shadow card-border>

        <div>
            <v-select :options="['foo','bar']" :dir="$vs.rtl ? 'rtl' : 'ltr'" />

            <p class="my-5">When provided an array of objects, <code>vue-select</code> will display a single value of the object. By default, <code>vue-select</code> will look for a key named <code>label</code> on the object to use as display text</p>

            <v-select :options="[{label: 'Foo', value: 'foo'}]" :dir="$vs.rtl ? 'rtl' : 'ltr'" />
        </div>

        <template slot="codeContainer">
&lt;template&gt;
  &lt;v-select :options=&quot;['foo','bar']&quot; :dir=&quot;$vs.rtl ? 'rtl' : 'ltr'&quot; /&gt;
  &lt;v-select :options=&quot;[{label: 'Foo', value: 'foo'}]&quot; :dir=&quot;$vs.rtl ? 'rtl' : 'ltr'&quot; /&gt;
&lt;/template&gt;

&lt;script&gt;
import vSelect from 'vue-select'

export default{
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
    components: {
        'v-select': vSelect,
    }
}
</script>
