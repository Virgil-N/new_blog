var gulp = require("gulp");
var sass = require("gulp-sass-china");
var sourcemaps = require("gulp-sourcemaps");
var livereload = require('gulp-livereload');
var cssnano = require('gulp-cssnano');
var uglify = require('gulp-uglify');
var image = require('gulp-image');

gulp.task("css", function() {
	gulp.src("./static/sass/**/*.scss")
			.pipe(sourcemaps.init())
			.pipe(sass())
			.pipe(cssnano())
			.pipe(sourcemaps.write("./"))
			.pipe(gulp.dest("./public/css"))
			.pipe(livereload());
});

gulp.task("js", function() {
  gulp.src("./static/js/**/*.js")
  		.pipe(sourcemaps.init())
	    .pipe(uglify())
	    .pipe(sourcemaps.write("./"))
	    .pipe(gulp.dest("./public/js"))
	    .pipe(livereload())
});

gulp.task("img", function() {
  gulp.src("./static/img/**/*")
  		.pipe(image({
	      pngquant: true,
	      optipng: false,
	      zopflipng: true,
	      jpegRecompress: false,
	      jpegoptim: true,
	      mozjpeg: true,
	      gifsicle: true,
	      svgo: true,
	      concurrent: 10
	    }))
	    .pipe(gulp.dest("./public/img"))
	    .pipe(livereload())
});


gulp.task("watch", function() {
	livereload.listen();
	gulp.watch("./static/sass/**/*.scss", ["css"]);
	gulp.watch("./static/js/**/*.js", ["js"]);
	gulp.watch("./static/img/**/*", ["img"]);
});