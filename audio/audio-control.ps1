$SendKey = @"
using System;
using System.Runtime.InteropServices;

public class MediaControl {
    [DllImport("user32.dll")]
    public static extern void keybd_event(byte bVk, byte bScan, uint dwFlags, uint dwExtraInfo);
    
    public static void Next() { keybd_event(0xB1, 0, 0, 0); }
    public static void Prev() { keybd_event(0xB0, 0, 0, 0); }
    public static void PlayPause() { keybd_event(0xB3, 0, 0, 0); }
    public static void VolumeUp() { keybd_event(0xAF, 0, 0, 0); }
    public static void VolumeDown() { keybd_event(0xAE, 0, 0, 0); }
}
"@
Add-Type -TypeDefinition $SendKey

# --- COMMANDS ---

# Un-comment the line you want to run:

[MediaControl]::PlayPause()
# [MediaControl]::Next()
# [MediaControl]::VolumeUp()

Write-Host "Media Command Sent!" -ForegroundColor Green