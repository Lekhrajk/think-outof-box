<!-- =========================================================================================
	File Name: CarouselEffectFade.vue
	Description: Carousel with fade effect
	----------------------------------------------------------------------------------------
	Item Name: Vuexy - Vuejs, HTML & Laravel Admin Dashboard Template
	Author: Pixinvent
	Author URL: http://www.themeforest.net/user/pixinvent
========================================================================================== -->


<template>
    <vx-card title="Fade Effect" class="carousel-example" code-toggler>
        <!-- swiper -->
        <swiper :options="swiperOption" :dir="$vs.rtl ? 'rtl' : 'ltr'" :key="$vs.rtl">
            <swiper-slide>
              <img class="responsive" src="@/assets/images/pages/carousel/banner-16.jpg" alt="banner">
            </swiper-slide>
            <swiper-slide>
              <img class="responsive" src="@/assets/images/pages/carousel/banner-10.jpg" alt="banner">
            </swiper-slide>
            <swiper-slide>
              <img class="responsive" src="@/assets/images/pages/carousel/banner-15.jpg" alt="banner">
            </swiper-slide>
            <swiper-slide>
              <img class="responsive" src="@/assets/images/pages/carousel/banner-18.jpg" alt="banner">
            </swiper-slide>
            <div class="swiper-pagination swiper-pagination-white" slot="pagination"></div>
            <div class="swiper-button-prev swiper-button-white" slot="button-prev"></div>
            <div class="swiper-button-next swiper-button-white" slot="button-next"></div>
        </swiper>

        <template slot="codeContainer">
&lt;template&gt;
  &lt;swiper :options=&quot;swiperOption&quot; :dir=&quot;$vs.rtl ? 'rtl' : 'ltr'&quot; :key=&quot;$vs.rtl&quot;&gt;
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
    &lt;div class=&quot;swiper-pagination swiper-pagination-white&quot; slot=&quot;pagination&quot;&gt;&lt;/div&gt;
    &lt;div class=&quot;swiper-button-prev swiper-button-white&quot; slot=&quot;button-prev&quot;&gt;&lt;/div&gt;
    &lt;div class=&quot;swiper-button-next swiper-button-white&quot; slot=&quot;button-next&quot;&gt;&lt;/div&gt;
  &lt;/swiper&gt;
&lt;/template&gt;

&lt;script&gt;
import 'swiper/dist/css/swiper.min.css'
import { swiper, swiperSlide } from 'vue-awesome-swiper'

export default {
  data() {
    return {
      swiperOption: {
        spaceBetween: 30,
        effect: 'fade',
        pagination: {
            el: '.swiper-pagination',
            clickable: true
        },
        navigation: {
            nextEl: '.swiper-button-next',
            prevEl: '.swiper-button-prev'
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
                spaceBetween: 30,
                effect: 'fade',
                pagination: {
                    el: '.swiper-pagination',
                    clickable: true
                },
                navigation: {
                    nextEl: '.swiper-button-next',
                    prevEl: '.swiper-button-prev'
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
