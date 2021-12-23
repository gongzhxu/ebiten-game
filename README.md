# ebiten-game

## 桌面终端编译

```
go run github.com/gongzhxu/ebiten-game
```

## Android编译

```
export ANDROID_NDK_HOME=/Users/gongzhangxun/Library/Android/android-ndk-r22b
export ANDROID_HOME=/Users/gongzhangxun/Library/Android/sdk

git clone https://github.com/gongzhxu/ebiten-game
cd ebiten-game
go run github.com/hajimehoshi/ebiten/v2/cmd/ebitenmobile bind -target android -javapkg com.gongzhxu.game -o ./mobile/android/gamelib/gamelib.aar ./mobile
```

and run the Android Studio project in `./mobile/android`.


