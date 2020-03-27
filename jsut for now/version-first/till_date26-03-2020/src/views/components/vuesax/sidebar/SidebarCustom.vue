<!-- =========================================================================================
    File Name: SidebarCustom.vue
    Description: Rendering of custom sidebar with vx-sidebar(replaces vs-sidebar-item) and vx-sidebar-group(vs-sidebar-group)
    ----------------------------------------------------------------------------------------
    Item Name: Vuexy - Vuejs, HTML & Laravel Admin Dashboard Template
      Author: Pixinvent
    Author URL: http://www.themeforest.net/user/pixinvent
========================================================================================== -->


<template>
    <vx-card title="Custom Sidebar" code-toggler>

        <p>With custom sidebar you can create sidebar items and groups using <strong> list of object</strong>.</p>
        <p>You can also use feather icons in sidebar group and sidebar items using icon property. Check out all feather icons <a href="https://vue-feather-icons.netlify.com/" target="_blank" rel="nofollow">here</a></p>

        <p class="mt-3">Main sidebar on left side is demo of Custom Sidebar. You can check full code by clicking on code-toggler icon of this card.</p>

        <template slot="codeContainer">
&lt;template&gt;
  &lt;v-nav-menu
    :sidebarItems=&quot;sidebarItems&quot;
    title=&quot;Custom&quot;
    :logo=&quot;navMenuLogo&quot;&gt;
&lt;/template&gt;

&lt;script&gt;
import VNavMenu from &quot;@/layouts/components/vertical-nav-menu/VerticalNavMenu.vue&quot;;

export default {
  data: () =&gt; ({
    showSidebar: false,
    navMenuLogo: require('@/assets/images/logo/logo.png'),
    sidebarItems: [
      {
        url: &quot;/&quot;,
        name: &quot;Dashboard&quot;,
        slug: &quot;dashboardAnalytics&quot;,
        icon: &quot;HomeIcon&quot;,
        i18n: &quot;Dashboard&quot;,
      },
      {
        url: &quot;/ui-elements/colors&quot;,
        name: &quot;Colors&quot;,
        slug: &quot;colors&quot;,
        icon: &quot;DropletIcon&quot;,
        i18n: &quot;Colors&quot;,
      },
      {
        header: &quot;Extensions&quot;,
        i18n: &quot;Extensions&quot;,
      },
      {
        url: '/extensions/select',
        name: &quot;Select&quot;,
        icon: &quot;PocketIcon&quot;,
        slug: &quot;extraComponentSelect&quot;,
        i18n: &quot;Select&quot;,
      },
      {
        url: '/extensions/quill-editor',
        name: &quot;Quill Editor&quot;,
        icon: &quot;EditIcon&quot;,
        slug: &quot;extraComponentQuillEditor&quot;,
        i18n: &quot;QuillEditor&quot;,
      },
      {
        url: '/extensions/drag-and-drop',
        name: &quot;Drag &amp; Drop&quot;,
        icon: &quot;DropletIcon&quot;,
        slug: &quot;extraComponentDragAndDrop&quot;,
        i18n: &quot;DragAndDrop&quot;,
      },
      {
        url: '/extensions/datepicker',
        name: &quot;Datepicker&quot;,
        icon: &quot;CalendarIcon&quot;,
        slug: &quot;extraComponentDatepicker&quot;,
        i18n: &quot;Datepicker&quot;,
      },
    ]
  }),
  components: {
    VNavMenu
  }
}
&lt;/script&gt;
        </template>

    </vx-card>
</template>
