<!-- =========================================================================================
    File Name: DialogsAlert.vue
    Description: Create alert dialogs
    ----------------------------------------------------------------------------------------
    Item Name: Vuexy - Vuejs, HTML & Laravel Admin Dashboard Template
      Author: Pixinvent
    Author URL: http://www.themeforest.net/user/pixinvent
========================================================================================== -->


<template>
    <vx-card title="Alert">

        <p>To generate a dialog we have the global function $vs.dialog the parameters to generate the dialog are</p>
        <vs-list>
            <vs-list-item title="title" subtitle="title of the dialogue - String"></vs-list-item>
            <vs-list-item title="text" subtitle="internal text of the dialog - String"></vs-list-item>
            <vs-list-item title="color" subtitle="color of the component - Colors, RGB, HEX"></vs-list-item>
            <vs-list-item title="accept" subtitle="function that executes the user accept the dialog - Function"></vs-list-item>
            <vs-list-item title="cancel" subtitle="function that executes the user to cancel the dialog (if it is a confirm type) - Function"></vs-list-item>
            <vs-list-item title="type" subtitle="determines the type of dialog (alert or confirm) - alert, confirm"></vs-list-item>
            <vs-list-item title="buttonAccept" subtitle="determines the type of accept button - Boolean"></vs-list-item>
            <vs-list-item title="buttonCancel" subtitle="determines the type of cancel button - Boolean"></vs-list-item>
        </vs-list>

        <div class="demo-alignment">
            <vs-button @click="openAlert('primary')" color="primary" type="flat">Alert Primary</vs-button>
            <vs-button @click="openAlert('success')" color="success" type="flat">Alert Success</vs-button>
            <vs-button @click="openAlert('danger')" color="danger" type="flat">Alert Danger</vs-button>
            <vs-button @click="openAlert('warning')" color="warning" type="flat">Alert Warning</vs-button>
            <vs-button @click="openAlert('dark')" color="dark" type="flat">Alert Dark</vs-button>
            <vs-button @click="openAlert('#FF6F91')" color="primary" type="gradient">Alert Color RGB | HEX</vs-button>
        </div>
    </vx-card>
</template>

<script>
export default {
    data: () => ({
        colorAlert: 'primary',
        titleAlert: 'Alert',
        activeAlert: false,
        valueInput: '',
    }),
    methods: {
        openAlert(color) {
            this.colorAlert = color
            this.$vs.dialog({
                color: this.colorAlert,
                title: `Dialog - ${this.colorAlert}`,
                text: 'I love soufflé lollipop liquorice wafer jelly-o halvah sesame snaps. Pastry chocolate cake jelly-o carrot cake jelly topping croissant ice cream.',
                accept: this.acceptAlert
            })
        },
        acceptAlert() {
            this.$vs.notify({
                color: this.colorAlert,
                title: 'Accept Selected',
                text: 'Gingerbread soufflé biscuit oat cake.'
            })
        }
    }
}
</script>
