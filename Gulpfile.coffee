gulp    = require 'gulp'
gulpif  = require 'gulp-if'
stylus  = require 'gulp-stylus'
pug     = require 'gulp-pug'
coffee  = require 'gulp-coffee'
htmlmin = require 'gulp-htmlmin'
uglify  = require 'gulp-uglify-es'
csso    = require 'gulp-csso'
rcs     = require 'gulp-rcs'
del     = require 'del'

gulp.task 'clean', -> del "dist"

if gulp.parallel
	clean = gulp.parallel 'clean'
else
	clean = ['clean']

gulp.task 'build', clean, ->
	gulp.src ["app/**/*.styl", "app/**/*.coffee", "app/**/*.pug"]
		.pipe gulpif "*.styl",   stylus()
		.pipe gulpif "*.coffee", coffee(bare: true)
		.pipe gulpif "*.pug",    pug()

		.pipe rcs()

		.pipe gulpif "*.css",  csso()
		.pipe gulpif "*.js",   uglify.default()
		.pipe gulpif "*.html", htmlmin(collapseWhitespace: true)

		.pipe gulp.dest "dist"

if gulp.parallel
	gulp.task 'default', gulp.parallel('build')
else
	gulp.task 'default', ['build']
