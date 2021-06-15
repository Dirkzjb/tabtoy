set DIR="%cd%"
.\tabtoy.exe ^
--mode=v2 ^
--csharp_out=.\csharp\Example\Config.cs ^
--binary_out=.\csharp\Example\Config.bin ^
--go_out=.\golang\table\table_gen.go ^
--combinename=Config ^
--lan=zh_cn ^
--InputDir=%DIR%

@IF %ERRORLEVEL% NEQ 0 pause