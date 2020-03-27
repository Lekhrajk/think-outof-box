<!-- =========================================================================================
  File Name: ButtonGroupSize.vue.vue
  Description: Buttons Group Size
  ----------------------------------------------------------------------------------------
  Item Name: Vuexy - Vuejs, HTML & Laravel Admin Dashboard Template
    Author: Pixinvent
  Author URL: http://www.themeforest.net/user/pixinvent
========================================================================================== -->


<template>
  <vx-card title="Sizing" code-toggler>

    <p class="mb-4">You can use <code>size</code> prop of button component to change the size of button group.</p>

    <div class="btn-group">
      <vs-button size="large">First</vs-button>
      <vs-button size="large">Second</vs-button>
      <vs-button size="large">Third</vs-button>
    </div>

    <div class="btn-group my-4">
      <vs-button>First</vs-button>
      <vs-button>Second</vs-button>
      <vs-button>Third</vs-button>
    </div>

    <div class="btn-group">
      <vs-button size="small">First</vs-button>
      <vs-button size="small">Second</vs-button>
      <vs-button size="small">Third</vs-button>
    </div>

    <template slot="codeContainer">
&lt;div class=&quot;btn-group&quot;&gt;
  &lt;vs-button size=&quot;large&quot;&gt;First&lt;/vs-button&gt;
  &lt;vs-button size=&quot;large&quot;&gt;Second&lt;/vs-button&gt;
  &lt;vs-button size=&quot;large&quot;&gt;Third&lt;/vs-button&gt;
&lt;/div&gt;

&lt;div class=&quot;btn-group my-4&quot;&gt;
  &lt;vs-button&gt;First&lt;/vs-button&gt;
  &lt;vs-button&gt;Second&lt;/vs-button&gt;
  &lt;vs-button&gt;Third&lt;/vs-button&gt;
&lt;/div&gt;

&lt;div class=&quot;btn-group&quot;&gt;
  &lt;vs-button size=&quot;small&quot;&gt;First&lt;/vs-button&gt;
  &lt;vs-button size=&quot;small&quot;&gt;Second&lt;/vs-button&gt;
  &lt;vs-button size=&quot;small&quot;&gt;Third&lt;/vs-button&gt;
&lt;/div&gt;
    </template>
  </vx-card>
</template>

