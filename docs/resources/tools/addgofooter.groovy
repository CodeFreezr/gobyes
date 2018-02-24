// Add one first line with the folder and filename

import static groovy.io.FileType.*
import static groovy.io.FileVisitResult.*
 
def rootDir = /D:\temp\gitemp\go-rosettacode/
def goDir = new File(rootDir)
def ant = new AntBuilder()
def countGoFiles = 0


goDir.traverse { it -> 
    if (it.name.endsWith('.go')) {
        println countGoFiles++ + " " + it
        f = it.path.replace(rootDir,"")
        ant.concat(destfile: it, append: true, "\n//$f\n")
    }
}




