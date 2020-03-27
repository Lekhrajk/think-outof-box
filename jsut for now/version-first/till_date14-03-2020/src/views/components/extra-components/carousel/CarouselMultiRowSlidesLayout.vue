<!-- =========================================================================================
    File Name: CarouselMultiRowSlidesLayout.vue
    Description: Carousel with multiple row layout
    ----------------------------------------------------------------------------------------
    Item Name: Vuexy - Vuejs, HTML & Laravel Admin Dashboard Template
      Author: Pixinvent
    Author URL: http://www.themeforest.net/user/pixinvent
========================================================================================== -->


<template>
    <vx-card title="Multi Row Slides Layout" class="carousel-example" code-toggler>
          <!-- swiper -->
      <swiper :options="swiperOption" :dir="$vs.rtl ? 'rtl' : 'ltr'" :key="$vs.rtl">
        <swiper-slide>
          <img class="responsive" src="@/assets/images/pages/carousel/banner-31.jpg" alt="banner">
        </swiper-slide>
        <swiper-slide>
          <img class="responsive" src="@/assets/images/pages/carousel/banner-32.jpg" alt="banner">
        </swiper-slide>
        <swiper-slide>
          <img class="responsive" src="@/assets/images/pages/carousel/banner-33.jpg" alt="banner">
        </swiper-slide>
        <swiper-slide>
          <img class="responsive" src="@/assets/images/pages/carousel/banner-34.jpg" alt="banner">
        </swiper-slide>
        <swiper-slide>
          <img class="responsive" src="@/assets/images/pages/carousel/banner-35.jpg" alt="banner">
        </swiper-slide>
        <swiper-slide>
          <img class="responsive" src="@/assets/images/pages/carousel/banner-23.jpg" alt="banner">
        </swiper-slide>
        <swiper-slide>
          <img class="responsive" src="@/assets/images/pages/carousel/banner-37.jpg" alt="banner">
        </swiper-slide>
        <swiper-slide>
          <img class="responsive" src="@/assets/images/pages/carousel/banner-38.jpg" alt="banner">
        </swiper-slide>
        <swiper-slide>
          <img class="responsive" src="@/assets/images/pages/carousel/banner-39.jpg" alt="banner">
        </swiper-slide>
        <swiper-slide>
          <img class="responsive" src="@/assets/images/pages/carousel/banner-40.jpg" alt="banner">
        </swiper-slide>
        <div class="swiper-pagination" slot="pagination"></div>
      </swiper>

          <template slot="codeContainer">
&lt;template&gt;
    &lt;div class=&quot;carousel-example&quot;&gt;
        &lt;swiper :options=&quot;swiperOption&quot; :dir=&quot;$vs.rtl ? 'rtl' : 'ltr'&quot; :key=&quot;$vs.rtl&quot;&gt;
            &lt;swiper-slide&gt;
              &lt;img class=&quot;responsive&quot; src=&quot;@/assets/images/pages/carousel/banner-31.jpg&quot; alt=&quot;banner&quot;&gt;
            &lt;/swiper-slide&gt;
            &lt;swiper-slide&gt;
              &lt;img class=&quot;responsive&quot; src=&quot;@/assets/images/pages/carousel/banner-32.jpg&quot; alt=&quot;banner&quot;&gt;
            &lt;/swiper-slide&gt;
            &lt;swiper-slide&gt;
              &lt;img class=&quot;responsive&quot; src=&quot;@/assets/images/pages/carousel/banner-33.jpg&quot; alt=&quot;banner&quot;&gt;
            &lt;/swiper-slide&gt;
            &lt;swiper-slide&gt;
              &lt;img class=&quot;responsive&quot; src=&quot;@/assets/images/pages/carousel/banner-34.jpg&quot; alt=&quot;banner&quot;&gt;
            &lt;/swiper-slide&gt;
            &lt;swiper-slide&gt;
              &lt;img class=&quot;responsive&quot; src=&quot;@/assets/images/pages/carousel/banner-35.jpg&quot; alt=&quot;banner&quot;&gt;
            &lt;/swiper-slide&gt;
            &lt;swiper-slide&gt;
              &lt;img class=&quot;responsive&quot; src=&quot;@/assets/images/pages/carousel/banner-23.jpg&quot; alt=&quot;banner&quot;&gt;
            &lt;/swiper-slide&gt;
            &lt;swiper-slide&gt;
              &lt;img class=&quot;responsive&quot; src=&quot;@/assets/images/pages/carousel/banner-37.jpg&quot; alt=&quot;banner&quot;&gt;
            &lt;/swiper-slide&gt;
            &lt;swiper-slide&gt;
              &lt;img class=&quot;responsive&quot; src=&quot;@/assets/images/pages/carousel/banner-38.jpg&quot; alt=&quot;banner&quot;&gt;
            &lt;/swiper-slide&gt;
            &lt;swiper-slide&gt;
              &lt;img class=&quot;responsive&quot; src=&quot;@/assets/images/pages/carousel/banner-39.jpg&quot; alt=&quot;banner&quot;&gt;
            &lt;/swiper-slide&gt;
            &lt;swiper-slide&gt;
              &lt;img class=&quot;responsive&quot; src=&quot;@/assets/images/pages/carousel/banner-40.jpg&quot; alt=&quot;banner&quot;&gt;
            &lt;/swiper-slide&gt;
            &lt;div class=&quot;swiper-pagination&quot; slot=&quot;pagination&quot;&gt;&lt;/div&gt;
        &lt;/swiper&gt;
    &lt;/div&gt;
&lt;/template&gt;

&lt;script&gt;
import 'swiper/dist/css/swiper.min.css'
import { swiper, swiperSlide } from 'vue-awesome-swiper'

export default {
    data() {
        return {
            swiperOption: {
                slidesPerView: 3,
                slidesPerColumn: 2,
                spaceBetween: 30,
                pagination: {
                    el: '.swiper-pagination',
                    clickable: true
                },
                breakpoints: {
                    1024: {
                        slidesPerView: 3,
                        spaceBetween: 40
                    },
                    768: {
                        slidesPerView: 2,
                        spaceBetween: 30
                    },
                    640: {
                        slidesPerView: 1,
                        spaceBetween: 20
                    }
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

&lt;style lang=&quot;scss&quot;&gt;
.carousel-example .swiper-container .swiper-slide {
    text-align: center;
    font-size: 38px;
    font-weight: 700;
    background-color: #eee;
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex;
    -webkit-box-pack: center;
    -ms-flex-pack: center;
    justify-content: center;
    -webkit-box-align: center;
    -ms-flex-align: center;
    align-items: center;
    min-height: 300px;
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
                slidesPerView: 3,
                slidesPerColumn: 2,
                spaceBetween: 30,
                pagination: {
                    el: '.swiper-pagination',
                    clickable: true
                },
                breakpoints: {
                    1024: {
                        slidesPerView: 3,
                        spaceBetween: 40
                    },
                    768: {
                        slidesPerView: 2,
                        spaceBetween: 30
                    },
                    640: {
                        slidesPerView: 1,
                        spaceBetween: 20
                    }
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
