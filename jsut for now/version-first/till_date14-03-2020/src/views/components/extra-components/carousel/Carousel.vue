<!-- =========================================================================================
    File Name: Carousel.vue
    Description: Carousel demo
    ----------------------------------------------------------------------------------------
    Item Name: Vuexy - Vuejs, HTML & Laravel Admin Dashboard Template
      Author: Pixinvent
    Author URL: http://www.themeforest.net/user/pixinvent
========================================================================================== -->

<template>
    <div id="extra-component-vue-awesome-swiper-demo">
        <carousel-default                  />
        <carousel-navigation               />
        <carousel-pagination               />
        <carousel-progress                 />
        <carousel-multiple-slides-per-view />
        <carousel-multi-row-slides-layout  />
        <carousel-effect-fade              />
        <carousel-3d-effect                />
        <carousel-3d-coverflow-effect      />
        <carousel-autoplay                 />
        <carousel-gallery                  />
        <carousel-parallax                 />
        <carousel-lazy-loading             />
        <carousel-responsive-breakpoints   />
        <carousel-virtual-slides           />
    </div>
</template>

<script>
import CarouselDefault               from './CarouselDefault.vue'
import CarouselNavigation            from './CarouselNavigation.vue'
import CarouselPagination            from './CarouselPagination.vue'
import CarouselProgress              from './CarouselProgress.vue'
import CarouselMultipleSlidesPerView from './CarouselMultipleSlidesPerView.vue'
import CarouselMultiRowSlidesLayout  from './CarouselMultiRowSlidesLayout.vue'
import CarouselEffectFade            from './CarouselEffectFade.vue'
import Carousel3dEffect              from './Carousel3dEffect.vue'
import Carousel3dCoverflowEffect     from './Carousel3dCoverflowEffect.vue'
import CarouselAutoplay              from './CarouselAutoplay.vue'
import CarouselGallery               from './CarouselGallery.vue'
import CarouselParallax              from './CarouselParallax.vue'
import CarouselLazyLoading           from './CarouselLazyLoading.vue'
import CarouselResponsiveBreakpoints from './CarouselResponsiveBreakpoints.vue'
import CarouselVirtualSlides         from './CarouselVirtualSlides.vue'

export default {
    components: {
        CarouselDefault,
        CarouselNavigation,
        CarouselPagination,
        CarouselProgress,
        CarouselMultipleSlidesPerView,
        CarouselMultiRowSlidesLayout,
        CarouselEffectFade,
        Carousel3dEffect,
        Carousel3dCoverflowEffect,
        CarouselAutoplay,
        CarouselGallery,
        CarouselParallax,
        CarouselLazyLoading,
        CarouselResponsiveBreakpoints,
        CarouselVirtualSlides,
    }
}
</script>
