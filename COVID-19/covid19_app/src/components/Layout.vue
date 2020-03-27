<template>
    <div class="layout">
        <div id="mobile-menu-overlay"></div>
        <nav class="navbar navbar-expand-lg fixed-top py-md-2">
            <div class="container">
                <a class="navbar-brand" href="#">COVID-19 <i class="fas fa-viruses mx-2 fa-lg"></i></a>
                <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarTogglerDemo01" aria-controls="navbarTogglerDemo01" aria-expanded="false" aria-label="Toggle navigation">
                    <span class="navbar-toggler-icon"><i class="fas fa-bars"> </i></span>
                </button>
                <div class="collapse navbar-collapse" id="navbarTogglerDemo01">
                    <div class="d-lg-none d-flex justify-content-between px-4 py-3 align-items-center">
                        <a class="logo-mobile-men" href="#">COVID-19 <i class="fas fa-viruses mx-2 fa-lg"></i></a>
                        <a href="javascript:;" class="close-menu"><i class="mdi mdi-close"></i></a>
                    </div>
                    <ul class="navbar-nav ml-auto align-items-center">
                        <li class="nav-item active">
                            <a class="nav-link" href="#home"><i class="fas fa-hospital-user mx-2 fa-lg"></i>Home <span class="sr-only">(current)</span></a>
                        </li>
                        <li class="nav-item">
                            <a class="nav-link" href="#services"><i class="fas fa-hand-holding-medical mx-2 fa-lg"></i>Helpful links</a>
                        </li>
                        <li class="nav-item">
                            <a class="nav-link btn btn-success" href="https://api.whatsapp.com/send?phone=+91(9013)-151-515&text=Namaste"><i class="fab fa-whatsapp mx-2 fa-lg"></i>901-315-1515</a>
                        </li>
                    </ul>
                </div>
            </div>
        </nav>
        <!-- header section -->
        <div class="page-body-wrapper">
            <section id="home" class="home">
                <img src="/assets/images/banner-bg.svg" alt="">
                <div class="container">
                    <div class="row">
                        <div class="col-sm-12">
                            <div class="main-banner">
                                <div class="d-sm-flex justify-content-between">
                                    <div data-aos="zoom-in-up">
                                        <div class="banner-title">
                                            <h4 class="font-weight-medium">Coronavirus disease (COVID-19) is an infectious disease caused by a newly discovered coronavirus.
                                            </h4>
                                        </div>
                                        <a href="#" class="btn btn-secondary mt-3">Learn more</a>
                                    </div>
                                    <div class="mt-5 mt-lg-0">
                                        <img src="@/assets/images/medical1.svg" alt="medical" class="img-fluid" data-aos="zoom-in-up">
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </section>
            <!-- end of header section -->
            <!-- counter section -->
            <section>
                <div class="container">
                    <div class="row mb-5">
                        <div class="col-sm-12">
                            <div class="d-sm-flex justify-content-between align-items-center mb-2">
                                <h2 class="font-weight-medium text-dark ">INDIA COVID19 REPORTS</h2>
                            </div>
                        </div>
                    </div>
                    <div class="row">
                        <div class="col-sm-3">
                            <div class="d-flex py-3 my-3 my-lg-0 justify-content-center" data-aos="fade-down">
                                <span><i class="fas fa-hospital text-primary fa-3x mr-3"></i></span>
                                <div>
                                    <h4 class="font-weight-bold text-dark mb-0"><span class="scVal">0</span></h4>
                                    <h5 class="text-dark mb-0">Confirmed Cases</h5>
                                </div>
                            </div>
                        </div>
                        <div class="col-sm-3">
                            <div class="d-flex py-3 my-3 my-lg-0 justify-content-center" data-aos="fade-up">
                                <span><i class="fas fa-hospital-user text-danger fa-3x mr-3"></i></span>
                                <div>
                                    <h4 class="font-weight-bold text-dark mb-0"><span class="fpVal">0</span></h4>
                                    <h5 class="text-dark mb-0">Active Cases</h5>
                                </div>
                            </div>
                        </div>
                        <div class="col-sm-3">
                            <div class="d-flex py-3 my-3 my-lg-0 justify-content-center" data-aos="fade-down">
                                <span><i class="fas fa-shield-virus text-warning fa-3x mr-3"></i></span>
                                <div>
                                    <h4 class="font-weight-bold text-dark mb-0"><span class="tMVal">0</span></h4>
                                    <h5 class="text-dark mb-0">Recovered Cases</h5>
                                </div>
                            </div>
                        </div>
                        <div class="col-sm-3">
                            <div class="d-flex py-3 my-3 my-lg-0 justify-content-center" data-aos="fade-up">
                                <span><i class="fas fa-procedures fa-3x mr-3"></i></span>
                                <div>
                                    <h4 class="font-weight-bold text-dark mb-0"><span class="bPVal">0</span></h4>
                                    <h5 class="text-dark mb-0">Decreased Cases</h5>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </section>
            <!-- quotes and counter section -->
            <section class="our-projects" id="projects">
                <div class="container">
                    <div class="row mb-5">
                        <div class="col-sm-12">
                            <div class="d-sm-flex justify-content-between align-items-center mb-2">
                                <h2 class="font-weight-medium text-dark ">Please follow the rules</h2>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="mb-5 container">
                    <div class="row py-5 ">
                        <div class="col-md-7 col-12 mx-auto">
                            <img src="@/assets/images/carousel/slider1.svg" alt="take doctor advice" height="450" width="500" class="img-fluid">
                        </div>
                        <div class="col-md-5 col-12 mx-auto my-auto">
                            <div class="text-center">
                                <h1 class="text-dark mb-4">Doctor Advice</h1>
                                <p class="text-dark lead">Don't panic or try anything by yourself Take doctor advice first</p>
                            </div>
                        </div>
                    </div>
                    <div class="row flex-column-reverse flex-md-row py-5 my-5">
                        <div class="col-md-5 col-12 mx-auto my-auto">
                            <div class="text-center">
                                <h1 class="text-dark mb-4">Distance From Neighbours</h1>
                                <p class="text-dark lead">Please maintain atleast one metre distance from your neighbours while talking</p>
                            </div>
                        </div>
                        <div class="col-md-7 col-12 text-right">
                            <img src="@/assets/images/carousel/slider2.svg" alt="take doctor advice" height="450" width="500" class="img-fluid">
                        </div>
                    </div>
                    <div class="row pb-5 mb-5">
                        <div class="col-md-7 col-12 mx-auto">
                            <img src="@/assets/images/carousel/slider3.svg" alt="take doctor advice" height="400" width="400" class="img-fluid">
                        </div>
                        <div class="col-md-5 col-12 mx-auto my-auto">
                            <div class="text-center">
                                <h1 class="text-dark mb-4">Proper Handwash</h1>
                                <p class="text-dark lead">Please make sure you wash your hand again and again within a span of 15-20 min</p>
                            </div>
                        </div>
                    </div>
                    <div class="row flex-column-reverse flex-md-row pb-5 mb-5">
                        <div class="col-md-5 col-12 mx-auto my-auto ">
                            <div class="text-center">
                                <h1 class="text-dark mb-4">Avoid UnNecessary Outing</h1>
                                <p class="text-dark lead">Please try to avoid outing if very emergency then maintain distance and wear mask</p>
                            </div>
                        </div>
                        <div class="col-md-7 col-12 text-right">
                            <img src="@/assets/images/carousel/slider4.svg" alt="take doctor advice" height="450" width="500" class="img-fluid">
                        </div>
                    </div>
                    <div class="row pb-5 mb-5">
                        <div class="col-md-7 col-12 mx-auto">
                            <img src="@/assets/images/carousel/slider5.svg" alt="take doctor advice" height="400" width="400" class="img-fluid">
                        </div>
                        <div class="col-md-5 col-12 mx-auto my-auto">
                            <div class="text-center">
                                <h1 class="text-dark mb-4">Stay at Home</h1>
                                <p class="text-dark lead">Please stay at home and help in breaking the chain of spreading virus</p>
                            </div>
                        </div>
                    </div>
                    <div class="row flex-column-reverse flex-md-row pb-5 mb-5">
                        <div class="col-md-5 col-12 mx-auto my-auto">
                            <div class="text-center">
                                <h1 class="text-dark mb-4">Having any Symptoms?</h1>
                                <p class="text-dark lead">If you found any symptoms of virus in you or anybody near you go for checkup as early as possible</p>
                            </div>
                        </div>
                        <div class="col-md-7 col-12 text-right">
                            <img src="@/assets/images/carousel/slider7.svg" alt="take doctor advice" height="450" width="500" class="img-fluid">
                        </div>
                    </div>
                    <div class="row pb-5 mb-5">
                        <div class="col-md-7 col-12 mx-auto">
                            <img src="@/assets/images/carousel/slider11.svg" alt="take doctor advice" height="400" width="400" class="img-fluid">
                        </div>
                        <div class="col-md-5 col-12 mx-auto my-auto">
                            <div class="text-center">
                                <h1 class="text-dark mb-4">Fight together</h1>
                                <p class="text-dark lead">Stay at home and give contribution towards fighting againts corona virus</p>
                            </div>
                        </div>
                    </div>
                </div>
            </section>
        </div>
        <router-view />
    </div>
