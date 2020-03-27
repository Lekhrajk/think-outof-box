<!-- =========================================================================================
    File Name: CarouselDefault.vue
    Description: Default Carousel demo
    ----------------------------------------------------------------------------------------
    Item Name: Vuexy - Vuejs, HTML & Laravel Admin Dashboard Template
      Author: Pixinvent
    Author URL: http://www.themeforest.net/user/pixinvent
========================================================================================== -->


<template>
    <vx-card title="Default" class="carousel-example" code-toggler>
          <!-- swiper -->
          <swiper :dir="$vs.rtl ? 'rtl' : 'ltr'" :key="$vs.rtl">
            <swiper-slide>
              <img class="responsive" src="@/assets/images/pages/carousel/banner-18.jpg" alt="banner">
            </swiper-slide>
            <swiper-slide>
              <img class="responsive" src="@/assets/images/pages/carousel/banner-15.jpg" alt="banner">
            </swiper-slide>
            <swiper-slide>
              <img class="responsive" src="@/assets/images/pages/carousel/banner-10.jpg" alt="banner">
            </swiper-slide>
            <swiper-slide>
              <img class="responsive" src="@/assets/images/pages/carousel/banner-16.jpg" alt="banner">
            </swiper-slide>
          </swiper>

          <template slot="codeContainer">
&lt;template&gt;
    &lt;div class=&quot;carousel-example&quot;&gt;
          &lt;swiper :dir=&quot;$vs.rtl ? 'rtl' : 'ltr'&quot; :key=&quot;$vs.rtl&quot;&gt;
            &lt;swiper-slide&gt;
              &lt;img class=&quot;responsive&quot; src=&quot;@/assets/images/pages/carousel/banner-18.jpg&quot; alt=&quot;banner&quot;&gt;
            &lt;/swiper-slide&gt;
            &lt;swiper-slide&gt;
              &lt;img class=&quot;responsive&quot; src=&quot;@/assets/images/pages/carousel/banner-15.jpg&quot; alt=&quot;banner&quot;&gt;
            &lt;/swiper-slide&gt;
            &lt;swiper-slide&gt;
              &lt;img class=&quot;responsive&quot; src=&quot;@/assets/images/pages/carousel/banner-10.jpg&quot; alt=&quot;banner&quot;&gt;
            &lt;/swiper-slide&gt;
            &lt;swiper-slide&gt;
              &lt;img class=&quot;responsive&quot; src=&quot;@/assets/images/pages/carousel/banner-16.jpg&quot; alt=&quot;banner&quot;&gt;
            &lt;/swiper-slide&gt;
          &lt;/swiper&gt;
    &lt;/div&gt;
&lt;/template&gt;

&lt;script&gt;
import 'swiper/dist/css/swiper.min.css'
import { swiper, swiperSlide } from 'vue-awesome-swiper'

export default {
    components: {
        swiper,
        swiperSlide
    }
}
&lt;/script&gt;
          </template>
      </vx-card>
</template>

<script>
import 'swiper/dist/css/swiper.min.css'
import { swiper, swiperSlide } from 'vue-awesome-swiper'

export default{
    components: {
        swiper,
        swiperSlide
    }
}
</script>
