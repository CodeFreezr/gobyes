import static groovy.io.FileType.*
import static groovy.io.FileVisitResult.*
 
//def goDir = new File('D:\\dbt\\01\\git\\gobyes\\corpus\\codegangsta\\bwag')
def rootDir = /D:\temp\gitemp\filipevarjao\webcheck/
def goDir = new File(rootDir + "\\404")
File filePlaylist = new File(rootDir + "\\playlist.txt") 
//def playlist 
def countGoFiles = 0
goDir.traverse { it -> 
    if (it.name.endsWith('.go')) {
        
        def cmdPlay = 'goplay -openbrowser=false -run=false ' + it.path
        def proc = cmdPlay.execute()
            def b = new StringBuffer()
                proc.consumeProcessErrorStream(b)
                    print b.toString()
        //println proc.text
        //playlist = playlist + it.path + " => " + proc.text
        Playlist << proc.text + " => " + it.path 
        println countGoFiles++ + ": " + it.path 
    }
    
    //println it.name.endsWith('.go')
    //countFilesAndDirs++
}

println "Total go files and directories in ${goDir.name}: $countGoFiles"
 

file << playlist
//println playlist


