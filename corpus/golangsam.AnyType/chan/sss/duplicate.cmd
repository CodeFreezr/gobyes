@Setlocal EnableExtensions

@If "%1"=="" (
	@Call %0 Core
	@Call %0 DefineCore
	@Call %0 DotGoTmpl
	@Call %0 CleanUp
	@goto :EOF
)

@If "%2"=="" (
	@Echo Now: %1
	goto %1
)

@Set args= /I /E /V /C /H /R /K /O /X /B /Y /Q
@xcopy %1	%2		%args%
@goto :EOF

:CleanUp
@cd	..\s\
@Del 		Send{{.}}Proxy.dot.go.tmpl
@cd	..\l\
@Del 		Send{{.}}Proxy.dot.go.tmpl
@cd	..\sl\
@Del 		Send{{.}}Proxy.dot.go.tmpl
@cd	..\xs\
@Del 		Send{{.}}Proxy.dot.go.tmpl
@cd	..\xl\
@Del 		Send{{.}}Proxy.dot.go.tmpl

@cd ..\sss\
@Echo Done Cleanup

@goto :EOF

:Core
@xcopy _Core.nonil		Define.Core.tmpl			/I /Y /Q
@xcopy _Core.nonil		basic\type\Define.Core.tmpl		/I /Y /Q
@xcopy _Core.merge		basic\type\IsFloat\Define.Core.tmpl	/I /Y /Q
@xcopy _Core.merge		basic\type\IsInteger\Define.Core.tmpl	/I /Y /Q
@xcopy _Core.merge		basic\type\IsOrdered\Define.Core.tmpl	/I /Y /Q
@xcopy _Core.merge		basic\type\IsUnsigned\Define.Core.tmpl	/I /Y /Q
@xcopy _Core.all		container\Define.Core.tmpl		/I /Y /Q
@xcopy _Core.all		standard\archive\Define.Core.tmpl	/I /Y /Q
@xcopy _Core.all		standard\container\Define.Core.tmpl	/I /Y /Q
@goto :EOF

:DefineCore
@Call %0 *Define.Core.tmpl	..\ss\
@Call %0 *Define.Core.tmpl	..\ssss\
@goto :EOF

:DotGoTmpl
@Call %0 *dot.go.tmpl	..\s\
@Call %0 *dot.go.tmpl	..\ss\

@Call %0 *dot.go.tmpl	..\ssss\
@Call %0 *dot.go.tmpl	..\l\
@Call %0 *dot.go.tmpl	..\sl\
@Call %0 *dot.go.tmpl	..\xs\
@Call %0 *dot.go.tmpl	..\xl\
@goto :EOF
