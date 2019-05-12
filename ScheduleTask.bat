@echo off
set BW="D:\BingWallpapers\BingWallpapers.exe"
schtasks /create /tn "必应壁纸自动获取A" /tr %BW% /sc daily /st 14:00:00 /ed 2025/05/02
