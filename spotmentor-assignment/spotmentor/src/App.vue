<template>
    <div id="app">
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
/* common style */
#app {
    font-family: 'Nunito', sans-serif;
    padding: 0;
    margin: 0;
}
/* datacard */
.cardbox {
    -webkit-box-shadow: 0 3px 9px rgba(0, 0, 0, 0.3);
    box-shadow: 0 3px 9px rgba(0, 0, 0, 0.3);
    display: block !important;
    background-color: #ffffff !important;
    margin-top: 5% !important;
    margin-bottom: 15px;
    max-width: 380px;
    padding: 3px 8px;
    border-radius: 8px;
    -webkit-border-radius: 8px;
    -moz-border-radius: 8px;
    -ms-border-radius: 8px;
    -o-border-radius: 8px;
}

/* scrollbar */

::-webkit-scrollbar {
    width: 12px;
}

::-webkit-scrollbar-thumb {
    background: linear-gradient(transparent, #0000ff);
    border-radius: 6px;
}

::-webkit-scrollbar-thumb:hover {
    background: linear-gradient(#0000ff, #0000ff);
}

/* button common styles */
.lbtn {
    border: 0;
    border-radius: 10px;
    font-size: 15px;
    font-weight: 600;
    padding: 9px 26px;
    margin: 5px 10px;
    display: inline-block;
    white-space: nowrap;
    outline: none;
    -webkit-transition: all 0.3s ease 0s;
    transition: all 0.3s ease 0s;
    -webkit-transform: scale(1);
    -moz-transform: scale(1);
    -ms-transform: scale(1);
    transform: scale(1);
}

.lbtn:hover,
.lbtn:focus {
    text-decoration: none;
    -webkit-transform: scale(0.95);
    -moz-transform: scale(0.95);
    -ms-transform: scale(0.95);
    transform: scale(0.95);
}

.lbtn-accent-color {
    color: #ffffff;
    background-color: #3195ff;
}

.lbtn-accent-color:hover,
.lbtn-accent-color:focus {
    color: #ffffff;
    background-color: #1888ff;
}


.lbtn-blue-color {
    color: #ffffff;
    background-color: #0000ff;
}

.lbtn-blue-color:hover,
.lbtn-blue-color:focus {
    color: #ffffff;
    background-color: #00000f;
}

.lbtn-shadow {
    -webkit-box-shadow: 0 9px 18px rgba(0, 0, 0, 0.08);
    box-shadow: 0 9px 18px rgba(0, 0, 0, 0.08);
}

.lbtn-light-color {
    background-color: #E5E7E9;
    color: #000;
}

.lbtn-light-color:hover,
.lbtn-light-color:focus {
    color: #000000;
    background-color: #ffffff;
}

.lbtn-small {
    border-radius: 8px;
    font-size: 13px;
    font-weight: 600;
    text-transform: none;
    padding: 8px 15px;
    margin: 2px 4px;
}
</style>