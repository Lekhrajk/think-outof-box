<!-- =========================================================================================
  File Name: ContextMenuClickNotClose.vue
  Description: Click not close on click of context item
  ----------------------------------------------------------------------------------------
  Item Name: Vuexy - Vuejs, HTML & Laravel Admin Dashboard Template
  Author: Pixinvent
  Author URL: http://www.themeforest.net/user/pixinvent
========================================================================================== -->

<template>
  <vx-card title="Click Not Close" code-toggler>
    <vs-button @contextmenu.prevent="$refs.menu.open">Right click on me - Do not close</vs-button>
    <vue-context ref="menu" :closeOnClick="false">
        <li>
          <a href="#" @click="optionClicked($event.target.innerText)" class="flex items-center text-sm">
            <feather-icon :icon="$vs.rtl ? 'ChevronLeftIcon' : 'ChevronRightIcon'" svgClasses="w-5 h-5" />
            <span class="ml-2">Next</span>
          </a>
        </li>
        <li>
          <a href="#" @click="optionClicked($event.target.innerText)" class="flex items-center text-sm">
            <feather-icon :icon="$vs.rtl ? 'ChevronRightIcon' : 'ChevronLeftIcon'" svgClasses="w-5 h-5" />
            <span class="ml-2">Previous</span>
          </a>
        </li>
        <li>
          <a href="#" @click="optionClicked($event.target.innerText)" class="flex items-center text-sm">
            <feather-icon icon="PlusIcon" svgClasses="w-5 h-5" />
            <span class="ml-2">Increase</span>
          </a>
        </li>
        <li>
          <a href="#" @click="optionClicked($event.target.innerText)" class="flex items-center text-sm">
            <feather-icon icon="MinusIcon" svgClasses="w-5 h-5" />
            <span class="ml-2">Decrease</span>
          </a>
        </li>
        <li>
          <a href="#" @click="optionClicked($event.target.innerText)" class="flex items-center text-sm">
            <feather-icon icon="ShuffleIcon" svgClasses="w-5 h-5" />
            <span class="ml-2">Random</span>
          </a>
        </li>
    </vue-context>
    <template slot="codeContainer">
&lt;template&gt;
    &lt;vs-button @contextmenu.prevent=&quot;$refs.menu.open&quot;&gt;Right click on me - Do not close&lt;/vs-button&gt;
    &lt;vue-context ref=&quot;menu&quot; :closeOnClick=&quot;false&quot;&gt;
        &lt;ul class=&quot;bordered-items p-0&quot;&gt;
            &lt;li&gt;
              &lt;a href=&quot;#&quot; @click=&quot;optionClicked($event.target.innerText)&quot; class=&quot;flex items-center text-sm&quot;&gt;
                &lt;feather-icon icon=&quot;ChevronRightIcon&quot; svgClasses=&quot;w-5 h-5&quot; /&gt;
                &lt;feather-icon :icon=&quot;$vs.rtl ? 'ChevronLeftIcon' : 'ChevronRightIcon'&quot; svgClasses=&quot;w-5 h-5&quot; /&gt;
              &lt;/a&gt;
            &lt;/li&gt;
            &lt;li&gt;
              &lt;a href=&quot;#&quot; @click=&quot;optionClicked($event.target.innerText)&quot; class=&quot;flex items-center text-sm&quot;&gt;
                &lt;feather-icon icon=&quot;ChevronLeftIcon&quot; svgClasses=&quot;w-5 h-5&quot; /&gt;
                &lt;feather-icon :icon=&quot;$vs.rtl ? 'ChevronRightIcon' : 'ChevronLeftIcon'&quot; svgClasses=&quot;w-5 h-5&quot; /&gt;
              &lt;/a&gt;
            &lt;/li&gt;
            &lt;li&gt;
              &lt;a href=&quot;#&quot; @click=&quot;optionClicked($event.target.innerText)&quot; class=&quot;flex items-center text-sm&quot;&gt;
                &lt;feather-icon icon=&quot;PlusIcon&quot; svgClasses=&quot;w-5 h-5&quot; /&gt;
                &lt;span class=&quot;ml-2&quot;&gt;Increase&lt;/span&gt;
              &lt;/a&gt;
            &lt;/li&gt;
            &lt;li&gt;
              &lt;a href=&quot;#&quot; @click=&quot;optionClicked($event.target.innerText)&quot; class=&quot;flex items-center text-sm&quot;&gt;
                &lt;feather-icon icon=&quot;MinusIcon&quot; svgClasses=&quot;w-5 h-5&quot; /&gt;
                &lt;span class=&quot;ml-2&quot;&gt;Decrease&lt;/span&gt;
              &lt;/a&gt;
            &lt;/li&gt;
            &lt;li&gt;
              &lt;a href=&quot;#&quot; @click=&quot;optionClicked($event.target.innerText)&quot; class=&quot;flex items-center text-sm&quot;&gt;
                &lt;feather-icon icon=&quot;ShuffleIcon&quot; svgClasses=&quot;w-5 h-5&quot; /&gt;
                &lt;span class=&quot;ml-2&quot;&gt;Random&lt;/span&gt;
              &lt;/a&gt;
            &lt;/li&gt;
        &lt;/ul&gt;
    &lt;/vue-context&gt;
&lt;/template&gt;

&lt;script&gt;
import { VueContext } from 'vue-context';

export default{
  components: {
    VueContext
  },
  methods: {
    optionClicked(text) {
      this.$vs.notify({
        title:'Action Clicked',
        text: text,
        icon: 'feather',
        iconPack: 'icon-alert-circle',
        color:'primary'
      })
    }
  }
}
&lt;/script&gt;
    </template>
  </vx-card>
</template>

<script>
import { VueContext } from 'vue-context';

export default{
  components: {
    VueContext
  },
  methods: {
    optionClicked(text) {
      this.$vs.notify({
        title:'Action Clicked',
        text: text,
        icon: 'feather',
        iconPack: 'icon-alert-circle',
        color:'primary'
      })
    }
  }
}
</script>
