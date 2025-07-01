# Руководство по запуску wave2mp3

## 🇷🇺 Русская версия

### 📦 Подготовка
1. Скачайте подходящий бинарник для вашей ОС:
   - `wave2mp3-macos` (macOS)
   - `wave2mp3-linux` (Linux)  
   - `wave2mp3-win.exe` (Windows)
2. Поместите в одну папку:
   - Бинарный файл
   - WAV-файл для конвертации (напр. `music.wav`)

### 🖥 macOS
```bash
# 1. Открыть Терминал (Cmd+Пробел → "Terminal")
# 2. Перейти в папку (где лежит файл и бинарник)
cd ~/path/to/folder

# 3. Дать права
chmod +x wave2mp3-macos

# 4. Запустить
./wave2mp3-macos -i music.wav
# Результат: music.mp3
```


### 🐧 Linux
```bash
# 1. Откройте Терминал (Ctrl+Alt+T)
# 2. Перейти в папку (где лежит файл и бинарник)
cd ~/path/to/folder

# 3. Дать права
chmod +x wave2mp3-linux

# 4. Запустить
./wave2mp3-linux -i music.wav
# Результат: music.mp3
```

### 🪟 Windows
```cmd
:: 1. Откройте Командную Нажмите Win + R, введите cmd, нажмите Enter
:: 2. Перейти в папку (где лежит файл и бинарник, в проводнике можно выделить и копировать путь)
cd \path\to\folder


:: 3. Запустить
wave2mp3-win.exe -i music.wav
:: или
.\wave2mp3-macos -i music.wav
:: Результат: music.mp3
```

### 💡 Примечания
- Файлы должны быть в одной папке
- Для имен с пробелами: -i "my music.wav"
- Логи выводятся в консоль