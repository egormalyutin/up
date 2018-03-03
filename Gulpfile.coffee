gulp    = require 'gulp'
uglify  = require 'gulp-uglify'
csso    = require 'gulp-csso'
htmlmin = require 'gulp-htmlmin'
del     = require 'del'

gulp.task 'clean', -> del "dest"

gulp.task 'html', ->
	gulp.src "app/**/*.html"
		.pipe htmlmin collapseWhitespace: true
		.pipe gulp.dest "dist"

gulp.task 'css', ->
	gulp.src "app/**/*.css"
		.pipe csso()
		.pipe gulp.dest "dist"

gulp.task 'js', ->
	gulp.src "app/**/*.js"
		.pipe uglify()
		.pipe gulp.dest "dist"

if gulp.parallel
	gulp.task 'build', gulp.parallel('html', 'css', 'js')
	gulp.task 'default', gulp.parallel('build')
else
	gulp.task 'build', ['html', 'css', 'js']
	gulp.task 'default', ['build']
	