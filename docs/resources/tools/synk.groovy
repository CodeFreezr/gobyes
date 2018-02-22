// Stripped all Go-Files
System.setProperty("http.agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X x.y; rv:10.0) Gecko/20100101 Firefox/10.0")

ant = new AntBuilder()
url200 = "http://www.rosettacode.org/wiki/"
url404 = "http://www.rosettacode.org/wiki/Special:Search/"
f200 = new File('result200.log')
f404 = new File('result404.log')
def sDir = "../RosettaGoCode"


file = new File('all-folder.log')
file.eachLine { String line ->
  def tDir = ""
  def urlstring = url200 + line
  def pString = line.replaceAll('_','-').toLowerCase()
  def url = urlstring.toURL()  
  def connection = url.openConnection()
  def code = connection.responseCode

  println code + " <-> " + urlstring
  if (code == 200) {
    f200 << '\n' << urlstring
    tDir = "200\\"+pString
  } else {
    f404 << 'n' << url404 + line
    tDir = "404\\"+pString
  }

  ant.sequential {
      def sFiles = pString + "*.go"
      mkdir(dir: tDir)
      copy(todir: tDir) {
        fileset(dir: sDir) {
          include(name: sFiles)
        }
      }
  }
}


