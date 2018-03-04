(function() {
  var API_PREFIX, byId, form, httpGet, input;
  //#### CONSTS #####
  API_PREFIX = "/api/";
  //#### HELPERS #####
  byId = function(id) {
    return document.getElementById(id);
  };
  httpGet = function(url, callback) {
    var xmlHttp;
    xmlHttp = new XMLHttpRequest;
    xmlHttp.onreadystatechange = function() {
      if (xmlHttp.readyState === 4 && xmlHttp.status === 200) {
        callback(xmlHttp.responseText);
      }
    };
    xmlHttp.open('GET', url, true);
    return xmlHttp.send(null);
  };
  //#### ELEMENTS #####
  form = byId('check');
  input = byId('n');
  //#### HANDLERS #####
  return form.onsubmit = function() {
    var addr;
    alert(input);
    addr = API_PREFIX + input.value;
    return httpGet(addr, function(data) {
      return alert(data);
    });
  };
})();
