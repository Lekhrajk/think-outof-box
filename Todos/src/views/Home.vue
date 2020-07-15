<template>
  <div class="home">
    <div class="container-fluid">
      <h1 class="text-center text-primary title mt-3">TODOS Card-View</h1>
      <div class="clearfix">
        <div class="float-right btn-group">
          <router-link :to="{ name: 'home' }" class="btn btn-dark">Card-view</router-link>
          <router-link :to="{ name: 'list' }" class="btn btn-primary">List-view</router-link>
        </div>
        <div class="float-left">
          <div class="cardbox shadow-lg">
            <p class="total_count border-right" v-if="TodoCount">Total <span
                class="badge badge-primary">{{ TodoCount }}</span></p>
            <p class="total_count border-right" v-if="TrueLength">Completed <span class="badge badge-warning">
                {{ TrueLength.length }}</span> </p>
            <p class="total_count" v-if="FalseLength">InComplete <span
                class="badge badge-danger">{{ FalseLength.length }}</span></p>
          </div>
        </div>
      </div>
      <div class="row" v-if="todos">
        <div class="col-md-4 col-lg-3 col-12 mx-auto" v-for="todo in TData" :key="todo.id">
          <div class="card border-0 shadow-lg pb-0 mt-5">
            <router-link :to="{ path: '/todo/'+todo.id}">
              <Card :title="todo.title" :id="todo.id" :userId="todo.userId" :completed="todo.completed" />
              
            </router-link>
          </div>
        </div>
      </div>
      <div v-else class="text-center my-5">
        <img src="@/assets/load.gif" class="img-fluid" alt="wait.." width="200">
        <h3 class="text-center text-danger">Please wait...</h3>
      </div>
    </div>
    <div class="row">
        <div class="col-md-4 col-12 mx-auto">
          <div class="showmore border-0 shadow-lg">
              <button @click="limit = limit+8" class="btn btn-primary btn-block"><span class="load_more">Show more</span></button>
          </div>
        </div>
    </div>
  </div>
</template>

<script>
  import Card from './Card';
  export default {
    name: 'Home',
    data(){
      return{
        limit: 8,
      }
    },
    computed: {
      todos() {
        return this.$store.state.Todos;
      },
      TData(){
          return this.limit ? this.todos.slice(0,this.limit) : this.todos
      },
      TodoCount() {
        return this.$store.getters.TodoCount;
      },
      TrueLength() {
        return this.$store.state.Todos.filter(todo => todo.completed === true);
      },
      FalseLength() {
        return this.$store.state.Todos.filter(todo => todo.completed === false);
      }
    },
    mounted() {
      this.$store.dispatch('getTodos');
    },
    components: {
      Card
    },
  }
</script>
<style scoped>
  .card {
    border-bottom: 3px solid transparent !important;
    min-height: 200px;
  }

  .card:hover {
    border-bottom: 3px solid #0000ff !important;
    cursor: pointer;
  }

  a {
    text-decoration: none !important;
    color: #000;
  }

  a:hover {
    text-decoration: none !important;
    color: #000;
  }

  .btn {
    color: #fff !important;
  }
  .showmore{
    margin-top: 30px;
    margin-bottom: 20px;
  }
  @media (max-width:570px) {
    .title {
      font-size: 30px !important;
      margin-top: 30px;
      margin-bottom: 20px;
    }
  }

  .cardbox {
    padding: 10px 15px;
    display: inline-flex;
  }

  .total_count {
    margin: 4px 5px;
  }
  .load_more{
    font-size: 40px;
  }
</style>