// Copy all go files from one director-tree into another
// flatten everything into one directory

import static groovy.io.FileType.*
import static groovy.io.FileVisitResult.*
 
def rootDir = /D:\dbt\01\git\RosettaCodeData\Task/
def goDir = new File(rootDir)
def toDir = /D:\temp\gitemp\go-rosettacode/
def ant = new AntBuilder()
def countGoFiles = 0

goDir.traverse { it -> 
    if (it.name.endsWith('.go')) {
        println countGoFiles++ + " " + it
        ant.copy(file: it, todir: toDir) 
    }
}

