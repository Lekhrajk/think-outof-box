<!-- =========================================================================================
    File Name: CarouselGallery.vue
    Description: Gallery Carousel demo
    ----------------------------------------------------------------------------------------
    Item Name: Vuexy - Vuejs, HTML & Laravel Admin Dashboard Template
      Author: Pixinvent
    Author URL: http://www.themeforest.net/user/pixinvent
========================================================================================== -->


<template>
    <vx-card title="Gallery" class="carousel-example-gallery" code-toggler>
        <div>
            <!-- swiper1 -->
            <swiper :options="swiperOptionTop" class="gallery-top" ref="swiperTop" :dir="$vs.rtl ? 'rtl' : 'ltr'" :key="$vs.rtl + ' gallery'">
                <swiper-slide>
                  <img class="responsive" src="@/assets/images/pages/carousel/banner-10.jpg" alt="banner">
                </swiper-slide>
                <swiper-slide>
                  <img class="responsive" src="@/assets/images/pages/carousel/banner-16.jpg" alt="banner">
                </swiper-slide>
                <swiper-slide>
                  <img class="responsive" src="@/assets/images/pages/carousel/banner-15.jpg" alt="banner">
                </swiper-slide>
                <swiper-slide>
                  <img class="responsive" src="@/assets/images/pages/carousel/banner-13.jpg" alt="banner">
                </swiper-slide>
                <div class="swiper-button-next swiper-button-white" slot="button-next"></div>
                <div class="swiper-button-prev swiper-button-white" slot="button-prev"></div>
            </swiper>

            <!-- swiper2 Thumbs -->
            <swiper :options="swiperOptionThumbs" class="gallery-thumbs" ref="swiperThumbs" :dir="$vs.rtl ? 'rtl gallery' : 'ltr gallery'" :key="$vs.rtl + ' thumbs'">
                <swiper-slide>
                  <img class="responsive" src="@/assets/images/pages/carousel/banner-10.jpg" alt="banner">
                </swiper-slide>
                <swiper-slide>
                  <img class="responsive" src="@/assets/images/pages/carousel/banner-16.jpg" alt="banner">
                </swiper-slide>
                <swiper-slide>
                  <img class="responsive" src="@/assets/images/pages/carousel/banner-15.jpg" alt="banner">
                </swiper-slide>
                <swiper-slide>
                  <img class="responsive" src="@/assets/images/pages/carousel/banner-13.jpg" alt="banner">
                </swiper-slide>
            </swiper>
        </div>

          <template slot="codeContainer">
&lt;template&gt;
  &lt;div style=&quot;height: 500px&quot;&gt;

    &lt;!-- swiper1 --&gt;
        &lt;swiper :options=&quot;swiperOptionTop&quot; class=&quot;gallery-top&quot; ref=&quot;swiperTop&quot; :dir=&quot;$vs.rtl ? 'rtl' : 'ltr'&quot; :key=&quot;$vs.rtl&quot;&gt;
            &lt;swiper-slide&gt;
              &lt;img class=&quot;responsive&quot; src=&quot;@/assets/images/pages/carousel/banner-10.jpg&quot; alt=&quot;banner&quot;&gt;
            &lt;/swiper-slide&gt;
            &lt;swiper-slide&gt;
              &lt;img class=&quot;responsive&quot; src=&quot;@/assets/images/pages/carousel/banner-16.jpg&quot; alt=&quot;banner&quot;&gt;
            &lt;/swiper-slide&gt;
            &lt;swiper-slide&gt;
              &lt;img class=&quot;responsive&quot; src=&quot;@/assets/images/pages/carousel/banner-15.jpg&quot; alt=&quot;banner&quot;&gt;
            &lt;/swiper-slide&gt;
            &lt;swiper-slide&gt;
              &lt;img class=&quot;responsive&quot; src=&quot;@/assets/images/pages/carousel/banner-13.jpg&quot; alt=&quot;banner&quot;&gt;
            &lt;/swiper-slide&gt;
            &lt;div class=&quot;swiper-button-next swiper-button-white&quot; slot=&quot;button-next&quot;&gt;&lt;/div&gt;
            &lt;div class=&quot;swiper-button-prev swiper-button-white&quot; slot=&quot;button-prev&quot;&gt;&lt;/div&gt;
        &lt;/swiper&gt;

        &lt;!-- swiper2 Thumbs --&gt;
        &lt;swiper :options=&quot;swiperOptionThumbs&quot; class=&quot;gallery-thumbs&quot; ref=&quot;swiperThumbs&quot; :dir=&quot;$vs.rtl ? 'rtl' : 'ltr'&quot; :key=&quot;$vs.rtl&quot;&gt;
            &lt;swiper-slide&gt;
              &lt;img class=&quot;responsive&quot; src=&quot;@/assets/images/pages/carousel/banner-10.jpg&quot; alt=&quot;banner&quot;&gt;
            &lt;/swiper-slide&gt;
            &lt;swiper-slide&gt;
              &lt;img class=&quot;responsive&quot; src=&quot;@/assets/images/pages/carousel/banner-16.jpg&quot; alt=&quot;banner&quot;&gt;
            &lt;/swiper-slide&gt;
            &lt;swiper-slide&gt;
              &lt;img class=&quot;responsive&quot; src=&quot;@/assets/images/pages/carousel/banner-15.jpg&quot; alt=&quot;banner&quot;&gt;
            &lt;/swiper-slide&gt;
            &lt;swiper-slide&gt;
              &lt;img class=&quot;responsive&quot; src=&quot;@/assets/images/pages/carousel/banner-13.jpg&quot; alt=&quot;banner&quot;&gt;
            &lt;/swiper-slide&gt;
        &lt;/swiper&gt;
  &lt;/div&gt;
