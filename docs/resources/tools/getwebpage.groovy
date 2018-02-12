System.setProperty("http.agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X x.y; rv:10.0) Gecko/20100101 Firefox/10.0")
//def url = "http://www.rosettacode.org/wiki/SHAb".toURL()


result = ""
file = new File('all-folder.log')
file.eachLine { String line ->
  def urlstring = "http://www.rosettacode.org/wiki/" + line
  def url = urlstring.toURL()  
  def connection = url.openConnection()

  println connection.responseCode + " <-> " + urlstring
  result = result + connection.responseCode + " <-> " + urlstring
  
}


