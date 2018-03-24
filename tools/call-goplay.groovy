def proc = /goplay -openbrowser=false -run=false D:\dbt\01\git\gobyes\corpus\codegangsta\bwag\controllers\example.go/.execute()
def b = new StringBuffer()
proc.consumeProcessErrorStream(b)

println proc.text
println b.toString()