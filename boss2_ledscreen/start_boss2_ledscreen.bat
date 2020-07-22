rem start boss2_ledscreen.exe
@echo off
Title [start boss2_ledscreen]
if "%1" == "h" goto begin
　　mshta vbscript:createobject("wscript.shell").run("%~nx0 h",0)(window.close)&&exit
:begin
:: 执行TIM
set TIM_HOME="D:\Program Files (x86)\Tencent\Boss2_ledscreen\"
cd /d %TIM_HOME%
start boss2_ledscreen.exe
::end-----------------------------------
pause>nul