</template>
<script>
export default {
    name: 'Layout',
}
</script>
<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.navbar {
    width: 100%;
    padding: 1.875rem 1rem;
    background: #3c37f1;
}

@media (max-width: 991px) {
    .navbar {
        padding: .875rem 1rem;
        height: 71px;
    }
}

.navbar .navbar-brand {
    font-size: 1.75rem;
    color: #fff;
    font-weight: 500;
}

.navbar .navbar-brand img {
    width: 200px;
}

@media (max-width: 991px) {
    .navbar .navbar-brand img {
        width: 150px;
    }
}

@media (max-width: 991px) {
    .navbar .navbar-nav {
        padding: 1rem 1.5rem;
        height: 100%;
    }
}

.navbar .navbar-nav .nav-item {
    margin-right: .25rem;
    margin-right: .25rem;
}

@media (max-width: 991px) {
    .navbar .navbar-nav .nav-item {
        display: block;
        width: 100%;
        margin-right: 0;
        margin-bottom: .25rem;
    }
}

.navbar .navbar-nav .nav-item:last-child {
    margin-right: 0;
    margin-left: 1rem;
}

@media (max-width: 991px) {
    .navbar .navbar-nav .nav-item:last-child {
        margin-left: 0;
    }
}

.navbar .navbar-nav .nav-item .nav-link {
    color: #fff;
    font-size: 0.937rem;
    font-weight: 500;
    padding: 0.625rem 1.25rem;
    -webkit-transition: all 1s;
    -moz-transition: all 1s;
    -ms-transition: all 1s;
    -o-transition: all 1s;
    transition: all 1s;
    border-radius: 50px;
}

