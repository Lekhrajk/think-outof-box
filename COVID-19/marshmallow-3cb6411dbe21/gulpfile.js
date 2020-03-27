'use strict'

var gulp = require('gulp');
var browserSync = require('browser-sync').create();
var sass = require('gulp-sass');
var del = require('del');
var sourcemaps = require('gulp-sourcemaps');
var concat = require('gulp-concat');
var merge = require('merge-stream');



gulp.task('sass', function () {
    return gulp.src('./scss/**/style.scss')
        .pipe(sourcemaps.init())
        .pipe(sass({outputStyle: 'expanded'}).on('error', sass.logError))
        .pipe(sourcemaps.write('./maps'))
        .pipe(gulp.dest('./css'))
        .pipe(browserSync.stream());
});

// Static Server + watching scss/html files
gulp.task('serve', gulp.series('sass', function() {
    browserSync.init({
        port: 3001,
        server: "./",
        ghostMode: false,
        notify: false
    });

    gulp.watch('scss/**/*.scss', gulp.series('sass'));
    gulp.watch('**/*.html').on('change', browserSync.reload);
    gulp.watch('js/**/*.js').on('change', browserSync.reload);

}));

// Static Server without watching scss files
gulp.task('serve:lite', function() {

    browserSync.init({
        server: "./",
        ghostMode: false,
        notify: false
    });

    gulp.watch('**/*.css').on('change', browserSync.reload);
    gulp.watch('**/*.html').on('change', browserSync.reload);
    gulp.watch('js/**/*.js').on('change', browserSync.reload);

});

gulp.task('clean:vendors', function () {
    return del([
      'vendors/**/*'
    ]);
});

/*Building vendor scripts needed for basic template rendering*/
gulp.task('buildBaseVendorScripts', function() {
    return gulp.src([
        './node_modules/jquery/dist/jquery.min.js', 
        './node_modules/popper.js/dist/umd/popper.min.js', 
        './node_modules/bootstrap/dist/js/bootstrap.min.js',
    ])
      .pipe(concat('vendor.bundle.base.js'))
      .pipe(gulp.dest('./vendors/base'));
});


gulp.task('copyAddonsScripts', function() {
    var aScript1 = gulp.src(['./node_modules/owl.carousel/dist/owl.carousel.js'])
        .pipe(gulp.dest('./vendors/owl.carousel/js'));
    var aScript2 = gulp.src(['./node_modules/aos/dist/aos.js'])
        .pipe(gulp.dest('./vendors/aos/js'));
    var aScript3 = gulp.src(['./node_modules/jquery.flipster/dist/jquery.flipster.min.js'])
        .pipe(gulp.dest('./vendors/jquery-flipster/js'));
    return merge(aScript1, aScript2, aScript3);
});

gulp.task('copyAddonsStyles', function() {
    var aStyle1 = gulp.src(['./node_modules/owl.carousel/dist/assets/owl.carousel.css'])
        .pipe(gulp.dest('./vendors/owl.carousel/css'));
    var aStyle2 = gulp.src(['./node_modules/owl.carousel/dist/assets/owl.theme.default.min.css'])
        .pipe(gulp.dest('./vendors/owl.carousel/css'));
    var aStyle3 = gulp.src(['./node_modules/aos/dist/aos.css'])
        .pipe(gulp.dest('./vendors/aos/css'));
    var aStyle4 = gulp.src(['./node_modules/jquery.flipster/dist/jquery.flipster.css'])
        .pipe(gulp.dest('./vendors/jquery-flipster/css'));
    var aStyle5 = gulp.src(['./node_modules/@mdi/font/css/materialdesignicons.min.css'])
        .pipe(gulp.dest('./vendors/mdi/css'));
    var aStyle6 = gulp.src(['./node_modules/@mdi/font/fonts/*'])
        .pipe(gulp.dest('./vendors/mdi/fonts'));
    return merge(aStyle1, aStyle2, aStyle3, aStyle4, aStyle5, aStyle6);
});

/*sequence for building vendor scripts and styles*/
gulp.task('bundleVendors', gulp.series('clean:vendors','buildBaseVendorScripts','copyAddonsScripts','copyAddonsStyles'));

gulp.task('default', gulp.series('serve'));
