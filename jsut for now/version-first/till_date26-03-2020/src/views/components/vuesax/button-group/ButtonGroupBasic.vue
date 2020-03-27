<!-- =========================================================================================
  File Name: ButtonGroupBasic.vue.vue
  Description: Buttons Group Basic
  ----------------------------------------------------------------------------------------
  Item Name: Vuexy - Vuejs, HTML & Laravel Admin Dashboard Template
    Author: Pixinvent
  Author URL: http://www.themeforest.net/user/pixinvent
========================================================================================== -->


<template>
  <vx-card title="Basic" code-toggler>

    <p class="mb-4">Add <code>btn-group</code> class to wrapper to make all contained button acts as grouped buttons.</p>

    <div class="btn-group">
      <vs-button>First</vs-button>
      <vs-button>Second</vs-button>
      <vs-button>Third</vs-button>
    </div>

    <template slot="codeContainer">
&lt;div class=&quot;btn-group&quot;&gt;
  &lt;vs-button&gt;First&lt;/vs-button&gt;
  &lt;vs-button&gt;Second&lt;/vs-button&gt;
  &lt;vs-button&gt;Third&lt;/vs-button&gt;
&lt;/div&gt;
    </template>
  </vx-card>
</template>

