import static groovy.io.FileType.*
import static groovy.io.FileVisitResult.*
 
def goDir = new File('D:\\dbt\\01\\git\\gobyes\\corpus\\codegangsta\\bwag')

def playlist 
def countGoFiles = 0
goDir.traverse { it -> 
    if (it.name.endsWith('.go')) {
        println it.path
        def cmdPlay = 'goplay -openbrowser=false -run=false ' + it.path
        def proc = cmdPlay.execute()
            def b = new StringBuffer()
                proc.consumeProcessErrorStream(b)
                    print b.toString()
        //println proc.text
        playlist = playlist + it.path + " => " + proc.text
        //println it.name
        countGoFiles++
    }
    
    //println it.name.endsWith('.go')
    //countFilesAndDirs++
}

println "Total go files and directories in ${goDir.name}: $countGoFiles"
 
println playlist

/* 
def totalFileSize = 0
def groovyFileCount = 0
def sumFileSize = {
    totalFileSize += it.size()
    groovyFileCount++
}
def filterGroovyFiles = ~/.*\.groovy$/
groovySrcDir.traverse type: FILES, visit: sumFileSize, nameFilter: filterGroovyFiles
println "Total file size for $groovyFileCount Groovy source files is: $totalFileSize"
 
def countSmallFiles = 0
def postDirVisitor = {
    if (countSmallFiles > 0) {
     println "Found $countSmallFiles files with small filenames in ${it.name}"
 }
    countSmallFiles = 0
}
groovySrcDir.traverse(type: FILES, postDir: postDirVisitor, nameFilter: ~/.*\.groovy$/) {
    if (it.name.size() < 15) {
     countSmallFiles++
    }
}
*/