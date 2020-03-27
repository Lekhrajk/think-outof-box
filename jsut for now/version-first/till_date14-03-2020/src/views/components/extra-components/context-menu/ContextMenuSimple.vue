<!-- =========================================================================================
  File Name: ContextMenuSimple.vue
  Description: Simple context menu Demo
  ----------------------------------------------------------------------------------------
  Item Name: Vuexy - Vuejs, HTML & Laravel Admin Dashboard Template
  Author: Pixinvent
  Author URL: http://www.themeforest.net/user/pixinvent
========================================================================================== -->

<template>
  <vx-card title="Simple" code-toggler>
    <vs-button @contextmenu.prevent="$refs.menu.open">Right click on me</vs-button>
    <vue-context ref="menu">
        <li>
          <a href="#" @click="optionClicked($event.target.innerText)" class="flex items-center text-sm">
            <feather-icon icon="PlusIcon" svgClasses="w-5 h-5" />
            <span class="ml-2">New</span>
          </a>
        </li>
        <li>
          <a href="#" @click="optionClicked($event.target.innerText)" class="flex items-center text-sm">
            <feather-icon icon="FileIcon" svgClasses="w-5 h-5" />
            <span class="ml-2">Open</span>
          </a>
        </li>
        <li>
          <a href="#" @click="optionClicked($event.target.innerText)" class="flex items-center text-sm">
            <feather-icon icon="SaveIcon" svgClasses="w-5 h-5" />
            <span class="ml-2">Save</span>
          </a>
        </li>
        <li>
          <a href="#" @click="optionClicked($event.target.innerText)" class="flex items-center text-sm">
            <feather-icon icon="SaveIcon" svgClasses="w-5 h-5" />
            <span class="ml-2">Save As</span>
          </a>
        </li>
        <li>
          <a href="#" @click="optionClicked($event.target.innerText)" class="flex items-center text-sm">
            <feather-icon icon="XIcon" svgClasses="w-5 h-5" />
            <span class="ml-2">Close</span>
          </a>
        </li>
    </vue-context>
    <template slot="codeContainer">
&lt;template&gt;
    &lt;vs-button @contextmenu.prevent=&quot;$refs.menu.open&quot;&gt;Right click on me&lt;/vs-button&gt;
    &lt;vue-context ref=&quot;menu&quot;&gt;
        &lt;li&gt;
          &lt;a href=&quot;#&quot; @click=&quot;optionClicked($event.target.innerText)&quot; class=&quot;flex items-center text-sm&quot;&gt;
            &lt;feather-icon icon=&quot;PlusIcon&quot; svgClasses=&quot;w-5 h-5&quot; /&gt;
            &lt;span class=&quot;ml-2&quot;&gt;New&lt;/span&gt;
          &lt;/a&gt;
        &lt;/li&gt;
        &lt;li&gt;
          &lt;a href=&quot;#&quot; @click=&quot;optionClicked($event.target.innerText)&quot; class=&quot;flex items-center text-sm&quot;&gt;
            &lt;feather-icon icon=&quot;FileIcon&quot; svgClasses=&quot;w-5 h-5&quot; /&gt;
            &lt;span class=&quot;ml-2&quot;&gt;Open&lt;/span&gt;
          &lt;/a&gt;
        &lt;/li&gt;
        &lt;li&gt;
          &lt;a href=&quot;#&quot; @click=&quot;optionClicked($event.target.innerText)&quot; class=&quot;flex items-center text-sm&quot;&gt;
            &lt;feather-icon icon=&quot;SaveIcon&quot; svgClasses=&quot;w-5 h-5&quot; /&gt;
            &lt;span class=&quot;ml-2&quot;&gt;Save&lt;/span&gt;
          &lt;/a&gt;
        &lt;/li&gt;
        &lt;li&gt;
          &lt;a href=&quot;#&quot; @click=&quot;optionClicked($event.target.innerText)&quot; class=&quot;flex items-center text-sm&quot;&gt;
            &lt;feather-icon icon=&quot;SaveIcon&quot; svgClasses=&quot;w-5 h-5&quot; /&gt;
            &lt;span class=&quot;ml-2&quot;&gt;Save As&lt;/span&gt;
          &lt;/a&gt;
        &lt;/li&gt;
        &lt;li&gt;
          &lt;a href=&quot;#&quot; @click=&quot;optionClicked($event.target.innerText)&quot; class=&quot;flex items-center text-sm&quot;&gt;
            &lt;feather-icon icon=&quot;XIcon&quot; svgClasses=&quot;w-5 h-5&quot; /&gt;
            &lt;span class=&quot;ml-2&quot;&gt;Close&lt;/span&gt;
          &lt;/a&gt;
        &lt;/li&gt;
    &lt;/vue-context&gt;
&lt;/template&gt;

&lt;script&gt;
import { VueContext } from 'vue-context';

export default {
    components: {
        VueContext
    },
    methods: {
        optionClicked(text) {
            this.$vs.notify({
                title: 'Action Clicked',
                text: text,
                icon: 'feather',
                iconPack: 'icon-alert-circle',
                color: 'primary'
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
