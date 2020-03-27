<!-- =========================================================================================
    File Name: ButtonLineDown.vue
    Description: Buttons with underline
    ----------------------------------------------------------------------------------------
    Item Name: Vuexy - Vuejs, HTML & Laravel Admin Dashboard Template
      Author: Pixinvent
    Author URL: http://www.themeforest.net/user/pixinvent
========================================================================================== -->


<template>
    <vx-card title="Line Down" code-toggler>

        <p>To add a type of button with edges we change the value of <code>type</code> by: <code>line</code></p>

        <vs-alert color="primary" icon="new_releases" active="true" class="mt-5">
            <span>you can change the position of the line up with the property line-position="top". You can also change the way the line appears on mouse hover with the property line-origin using the allowed values: <strong>left, center(default) or right</strong></span>
        </vs-alert>

        <div class="demo-alignment">

            <vs-button type="line">Primary</vs-button>
            <vs-button type="line" line-origin="left" color="success">Success</vs-button>
            <vs-button type="line" line-origin="right" color="danger">Danger</vs-button>
            <vs-button type="line" line-position="top" line-origin="left" color="warning">Warning</vs-button>
            <vs-button type="line" line-position="top" line-origin="right" color="dark">Dark</vs-button>
            <vs-button type="line" disabled>Disabled</vs-button>

        </div>

        <template slot="codeContainer">
&lt;vs-button type=&quot;line&quot;&gt;Primary&lt;/vs-button&gt;
&lt;vs-button type=&quot;line&quot; line-origin=&quot;left&quot; color=&quot;success&quot;&gt;Success&lt;/vs-button&gt;
&lt;vs-button type=&quot;line&quot; line-origin=&quot;right&quot; color=&quot;danger&quot;&gt;Danger&lt;/vs-button&gt;
&lt;vs-button type=&quot;line&quot; line-position=&quot;top&quot; line-origin=&quot;left&quot; color=&quot;warning&quot;&gt;Warning&lt;/vs-button&gt;
&lt;vs-button type=&quot;line&quot; line-position=&quot;top&quot; line-origin=&quot;right&quot; color=&quot;dark&quot;&gt;Dark&lt;/vs-button&gt;
&lt;vs-button type=&quot;line&quot; disabled&gt;Disabled&lt;/vs-button&gt;
        </template>

    </vx-card>
</template>
