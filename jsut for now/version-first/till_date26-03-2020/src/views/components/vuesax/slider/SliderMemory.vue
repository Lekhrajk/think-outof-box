<template>
    <vx-card title="Storage memory">
        <div class="mt-5">
            <label>Select Storage(Disks)</label>
            <vs-slider color="dark" text-fixed=GB ticks max="1024" v-model="value2"/>
            <p>{{ value2 }} GB</p><span class="hidden">{{ value2*1024 }}</span>
        </div>
    </vx-card>
</template>

<script>
export default {
    data() {
        return {
            value2:30,
        }
    }
}
</script>
