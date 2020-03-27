<!-- =========================================================================================
  File Name: ButtonGroupVertical.vue.vue
  Description: Buttons Group in vertical variant
  ----------------------------------------------------------------------------------------
  Item Name: Vuexy - Vuejs, HTML & Laravel Admin Dashboard Template
    Author: Pixinvent
  Author URL: http://www.themeforest.net/user/pixinvent
========================================================================================== -->


<template>
  <vx-card title="Vertical Variant" code-toggler>

    <p class="mb-4">You can create vertical variant of button group using <code>btn-group-vertical</code> class</p>

    <div class="btn-group-vertical">
      <vs-button>Top</vs-button>
      <vs-button>Middle</vs-button>
      <vs-button>Bottom</vs-button>
    </div>

    <template slot="codeContainer">
&lt;div class=&quot;btn-group-vertical&quot;&gt;
  &lt;vs-button&gt;Top&lt;/vs-button&gt;
  &lt;vs-button&gt;Middle&lt;/vs-button&gt;
  &lt;vs-button&gt;Bottom&lt;/vs-button&gt;
&lt;/div&gt;
    </template>
  </vx-card>
</template>

