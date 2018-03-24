@Grab('com.xlson.groovycsv:groovycsv:1.1')
import static com.xlson.groovycsv.CsvParser.parseCsv

counter=0
for(line in parseCsv(new FileReader('goplay200.csv'), separator: ',')) {
    sFile = new File("dirtree200.html")
    def txt = sFile.text
    txt = txt.replaceAll(line.searchString, line.replaceString)
    sFile.write(txt)
    println counter++ + ". " + line.searchString + " replaced."
}

println "done."
