System.setProperty("http.agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X x.y; rv:10.0) Gecko/20100101 Firefox/10.0")

def NEW_LINE = System.getProperty("line.separator")
if (args.length < 1)
{
   println "Enter a URL as an argument."
   System.exit(-1)
}
def address = args[0]
def urlInfo = address.toURL()
println "URL: ${address}${NEW_LINE}"
println "Host/Port: ${urlInfo.host}/${urlInfo.port}${NEW_LINE}"
println "Protocol: ${urlInfo.protocol}${NEW_LINE}"

def connection = urlInfo.openConnection()
println "Connection Type: ${connection.class}"
println "Content Type: ${connection.contentType}"
println "Response Code/Message: ${connection.responseCode} / ${connection.responseMessage}"
println "Request Method: ${connection.requestMethod}"
println "Date: ${connection.date}"
println "Last-Modified: ${connection.lastModified}"
println "Content Length: ${connection.contentLength}"


new File("download.html") << urlInfo.getText()