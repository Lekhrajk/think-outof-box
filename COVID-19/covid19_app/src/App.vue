<template>
    <div id="app" data-spy="scroll" data-target=".navbar" data-offset="50">
        <Layout />
    </div>
</template>
<script>
import Layout from '@/components/Layout.vue'

export default {
    name: 'App',
    components: {
        Layout
    }
}
</script>
<style>
    *{
        font-family: Nunito,sans-serif;
    }
</style>