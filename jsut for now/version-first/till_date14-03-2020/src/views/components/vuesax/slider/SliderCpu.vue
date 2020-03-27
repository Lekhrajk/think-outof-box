<template>
    <vx-card title="CPU">
        <div class="mt-5">
            <label>Select VCPU</label>
            <vs-slider color="dark" text-fixed=vCPU ticks max=4 v-model="value3"/>
            <p>{{ value3 }} VCPU</p>
        </div>
    </vx-card>
</template>

<script>
export default {
    data() {
        return {
            value3:1,
        }
    }
}
</script>
