
// set location (for e.g. d:\\gitemp)
// set user / repo
// set corpus.dir

// Download  repo in location
// Delete location/user/repo/.git
// Delete corpus/user/repo
// Move location/user/repo to corpus/user/repo
// Run sloc / headcom / pgrounder etc. pp.
// checkin stuff

def tempDir = 'd:\\gitemp'
def corpusDir = 'd:\\dbt\\01\\git\\gobeys\\corpus'



println "1: " + this.args[0]
println "2: " + this.args[1]


//def gitDir = new File('.git') 

//def result = gitDir.deleteDir()  // Returns true if all goes well, false otherwise.
//assert result

// cd gitemp
//git clone git://repo.org/fossproject.git && rm -rf fossproject/.git



//print "sloc".execute().text
//print "cmd /c dir /b /s *.go".execute().text