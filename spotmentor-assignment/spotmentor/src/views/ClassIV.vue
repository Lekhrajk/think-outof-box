<template>

    <!-- class first component  -->

    <div class="classI">
        <!--Header  -->
        <div class="row border-bottom border-primary clearfix my-md-4">
            <div class="col-md-4 col-12 text-left">
                <h3>{{ getData.classname }}-Performance</h3>
                <span class="text-primary">Total Students-{{ getData.students.length }}</span>
            </div>
            <div class="col-md-8 col-12 text-right">
                <button class="lbtn lbtn-blue-color" @click="calculateAvg">{{ text }} <i :class="[!isActive ? 'fa-eye' : 'fa-eye-slash', 'fa']"></i></button>
            </div>
            <div class="col-10 mx-auto mt-3">
                <transition name="fade">
                    <div class="row" v-if="isActive">
                        <div class="card col-md-6 mx-auto p-3 border-0 mb-1">
                            <span class="text-center h4"> Class Average score is {{ avg }}% <i :class="[failure ? 'fa-frown' : 'fa-smile', 'far','fa-2x','text-success']"></i></span>
                            <small class="text-center">{{ success }}</small>
                        </div>
                        <div class="card col-md-6 float-right text-right p-3 border-0 mb-1">
                            <h5>Maximum Marks - {{ maxVal }}</h5>
                            <h5 class="text-danger">Minimum Marks - {{ minVal }}</h5>
                            <div class="col float-right">
                              <button class="lbtn lbtn-light-color lbtn-small" data-toggle="modal" data-target="#analysis">Analysis</button>
                            </div>
                        </div>
                    </div>
                </transition>
            </div>
        </div>
        <!-- The Modal -->
        <div class="modal fade" id="analysis">
          <div class="modal-dialog modal-dialog-centered">
            <div class="modal-content border-0">
              <div class="modal-body">
                <img class="img-fluid " src="@/assets/img/report.svg"  width="500"/>  
                <h3 class="text-left">Student Analysis Report</h3>
                <div class="row mx-auto mt-3 text-center">
                    <div class="col-md-8 col-10 p-0">
                        <div class="p-md-2 p-0" v-for="user in getData.students" v-bind:key="user.id">
                            Total percentage for {{ user.name }} is
                        </div>
                    </div>
                    <div class="col-2 p-0">
                        <div class="p-md-2 p-0" v-for="mark in marklist" :key="mark.id">
                            {{ mark }}%
                        </div>
                    </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <!-- end of Header -->

        <!-- Body starts here -->

        <div class="container">
            <div class="row">
                <div class="col-md-6 col-12" v-for="user in getData.students" v-bind:key="user.id">
                    <div class="cardbox mx-auto">
                        <div class="p-3">
                            <div class="mb-4 border-bottom border-primary">
                                <span class="text-capitalize h5">{{ user.name }}</span>
                            </div>
                            <div class="row my-3">
                                <div class="col-3">
                                    <div class="text-left ">
                                        <span>Math</span>
                                    </div>
                                </div>
                                <div class="col-7 mt-1">
                                    <div class="progress" style="height:12px">
                                        <div class="progress-bar" v-bind:style="{width: user.marks.Maths + '%'}"></div>
                                    </div>
                                </div>
                                <div class="col-2">
                                    <div>{{ user.marks.Maths }}%</div>
                                </div>
                            </div>
                            <div class="row mb-3">
                                <div class="col-3">
                                    <div class="text-left ">
                                        <span>Science</span>
                                    </div>
                                </div>
                                <div class="col-7 mt-1">
                                    <div class="progress" style="height:12px">
                                        <div class="progress-bar" v-bind:style="{width: user.marks.Science + '%'}"></div>
                                    </div>
                                </div>
                                <div class="col-2">
                                    <div v-bind="science">{{ user.marks.Science }}%</div>
                                </div>
                            </div>
                            <div class="row">
                                <div class="col-3">
                                    <div class="text-left ">
                                        <span>English</span>
                                    </div>
                                </div>
                                <div class="col-7 mt-1">
                                    <div class="progress" style="height:12px">
                                        <div class="progress-bar" v-bind:style="{width: user.marks.English + '%'}"></div>
                                    </div>
                                </div>
                                <div class="col-2">
                                    <div v-bind="english">{{ user.marks.English}}%</div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>

        <!-- end of body -->

    </div>
</template>
<script>
// import {mapGetters} from 'vuex';
import store from '../store/index';
export default {
    name: 'classI',
    data() {
        return {
            marklist: [],
            math: '',
            science: '',
            english: '',
            avg: '',
            maxVal: '',
            minVal: '',
            success: '',
            failure: false,
            isActive: false,
            seen: true,
            text: 'Show Average',
        }
    },
    computed: {
        getData() {
            var temp_class = this.$route.path;
            var classurl = temp_class.substring(temp_class.lastIndexOf("/") + 1);
            return this.$store.getters.getClassData(classurl);
        }
    },
    mounted() {
        setTimeout(() => {
            this.averageMarks();
            this.insertPercentage(this.marklist);
            this.MinMax();
        }, 1000);
    },
    methods: {
        averageMarks() {
            for (var i = 0; i < this.getData.students.length; i++) {
                var rdata = this.getData.students[i];
                var rmath = rdata.marks.Maths;
                var rscience = rdata.marks.Science;
                var renglish = rdata.marks.English;
                var total = parseInt((rmath + rscience + renglish) / 3);
                this.marklist.push(total);
            }
        },
        insertPercentage(marklist) {
            store.dispatch("updatePercentage", { marklist });
        },
        calculateAvg() {
            this.isActive = !this.isActive;
            this.seen = !this.seen;
            this.text = this.seen ? 'Show Average' : 'Hide Average';
            var sum = 0;
            for (var i = 0; i < this.marklist.length; i++) {
                sum = sum + this.marklist[i];
            }
            var avg = sum / this.marklist.length;
            this.avg = avg.toFixed(2);
            if (this.avg < 40) {
                this.failure = true;
                this.success = "Oops better luck next time";
            } else {
                this.success = 'WOW!.. Great work by Class';
            }
        },
        MinMax() {
            this.maxVal = Math.max.apply(Math, this.marklist) + '%';
            this.minVal = Math.min.apply(Math, this.marklist) + '%';
        }
    }
}
</script>
<style scoped>

.fade-enter-active,
.fade-leave-active {
    transition: opacity .5s;
}

.fade-enter,
.fade-leave-to {
    opacity: 0;
}
</style>