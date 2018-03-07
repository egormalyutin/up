gulp    = require 'gulp'
lazy    = require 'lazypipe'
watch   = require 'gulp-watch'
plumber = require 'gulp-plumber'
gulpif  = require 'gulp-if'
filter  = require 'gulp-filter'
debug   = require 'gulp-debug'
stylus  = require 'gulp-stylus'
pug     = require 'gulp-pug'
coffee  = require 'gulp-coffee'
htmlmin = require 'gulp-htmlmin'
uglify  = require 'gulp-uglify-es'
csso    = require 'gulp-csso'
fontmin = require 'gulp-fontmin'
del     = require 'del'

SOURCES = ["app/**/*.styl", "app/**/*.coffee", "app/**/*.pug", "app/**/*.ttf"]

gulp.task 'clean', -> del "dist"

if gulp.parallel
	clean = gulp.parallel 'clean'
else
	clean = ['clean']

builders = (x, p) ->
	if p
		x = x.pipe plumber()

	fonts = lazy()
		.pipe fontmin, text: "IS HOST UP OR DOWN?"
		.pipe filter, ["**/*.ttf"]

	x.pipe gulpif "*.styl",   stylus()
		.pipe gulpif "*.coffee", coffee(bare: true)
		.pipe gulpif "*.pug",    pug()

		.pipe gulpif "*.css",  csso()
		.pipe gulpif "*.js",   uglify.default()
		.pipe gulpif "*.html", htmlmin(collapseWhitespace: true)

		.pipe gulpif "*.ttf",  fonts()

		.pipe debug title: "Compiled"

gulp.task 'build', clean, ->
	builders(gulp.src(SOURCES))
		.pipe gulp.dest "dist"

if gulp.parallel
	build = gulp.parallel 'build'
else
	build = ['build']


gulp.task 'watch', build, ->
	builders(watch(SOURCES), true)
		.pipe gulp.dest "dist"


gulp.task 'default', build