.navbar .navbar-nav .nav-item .nav-link.btn {
    padding: 0.5rem 1.25rem;
    font-size: 14px;
}

@media (max-width: 991px) {
    .navbar .navbar-nav .nav-item .nav-link {
        display: block;
        padding: .625rem;
    }
}

.navbar .navbar-nav .nav-item .nav-link.active {
    background: rgba(0, 0, 0, 0.2);
    border-radius: 50px;
}

.navbar .navbar-nav .nav-item .nav-link:hover,
.navbar .navbar-nav .nav-item .nav-link:focus {
    border-radius: 50px;
}

.navbar .navbar-nav .nav-item:last-child .nav-link.active {
    background: #00d663;
    border-radius: 50px;
}

.navbar .navbar-toggler {
    color: #ffffff;
}

.navbar .navbar-toggler-icon {
    height: auto;
}

@media (max-width: 991px) {
    .navbar .navbar-collapse {
        position: fixed;
        right: -280px;
        width: 280px;
        top: 0;
        height: 100%;
        background: #ffffff;
        z-index: 999999;
        overflow-y: scroll;
        -webkit-transition: all 0.25s ease-out;
        -o-transition: all 0.25s ease-out;
        transition: all 0.25s ease-out;
    }

    .navbar .navbar-collapse.show {
        right: 0;
    }

    .navbar .navbar-collapse .logo-mobile-menu {
        width: 120px;
    }

    .navbar .navbar-collapse .navbar-nav .nav-item .nav-link {
        color: #000000;
    }
}

