<template>
    <vx-card title="Memory">
        <div class="mt-5">
            <label>Select Memory</label>
            <vs-slider color="dark" text-fixed=GB ticks max="8" v-model="value1"/>
            <p>{{ value1 }} GB-RAM</p><span class="hidden">{{ value1*1024 }}</span>
        </div>
    </vx-card>
</template>

<script>
export default {
    data() {
        return {
            value1:2,
        }
    }
}
</script>
