@Echo	@Del Define.Core.tmpl	/Q /S
	@Del Define.Core.tmpl	/Q /S >NUL:

@Echo	@Call duplicate.cmd
@cd .\sss
	@Call duplicate.cmd
@cd ..

@Echo Please Check git
@Pause

@Echo	@Del *.dot.go		/Q /S
	@Del *.dot.go		/Q /S >NUL:
@Echo	@Del *.ugo		/Q /S
	@Del *.ugo		/Q /S >NUL:

@Echo	@Call dot -ugo -x .....
	@Call dot -ugo -x .....
@Echo	@Call go test .\...
	@Call go test .\...