&lt;/template&gt;

&lt;script&gt;
import 'swiper/dist/css/swiper.min.css'
import { swiper, swiperSlide } from 'vue-awesome-swiper'

export default {
  data() {
    return {
      swiperOptionTop: {
        spaceBetween: 10,
        navigation: {
          nextEl: '.swiper-button-next',
          prevEl: '.swiper-button-prev'
        }
      },
      swiperOptionThumbs: {
        spaceBetween: 10,
        centeredSlides: true,
        slidesPerView: 'auto',
        touchRatio: 0.2,
        slideToClickedSlide: true
      }
    }
  },
  mounted() {
    this.$nextTick(() =&gt; {
      const swiperTop = this.$refs.swiperTop.swiper
      const swiperThumbs = this.$refs.swiperThumbs.swiper
      swiperTop.controller.control = swiperThumbs
      swiperThumbs.controller.control = swiperTop
    })
  },
  components: {
    swiper,
    swiperSlide
  }
}
&lt;/script&gt;


&lt;style lang=&quot;scss&quot; scoped&gt;
  .swiper-container {
    background-color: #000;
  }
  .gallery-top {
    height: 80%!important;
    width: 100%;
  }
  .gallery-thumbs {
    height: 20%!important;
    box-sizing: border-box;
    padding: 10px 0;
  }
  .gallery-thumbs .swiper-slide {
    width: 25%;
    height: 100%;
    opacity: 0.4;
  }
  .gallery-thumbs .swiper-slide-active {
    opacity: 1;
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
            swiperOptionTop: {
                spaceBetween: 10,
                navigation: {
                    nextEl: '.swiper-button-next',
                    prevEl: '.swiper-button-prev'
                }
            },
            swiperOptionThumbs: {
                spaceBetween: 10,
                centeredSlides: true,
                slidesPerView: 'auto',
                touchRatio: 0.2,
                slideToClickedSlide: true
            }
        }
    },
    mounted() {
        this.$nextTick(() => {
            const swiperTop = this.$refs.swiperTop.swiper
            const swiperThumbs = this.$refs.swiperThumbs.swiper
            swiperTop.controller.control = swiperThumbs
            swiperThumbs.controller.control = swiperTop
        })
    },
    components: {
        swiper,
        swiperSlide
    }
}
</script>


<style lang="scss" scoped>
  .swiper-container {
    background-color: #000;
  }
  .gallery-top {
    height: 80%!important;
    width: 100%;
  }
  .gallery-thumbs {
    height: 20%!important;
    box-sizing: border-box;
    padding: 10px 0;
  }
  .gallery-thumbs .swiper-slide {
    width: 25%;
    height: 100%;
    opacity: 0.4;
  }
  .gallery-thumbs .swiper-slide-active {
    opacity: 1;
  }
</style>
