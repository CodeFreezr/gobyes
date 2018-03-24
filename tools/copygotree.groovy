// Copy all go files from one director-tree into another
// Path will be shifted, rootDir will be replaced with toDir
// replace the go-path in between
// 667 Folder and 997 Go-Files

import static groovy.io.FileType.*
import static groovy.io.FileVisitResult.*
 
def rootDir = /D:\dbt\01\git\RosettaCodeData\Task/
def goDir = new File(rootDir)
def toDir = /D:\temp\gitemp\go-rosettacode/
def ant = new AntBuilder()
def countGoFiles = 0
def target = ""

goDir.traverse { it -> 
    if (it.name.endsWith('.go')) {
        target = it.path.replace(rootDir,toDir)
        target = target.replace("\\Go\\","\\")
        target = target.replace("\\$it.name","")
        println countGoFiles++ + ": " + it.name
        ant.copy(file: it, todir: target) 
    }
}


