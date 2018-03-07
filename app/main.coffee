showLoader = undefined
hideLoader = undefined

do ->
	##### CONSTS #####

	API_PREFIX = "/api/"

	##### HELPERS #####

	byId = (id) -> document.getElementById id

	httpGet = (url, callback) ->
		xmlHttp = new XMLHttpRequest

		xmlHttp.onreadystatechange = ->
			if (xmlHttp.readyState is 4) and (xmlHttp.status is 200)
				callback xmlHttp.responseText
			return

		xmlHttp.open 'GET', url, true
		xmlHttp.send null

	##### ELEMENTS #####

	form        = byId('check')
	formContent = byId('form-content')
	loader      = byId('loader-container')
	status      = byId('status')

	loader.style.display = "none"

	##### HANDLERS #####

	showTimeout = undefined
	hideTimeout = undefined

	showLoader = ->
		loader.style.display = "flex"
		loader.className = "show"
		formContent.className = ""
		clearTimeout showTimeout
		clearTimeout hideTimeout
		showTimeout = setTimeout (-> formContent.style.display = "none"), 400
		return

	hideLoader = ->
		loader.className = ""
		clearTimeout showTimeout
		clearTimeout hideTimeout
		hideTimeout = setTimeout (-> loader.style.display = "none"), 400
		formContent.style.display = "flex"
		formContent.className = "show"
		return

	parser = (data) ->
		obj = JSON.parse data

		if obj.up is true
			status.className = "status-up"
			status.innerHTML = "UP"
		else if obj.up is false
			status.className = "status-down"
			status.innerHTML = "DOWN"
		else if obj.error?
			status.className = "status-error"
			status.innerHTML = "ERROR: " + obj.error.toUpperCase()

		return



	form.onsubmit = ->
		url = API_PREFIX + encodeURIComponent(form.url.value)
		showLoader()
		httpGet url, (data) ->
			parser data
			hideLoader()


