set DIR="%cd%"
.\tabtoy.exe ^
--mode=v2 ^
--csharp_client_out=.\csharp\Example\Config.cs ^
--binary_client_out=.\csharp\Example\Config.bin ^
--go_server_out=.\golang\table\table_gen.go ^
--pbt_server_out=.\pb\data.pbt ^
--combinename=Config ^
--lan=zh_cn ^
--InputDir=%DIR%

@IF %ERRORLEVEL% NEQ 0 pause