@media (max-width: 991px) and (max-width: 991px) {
    .navbar .navbar-collapse .close-menu {
        color: #000000;
        padding: 2px 4px;
        background: transparent;
        display: inline-block;
        cursor: pointer;
        font-size: 20px;
        width: 30px;
        height: 30px;
        text-align: center;
    }
}

.sidebar-overlay .home {
    padding-top: 0px;
    background-position: 0px -121px;
}

.sidebar-overlay #mobile-menu-overlay {
    position: fixed;
    width: 100%;
    height: 100%;
    background: rgba(0, 0, 0, 0.8);
    z-index: 99;
    -webkit-transition: all .9s;
    -moz-transition: all .8s;
    -ms-transition: all .8s;
    -o-transition: all .8s;
    transition: all .8s;
    transition-delay: .35s;
    left: 0;
}

.sidebar-overlay .navbar {
    padding: .875rem .1rem;
    position: static;
    z-index: auto;
    background: #3c37f1;
}

@media (max-width: 991px) {
    .sidebar-overlay .navbar {
        padding: .875rem 1rem;
    }
}

.sidebar-overlay .navbar .navbar-collapse {
    z-index: 100;
}

.sidebar-overlay.fixed-on-scroll {
    z-index: auto;
    position: absolute;
    top: 0;
}

body,
html {
    overflow-x: hidden;
}

.home {
    background: url("/assets/images/banner-bg.svg");
    background-repeat: no-repeat;
    background-position: 0 -50px;
    padding-top: 103px;
}

@media (max-width: 991px) {
    .home {
        padding-top: 71px;
    }
}

.main-banner {
    padding: 4.2rem 0 14rem 0;
    color: #ffffff;
    background-size: cover;
}

@media (max-width: 991px) {
    .main-banner {
        padding: 1rem;
        margin-bottom: 50px;
    }
}

.main-banner .banner-title {
    width: 56%;
}

@media (max-width: 991px) {
    .main-banner .banner-title {
        width: 100%;
    }
}

section {
    padding: 120px 0 0 0;
}

@media (max-width: 1024px) {
    section {
        padding: 130px 0 0 0;
    }
}

@media (max-width: 991px) {
    section {
        padding: 70px 0 0 0;
    }
}

.services-box {
    padding: 10px 100px 40px 0;
}

@media (max-width: 991px) {
    .services-box {
        padding: 10px 10px 40px 0;
    }
}

.services-box:hover img {
    -webkit-filter: grayscale(100%);
    /* Safari 6.0 - 9.0 */
    filter: grayscale(100%);
    -webkit-transition: all 2s;
    -moz-transition: all 2s;
    -ms-transition: all 2s;
    -o-transition: all 2s;
    transition: all 2s;
}

.tick-icon {
    width: 20px;
    height: 20px;
}

.testimonial {
    background: #a70ff6 url("/assets/images/testimonial-bg.svg");
    padding: 100px 0 70px 0;
    background-size: cover;
}

@media (max-width: 991px) {
    .testimonial {
        padding: 20px 0;
    }
}

