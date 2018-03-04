do ->
	##### CONSTS #####

	API_PREFIX = "/api/"

	##### HELPERS #####

	byId = (id) -> document.getElementById id

	httpGet = (url, callback) ->
		xmlHttp = new XMLHttpRequest

		xmlHttp.onreadystatechange = ->
			if xmlHttp.readyState == 4 and xmlHttp.status == 200
				callback xmlHttp.responseText
			return

		xmlHttp.open 'GET', url, true
		xmlHttp.send null

	##### ELEMENTS #####

	form  = byId('check')
	input = byId('address')

	##### HANDLERS #####

	form.onsubmit = ->
		alert input
		addr = API_PREFIX + input.value
		httpGet addr, (data) ->
			alert data
