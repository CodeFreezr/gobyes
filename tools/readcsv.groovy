@Grab('com.xlson.groovycsv:groovycsv:1.1')
import static com.xlson.groovycsv.CsvParser.parseCsv

for(line in parseCsv(new FileReader('goplay200.csv'), separator: ',')) {
    println "searchString=$line.searchString, replaceString=$line.replaceString"
}