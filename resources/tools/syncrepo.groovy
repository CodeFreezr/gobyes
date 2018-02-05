
// set user arg[0] and repo arg[1]
// 2do: User os.FileSeperator ...
// 2do: run headcom / playgrounder etc. pp.
// 2do: change "gobeyd"

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
  move (todir: corpusDir + "\\" + user + "\\" + repo) {
    fileset (dir: repo)
  }
}

println "Sync: done."
