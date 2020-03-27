<!-- =========================================================================================
    File Name: Carousel3dCoverflowEffect.vue
    Description: Carousel with 3D Coverflow effect
    ----------------------------------------------------------------------------------------
    Item Name: Vuexy - Vuejs, HTML & Laravel Admin Dashboard Template
      Author: Pixinvent
    Author URL: http://www.themeforest.net/user/pixinvent
========================================================================================== -->


<template>
    <vx-card title="3d Effect Coverflow Effect" class="carousel-example" code-toggler>
        <div class="swiper-inner">
            <!-- swiper -->
            <swiper :options="swiperOption" :dir="$vs.rtl ? 'rtl' : 'ltr'" :key="$vs.rtl">
                <swiper-slide>
                  <img class="responsive" src="@/assets/images/pages/carousel/banner-39.jpg" alt="banner">
                </swiper-slide>
                <swiper-slide>
                  <img class="responsive" src="@/assets/images/pages/carousel/banner-35.jpg" alt="banner">
                </swiper-slide>
                <swiper-slide>
                  <img class="responsive" src="@/assets/images/pages/carousel/banner-37.jpg" alt="banner">
                </swiper-slide>
                <swiper-slide>
                  <img class="responsive" src="@/assets/images/pages/carousel/banner-38.jpg" alt="banner">
                </swiper-slide>
                <swiper-slide>
                  <img class="responsive" src="@/assets/images/pages/carousel/banner-33.jpg" alt="banner">
                </swiper-slide>
                <swiper-slide>
                  <img class="responsive" src="@/assets/images/pages/carousel/banner-40.jpg" alt="banner">
                </swiper-slide>
                <swiper-slide>
                  <img class="responsive" src="@/assets/images/pages/carousel/banner-32.jpg" alt="banner">
                </swiper-slide>
                <div class="swiper-pagination" slot="pagination"></div>
            </swiper>
        </div>
        <template slot="codeContainer">
&lt;template&gt;
    &lt;div class=&quot;carousel-example&quot;&gt;
        &lt;div class=&quot;swiper-inner&quot;&gt;
            &lt;!-- swiper --&gt;
            &lt;swiper :options=&quot;swiperOption&quot; :dir=&quot;$vs.rtl ? 'rtl' : 'ltr'&quot; :key=&quot;$vs.rtl&quot;&gt;
                &lt;swiper-slide&gt;
                  &lt;img class=&quot;responsive&quot; src=&quot;@/assets/images/pages/carousel/banner-39.jpg&quot; alt=&quot;banner&quot;&gt;
                &lt;/swiper-slide&gt;
                &lt;swiper-slide&gt;
                  &lt;img class=&quot;responsive&quot; src=&quot;@/assets/images/pages/carousel/banner-35.jpg&quot; alt=&quot;banner&quot;&gt;
                &lt;/swiper-slide&gt;
                &lt;swiper-slide&gt;
                  &lt;img class=&quot;responsive&quot; src=&quot;@/assets/images/pages/carousel/banner-37.jpg&quot; alt=&quot;banner&quot;&gt;
                &lt;/swiper-slide&gt;
                &lt;swiper-slide&gt;
                  &lt;img class=&quot;responsive&quot; src=&quot;@/assets/images/pages/carousel/banner-38.jpg&quot; alt=&quot;banner&quot;&gt;
                &lt;/swiper-slide&gt;
                &lt;swiper-slide&gt;
                  &lt;img class=&quot;responsive&quot; src=&quot;@/assets/images/pages/carousel/banner-33.jpg&quot; alt=&quot;banner&quot;&gt;
                &lt;/swiper-slide&gt;
                &lt;swiper-slide&gt;
                  &lt;img class=&quot;responsive&quot; src=&quot;@/assets/images/pages/carousel/banner-40.jpg&quot; alt=&quot;banner&quot;&gt;
                &lt;/swiper-slide&gt;
                &lt;swiper-slide&gt;
                  &lt;img class=&quot;responsive&quot; src=&quot;@/assets/images/pages/carousel/banner-32.jpg&quot; alt=&quot;banner&quot;&gt;
                &lt;/swiper-slide&gt;
                &lt;div class=&quot;swiper-pagination&quot; slot=&quot;pagination&quot;&gt;&lt;/div&gt;
            &lt;/swiper&gt;
        &lt;/div&gt;
    &lt;/div&gt;
&lt;/template&gt;

&lt;script&gt;
import 'swiper/dist/css/swiper.min.css'
import { swiper, swiperSlide } from 'vue-awesome-swiper'

export default {
    data() {
        return {
            swiperOption: {
                effect: 'cube',
                grabCursor: true,
                cubeEffect: {
                    shadow: true,
                    slideShadows: true,
                    shadowOffset: 20,
                    shadowScale: 0.94
                },
                pagination: {
                    el: '.swiper-pagination'
                }
            }
        }
    },
    components: {
        swiper,
        swiperSlide
    }
}
&lt;/script&gt;

&lt;style scoped&gt;
.swiper-inner {
    position: relative;
    overflow: hidden;
    height: 330px;
    padding: 15px;
}

.swiper-container {
    width: 300px !important;
    height: 300px;
    position: absolute;
    left: 50%;
    top: 50%;
    margin-left: -150px;
    margin-top: -150px;
}

.swiper-slide {
    background-position: center;
    background-size: cover;
}
&lt;/style&gt;
        </template>
    </vx-card>
</template>

<script>
import 'swiper/dist/css/swiper.min.css'
import { swiper, swiperSlide } from 'vue-awesome-swiper'

export default {
    data() {
        return {
            swiperOption: {
                effect: 'coverflow',
                grabCursor: true,
                centeredSlides: true,
                slidesPerView: 'auto',
                coverflowEffect: {
                    rotate: 50,
                    stretch: 0,
                    depth: 100,
                    modifier: 1,
                    slideShadows: true
                },
                pagination: {
                    el: '.swiper-pagination'
                }
            }
        }
    },
    components: {
        swiper,
        swiperSlide
    }
}
</script>

<style scoped>
  .swiper-inner {
    width: 100%;
    height: 400px;
    padding-top: 50px;
    padding-bottom: 50px;
  }
  .swiper-slide {
    background-position: center;
    background-size: cover;
    width: 300px;
    height: 300px;
  }
</style>
