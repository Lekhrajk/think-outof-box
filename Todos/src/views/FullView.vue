<template>
    <div class="fullview">
        <div class="clearfix my-4"><button class="btn btn-outline-dark btn-sm float-right" @click="$router.go(-1)">GO
                back</button></div>
        <div class="container">
            <div class="row">
                <div class="col-md-7 col-12 mx-auto">
                    <div class="card shadow-lg border-0">
                        <div class="row">
                            <div class="col-md-6 col-12">
                                <img src="@/assets/images/todo.svg" class="img-fluid my-auto" />
                            </div>
                            <div class="col-md-6 col-12" v-if="Todo">
                                <div class="Top-heading">
                                    <h3>TASK NO : <span>{{ Todo.id }}</span></h3>
                                </div>
                                <div class="main-content px-4">
                                    <div class="row my-4">
                                        <div class="col-sm-4 col-6">
                                            <h6 class="font-weight-bold">User-id</h6>
                                        </div>
                                        <div class="col-sm-6 col-6 float-left">
                                            <h6>{{ Todo.userId }}</h6>
                                        </div>
                                    </div>
                                    <div class="row mb-4">
                                        <div class="col-sm-4 col-4">
                                            <h6 class="font-weight-bold">Title</h6>
                                        </div>
                                        <div class="col-sm-8 col-8">
                                            <h6>{{ Todo.title }}</h6>
                                        </div>
                                    </div>
                                    <div class="row">
                                        <div class="col-sm-4 col-4">
                                            <h6 class="font-weight-bold">Status</h6>
                                        </div>
                                        <div class="col-sm-8 col-8">
                                            <h6>{{ Todo.completed }}</h6>
                                        </div>
                                    </div>
                                </div>
                            </div>
                            <div v-else class="text-center my-5">
                                <img src="@/assets/load.gif" class="img-fluid" alt="wait.." width="200">
                                <h3 class="text-center text-danger">No Data Found</h3>
                            </div>
                        </div>
                        <div class="table-responsive card border-0 shadow-lg">
                            <table class="table">
                                <thead class="thead-light">
                                    <tr v-if="UserTask">
                                        <th scope="col">TOTAL</th>
                                        <td>{{ UserTask.length }} </td>
                                    </tr>
                                    <tr v-if="Completed">
                                        <th scope="col">COMPLETED</th>
                                        <td>{{ Completed.length }}</td>
                                    </tr>
                                    <tr v-if="InCompleted">
                                        <th scope="col">REMAINING</th>
                                        <td>{{ InCompleted.length }} </td>
                                    </tr>
                                </thead>
                            </table>
                            <button class="btn btn-danger btn-block" @click="ShowData">{{ text }} <i :class="[!isActive ? 'fa-eye' : 'fa-eye-slash', 'fa','ml-3']"></i></button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div class="row" v-if="show">
            <div class="col-md-4 col-lg-3 col-12 mx-auto" v-for="task in UserTask" :key="task.id">
                <div class="card border-0 shadow-lg pb-0 mt-5">
                    <Card :title="task.title" :id="task.id" :userId="task.userId" :completed="task.completed" />
                </div>
            </div>
        </div>
        <!-- <div v-else class="text-center my-5">
                <img src="@/assets/load.gif" class="img-fluid" alt="wait.." width="200">
                <h3 class="text-center text-danger">Please wait...</h3>
        </div> -->
    </div>
</template>

<script>
import Card from './Card';
    export default {
        name: "fullview",
        data() {
            return {
                Req_id: "",
                text: "Show Tasks",
                show: false,
                isActive: false,
                seen: true,
            }
        },
        beforeMount() {
            this.Req_id = this.$route.params.id;
        },
        computed: {
            Todo() {
                return this.$store.state.Todo;
            },
            UserTask() {
                return this.$store.state.Todos.filter(todo => todo.userId === this.Todo.userId);
            },
            Completed() {
                return this.UserTask.filter(com => com.completed === true);
            },
            InCompleted() {
                return this.UserTask.filter(com => com.completed === false);
            }
        },
        mounted() {
            this.$store.dispatch("getTodo", this.Req_id);
        },
        methods:{
            ShowData(){
                this.isActive = !this.isActive;
                this.show = !this.show;
                this.seen = !this.seen;
                this.text = this.seen ? 'Show Tasks' : 'Hide Tasks';
            }
        },
        components:{
            Card
        }
    }
</script>
<style scoped>
    .fullview {
        min-height: 70vh;
    }

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