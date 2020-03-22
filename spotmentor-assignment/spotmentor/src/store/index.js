import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        totalStudents: [],
        classes: [{
                "classname": "Class I",
                "percentage": [],
                "students": [

                    {
                        "name": "student 1",
                        "marks": {
                            "Maths": 35,
                            "Science": 64,
                            "English": 42,
                        },
                    },
                    {
                        "name": "student 2",
                        "marks": {
                            "Maths": 20,
                            "Science": 100,
                            "English": 62
                        },
                    },
                    {
                        "name": "student 3",
                        "marks": {
                            "Maths": 60,
                            "Science": 90,
                            "English": 62
                        },
                    },
                    {
                        "name": "student 4",
                        "marks": {
                            "Maths": 90,
                            "Science": 99,
                            "English": 66
                        },
                    },
                ]
            },
            {
                "classname": "Class II",
                "percentage": [],
                "students": [{
                        "name": "student 5",
                        "marks": {
                            "Maths": 34,
                            "Science": 10,
                            "English": 42
                        },
                    },
                    {
                        "name": "student 6",
                        "marks": {
                            "Maths":  33,
                            "Science": 27,
                            "English": 42
                        },
                    },
                    {
                        "name": "student 7",
                        "marks": {
                            "Maths": 23,
                            "Science": 24,
                            "English": 27
                        },
                    },
                    {
                        "name": "student 8",
                        "marks": {
                            "Maths": 63,
                            "Science": 34,
                            "English": 29
                        },
                    }
                ]
            },
            {
                "classname": "Class III",
                "percentage": [],
                "students": [{
                        "name": "student 9",
                        "marks": {
                            "Maths": 74,
                            "Science": 90,
                            "English": 62
                        },
                    },
                    {
                        "name": "student 10",
                        "marks": {
                            "Maths": 63,
                            "Science": 57,
                            "English": 92
                        },
                    },
                    {
                        "name": "student 11",
                        "marks": {
                            "Maths": 53,
                            "Science": 34,
                            "English": 77
                        },
                    }
                ]
            },
            {
                "classname": "Class IV",
                "percentage": [],
                "students": [{
                        "name": "student 12",
                        "marks": {
                            "Maths": 54,
                            "Science": 70,
                            "English": 32
                        },
                    },
                    {
                        "name": "student 13",
                        "marks": {
                            "Maths": 93,
                            "Science": 87,
                            "English": 82
                        },
                    },
                    {
                        "name": "student 14",
                        "marks": {
                            "Maths": 68,
                            "Science": 97,
                            "English": 61
                        },
                    }
                ]
            },
            {
                "classname": "Class V",
                "percentage": [],
                "students": [{
                        "name": "student 15",
                        "marks": {
                            "Maths": 65,
                            "Science": 88,
                            "English": 46
                        },
                    },
                    {
                        "name": "student 16",
                        "marks": {
                            "Maths": 43,
                            "Science": 47,
                            "English": 52
                        },
                    },
                    {
                        "name": "student 17",
                        "marks": {
                            "Maths": 73,
                            "Science": 95,
                            "English": 62
                        },
                    }
                ]
            },
            {
                "classname": "Class VI",
                "percentage": [],
                "students": [{
                        "name": "student 18",
                        "marks": {
                            "Maths": 24,
                            "Science": 30,
                            "English": 42
                        },
                    },
                    {
                        "name": "student 19",
                        "marks": {
                            "Maths": 59,
                            "Science": 70,
                            "English": 20
                        },
                    },
                    {
                        "name": "student 20",
                        "marks": {
                            "Maths": 67,
                            "Science": 74,
                            "English": 67
                        },
                    }
                ]
            },
            {
                "classname": "Class VII",
                "percentage": [],
                "students": [{
                        "name": "student 21",
                        "marks": {
                            "Maths": 84,
                            "Science": 50,
                            "English": 82
                        },
                    },
                    {
                        "name": "student 22",
                        "marks": {
                            "Maths": 59,
                            "Science": 79,
                            "English": 70
                        },
                    },
                ]
            }
        ]
    },
    mutations: {
        updatePercentage(state, payload) {
            state.classes[0].percentage.push(payload);
        },
    },
    actions: {
        updatePercentage(context, data) {
            context.commit('updatePercentage', data);
        },
    },
    getters: {
        getClassData(state) {
            return keyword => state.classes.find(classes =>{
                return classes.classname === keyword
            });
        },
        getClassCount(state) {
            return state.classes.length;
        },
    }
})