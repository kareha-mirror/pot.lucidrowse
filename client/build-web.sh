#!/bin/sh
flutter build web --release --pwa-strategy=none --dart-define-from-file=env/demo.json
date +%s > build/web/version.txt
