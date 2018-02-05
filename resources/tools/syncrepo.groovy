
// set user arg[0] and repo arg[1]
// 2do: User os.FileSeperator ...
// 2do: run headcom / playgrounder etc. pp.

// Move location/user/repo to corpus/user/repo

// checkin stuff

def tempDir = 'd:\\temp\\gitemp'
def corpusDir = 'd:\\dbt\\01\\git\\gobeys\\corpus'
def provider = 'https://github.com/'
def cloneCmd = 'git clone '
def user = this.args[0]
def repo = this.args[1]

println "Sync: " + provider + user + "/" + repo


//clone
cloneCmd = cloneCmd + provider + user + "/" + repo
def proc = cloneCmd.execute()
proc.waitFor()

/*
//desync
def gitDir = new File(repo + '\\.git') 
def result = gitDir.deleteDir()  // Returns true if all goes well, false otherwise.
assert result
*/

//sloc
def slocCmd = 'cmd /c sloc ' + repo + ' > ' + repo + '\\sloc.log'
slocCmd.execute()


ant = new AntBuilder()
ant.sequential {
  delete(file:repo + '\\.gitignore')  
  delete(dir:repo + '\\.git')  
  move (todir: tempDir) {
    fileset (dir: repo)
  }
}



println "done."
