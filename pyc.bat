@echo off

if [%1] == [] goto show_pys
if exist c:\%1 goto set_path

exit /B 1


:show_pys
for /f "delims=|" %%i in ('pyc_get_path.exe %1') do echo %%i
exit /B

:set_path
for /f "delims=|" %%i in ('pyc_get_path.exe %1') do set PATH=%%i
exit /B
