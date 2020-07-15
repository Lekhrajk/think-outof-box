<template>
  <div class="CARD">
    <div class="cards">
      <div class="Top-heading">
        <h3>TASK NO : <span>{{ id }}</span></h3>
      </div>
      <div class="main-content px-4">
        <div class="row">
          <div class="col-sm-4 col-6">
            <h6 class="font-weight-bold">User-id</h6>
          </div>
          <div class="col-sm-6 col-6 float-left">
            <h6>{{ userId }}</h6>
          </div>
        </div>
        <div class="row">
          <div class="col-sm-4 col-4">
            <h6 class="font-weight-bold">Title</h6>
          </div>
          <div class="col-sm-8 col-8">
            <h6>{{ title }}</h6>
          </div>
        </div>
        <div class="row">
          <div class="col-sm-4 col-4">
            <h6 class="font-weight-bold">Status</h6>
          </div>
          <div class="col-sm-8 col-8">
            <h6>{{ completed }}</h6>
          </div>
        </div>
        <div class="float-right mb-2">
          <a href="#" class="badge badge-danger" @click.prevent="removeTodo(id)">Remove <i class="fas fa-trash"></i></a>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    name: "name",
    props: ['title', 'id', 'completed', 'userId'],
    methods: {
      removeTodo(id) {
        this.$store.dispatch('removeTodo', id);
      },
    }
  }
</script>
<style scoped>
  .Top-heading {
    background: #1C2833;
    color: #ffffff;
    padding: 8px 15px;
  }

  .main-content {
    padding: 5px 8px;
    margin-top: 5px;
  }
</style>