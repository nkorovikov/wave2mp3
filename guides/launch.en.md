# wave2mp3 User Guide

## 🇬🇧 English Version

### 📦 Preparation
1. Download the appropriate binary for your OS:
   - `wave2mp3-macos` (macOS)
   - `wave2mp3-linux` (Linux)
   - `wave2mp3-win.exe` (Windows)
2. Place in the same folder:
   - The binary file
   - WAV file to convert (e.g. `music.wav`)

### 🍎 macOS
```bash
# 1. Open Terminal (Cmd+Space → type "Terminal")
# 2. Navigate to folder (where the binary and file are located)
cd ~/path/to/folder

# 3. Make executable
chmod +x wave2mp3-macos

# 4. Run
./wave2mp3-macos -i music.wav
# Output: music.mp3
```

### 🐧 Linux
```bash
# 1. Open Terminal (Ctrl+Alt+T)
# 2. Navigate to folder (where the binary and file are located)
cd ~/path/to/folder

# 3. Make executable
chmod +x wave2mp3-linux

# 4. Run
./wave2mp3-linux -i music.wav
# Output: music.mp3
```

### 🪟 Windows
```cmd
:: 1. Open Command Prompt (Win+R → type "cmd")
:: 2. Navigate to folder (you can copy path from File Explorer)
cd \path\to\folder

:: 3. Run
wave2mp3-win.exe -i music.wav
:: or
.\wave2mp3-win.exe -i music.wav
:: Output: music.mp3
```

### 💡 Notes
- Files must be in the same folder
- For filenames with spaces: -i "my music.wav"
- Logs are displayed in the console