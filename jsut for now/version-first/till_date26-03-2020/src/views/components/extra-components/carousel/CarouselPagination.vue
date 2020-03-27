<!-- =========================================================================================
    File Name: CarouselPagination.vue
    Description: Carousel with pagination
    ----------------------------------------------------------------------------------------
    Item Name: Vuexy - Vuejs, HTML & Laravel Admin Dashboard Template
      Author: Pixinvent
    Author URL: http://www.themeforest.net/user/pixinvent
========================================================================================== -->


<template>
    <vx-card title="Pagination" class="carousel-example" code-toggler>
          <!-- swiper -->
          <swiper :options="swiperOption" :dir="$vs.rtl ? 'rtl' : 'ltr'" :key="$vs.rtl">
            <swiper-slide>
              <img class="responsive" src="@/assets/images/pages/carousel/banner-16.jpg" alt="banner">
            </swiper-slide>
            <swiper-slide>
              <img class="responsive" src="@/assets/images/pages/carousel/banner-19.jpg" alt="banner">
            </swiper-slide>
            <swiper-slide>
              <img class="responsive" src="@/assets/images/pages/carousel/banner-5.jpg" alt="banner">
            </swiper-slide>
            <swiper-slide>
              <img class="responsive" src="@/assets/images/pages/carousel/banner-9.jpg" alt="banner">
            </swiper-slide>
            <div class="swiper-pagination" slot="pagination"></div>
          </swiper>

          <template slot="codeContainer">
&lt;template&gt;
    &lt;div class=&quot;carousel-example&quot; :dir=&quot;$vs.rtl ? 'rtl' : 'ltr'&quot; :key=&quot;$vs.rtl&quot;&gt;
        &lt;swiper :options=&quot;swiperOption&quot;&gt;
            &lt;swiper-slide&gt;
              &lt;img class=&quot;responsive&quot; src=&quot;@/assets/images/pages/carousel/banner-16.jpg&quot; alt=&quot;banner&quot;&gt;
            &lt;/swiper-slide&gt;
            &lt;swiper-slide&gt;
              &lt;img class=&quot;responsive&quot; src=&quot;@/assets/images/pages/carousel/banner-19.jpg&quot; alt=&quot;banner&quot;&gt;
            &lt;/swiper-slide&gt;
            &lt;swiper-slide&gt;
              &lt;img class=&quot;responsive&quot; src=&quot;@/assets/images/pages/carousel/banner-5.jpg&quot; alt=&quot;banner&quot;&gt;
            &lt;/swiper-slide&gt;
            &lt;swiper-slide&gt;
              &lt;img class=&quot;responsive&quot; src=&quot;@/assets/images/pages/carousel/banner-9.jpg&quot; alt=&quot;banner&quot;&gt;
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