.testimonial .testimonial-item {
    width: 450px;
    padding: 34px 34px 34px 94px;
    background: #ffffff;
    color: #000000;
    border-radius: 10px;
    position: relative;
    box-shadow: 0 20px 70px -10px rgba(109, 121, 141, 0.15);
}

@media (max-width: 1024px) {
    .testimonial .testimonial-item {
        width: 400px;
    }
}

@media (max-width: 991px) {
    .testimonial .testimonial-item {
        width: 280px;
    }
}

.testimonial .testimonial-item .testimonial-icons {
    width: 50px;
    height: 50px;
    border-radius: 100%;
    position: absolute;
    top: 40px;
    left: 28px;
}

.testimonial .testimonial-item .testimonial-author {
    margin-top: 1rem;
    margin-bottom: .5rem;
    position: relative;
    font-weight: 500;
}

.testimonial .testimonial-item .testimonial-author:before {
    content: "\f756";
    position: absolute;
    left: -40px;
    color: transparent;
    font-size: 1.75rem;
    background: url(/assets/images/quote.png) no-repeat;
}

.testimonial .testimonial-item .testimonial-destination {
    color: #506690;
    font-weight: normal;
    font-size: 14px;
}

.testimonial .flipster__button--prev {
    left: auto;
    right: 0;
    top: 86%;
    transform: rotate(270deg);
}

.testimonial .flipster__button--next {
    transform: rotate(270deg);
    top: 78%;
}

.testimonial .flipster__button svg {
    width: .75em;
}

/*-------------------------------------------------------------------*/
/* === Components === */
.btn {
    font-weight: bold;
}

/* Forms */
.form-group {
    margin-bottom: 1.5rem;
}

.form-control {
    border: 1px solid #f1efef;
    font-weight: 400;
    font-size: 14px;
    height: auto;
}

/* Toggle Switch */
.toggle-switch {
    position: relative;
    display: inline-block;
    width: 45px;
    height: 23px;
}

.toggle-switch .toggle-slider {
    position: absolute;
    cursor: pointer;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background-color: #ccc;
    -webkit-transition: .4s;
    transition: .4s;
    border-radius: 34px;
    border: 1px solid #f1efef;
}

.toggle-switch .toggle-slider:before {
    position: absolute;
    content: "";
    height: 18px;
    width: 18px;
    left: 2px;
    bottom: 2px;
    background-color: #ffffff;
    -webkit-transition: .4s;
    transition: .4s;
    border-radius: 50%;
}

.toggle-switch input {
    visibility: hidden;
}

.toggle-switch input:checked+.toggle-slider {
    background-color: #3c37f1;
}

.toggle-switch input:checked+.toggle-slider:before {
    -webkit-transform: translateX(21px);
    -ms-transform: translateX(21px);
    transform: translateX(21px);
}

.text-active {
    color: #000000;
}


.flipster--coverflow .flipster__container {
    padding-bottom: 0;
}

.flipster--coverflow .flipster__item__content {
    -webkit-transform-origin: 50% 100%;
    -ms-transform-origin: 50% 100%;
    transform-origin: 50% 100%;
    -webkit-box-reflect: below 0 -webkit-gradient(linear, left bottom, left top, color-stop(0.05, rgba(255, 255, 255, 0)), color-stop(0.2, transparent));
}

.flipster-custom-nav {
    padding-left: 0;
    list-style: none;
    margin-top: 2rem;
}

.flipster-custom-nav .flipster-custom-nav-item {
    display: inline-block;
}

.flipster-custom-nav .flipster-custom-nav-item .flipster-custom-nav-link {
    display: block;
    width: 10px;
    height: 10px;
    border-radius: 100%;
    border: 1px solid #f1efef;
    margin-right: 1rem;
}

.flipster-custom-nav .flipster-custom-nav-item .flipster-custom-nav-link.active {
    background: #f1efef;
}
</style>