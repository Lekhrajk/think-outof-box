<template>
    <div class="create">
        <div class="modal" id="create">
            <div class="modal-dialog">
                <div class="modal-content">
                    <!-- Modal Header -->
                    <div class="modal-header">
                        <h2 class="text-center modal-title">Add New Task</h2>
                        <button type="button" class="close" data-dismiss="modal">&times;</button>
                    </div>
                    <div class="modal-body">
                        <div class="row">
                            <div class="col-md-12 col-12 mx-auto">
                                <div class="card shadow-lg border-0 p-4">
                                    <form>
                                        <div class="form-row my-4">
                                            <div class="col">
                                                <div class="input-group">
                                                    <div class="input-group-prepend">
                                                        <div class="input-group-text"><i class="fas fa-id-card"></i>
                                                        </div>
                                                    </div>
                                                    <input type="number" class="form-control" name="id" readonly
                                                        v-model="number" required />
                                                </div>
                                            </div>
                                            <div class="col">
                                                <div class="input-group">
                                                    <div class="input-group-prepend">
                                                        <div class="input-group-text"><i class="fas fa-id-badge"></i>
                                                        </div>
                                                    </div>
                                                    <input type="text" v-model="userId" class="form-control"
                                                        placeholder="Enter user id" name="userid" required />
                                                </div>
                                            </div>
                                        </div>
                                        <div class="form-row mb-4">
                                            <div class="col">
                                                <div class="input-group">
                                                    <div class="input-group-prepend">
                                                        <div class="input-group-text"><i
                                                                class="fas fa-clipboard-list"></i></div>
                                                    </div>
                                                    <input type="title" v-model="title" class="form-control"
                                                        placeholder="Enter title" name="title" required />
                                                </div>
                                            </div>
                                            <div class="col">
                                                <div class="input-group">
                                                    <div class="input-group-prepend">
                                                        <div class="input-group-text"><i class="fas fa-question"></i>
                                                        </div>
                                                    </div>
                                                    <select name="status" v-model="status" class="form-control"
                                                        required>
                                                        <option value="true" selected>Completed</option>
                                                        <option value="false">Not Complete</option>
                                                    </select>
                                                </div>
                                            </div>
                                        </div>
                                        <div class="text-center">
                                            <button class="btn btn-dark" @click.prevent="createTodo" data-dismiss="modal">Add</button>
                                        </div>
                                    </form>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    export default {
        name: "create",
        data() {
            return {
                total: "",
                userId: "",
                title: "",
                status: "",
                number: "",
                max: 9999,
                min: 200
            }
        },
        mounted() {
            this.getRandomNumber();
        },
        methods: {
            createTodo() {
                this.$store.dispatch("addTodo", {
                    id: this.number,
                    userId: this.userId,
                    title: this.title,
                    completed: this.status
                })
            },
            getRandomNumber: function () {
                this.number = this.generateNumber();
            },
            generateNumber: function () {
                return Math.floor(Math.random() * (this.max - this.min + 1) + this.min);
            }
        },
    }
</script>

<style scoped>
    .form-control {
        border: 2px solid #000;
    }

    .form-control:focus,
    .form-control.active {
        border: 2px solid #000000 !important;
    }

    textarea:focus,
    textarea.form-control:focus,
    input.form-control:focus,
    input[type=text]:focus,
    input[type=password]:focus,
    input[type=email]:focus,
    input[type=number]:focus,
    [type=text].form-control:focus,
    [type=password].form-control:focus,
    [type=email].form-control:focus,
    [type=tel].form-control:focus,
    [contenteditable].form-control:focus {
        box-shadow: inset 0 -1px 0 #ddd;
    }
</style>