<template>
    <div class="list container">
        <h1 class="text-center text-primary title mt-3">TODOS List-View</h1>
        <div class="clearfix">
            <div class="float-right btn-group">
                <router-link :to="{ name: 'home' }" class="btn btn-dark">Card-view</router-link>
                <router-link :to="{ name: 'list' }" class="btn btn-primary">List-view</router-link>
            </div>
            <div class="float-left">
                <div class="cardbox shadow-lg">
                    <p class="total_count border-right">Total <span class="badge badge-primary">{{ TodoCount }}</span>
                    </p>
                    <p class="total_count border-right">Completed <span class="badge badge-warning">
                            {{ TrueLength.length }}</span> </p>
                    <p class="total_count">InComplete <span class="badge badge-danger">{{ FalseLength.length }}</span>
                    </p>
                </div>
            </div>
        </div>
        <div class="table-responsive card border-0 shadow-lg">
            <table class="table">
                <thead class="thead-light">
                    <tr>
                        <th scope="col">SNo.</th>
                        <th scope="col">User-ID</th>
                        <th scope="col">Title</th>
                        <th scope="col">Status</th>
                        <th scope="col">Action</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="todo in todos" :key="todo.id" class="todo">
                        <td>{{ todo.id }}</td>
                        <td>{{ todo.userId }}</td>
                        <td>{{ todo.title }}</td>
                        <td>{{ todo.completed }}</td>
                        <td>
                            <router-link :to="{ path: '/todo/'+todo.id}" class="view"><i class="fas fa-eye"></i>
                            </router-link>
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
</template>

<script>
    export default {
        name: "list",
        computed: {
            todos() {
                return this.$store.state.Todos;
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
    }
</script>

<style scoped>
    .view {
        opacity: 0;
    }

    .todo:hover .view {
        opacity: 1;
    }

    @media (max-width:570px) {
        .title {
            font-size: 30px !important;
            margin-top: 30px;
            margin-bottom: 20px;
        }

        .container {
            margin: 0px !important;
        }
    }

    .cardbox {
        padding: 10px 15px;
        display: inline-flex;
    }

    .total_count {
        margin: 4px 5px;
    }
</style>