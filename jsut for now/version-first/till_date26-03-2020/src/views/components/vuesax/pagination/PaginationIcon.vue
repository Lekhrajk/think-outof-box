<!-- =========================================================================================
    File Name: PaginationDefault.vue
    Description: Changing icon of next and previous arrow
    ----------------------------------------------------------------------------------------
    Item Name: Vuexy - Vuejs, HTML & Laravel Admin Dashboard Template
      Author: Pixinvent
    Author URL: http://www.themeforest.net/user/pixinvent
========================================================================================== -->


<template>
    <vx-card title="Icon" code-toggler>

        <p>To customize the previous and next page icons change the value of the prev-icon and prev-icon directives. By default, their values are respectively arrow_back and arrow_forward</p>

        <vs-alert color="primary" icon="new_releases" active="true" class="mt-5">
            <p>Vuesax uses the Google Material Icons font library by default. For a list of all available icons, visit the official <a href="https://material.io/icons/" target="_blank">Material Icons page</a>.</p>
            <p>FontAwesome and other fonts library are supported. Simply use the icon-pack with fa or fas. You still need to include the Font Awesome icons in your project.</p>
        </vs-alert>

        <div class="my-5">
            <vs-pagination
              :total="100"
              v-model="currentx"
              :prev-icon="$vs.rtl ? 'arrow_forward' : 'arrow_back'"
              :next-icon="$vs.rtl ? 'arrow_back' : 'arrow_forward'" />
        </div>

        <template slot="codeContainer">
&lt;template&gt;
  &lt;vs-pagination
    :total=&quot;100&quot;
    v-model=&quot;currentx&quot;
    :prev-icon=&quot;$vs.rtl ? 'arrow_forward' : 'arrow_back'&quot;
    :next-icon=&quot;$vs.rtl ? 'arrow_back' : 'arrow_forward'&quot; /&gt;
&lt;/template&gt;

&lt;script&gt;
export default {
  data: () =&gt; ({
    currentx: 64
  })
}
&lt;/script&gt;
        </template>

    </vx-card>
</template>

<script>
export default {
    data: () => ({
        currentx: 64
    })
}
</script>
