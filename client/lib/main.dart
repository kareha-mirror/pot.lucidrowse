import 'package:flutter/material.dart';

import 'package:client/app.dart';
import 'package:client/state/app_state.dart';

void main() {
  final state = AppState();
  runApp(MyApp(state: state));
}
