// Downloads a repo, decoupled it, sloc it and copy to a destination

// 2do: User os.FileSeperator ...
// 2do: run headcom / playgrounder etc. pp.
// 2do: change "gobeyd" in readme.md
// 2do: Solve Problme with ".files and folders"
// try git checkout --orphant
// try utree
// try goplay
// modify dirtree.html
// 

def tempDir = 'd:\\temp\\gitemp'
def corpusDir = 'd:\\dbt\\01\\git\\gobyes\\corpus'
def provider = 'https://github.com/'
def cloneCmd = 'git clone '
def user = this.args[0]
def repo = this.args[1]

println "Sync: " + provider + user + "/" + repo

//clone
cloneCmd = cloneCmd + provider + user + "/" + repo
def proc = cloneCmd.execute()
proc.waitFor()
println "Clone: done."

//sloc
def slocCmd = 'cmd /c sloc ' + repo + ' > ' + repo + '\\sloc.log'
def procSloc = slocCmd.execute()
procSloc.waitFor()
println "Sloc: done."

ant = new AntBuilder()
ant.sequential {
  delete(file:repo + '\\.gitignore')  
  delete(dir:repo + '\\.git')  
  move (todir: tempDir + "\\" + user + "\\" + repo) {
    fileset (dir: repo)
  }
}

println "Sync: done."
