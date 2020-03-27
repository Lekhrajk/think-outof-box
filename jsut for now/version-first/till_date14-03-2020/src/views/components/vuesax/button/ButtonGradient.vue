<!-- =========================================================================================
    File Name: ButtonGradient.vue
    Description: Give gradient to button background
    ----------------------------------------------------------------------------------------
    Item Name: Vuexy - Vuejs, HTML & Laravel Admin Dashboard Template
      Author: Pixinvent
    Author URL: http://www.themeforest.net/user/pixinvent
========================================================================================== -->


<template>
    <vx-card title="Gradient" code-toggler>

        <p>To add a type of button with edges we change the value of <code>type</code> by: <code>gradient</code></p>

        <div class="demo-alignment">

            <vs-button type="gradient">Primary</vs-button>
            <vs-button color="success" type="gradient">Success</vs-button>
            <vs-button color="danger" type="gradient">Danger</vs-button>
            <vs-button color="warning" type="gradient">Warning</vs-button>
            <vs-button color="dark" type="gradient">Dark</vs-button>
            <vs-button color="#3dd495" gradient-color-secondary="rgb(130, 207, 23)" type="gradient">Dark</vs-button>
            <vs-button disabled type="gradient">Disabled</vs-button>

        </div>

        <template slot="codeContainer">
&lt;vs-button type=&quot;gradient&quot;&gt;primary&lt;/vs-button&gt;
&lt;vs-button color=&quot;success&quot; type=&quot;gradient&quot;&gt;success&lt;/vs-button&gt;
&lt;vs-button color=&quot;danger&quot; type=&quot;gradient&quot;&gt;danger&lt;/vs-button&gt;
&lt;vs-button color=&quot;warning&quot; type=&quot;gradient&quot;&gt;warning&lt;/vs-button&gt;
&lt;vs-button color=&quot;dark&quot; type=&quot;gradient&quot;&gt;dark&lt;/vs-button&gt;
&lt;vs-button color=&quot;#3dd495&quot; gradient-color-secondary=&quot;rgb(130, 207, 23)&quot; type=&quot;gradient&quot;&gt;dark&lt;/vs-button&gt;
&lt;vs-button disabled type=&quot;gradient&quot;&gt;Disabled&lt;/vs-button&gt;
        </template>

    </vx-card>
</template>
