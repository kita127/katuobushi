# katuobushi
Katsuobushi is a usb serial communication device that runs on your OS terminal.

## Description

Katsuobushi はあなたのマシンのターミナルで USB シリアル通信を実現する CLI ツールです.
起動時に通信するポートとボーレートを設定します.
ポートは `--port` オプションで指定します. 通信対象のポートの調べ方は Windows 環境であれば `MODE` コマンドで調べることが可能です.
ボーレートは `--baud-rate` で設定します(デフォルトは 9600).

### Send
標準入力から1行ごとに入力したテキストをアスキーで送信します.
1行ごとのテキストは改行で区切られます.

### Receive
受信したデータをアスキーで標準出力します.
一度に受信可能なデータサイズは 128 byte です.
受信データはポーリングにより読み出します.
ポーリング周期は `--read-time` オプションで設定可能です(デフォルトは 100ms).

## Installation

## Usage

    usage: katuobushi.exe --port=PORT [<flags>]

    Flags:
      --help            Show context-sensitive help (also try --help-long and
                        --help-man).
      --port=PORT       port name (--port=COM3)
      --baud-rate=9600  baud rate (--baud-rate=9600)
      --read-time=100   read cycle time(ms)


## License
This software is released under the MIT License, see LICENSE.
